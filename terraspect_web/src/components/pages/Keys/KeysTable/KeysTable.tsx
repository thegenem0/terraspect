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

export type Key = {
  id: string
  name: string
  description: string
  value: string
  createdAt: string
  actions?: string
}

type KeysTableProps = {
  data: Key[]
}

export function KeysTable({ data }: KeysTableProps) {
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
      accessorKey: 'createdAt',
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
                onSelect={() => console.log('Copy Key', apiKey.value)}
              >
                Copy Key
              </DropdownMenuItem>
              <DropdownMenuSeparator />
              <DropdownMenuItem
                onSelect={() => console.log('View customer', apiKey.name)}
              >
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
