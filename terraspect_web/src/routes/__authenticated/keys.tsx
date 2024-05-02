import { createFileRoute } from '@tanstack/react-router'

import PageContainer from '@/components/common/PageContainer/PageContainer'
import CreateKeyPopover from '@/components/pages/Keys/CreateKeyPopover/CreateKeyPopover'
import KeysTable from '@/components/pages/Keys/KeysTable/KeysTable'
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle
} from '@/components/ui/card'
import { useGetKeysQuery } from '@/hooks/queries/useGetKeysQuery'

export const Route = createFileRoute('/__authenticated/keys')({
  component: () => <KeysComponent />
})

const KeysComponent = () => {
  const { data } = useGetKeysQuery()

  return (
    <PageContainer className="w-3/4">
      <div className="flex justify-center gap-4">
        <Card className="h-fit w-full bg-slate-200">
          <CardHeader className="flex flex-row justify-between">
            <div>
              <CardTitle>API Keys</CardTitle>
              <CardDescription>Manage your API keys here.</CardDescription>
            </div>
            <CreateKeyPopover />
          </CardHeader>
          <CardContent>
            <KeysTable data={data?.keys} />
          </CardContent>
        </Card>
      </div>
    </PageContainer>
  )
}
