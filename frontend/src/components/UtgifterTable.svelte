<script lang="ts">
  import { onMount } from "svelte";
  import * as Table from "$lib/components/ui/table/index.js";
  // Svelte checker may fail to resolve this package types, but Vite resolves it at build/runtime.
  // @ts-ignore
  import { createTable, getCoreRowModel, getFilteredRowModel, getSortedRowModel, type ColumnDef } from "@tanstack/table-core";
  import { Trash2, ArrowUp, ArrowDown, ChevronsUpDown } from "@lucide/svelte";
  import { GetUtgifter, DeleteUtgift } from "../../wailsjs/go/main/App.js";
  import type { models } from "../../wailsjs/go/models.js";

  let utgifter = $state<models.Utgift[]>([]);
  let loading = $state(true);
  let error = $state<string | null>(null);

  // Filters
  let ownerFilter = $state("all");
  let fastFilter = $state("all");
  let dateFrom = $state("");
  let dateTo = $state("");

  // Sorting
  let sorting = $state<{ id: string; desc: boolean }[]>([]);

  let { onDeleted = () => {} } = $props();

  async function fetchUtgifter() {
    try {
      loading = true;
      error = null;
      utgifter = await GetUtgifter();
    } catch (e) {
      error = e instanceof Error ? e.message : String(e);
    } finally {
      loading = false;
    }
  }

  onMount(fetchUtgifter);

  export async function refetch() {
    await fetchUtgifter();
  }

  async function handleDelete(id: number) {
    try {
      await DeleteUtgift(id);
      await fetchUtgifter();
      onDeleted(id);
    } catch (e) {
      error = e instanceof Error ? e.message : String(e);
    }
  }

  function formatCurrency(amount: number): string {
    return new Intl.NumberFormat("nb-NO", {
      style: "currency",
      currency: "NOK",
    }).format(amount);
  }

  function formatDate(dateStr: any): string {
    if (!dateStr) return "";
    const date = new Date(dateStr);
    return new Intl.DateTimeFormat("nb-NO").format(date);
  }

  type CellCtx = { row: { original: models.Utgift } };
  type FilterRow = { original: models.Utgift };

  const columns: ColumnDef<models.Utgift>[] = [
    {
      accessorKey: "beløp",
      header: "Beløp",
      cell: (ctx: CellCtx) => formatCurrency(ctx.row.original["beløp"]),
    },
    {
      accessorKey: "navn",
      header: "Navn",
      cell: (ctx: CellCtx) => ctx.row.original.navn,
    },
    {
      id: "dato",
      header: "Dato",
      accessorFn: (row: models.Utgift) =>
        row.dato ? new Date(row.dato).getTime() : 0,
      cell: (ctx: CellCtx) => formatDate(ctx.row.original.dato),
      filterFn: (row: FilterRow, _columnId: string, filterValue: unknown) => {
        const { from, to } = filterValue as { from: string; to: string };
        const d = row.original.dato ? new Date(row.original.dato) : null;
        if (!d) return false;
        if (from && d < new Date(from)) return false;
        if (to) {
          const end = new Date(to);
          end.setHours(23, 59, 59, 999);
          if (d > end) return false;
        }
        return true;
      },
    },
    {
      accessorKey: "fast",
      header: "Fast",
      cell: (ctx: CellCtx) => (ctx.row.original.fast ? "Ja" : "Nei"),
      filterFn: (row: FilterRow, _columnId: string, filterValue: unknown) => {
        if (filterValue === "all") return true;
        return row.original.fast === (filterValue === "ja");
      },
    },
    {
      id: "eier",
      header: "Eier",
      accessorFn: (row: models.Utgift) => row.eier?.navn ?? "",
      cell: (ctx: CellCtx) => ctx.row.original.eier?.navn ?? "-",
      filterFn: (row: FilterRow, _columnId: string, filterValue: unknown) => {
        if (filterValue === "all") return true;
        return String(row.original.eier_id ?? "") === String(filterValue);
      },
    },
    {
      id: "actions",
      header: "",
      cell: (ctx: CellCtx) => ctx.row.original,
      enableSorting: false,
    },
  ];

  const owners = $derived.by(() => {
    const map = new Map<number, string>();
    for (const u of utgifter) {
      if (u.eier_id != null && u.eier?.navn) {
        map.set(u.eier_id, u.eier.navn);
      }
    }
    return Array.from(map, ([id, navn]) => ({ id, navn })).sort((a, b) =>
      a.navn.localeCompare(b.navn, "nb"),
    );
  });

  const columnFilters = $derived.by(() => {
    const filters: { id: string; value: unknown }[] = [];
    if (ownerFilter !== "all") filters.push({ id: "eier", value: ownerFilter });
    if (fastFilter !== "all") filters.push({ id: "fast", value: fastFilter });
    if (dateFrom || dateTo)
      filters.push({ id: "dato", value: { from: dateFrom, to: dateTo } });
    return filters;
  });

  function resetFilters() {
    ownerFilter = "all";
    fastFilter = "all";
    dateFrom = "";
    dateTo = "";
  }

  const hasActiveFilters = $derived(
    ownerFilter !== "all" || fastFilter !== "all" || !!dateFrom || !!dateTo,
  );

  const table = $derived(
    createTable<models.Utgift>({
      data: utgifter,
      columns,
      state: {
        columnPinning: { left: [], right: [] },
        columnFilters,
        sorting,
      },
      onSortingChange: (updater: unknown) => {
        sorting =
          typeof updater === "function" ? (updater as any)(sorting) : updater;
      },
      getCoreRowModel: getCoreRowModel(),
      getFilteredRowModel: getFilteredRowModel(),
      getSortedRowModel: getSortedRowModel(),
    }),
  );

  const selectClass =
    "flex h-9 rounded-md border border-input bg-transparent px-3 py-1 text-sm shadow-sm transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring";
</script>

<div class="w-full space-y-4">
  <div class="rounded-lg border bg-card text-card-foreground shadow-sm">
    <div class="flex flex-col space-y-1.5 p-6">
      <h3 class="text-2xl font-semibold leading-none tracking-tight">
        Utgifter
      </h3>
    </div>
    <div class="p-6 pt-0">
      {#if loading}
        <p class="text-sm text-muted-foreground">Laster...</p>
      {:else if error}
        <div class="rounded-md bg-destructive/15 p-3">
          <p class="text-sm font-medium text-destructive">Feil: {error}</p>
        </div>
      {:else}
        <div
          class="mb-4 flex flex-wrap items-end gap-3 rounded-md border bg-muted/30 p-3"
        >
          <div class="flex flex-col gap-1">
            <label
              for="filter-eier"
              class="text-xs font-medium text-muted-foreground">Eier</label
            >
            <select id="filter-eier" bind:value={ownerFilter} class={selectClass}>
              <option value="all">Alle</option>
              {#each owners as owner}
                <option value={String(owner.id)}>{owner.navn}</option>
              {/each}
            </select>
          </div>

          <div class="flex flex-col gap-1">
            <label
              for="filter-fast"
              class="text-xs font-medium text-muted-foreground">Fast</label
            >
            <select id="filter-fast" bind:value={fastFilter} class={selectClass}>
              <option value="all">Alle</option>
              <option value="ja">Ja</option>
              <option value="nei">Nei</option>
            </select>
          </div>

          <div class="flex flex-col gap-1">
            <label
              for="filter-date-from"
              class="text-xs font-medium text-muted-foreground">Fra dato</label
            >
            <input
              id="filter-date-from"
              type="date"
              bind:value={dateFrom}
              class={selectClass}
            />
          </div>

          <div class="flex flex-col gap-1">
            <label
              for="filter-date-to"
              class="text-xs font-medium text-muted-foreground">Til dato</label
            >
            <input
              id="filter-date-to"
              type="date"
              bind:value={dateTo}
              class={selectClass}
            />
          </div>

          {#if hasActiveFilters}
            <button
              type="button"
              onclick={resetFilters}
              class="inline-flex h-9 items-center rounded-md px-3 text-sm font-medium text-muted-foreground transition-colors hover:bg-muted hover:text-foreground"
            >
              Nullstill
            </button>
          {/if}
        </div>

        <div class="overflow-hidden rounded-md border">
          <Table.Root>
            <Table.Header class="bg-muted/50">
              {#each table.getHeaderGroups() as headerGroup (headerGroup.id)}
                <Table.Row class="hover:bg-transparent">
                  {#each headerGroup.headers as header (header.id)}
                    <Table.Head
                      class={header.column.id === "beløp"
                        ? "text-end"
                        : header.column.id === "actions"
                          ? "w-15"
                          : ""}
                    >
                      {#if !header.isPlaceholder}
                        {#if header.column.getCanSort()}
                          <button
                            type="button"
                            onclick={header.column.getToggleSortingHandler()}
                            class="flex w-full items-center gap-1 select-none hover:text-foreground {header.column.id ===
                            'beløp'
                              ? 'flex-row-reverse'
                              : 'justify-start'}"
                          >
                            {header.column.columnDef.header as string}
                            {#if header.column.getIsSorted() === "asc"}
                              <ArrowUp class="h-3.5 w-3.5" />
                            {:else if header.column.getIsSorted() === "desc"}
                              <ArrowDown class="h-3.5 w-3.5" />
                            {:else}
                              <ChevronsUpDown
                                class="h-3.5 w-3.5 text-muted-foreground/50"
                              />
                            {/if}
                          </button>
                        {:else}
                          {header.column.columnDef.header as string}
                        {/if}
                      {/if}
                    </Table.Head>
                  {/each}
                </Table.Row>
              {/each}
            </Table.Header>
            <Table.Body>
              {#if table.getRowModel().rows.length === 0}
                <Table.Row>
                  <Table.Cell class="h-24 text-center" colspan={columns.length}
                    >Ingen utgifter funnet.</Table.Cell
                  >
                </Table.Row>
              {:else}
                {#each table.getRowModel().rows as row (row.id)}
                  <Table.Row>
                    {#each row.getVisibleCells() as cell (cell.id)}
                      <Table.Cell
                        class={cell.column.id === "beløp"
                          ? "text-end font-medium tabular-nums"
                          : cell.column.id === "actions"
                            ? "w-15 text-left"
                            : "text-left"}
                      >
                        {#if cell.column.id === "actions"}
                          <button
                            onclick={() => handleDelete(row.original.id)}
                            class="inline-flex items-center justify-center whitespace-nowrap rounded-md text-xs font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 disabled:pointer-events-none disabled:opacity-50 h-5 w-5 shrink-0 bg-destructive/10 text-destructive hover:bg-destructive/20"
                          >
                            <Trash2 class="h-3.5 w-3.5" />
                          </button>
                        {:else if cell.column.id === "navn"}
                          <span class="font-medium">{row.original.navn}</span>
                        {:else if cell.column.id === "eier"}
                          {row.original.eier?.navn ?? "-"}
                        {:else if cell.column.id === "dato"}
                          {formatDate(row.original.dato)}
                        {:else if cell.column.id === "fast"}
                          {row.original.fast ? "Ja" : "Nei"}
                        {:else if cell.column.id === "beløp"}
                          {formatCurrency(row.original["beløp"])}
                        {/if}
                      </Table.Cell>
                    {/each}
                  </Table.Row>
                {/each}
              {/if}
            </Table.Body>
          </Table.Root>
        </div>
      {/if}
    </div>
  </div>
</div>
