<script setup lang="ts">
import Chat from "@/components/Chat.vue";
import type {
  ColumnDef,
  ColumnFiltersState,
  SortingState,
  VisibilityState,
} from "@tanstack/vue-table";
import {
  FlexRender,
  getCoreRowModel,
  getFilteredRowModel,
  getPaginationRowModel,
  getSortedRowModel,
  useVueTable,
} from "@tanstack/vue-table";
import {
  ArrowUpDown,
  ChevronDown,
  MessageSquare,
  Play,
  Pause,
} from "lucide-vue-next";
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip";
import { Skeleton } from "@/components/ui/skeleton";
import { h, ref, watch, onMounted, computed } from "vue";
import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";
import {
  DropdownMenu,
  DropdownMenuCheckboxItem,
  DropdownMenuContent,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Input } from "@/components/ui/input";
import {
  Table as UiTable,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import {
  Select,
  SelectTrigger,
  SelectContent,
  SelectItem,
  SelectValue,
} from "@/components/ui/select";
import { valueUpdater } from "@/utils/utils";
import { performHttpCall } from "@/utils/http";
import type { Patient } from "@/types/patient";

const emit = defineEmits(["notify"]);

const data = ref<Patient[]>([]);

async function fetchData() {
  const response = await performHttpCall<Patient[]>("patients", "GET");
  data.value = response.patients;
}

async function updatePatientStatus({
  status,
  keywords,
  patient,
}: {
  status: string;
  keywords: string[];
  patient: Patient;
}) {
  console.log("Updating patient status...");
  await performHttpCall(`patients/${patient.uuid}`, "PUT", {
    status,
    keywords,
  });
}

const showTable = ref(false);

const sorting = ref<SortingState>([]);
const columnFilters = ref<ColumnFiltersState>([]);
const columnVisibility = ref<VisibilityState>({});
const rowSelection = ref({});
const statusFilter = ref("all");
const selectedPatient = ref<Patient>({});

const columns: ColumnDef<Patient>[] = [
  {
    id: "select",
    header: ({ table }) =>
      h(Checkbox, {
        checked:
          table.getIsAllPageRowsSelected() ||
          (table.getIsSomePageRowsSelected() && "indeterminate"),
        "onUpdate:checked": (value: boolean) =>
          table.toggleAllPageRowsSelected(!!value),
        ariaLabel: "Select all",
      }),
    cell: ({ row }) =>
      h(Checkbox, {
        checked: row.getIsSelected(),
        "onUpdate:checked": (value: boolean) => row.toggleSelected(!!value),
        ariaLabel: "Select row",
      }),
    enableSorting: false,
    enableHiding: false,
  },
  {
    accessorKey: "status",
    header: ({ column }) => {
      return h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        {
          default: () => ["Status", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })],
        }
      );
    },
    cell: ({ row }) => {
      const status = row.getValue("status");
      return h("div", { class: "flex items-center gap-2" }, [
        h("div", {
          class: `status-circle w-4 h-4 rounded-full mr-4 bg-${status}-500`,
        }),
        h(
          Button,
          {
            variant: "outline",
            size: "icon",
            onClick: () => console.log("Play clicked"),
          },
          {
            default: () => h(Play, { class: "w-4 h-4" }),
          }
        ),
        h(
          Button,
          {
            variant: "outline",
            size: "icon",
            onClick: () => console.log("Pause clicked"),
          },
          {
            default: () => h(Pause, { class: "w-4 h-4" }),
          }
        ),
      ]);
    },
    sortingFn: (a, b) => {
      return (
        statusOrder.indexOf(a.original.status) -
        statusOrder.indexOf(b.original.status)
      );
    },
  },
  {
    accessorKey: "phone",
    header: ({ column }) => {
      return h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        {
          default: () => [
            "N° de téléphone",
            h(ArrowUpDown, { class: "ml-2 h-4 w-4" }),
          ],
        }
      );
    },
    cell: ({ row }) => h("div", { class: "lowercase" }, row.getValue("phone")),
  },
  {
    id: "actions",
    header: "Actions",
    cell: ({ row }) => {
      const patient = row.original;
      return h(
        Button,
        {
          variant: "outline",
          size: "icon",
          onClick: () => {
            chatIsOpened.value = true;
            selectedPatient.value = patient;
          },
        },
        {
          default: () => h(MessageSquare, { class: "w-4 h-4" }),
        }
      );
    },
    enableHiding: false,
  },
];

const table = ref(
  useVueTable({
    data: data.value,
    columns,
    getCoreRowModel: getCoreRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
    getSortedRowModel: getSortedRowModel(),
    getFilteredRowModel: getFilteredRowModel(),
    onSortingChange: (updaterOrValue) => valueUpdater(updaterOrValue, sorting),
    onColumnFiltersChange: (updaterOrValue) =>
      valueUpdater(updaterOrValue, columnFilters),
    onColumnVisibilityChange: (updaterOrValue) =>
      valueUpdater(updaterOrValue, columnVisibility),
    onRowSelectionChange: (updaterOrValue) =>
      valueUpdater(updaterOrValue, rowSelection),
    state: {
      get sorting() {
        return sorting.value;
      },
      get columnFilters() {
        return columnFilters.value;
      },
      get columnVisibility() {
        return columnVisibility.value;
      },
      get rowSelection() {
        return rowSelection.value;
      },
      set rowSelection(value) {
        rowSelection.value = value;
      },
    },
  })
);

onMounted(async () => {
  await fetchData();
  await new Promise((resolve) => setTimeout(resolve, 2000));
  table.value = useVueTable({
    data: data.value,
    columns,
    getCoreRowModel: getCoreRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
    getSortedRowModel: getSortedRowModel(),
    getFilteredRowModel: getFilteredRowModel(),
    onSortingChange: (updaterOrValue) => valueUpdater(updaterOrValue, sorting),
    onColumnFiltersChange: (updaterOrValue) =>
      valueUpdater(updaterOrValue, columnFilters),
    onColumnVisibilityChange: (updaterOrValue) =>
      valueUpdater(updaterOrValue, columnVisibility),
    onRowSelectionChange: (updaterOrValue) =>
      valueUpdater(updaterOrValue, rowSelection),
    state: {
      get sorting() {
        return sorting.value;
      },
      get columnFilters() {
        return columnFilters.value;
      },
      get columnVisibility() {
        return columnVisibility.value;
      },
      get rowSelection() {
        return rowSelection.value;
      },
      set rowSelection(value) {
        rowSelection.value = value;
      },
    },
  });
  showTable.value = true;
});

const chatIsOpened = ref(false);
watch(statusFilter, (newValue) => {
  const statusColumn = table.value.getColumn("status");
  if (statusColumn) {
    statusColumn.setFilterValue(newValue === "all" ? undefined : newValue);
  }
});

const statusOrder = ["red", "orange", "yellow", "blue", "gray"];

const mutatePatientStatus = async ({
  status,
  keywords,
  patient,
}: {
  status: string;
  keywords: string[];
  patient: Patient;
}) => {
  const index = data.value.findIndex((p) => p.uuid === patient.uuid);
  if (index !== -1) {
    data.value[index].status = status as Patient["status"];
    data.value[index].keywords = keywords;

    await updatePatientStatus({ status, keywords, patient });

    table.value = useVueTable({
      data: data.value,
      columns,
      getCoreRowModel: getCoreRowModel(),
      getPaginationRowModel: getPaginationRowModel(),
      getSortedRowModel: getSortedRowModel(),
      getFilteredRowModel: getFilteredRowModel(),
      onSortingChange: (updaterOrValue) =>
        valueUpdater(updaterOrValue, sorting),
      onColumnFiltersChange: (updaterOrValue) =>
        valueUpdater(updaterOrValue, columnFilters),
      onColumnVisibilityChange: (updaterOrValue) =>
        valueUpdater(updaterOrValue, columnVisibility),
      onRowSelectionChange: (updaterOrValue) =>
        valueUpdater(updaterOrValue, rowSelection),
      state: {
        get sorting() {
          return sorting.value;
        },
        get columnFilters() {
          return columnFilters.value;
        },
        get columnVisibility() {
          return columnVisibility.value;
        },
        get rowSelection() {
          return rowSelection.value;
        },
        set rowSelection(value) {
          rowSelection.value = value;
        },
      },
    });
  }
};

const formattedKeywords = (patientId: string) => {
  console.log("id :" + patientId);
  return data.value.find((p) => p.uuid === patientId)?.keywords.join(", ");
};
</script>

<template>
  <div class="w-full">
    <div class="flex gap-2 items-center py-4">
      <Input
        class="max-w-sm"
        placeholder="Filtrer les numéros de téléphone..."
        :model-value="table.getColumn('phone')?.getFilterValue() ?? ''"
        @update:model-value="
          (value) => table.getColumn('phone')?.setFilterValue(value)
        "
      />

      <Select v-model="statusFilter">
        <SelectTrigger class="w-[180px]">
          <SelectValue placeholder="Filter status" />
        </SelectTrigger>
        <SelectContent>
          <SelectItem value="all">Status</SelectItem>
          <SelectItem value="gray">Gris</SelectItem>
          <SelectItem value="red">Rouge</SelectItem>
          <SelectItem value="orange">Orange</SelectItem>
          <SelectItem value="yellow">Jaune</SelectItem>
          <SelectItem value="blue">Bleu</SelectItem>
        </SelectContent>
      </Select>
      <DropdownMenu>
        <DropdownMenuTrigger as-child>
          <Button variant="outline" class="ml-auto">
            Colonnes <ChevronDown class="ml-2 h-4 w-4" />
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent align="end">
          <DropdownMenuCheckboxItem
            v-for="column in table
              .getAllColumns()
              .filter((column) => column.getCanHide())"
            :key="column.id"
            class="capitalize"
            :checked="column.getIsVisible()"
            @update:checked="
              (value: boolean) => {
                column.toggleVisibility(!!value);
              }
            "
          >
            {{ column.id }}
          </DropdownMenuCheckboxItem>
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
    <div class="rounded-md border">
      <UiTable>
        <TableHeader>
          <TableRow
            v-for="headerGroup in table.getHeaderGroups()"
            :key="headerGroup.id"
          >
            <TableHead v-for="header in headerGroup.headers" :key="header.id">
              <FlexRender
                v-if="!header.isPlaceholder"
                :render="header.column.columnDef.header"
                :props="header.getContext()"
              />
            </TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          <template v-if="table.getRowModel().rows?.length">
            <TableRow
              v-for="row in table.getRowModel().rows"
              :key="row.id"
              :data-state="row.getIsSelected() && 'selected'"
            >
              <TableCell v-for="cell in row.getVisibleCells()" :key="cell.id">
                <TooltipProvider>
                  <Tooltip>
                    <TooltipTrigger
                      ><FlexRender
                        :render="cell.column.columnDef.cell"
                        :props="cell.getContext()"
                    /></TooltipTrigger>
                    <TooltipContent v-if="cell.column.id === 'status'">
                      <p>
                        {{ formattedKeywords(row.original.uuid) }}
                      </p>
                    </TooltipContent>
                  </Tooltip>
                </TooltipProvider>
              </TableCell>
            </TableRow>
          </template>

          <TableRow v-else>
            <TableCell :colspan="columns.length" class="h-24 text-center">
              <Skeleton class="w-[100px] h-5 rounded-full" />
            </TableCell>
          </TableRow>
        </TableBody>
      </UiTable>
    </div>

    <div class="flex items-center justify-end space-x-2 py-4">
      <div class="flex-1 text-sm text-muted-foreground">
        {{ table.getFilteredSelectedRowModel().rows.length }} of
        {{ table.getFilteredRowModel().rows.length }} row(s) selected.
      </div>
      <div class="space-x-2">
        <Button
          variant="outline"
          size="sm"
          :disabled="!table.getCanPreviousPage()"
          @click="table.previousPage()"
        >
          Previous
        </Button>
        <Button
          variant="outline"
          size="sm"
          :disabled="!table.getCanNextPage()"
          @click="table.nextPage()"
        >
          Next
        </Button>
      </div>
    </div>
  </div>
  <Chat
    v-model="selectedPatient"
    :patient="selectedPatient"
    :isOpened="chatIsOpened"
    @new-status="mutatePatientStatus"
    @notify="emit('notify', $event)"
    @close="chatIsOpened = false"
  />
</template>
