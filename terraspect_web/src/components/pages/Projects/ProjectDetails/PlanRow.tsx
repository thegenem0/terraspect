import { Link } from '@tanstack/react-router'
import { Eye } from 'lucide-react'

import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { Plan } from '@/hooks/queries/useAllPlansQuery'

const PlanRow = ({ projectId, plan }: { projectId?: string; plan: Plan }) => {
  return (
    <div className="my-4 grid grid-cols-4 items-center gap-2 rounded-lg border-2 border-black p-2">
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
        <div className="flex flex-row items-center gap-2">
          <Link
            to="/graph"
            search={{
              planId: plan.id,
              projectId: projectId ?? ''
            }}
          >
            <Button type="button">
              <Eye size={16} />
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

export { PlanRow, PlanRowNoData, PlanRowSkeleton }
