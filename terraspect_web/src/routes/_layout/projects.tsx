import { createFileRoute } from '@tanstack/react-router'
import { useState } from 'react'

import PageContainer from '@/components/common/PageContainer/PageContainer'
import NewProject from '@/components/pages/Projects/NewProject/NewProject'
import ProjectDetails from '@/components/pages/Projects/ProjectDetails/ProjectDetails'
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle
} from '@/components/ui/card'
import {
  Project,
  useAllProjectsQuery
} from '@/hooks/queries/useAllProjectsQuery'

export const Route = createFileRoute('/_layout/projects')({
  component: () => <ProjectsComponent />
})

const ProjectsComponent = () => {
  const { data, isLoading } = useAllProjectsQuery()
  const [selectedProject, setSelectedProject] = useState<Project>()

  return (
    <PageContainer className="w-3/4">
      <Card className="w-full">
        <CardHeader className="flex flex-row justify-between">
          <div>
            <CardTitle>Projects</CardTitle>
            <CardDescription>Manage your projects here.</CardDescription>
          </div>
          <NewProject />
        </CardHeader>
        <CardContent>
          <div className="flex flex-row flex-wrap justify-center gap-8">
            {isLoading ? (
              <SkeletonLoader />
            ) : (
              data?.projects &&
              data.projects.map((project) => (
                <Card
                  className="w-[350px] bg-ts-light-gray hover:bg-ts-dark-gray hover:shadow-lg"
                  key={project.id}
                  onClick={() => setSelectedProject(project)}
                >
                  <CardHeader>
                    <CardTitle>{project.name}</CardTitle>
                    <CardDescription>
                      {project.description || 'No description'}
                    </CardDescription>
                  </CardHeader>
                  <CardFooter className="flex justify-between">
                    <h4 className="text-sm text-black">
                      {project.planCount} plans
                    </h4>
                  </CardFooter>
                </Card>
              ))
            )}
            <ProjectDetails
              project={selectedProject}
              setProjectId={setSelectedProject}
            />
          </div>
        </CardContent>
      </Card>
    </PageContainer>
  )
}

const SkeletonLoader = () => {
  return (
    <>
      {Array.from({ length: 4 }, (_, index) => (
        <Card
          key={index}
          className="w-[350px] animate-pulse space-y-3 rounded-lg bg-slate-200 p-4"
        >
          <CardHeader className="h-6 rounded bg-slate-300"></CardHeader>
          <div className="h-4 rounded bg-slate-300"></div>
          <div className="flex justify-between">
            <div className="h-4 w-1/2 rounded bg-slate-300"></div>
          </div>
        </Card>
      ))}
    </>
  )
}
