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

export const Route = createFileRoute('/__authenticated/projects')({
  component: () => <ProjectsComponent />
})

const ProjectsComponent = () => {
  const { data, isLoading, isError } = useAllProjectsQuery()
  const [selectedProject, setSelectedProject] = useState<Project>()

  return (
    <>
      <div className="mx-auto flex w-2/3 flex-row justify-between rounded-lg bg-slate-300 p-3">
        <h1 className="text-center text-3xl font-semibold">Projects</h1>
        <NewProject />
      </div>
      <PageContainer className="">
        <div className="flex flex-row flex-wrap justify-center gap-8">
          {data?.projects &&
            data.projects.map((project) => (
              <Card
                className="w-[350px] hover:bg-slate-300 hover:shadow-lg"
                key={project.id}
                onClick={() => setSelectedProject(project)}
              >
                <CardHeader>
                  <CardTitle>{project.name}</CardTitle>
                  <CardDescription>
                    {project.description || 'No description'}
                  </CardDescription>
                </CardHeader>
                <CardContent></CardContent>
                <CardFooter className="flex justify-between">
                  <h4 className="text-sm text-slate-400">
                    {project.planCount} plans
                  </h4>
                </CardFooter>
              </Card>
            ))}
          <ProjectDetails
            project={selectedProject}
            setProjectId={setSelectedProject}
          />
        </div>
      </PageContainer>
    </>
  )
}
