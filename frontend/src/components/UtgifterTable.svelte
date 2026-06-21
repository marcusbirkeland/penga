<script lang="ts">
  import { onMount } from "svelte";
  import * as Table from "$lib/components/ui/table/index.js";
  import { GetUtgifter } from "../../wailsjs/go/main/App.js";

  let utgifter: Array<{
    id: number;
    navn: string;
    beløp: number;
    dato: any;
    type?: string;
    beskrivelse: string;
    fast: boolean;
    fast_årlig: boolean;
  }> = [];
  let loading = true;
  let error: string | null = null;

  onMount(async () => {
    try {
      utgifter = await GetUtgifter();
    } catch (e) {
      error = e instanceof Error ? e.message : String(e);
    } finally {
      loading = false;
    }
  });

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

<div class="space-y-4">
  <h2 class="text-2xl font-bold">Utgifter</h2>

  {#if loading}
    <p>Laster...</p>
  {:else if error}
    <p class="text-red-500">Feil: {error}</p>
  {:else}
    <Table.Root>
      <Table.Caption>En oversikt over alle dine utgifter.</Table.Caption>
      <Table.Header>
        <Table.Row>
          <Table.Head class="w-[150px]">Navn</Table.Head>
          <Table.Head>Beskrivelse</Table.Head>
          <Table.Head>Type</Table.Head>
          <Table.Head>Dato</Table.Head>
          <Table.Head>Fast</Table.Head>
          <Table.Head class="text-end">Beløp</Table.Head>
        </Table.Row>
      </Table.Header>
      <Table.Body>
        {#each utgifter as utgift (utgift.id)}
          <Table.Row>
            <Table.Cell class="font-medium">{utgift.navn}</Table.Cell>
            <Table.Cell>{utgift.beskrivelse}</Table.Cell>
            <Table.Cell>{utgift.type ?? "-"}</Table.Cell>
            <Table.Cell>{formatDate(utgift.dato)}</Table.Cell>
            <Table.Cell>{utgift.fast ? "Ja" : "Nei"}</Table.Cell>
            <Table.Cell class="text-end"
              >{formatCurrency(utgift["beløp"])}</Table.Cell
            >
          </Table.Row>
        {/each}
      </Table.Body>
    </Table.Root>
  {/if}
</div>
