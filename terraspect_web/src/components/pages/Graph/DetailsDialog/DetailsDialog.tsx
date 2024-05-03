import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle
} from '@/components/ui/dialog'
import { ScrollArea } from '@/components/ui/scroll-area'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs'
import { useTreeContext } from '@/contexts/TreeContextProvider'

import ChangesView from '../ChangesView/ChangesView'
import DetailsView from '../DetailsView/DetailsView'

const DetailsDialog = () => {
  const { activeNode, toggleActiveNodeById } = useTreeContext()
  return (
    <Dialog open={!!activeNode}>
      <DialogContent
        className=""
        onClose={() => toggleActiveNodeById(activeNode?.id ?? '')}
      >
        <DialogHeader>
          <DialogTitle>{activeNode?.label}</DialogTitle>
          <DialogDescription>{activeNode?.id}</DialogDescription>
        </DialogHeader>
        <Tabs defaultValue="account" className="w-full">
          <TabsList>
            <TabsTrigger value="account">Resource Overview</TabsTrigger>
            <TabsTrigger value="password">Changes Overview</TabsTrigger>
          </TabsList>
          <TabsContent value="account">
            <ScrollArea className="flex h-[500px] flex-row gap-32">
              <DetailsView activeNode={activeNode} />
            </ScrollArea>
          </TabsContent>
          <TabsContent value="password">
            <ChangesView activeNode={activeNode} />
          </TabsContent>
        </Tabs>
      </DialogContent>
    </Dialog>
  )
}

export default DetailsDialog
