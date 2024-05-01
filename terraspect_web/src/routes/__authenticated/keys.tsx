import { createFileRoute } from '@tanstack/react-router'

import KeysTable from '@/components/pages/Keys/KeysTable/KeysTable'
import { Button } from '@/components/ui/button'
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle
} from '@/components/ui/card'

export const Route = createFileRoute('/__authenticated/keys')({
  component: () => <KeysComponent />
})

const KeysComponent = () => {
  const keys = [
    {
      id: '1',
      name: 'Key 1',
      description: 'Key 1 Description',
      value: 'key1',
      createdAt: '2021-09-20'
    },
    {
      id: '2',
      name: 'Key 2',
      description: 'Key 2 Description',
      value: 'key2',
      createdAt: '2021-09-21'
    },
    {
      id: '3',
      name: 'Key 3',
      description: 'Key 3 Description',
      value: 'key3',
      createdAt: '2021-09-22'
    }
  ]

  return (
    <div className="flex h-full flex-row justify-center gap-4">
      <Card className="h-fit w-1/3 bg-slate-200">
        <CardHeader className="flex flex-row justify-between">
          <div>
            <CardTitle>API Keys</CardTitle>
            <CardDescription>Manage your API keys here.</CardDescription>
          </div>
          <Button variant="default">Create Key</Button>
        </CardHeader>
        <CardContent>
          <KeysTable data={keys} />
        </CardContent>
      </Card>
    </div>
  )
}
