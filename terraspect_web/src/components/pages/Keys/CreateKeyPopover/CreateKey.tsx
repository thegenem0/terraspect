import { zodResolver } from '@hookform/resolvers/zod'
import { useForm } from 'react-hook-form'
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
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue
} from '@/components/ui/select'
import {
  Sheet,
  SheetClose,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
  SheetTrigger
} from '@/components/ui/sheet'
import { useGenerateKeyMutation } from '@/hooks/mutations/useGenerateKeyMutation'
import { useAllProjectsQuery } from '@/hooks/queries/useAllProjectsQuery'

export const NewApiKeyValidationSchema = z.object({
  name: z.string({
    message: 'Name is required'
  }),
  description: z.string({
    message: 'Description is required'
  }),
  projectId: z.string({
    message: 'Project is required'
  })
})

const CreateKey = () => {
  const { mutateAsync } = useGenerateKeyMutation()

  const form = useForm<z.infer<typeof NewApiKeyValidationSchema>>({
    mode: 'onTouched',
    resolver: zodResolver(NewApiKeyValidationSchema)
  })

  function onSubmit(values: z.infer<typeof NewApiKeyValidationSchema>) {
    mutateAsync(values).then(() => {
      form.reset()
    })
  }

  return (
    <Sheet>
      <SheetTrigger asChild>
        <Button variant="default">Create Key</Button>
      </SheetTrigger>
      <SheetContent>
        <SheetHeader>
          <SheetTitle>name</SheetTitle>
          <SheetDescription>Desc</SheetDescription>
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
                    <Input placeholder="Key name" {...field} />
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
                    <Input placeholder="Key description" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <ProjectSelector form={form} />
            <SheetClose asChild>
              <Button type="submit">Submit</Button>
            </SheetClose>
          </form>
        </Form>
      </SheetContent>
    </Sheet>
  )
}

export default CreateKey

type ProjectSelectorProps = {
  form: ReturnType<typeof useForm<z.infer<typeof NewApiKeyValidationSchema>>>
}

const ProjectSelector = ({ form }: ProjectSelectorProps) => {
  const { data, isLoading } = useAllProjectsQuery()

  return (
    <FormField
      control={form.control}
      name="projectId"
      render={({ field }) => (
        <FormItem className="flex flex-col pb-6">
          <FormLabel>Project</FormLabel>
          <Select onValueChange={field.onChange} defaultValue={field.value}>
            <SelectTrigger className="w-full">
              <SelectValue placeholder="Select Project" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="Loading..." disabled>
                {isLoading ? 'Loading...' : 'Select Project'}
              </SelectItem>
              {data?.projects && data.projects.length > 0 && !isLoading ? (
                data.projects.map((project) => (
                  <SelectItem key={project.id} value={project.id}>
                    {project.name}
                  </SelectItem>
                ))
              ) : (
                <SelectItem value="none" disabled>
                  No projects found
                </SelectItem>
              )}
            </SelectContent>
          </Select>
          <FormMessage />
        </FormItem>
      )}
    />
  )
}
