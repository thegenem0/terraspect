import toast from 'react-hot-toast'

import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle
} from '@/components/ui/alert-dialog'
import { Button } from '@/components/ui/button'
import { useDeleteProjectMutation } from '@/hooks/mutations/useDeleteProjectMutation'
import { Project } from '@/hooks/queries/useAllProjectsQuery'

type DeleteProjectDialogProps = {
  project?: Project
  isOpen: boolean
  setOpen: (open: boolean) => void
}

const DeleteProjectDialog = ({
  project,
  isOpen,
  setOpen
}: DeleteProjectDialogProps) => {
  const { mutateAsync } = useDeleteProjectMutation()

  const deleteProject = async () => {
    await mutateAsync({ projectId: project?.id ?? '' })
      .then(() => {
        toast.success('Project deleted')
      })
      .catch(() => {
        toast.error('Failed to delete project')
      })
    setOpen(false)
  }

  return (
    <AlertDialog open={isOpen}>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Are you absolutely sure?</AlertDialogTitle>
          <AlertDialogDescription>
            This will permanently delete the project:
            <strong>{project?.name}</strong>. This project has
            {project?.planCount} associated plans.
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogCancel asChild>
            <Button variant="ghost" onClick={() => setOpen(false)}>
              Cancel
            </Button>
          </AlertDialogCancel>
          <AlertDialogAction asChild>
            <Button onClick={() => deleteProject()}>Delete Project</Button>
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  )
}

export default DeleteProjectDialog
