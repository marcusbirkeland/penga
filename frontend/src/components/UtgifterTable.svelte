<script lang="ts">
  import { onMount } from "svelte";
  import * as Table from "$lib/components/ui/table/index.js";
  import { Button } from "$lib/components/ui/button/index.js";
  import { Trash2 } from "@lucide/svelte";
  import { GetUtgifter, DeleteUtgift } from "../../wailsjs/go/main/App.js";
  import type { models } from "../../wailsjs/go/models.js";

  let utgifter = $state<models.Utgift[]>([]);
  let loading = $state(true);
  let error = $state<string | null>(null);

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
</script>

<div class="w-full space-y-4">
  <div class="rounded-lg border bg-card text-card-foreground shadow-sm">
    <div class="flex flex-col space-y-1.5 p-6">
      <h3 class="text-2xl font-semibold leading-none tracking-tight">
        Utgifter
      </h3>
      <p class="text-sm text-muted-foreground">
        En oversikt over alle dine utgifter.
      </p>
    </div>
    <div class="p-6 pt-0">
      {#if loading}
        <p class="text-sm text-muted-foreground">Laster...</p>
      {:else if error}
        <div class="rounded-md bg-destructive/15 p-3">
          <p class="text-sm font-medium text-destructive">Feil: {error}</p>
        </div>
      {:else}
        <Table.Root>
          <Table.Header>
            <Table.Row>
              <Table.Head class="w-[150px]">Navn</Table.Head>
              <Table.Head>Eier</Table.Head>
              <Table.Head>Beskrivelse</Table.Head>
              <Table.Head>Type</Table.Head>
              <Table.Head>Dato</Table.Head>
              <Table.Head>Fast</Table.Head>
              <Table.Head class="text-end">Beløp</Table.Head>
              <Table.Head class="w-[60px]">Aksjoner</Table.Head>
            </Table.Row>
          </Table.Header>
          <Table.Body>
            {#each utgifter as utgift (utgift.id)}
              <Table.Row>
                <Table.Cell class="font-medium">{utgift.navn}</Table.Cell>
                <Table.Cell>{utgift.eier?.navn ?? "-"}</Table.Cell>
                <Table.Cell>{utgift.beskrivelse || "-"}</Table.Cell>
                <Table.Cell>{utgift.type ?? "-"}</Table.Cell>
                <Table.Cell>{formatDate(utgift.dato)}</Table.Cell>
                <Table.Cell>{utgift.fast ? "Ja" : "Nei"}</Table.Cell>
                <Table.Cell class="text-end"
                  >{formatCurrency(utgift["beløp"])}</Table.Cell
                >
                <Table.Cell>
                  <button
                    onclick={() => handleDelete(utgift.id)}
                    class="inline-flex items-center justify-center whitespace-nowrap rounded-md text-xs font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 disabled:pointer-events-none disabled:opacity-50 h-5 w-5 shrink-0 bg-destructive/10 text-destructive hover:bg-destructive/20"
                  >
                    <Trash2 class="h-3.5 w-3.5" />
                  </button>
                </Table.Cell>
              </Table.Row>
            {/each}
          </Table.Body>
        </Table.Root>
      {/if}
    </div>
  </div>
</div>
