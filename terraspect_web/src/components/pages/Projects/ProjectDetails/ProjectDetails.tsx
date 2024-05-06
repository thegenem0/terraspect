import { useState } from 'react'

import { Button } from '@/components/ui/button'
import { ScrollArea } from '@/components/ui/scroll-area'
import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle
} from '@/components/ui/sheet'
import { useAllPlansQuery } from '@/hooks/queries/useAllPlansQuery'
import { Project } from '@/hooks/queries/useAllProjectsQuery'

import DeleteProjectDialog from './DeleteProjectDialog'
import { PlanRow, PlanRowNoData, PlanRowSkeleton } from './PlanRow'

type ProjectDetailsProps = {
  project?: Project
  setProjectId: (project?: Project) => void
}

const ProjectDetails = ({ project, setProjectId }: ProjectDetailsProps) => {
  const { data, isLoading } = useAllPlansQuery({
    projectId: project?.id.toString()
  })
  const [deleteOpen, setDeleteOpen] = useState(false)

  return (
    <Sheet open={!!project}>
      <SheetContent onClose={() => setProjectId(undefined)}>
        <SheetHeader>
          <SheetTitle>{project?.name}</SheetTitle>
          <SheetDescription>
            {project?.description || 'No description'}
          </SheetDescription>
        </SheetHeader>
        <h3 className="text-lg font-semibold">Plans</h3>
        <ScrollArea className="flex h-5/6 flex-col p-4">
          {isLoading ? (
            Array.from({ length: 5 }, (_, index) => (
              <PlanRowSkeleton key={index} />
            ))
          ) : data?.plans && data?.plans.length > 0 ? (
            data.plans.map((plan) => (
              <PlanRow key={plan.id} projectId={project?.id} plan={plan} />
            ))
          ) : (
            <PlanRowNoData />
          )}
        </ScrollArea>
        <Button
          variant="destructive"
          onClick={() => setDeleteOpen(true)}
          className="w-full"
        >
          Delete Project
        </Button>
      </SheetContent>
      <DeleteProjectDialog
        project={project}
        isOpen={deleteOpen}
        setOpen={setDeleteOpen}
      />
    </Sheet>
  )
}

export default ProjectDetails
