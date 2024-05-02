import { Link } from '@tanstack/react-router'
import { Edit, Eye, Trash } from 'lucide-react'

import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { ScrollArea } from '@/components/ui/scroll-area'
import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle
} from '@/components/ui/sheet'
import { Plan, useAllPlansQuery } from '@/hooks/queries/useAllPlansQuery'
import { Project } from '@/hooks/queries/useAllProjectsQuery'

type ProjectDetailsProps = {
  project?: Project
  setProjectId: (project?: Project) => void
}

const ProjectDetails = ({ project, setProjectId }: ProjectDetailsProps) => {
  const { data, isLoading } = useAllPlansQuery({
    projectId: project?.id.toString()
  })

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
            data.plans.map((plan) => <PlanRow key={plan.id} plan={plan} />)
          ) : (
            <PlanRowNoData />
          )}
        </ScrollArea>
      </SheetContent>
    </Sheet>
  )
}

export default ProjectDetails

const PlanRow = ({ plan }: { plan: Plan }) => {
  return (
    <div className="my-4 grid grid-cols-4 items-center gap-2 rounded-lg border-2 border-black py-2">
      <Label htmlFor="name" className="text-right">
        ID:
      </Label>
      <Input id="name" defaultValue={plan.id} className="col-span-3" />
      <Label htmlFor="username" className="text-right">
        Created At:
      </Label>
      <Input
        id="username"
        defaultValue={plan.createdAt}
        className="col-span-3"
      />
      <div className="col-span-4 grid w-full grid-cols-1 place-items-center gap-2 px-2">
        <div className="flex flex-row gap-2">
          <Link to="/plans/:id">
            <Button type="button">
              <Eye size={16} />
            </Button>
          </Link>
          <Link to="/plans/:id/edit">
            <Button type="button">
              <Edit size={16} />
            </Button>
          </Link>
          <Link to="/plans/:id/delete">
            <Button type="button" variant="destructive">
              <Trash size={16} />
            </Button>
          </Link>
        </div>
      </div>
    </div>
  )
}

const PlanRowSkeleton = () => {
  return (
    <div className="my-4 grid animate-pulse grid-cols-4 items-center gap-2 rounded-lg border-2 border-gray-300 py-2">
      <div className="col-span-1 h-4 bg-gray-200"></div>
      <div className="col-span-3 h-4 bg-gray-200"></div>
      <div className="col-span-1 h-4 bg-gray-200"></div>
      <div className="col-span-3 h-4 bg-gray-200"></div>
      <div className="col-span-4 flex justify-center gap-4 p-2">
        <div className="size-10 rounded-full bg-gray-200"></div>
        <div className="size-10 rounded-full bg-gray-200"></div>
        <div className="size-10 rounded-full bg-gray-200"></div>
      </div>
    </div>
  )
}

const PlanRowNoData = () => {
  return (
    <div className="my-4 grid grid-cols-4 items-center gap-2 rounded-lg border-2 border-gray-300 py-2">
      <div className="col-span-4 flex justify-center gap-4 p-2">
        <p>There are no plans for this project yet</p>
      </div>
    </div>
  )
}
