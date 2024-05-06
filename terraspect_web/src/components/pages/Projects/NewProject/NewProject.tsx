import { zodResolver } from '@hookform/resolvers/zod'
import { useForm } from 'react-hook-form'
import { toast } from 'react-hot-toast'
import { z } from 'zod'

import { Button } from '@/components/ui/button'
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage
} from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import {
  Sheet,
  SheetClose,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
  SheetTrigger
} from '@/components/ui/sheet'
import { useCreateProjectMutation } from '@/hooks/mutations/useCreateProjectMutation'

export const NewProjectValidationSchema = z.object({
  name: z.string({
    message: 'Name is required'
  }),
  description: z.string({
    message: 'Description is required'
  })
})

const NewProject = () => {
  const { mutateAsync } = useCreateProjectMutation()

  const form = useForm<z.infer<typeof NewProjectValidationSchema>>({
    mode: 'onTouched',
    resolver: zodResolver(NewProjectValidationSchema)
  })

  function onSubmit(values: z.infer<typeof NewProjectValidationSchema>) {
    mutateAsync(values)
      .then(() => {
        form.reset()
        toast.success('Project created')
      })
      .catch(() => {
        toast.error('Failed to create project')
      })
  }

  return (
    <Sheet>
      <SheetTrigger asChild>
        <Button>Create a new project</Button>
      </SheetTrigger>
      <SheetContent>
        <SheetHeader>
          <SheetTitle>Create a new project</SheetTitle>
          <SheetDescription>
            Create a new project and start adding plans to it.
          </SheetDescription>
        </SheetHeader>
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className="grid gap-2">
            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Key Name</FormLabel>
                  <FormControl>
                    <Input placeholder="Project name" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="description"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Key Description</FormLabel>
                  <FormControl>
                    <Input placeholder="Project description" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <SheetClose asChild>
              <Button type="submit">Submit</Button>
            </SheetClose>
          </form>
        </Form>
      </SheetContent>
    </Sheet>
  )
}

export default NewProject
