package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds"

	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type InfraStackProps struct {
	awscdk.StackProps
}

func NewInfraStack(scope constructs.Construct, id string, props *InfraStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	existingVpc := awsec2.Vpc_FromLookup(stack, jsii.String("VPC"), &awsec2.VpcLookupOptions{
		IsDefault: jsii.Bool(true),
	})

	dbSecurityGroup := awsec2.NewSecurityGroup(stack, jsii.String("DBSecurityGroup"), &awsec2.SecurityGroupProps{
		Vpc: existingVpc,
	})

	lbSecurityGroup := awsec2.NewSecurityGroup(stack, jsii.String("LBSecurityGroup"), &awsec2.SecurityGroupProps{
		Vpc: existingVpc,
	})

	lbSecurityGroup.AddIngressRule(
		awsec2.Peer_AnyIpv4(),
		awsec2.Port_Tcp(jsii.Number(80)),
		jsii.String("Allow HTTP traffic"),
		jsii.Bool(false),
	)

	lbSecurityGroup.AddIngressRule(
		awsec2.Peer_AnyIpv4(),
		awsec2.Port_Tcp(jsii.Number(3000)),
		jsii.String("Allow HTTP traffic"),
		jsii.Bool(false),
	)

	lbSecurityGroup.AddIngressRule(
		awsec2.Peer_AnyIpv4(),
		awsec2.Port_Tcp(jsii.Number(8501)),
		jsii.String("Allow HTTP traffic"),
		jsii.Bool(false),
	)

	lbSecurityGroup.AddIngressRule(
		awsec2.Peer_AnyIpv4(),
		awsec2.Port_Tcp(jsii.Number(5000)),
		jsii.String("Allow HTTP traffic"),
		jsii.Bool(false),
	)

	ecrRepo := awsecr.Repository_FromRepositoryName(stack, jsii.String("ECRRepo"), jsii.String("terraspect-api"))

	dbSecurityGroup.AddIngressRule(
		awsec2.Peer_SecurityGroupId(
			lbSecurityGroup.SecurityGroupId(),
			lbSecurityGroup.SecurityGroupVpcId(),
		),
		awsec2.Port_Tcp(jsii.Number(5432)),
		jsii.String("Allow Postgres traffic"),
		jsii.Bool(false),
	)

	dbInstance := awsrds.NewDatabaseInstance(
		stack,
		jsii.String("MyPostgresDB"),
		&awsrds.DatabaseInstanceProps{
			Engine: awsrds.DatabaseInstanceEngine_Postgres(&awsrds.PostgresInstanceEngineProps{
				Version: awsrds.PostgresEngineVersion_VER_12_5(),
			}),
			InstanceType: awsec2.InstanceType_Of(
				awsec2.InstanceClass_BURSTABLE2,
				awsec2.InstanceSize_MICRO,
			),
			Vpc: existingVpc,
			VpcSubnets: &awsec2.SubnetSelection{
				SubnetType: awsec2.SubnetType_PUBLIC,
			},
			MultiAz:            jsii.Bool(false),
			AllocatedStorage:   jsii.Number(20),
			StorageType:        awsrds.StorageType_GP2,
			BackupRetention:    awscdk.Duration_Days(jsii.Number(0)),
			DeletionProtection: jsii.Bool(false),
			PubliclyAccessible: jsii.Bool(false),
			SecurityGroups:     &[]awsec2.ISecurityGroup{dbSecurityGroup},
			DatabaseName:       jsii.String("terraspect_db"),
		})

	cluster := awsecs.NewCluster(stack, jsii.String("Cluster"), &awsecs.ClusterProps{
		Vpc: existingVpc,
	})

	taskRole := awsiam.NewRole(stack, jsii.String("TaskRole"), &awsiam.RoleProps{
		AssumedBy: awsiam.NewServicePrincipal(
			jsii.String("ecs-tasks.amazonaws.com"),
			&awsiam.ServicePrincipalOpts{},
		),
		ManagedPolicies: &[]awsiam.IManagedPolicy{
			awsiam.ManagedPolicy_FromManagedPolicyArn(
				stack,
				jsii.String("AmazonECSTaskExecutionRolePolicy"),
				jsii.String("arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"),
			),
			awsiam.ManagedPolicy_FromAwsManagedPolicyName(
				jsii.String("AmazonEC2ContainerRegistryReadOnly"),
			),
		},
	})

	taskDefinition := awsecs.NewFargateTaskDefinition(
		stack,
		jsii.String("TaskDefinition"),
		&awsecs.FargateTaskDefinitionProps{
			Cpu:            jsii.Number(256),
			MemoryLimitMiB: jsii.Number(512),
			ExecutionRole:  taskRole,
		},
	)

	container := taskDefinition.AddContainer(
		jsii.String("terraspect-container"),
		&awsecs.ContainerDefinitionOptions{
			Image: awsecs.ContainerImage_FromRegistry(
				ecrRepo.RepositoryUriForTag(jsii.String("latest")),
				&awsecs.RepositoryImageProps{}),
		},
	)

	container.AddPortMappings(&awsecs.PortMapping{
		ContainerPort: jsii.Number(8000),
		Protocol:      awsecs.Protocol_TCP,
	})

	container.AddEnvironment(jsii.String("DATABASE_URL"), dbInstance.DbInstanceEndpointAddress())
	container.AddEnvironment(jsii.String("DATABASE_USER"), jsii.String("terraspect_root"))
	container.AddEnvironment(jsii.String("DATABASE_PASSWORD"), jsii.String("SuperSecretPassword"))
	container.AddEnvironment(jsii.String("DATABASE_NAME"), jsii.String("terraspect_db"))

	ecsService := awsecs.NewFargateService(stack, jsii.String("ECSService"), &awsecs.FargateServiceProps{
		Cluster:        cluster,
		TaskDefinition: taskDefinition,
		AssignPublicIp: jsii.Bool(true),
		DesiredCount:   jsii.Number(1),
		SecurityGroups: &[]awsec2.ISecurityGroup{
			lbSecurityGroup,
		},
		VpcSubnets: &awsec2.SubnetSelection{
			SubnetType: awsec2.SubnetType_PUBLIC,
		},
	})

	alb := awselasticloadbalancingv2.NewApplicationLoadBalancer(
		stack,
		jsii.String("ALB"),
		&awselasticloadbalancingv2.ApplicationLoadBalancerProps{
			Vpc:            existingVpc,
			InternetFacing: jsii.Bool(true),
			VpcSubnets: &awsec2.SubnetSelection{
				SubnetType: awsec2.SubnetType_PUBLIC,
			},
			SecurityGroup: lbSecurityGroup,
		},
	)

	listener := alb.AddListener(jsii.String("Listener"), &awselasticloadbalancingv2.BaseApplicationListenerProps{
		Port: jsii.Number(80),
		Open: jsii.Bool(true),
	})

	listener.AddTargets(jsii.String("Target"), &awselasticloadbalancingv2.AddApplicationTargetsProps{
		Targets: &[]awselasticloadbalancingv2.IApplicationLoadBalancerTarget{ecsService},
		Port:    jsii.Number(8000),
		//HealthCheck: &awselasticloadbalancingv2.HealthCheck{
		//	Interval:                awscdk.Duration_Seconds(jsii.Number(30)),
		//	Path:                    jsii.String("/"),
		//	Timeout:                 awscdk.Duration_Seconds(jsii.Number(5)),
		//	UnhealthyThresholdCount: jsii.Number(2),
		//	HealthyThresholdCount:   jsii.Number(2),
		//},
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewInfraStack(app, "InfraStack", &InfraStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return &awscdk.Environment{
		Account: jsii.String("211125303136"),
		Region:  jsii.String("eu-west-2"),
	}
}
