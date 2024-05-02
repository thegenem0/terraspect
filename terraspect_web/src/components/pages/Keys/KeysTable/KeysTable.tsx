import { ColumnDef } from '@tanstack/react-table'
import { MoreHorizontal } from 'lucide-react'

import { DataTable } from '@/components/common/DataTable/DataTable'
import { Button } from '@/components/ui/button'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger
} from '@/components/ui/dropdown-menu'
import { useDeleteKeyMutation } from '@/hooks/mutations/useDeleteKeyMutation'
import { Key } from '@/hooks/queries/useGetKeysQuery'

type KeysTableProps = {
  data?: Key[]
}

export function KeysTable({ data }: KeysTableProps) {
  const { mutateAsync } = useDeleteKeyMutation()

  const deleteKey = async (key: string) => {
    await mutateAsync({ key }).then(() => {
      console.log('Key deleted')
    })
  }

  const columns: ColumnDef<Key>[] = [
    {
      accessorKey: 'name',
      header: 'Name'
    },
    {
      accessorKey: 'description',
      header: 'Description'
    },
    {
      accessorKey: 'created_at',
      header: 'Created At'
    },
    {
      id: 'actions',
      cell: ({ row }) => {
        const apiKey = row.original

        return (
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button variant="ghost">
                <MoreHorizontal />
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent>
              <DropdownMenuLabel>Key Actions</DropdownMenuLabel>
              <DropdownMenuItem
                onSelect={() => navigator.clipboard.writeText(apiKey.key)}
              >
                Copy Key
              </DropdownMenuItem>
              <DropdownMenuSeparator />
              <DropdownMenuItem onSelect={() => deleteKey(apiKey.key)}>
                Delete Key
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        )
      }
    }
  ]

  return <DataTable data={data} columns={columns} />
}
export default KeysTable
