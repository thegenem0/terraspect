package main

import (
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfrontorigins"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecspatterns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53targets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3deployment"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"infra/config"
)

type InfraStackProps struct {
	awscdk.StackProps
}

func NewInfraStack(scope constructs.Construct, id string, props *InfraStackProps) awscdk.Stack {
	var stackProps awscdk.StackProps
	if props != nil {
		stackProps = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &stackProps)

	// DEFAULTS

	dbName := jsii.String("terraspect_db")
	clerkSecretName := jsii.String("prod/clerk_api_key")

	// DNS SETTINGS

	hostedZone := awsroute53.HostedZone_FromLookup(stack, jsii.String("Terraspect_ZONE"), &awsroute53.HostedZoneProviderProps{
		DomainName: jsii.String("terraspect.genem0.com"),
	})

	wildcardCertificate := awscertificatemanager.NewCertificate(
		stack,
		jsii.String("Terraspect-wildcard"),
		&awscertificatemanager.CertificateProps{
			DomainName: jsii.String("*.terraspect.genem0.com"),
			Validation: awscertificatemanager.CertificateValidation_FromDns(hostedZone),
		})

	wwwCertificate := awscertificatemanager.Certificate_FromCertificateArn(
		stack,
		jsii.String("Terraspect_Cloudfront_Cert"),
		jsii.String("arn:aws:acm:us-east-1:211125303136:certificate/cae9f667-892e-4212-a526-ab64ca070d99"),
	)

	// VPC

	vpc := awsec2.NewVpc(
		stack,
		jsii.String("prod-vpc"),
		&awsec2.VpcProps{
			MaxAzs:      jsii.Number(3),
			NatGateways: jsii.Number(1),
			SubnetConfiguration: &[]*awsec2.SubnetConfiguration{
				{
					Name:       jsii.String("public"),
					SubnetType: awsec2.SubnetType_PUBLIC,
					CidrMask:   jsii.Number(24),
				},
				{
					Name:       jsii.String("private"),
					SubnetType: awsec2.SubnetType_PRIVATE_WITH_EGRESS,
					CidrMask:   jsii.Number(24),
				},
			},
		},
	)

	// SECRETS

	dbSecret := awssecretsmanager.NewSecret(stack, jsii.String("DBSecret"), &awssecretsmanager.SecretProps{
		SecretName: jsii.String(*stack.StackName() + "-Secret"),
		GenerateSecretString: &awssecretsmanager.SecretStringGenerator{
			SecretStringTemplate: jsii.String(fmt.Sprintf(`{"username":"%s"}`, config.MasterUser(stack))),
			ExcludePunctuation:   jsii.Bool(true),
			IncludeSpace:         jsii.Bool(false),
			GenerateStringKey:    jsii.String("password"),
		},
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	})

	clerkApiKey := awssecretsmanager.Secret_FromSecretNameV2(
		stack,
		jsii.String("CLERK_API_KEY"),
		clerkSecretName,
	)

	// SECURITY

	dbSecurityGroup := awsec2.NewSecurityGroup(stack, jsii.String("PostgresSecurityGroup"), &awsec2.SecurityGroupProps{
		Vpc:              vpc,
		AllowAllOutbound: jsii.Bool(true),
	})

	ecsSecurityGroup := awsec2.NewSecurityGroup(stack, jsii.String("ECSSecurityGroup"), &awsec2.SecurityGroupProps{
		Vpc:              vpc,
		Description:      jsii.String("Security group for ECS tasks"),
		AllowAllOutbound: jsii.Bool(true),
	})

	dbSecurityGroup.AddIngressRule(
		awsec2.SecurityGroup_FromSecurityGroupId(
			stack,
			jsii.String("ECSIngress"),
			ecsSecurityGroup.SecurityGroupId(),
			&awsec2.SecurityGroupImportOptions{},
		),
		awsec2.Port_Tcp(jsii.Number(5432)),
		jsii.String("Allow PostgreSQL access from ECS"),
		jsii.Bool(true),
	)

	dbSecurityGroup.Node().AddDependency(ecsSecurityGroup)

	// DATABASE

	dbInstance := awsrds.NewDatabaseInstance(
		stack,
		jsii.String("terraspect-postgres"),
		&awsrds.DatabaseInstanceProps{
			Engine: awsrds.DatabaseInstanceEngine_Postgres(&awsrds.PostgresInstanceEngineProps{
				Version: awsrds.PostgresEngineVersion_VER_14_11(),
			}),
			InstanceType: awsec2.InstanceType_Of(
				awsec2.InstanceClass_T3,
				awsec2.InstanceSize_MICRO,
			),
			Vpc: vpc,
			VpcSubnets: &awsec2.SubnetSelection{
				SubnetType: awsec2.SubnetType_PUBLIC,
			},
			Credentials:        awsrds.Credentials_FromSecret(dbSecret, jsii.String(config.MasterUser(stack))),
			MultiAz:            jsii.Bool(false),
			AllocatedStorage:   jsii.Number(20),
			StorageType:        awsrds.StorageType_GP2,
			BackupRetention:    awscdk.Duration_Days(jsii.Number(0)),
			DeletionProtection: jsii.Bool(false),
			PubliclyAccessible: jsii.Bool(false),
			SecurityGroups:     &[]awsec2.ISecurityGroup{dbSecurityGroup},
			DatabaseName:       dbName,
		})

	dbInstance.Node().AddDependency(dbSecurityGroup)

	// API CONTAINER SERVICE

	ecsService := awsecspatterns.NewApplicationLoadBalancedFargateService(
		stack,
		jsii.String("terraspect-service"),
		&awsecspatterns.ApplicationLoadBalancedFargateServiceProps{
			Vpc:                vpc,
			AssignPublicIp:     jsii.Bool(true),
			PublicLoadBalancer: jsii.Bool(true),
			RedirectHTTP:       jsii.Bool(false),
			Cpu:                jsii.Number(256),
			MemoryLimitMiB:     jsii.Number(512),
			ListenerPort:       jsii.Number(8080),
			SecurityGroups:     &[]awsec2.ISecurityGroup{ecsSecurityGroup},
			TaskImageOptions: &awsecspatterns.ApplicationLoadBalancedTaskImageOptions{
				Image: awsecs.ContainerImage_FromAsset(
					jsii.String("../terraspect_server"),
					&awsecs.AssetImageProps{},
				),
				ContainerPort: jsii.Number(8080),
				Environment: &map[string]*string{
					"DATABASE_HOST":     dbInstance.DbInstanceEndpointAddress(),
					"DATABASE_NAME":     dbName,
					"DATABASE_USER":     dbSecret.SecretValueFromJson(jsii.String("username")).UnsafeUnwrap(),
					"DATABASE_PASSWORD": dbSecret.SecretValueFromJson(jsii.String("password")).UnsafeUnwrap(),
					"CLERK_API_KEY":     clerkApiKey.SecretValue().UnsafeUnwrap(),
				},
			},
		},
	)

	ecsService.LoadBalancer().AddListener(
		jsii.String("HTTP-HTTPS-Rewrite"),
		&awselasticloadbalancingv2.BaseApplicationListenerProps{
			Port: jsii.Number(80),
			DefaultAction: awselasticloadbalancingv2.ListenerAction_Redirect(
				&awselasticloadbalancingv2.RedirectOptions{
					Port:     jsii.String("443"),
					Protocol: jsii.String("HTTPS"),
				},
			),
		},
	)

	ecsService.LoadBalancer().AddListener(
		jsii.String("HTTPSListener"),
		&awselasticloadbalancingv2.BaseApplicationListenerProps{
			Port: jsii.Number(443),
			Certificates: &[]awselasticloadbalancingv2.IListenerCertificate{
				wildcardCertificate,
			},
			DefaultTargetGroups: &[]awselasticloadbalancingv2.IApplicationTargetGroup{
				ecsService.TargetGroup(),
			},
		},
	)

	awsroute53.NewARecord(stack, jsii.String("API_A_RECORD"), &awsroute53.ARecordProps{
		Zone: hostedZone,
		Target: awsroute53.RecordTarget_FromAlias(
			awsroute53targets.NewLoadBalancerTarget(ecsService.LoadBalancer())),
		RecordName: jsii.String("api.terraspect.genem0.com"), // or use the apex domain
	})

	webBucketName := "terraspect-web-bucket"

	webBucket := awss3.NewBucket(
		stack,
		jsii.String("WebBucket"),
		&awss3.BucketProps{
			BucketName: jsii.String(webBucketName),
			BlockPublicAccess: awss3.NewBlockPublicAccess(
				&awss3.BlockPublicAccessOptions{
					BlockPublicAcls:       jsii.Bool(true),
					BlockPublicPolicy:     jsii.Bool(true),
					IgnorePublicAcls:      jsii.Bool(true),
					RestrictPublicBuckets: jsii.Bool(true),
				},
			),
			RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
		},
	)

	awss3deployment.NewBucketDeployment(
		stack,
		jsii.String("WebDeploymentFiles"),
		&awss3deployment.BucketDeploymentProps{
			DestinationBucket: webBucket,
			Sources: &[]awss3deployment.ISource{
				awss3deployment.Source_Asset(
					jsii.String("../terraspect_web/dist"),
					&awss3assets.AssetOptions{},
				),
			},
		},
	)

	webCloudfrontDistro := awscloudfront.NewDistribution(
		stack,
		jsii.String("WebCloudfrontDistribution"),
		&awscloudfront.DistributionProps{
			Enabled: jsii.Bool(true),
			DefaultBehavior: &awscloudfront.BehaviorOptions{
				Origin: awscloudfrontorigins.NewS3Origin(
					webBucket,
					&awscloudfrontorigins.S3OriginProps{
						OriginAccessIdentity: awscloudfront.NewOriginAccessIdentity(
							stack,
							jsii.String("terraspect-web-OAI"),
							&awscloudfront.OriginAccessIdentityProps{},
						),
					},
				),
				ViewerProtocolPolicy: awscloudfront.ViewerProtocolPolicy_REDIRECT_TO_HTTPS,
			},
			Certificate:       wwwCertificate,
			PriceClass:        awscloudfront.PriceClass_PRICE_CLASS_ALL,
			Comment:           jsii.String("terraspect web distribution"),
			DefaultRootObject: jsii.String("index.html"),
			DomainNames: &[]*string{
				jsii.String("terraspect.genem0.com"),
				jsii.String("www.terraspect.genem0.com"),
			},
			SslSupportMethod: awscloudfront.SSLMethod_SNI,
			ErrorResponses: &[]*awscloudfront.ErrorResponse{
				{
					HttpStatus:         jsii.Number(403),
					ResponseHttpStatus: jsii.Number(200),
					ResponsePagePath:   jsii.String("/index.html"),
				},
				{
					HttpStatus:         jsii.Number(404),
					ResponseHttpStatus: jsii.Number(200),
					ResponsePagePath:   jsii.String("/index.html"),
				},
			},
		},
	)

	webBucket.AddToResourcePolicy(
		awsiam.NewPolicyStatement(
			&awsiam.PolicyStatementProps{
				Actions: &[]*string{jsii.String("s3:GetObject")},
				Effect:  awsiam.Effect_ALLOW,
				Resources: &[]*string{
					jsii.Sprintf("arn:aws:s3:::%s", webBucketName),
					jsii.Sprintf("arn:aws:s3:::%s/*", webBucketName),
				},
				Principals: &[]awsiam.IPrincipal{
					awsiam.NewServicePrincipal(jsii.String("cloudfront.amazonaws.com"), nil),
				},
			},
		),
	)

	awsroute53.NewARecord(stack, jsii.String("WEB_A_RECORD"), &awsroute53.ARecordProps{
		Zone: hostedZone,
		Target: awsroute53.RecordTarget_FromAlias(
			awsroute53targets.NewCloudFrontTarget(webCloudfrontDistro),
		),
		RecordName: jsii.String("terraspect.genem0.com"),
	})

	awsroute53.NewARecord(stack, jsii.String("WEB_WWW_A_RECORD"), &awsroute53.ARecordProps{
		Zone: hostedZone,
		Target: awsroute53.RecordTarget_FromAlias(
			awsroute53targets.NewCloudFrontTarget(webCloudfrontDistro),
		),
		RecordName: jsii.String("www.terraspect.genem0.com"),
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
