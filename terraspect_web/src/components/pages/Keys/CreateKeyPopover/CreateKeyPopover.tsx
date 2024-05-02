import { zodResolver } from '@hookform/resolvers/zod'
import { PopoverClose } from '@radix-ui/react-popover'
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
  Popover,
  PopoverContent,
  PopoverTrigger
} from '@/components/ui/popover'
import { useGenerateKeyMutation } from '@/hooks/mutations/useGenerateKeyMutation'

export const NewApiKeyValidationSchema = z.object({
  name: z.string({
    message: 'Name is required'
  }),
  description: z.string({
    message: 'Description is required'
  })
})

const CreateKeyPopover = () => {
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
    <Popover>
      <PopoverTrigger asChild>
        <Button variant="default">New Key</Button>
      </PopoverTrigger>
      <PopoverContent className="w-96 border border-black">
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
            <PopoverClose asChild>
              <Button type="submit">Submit</Button>
            </PopoverClose>
          </form>
        </Form>
      </PopoverContent>
    </Popover>
  )
}

export default CreateKeyPopover
