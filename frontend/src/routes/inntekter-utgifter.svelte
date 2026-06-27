<script lang="ts">
  import CreateUtgiftForm from "../components/CreateUtgiftForm.svelte";
  import UtgifterTable from "../components/UtgifterTable.svelte";

  let table: any;

  async function handleSaved() {
    table?.refetch();
  }
</script>

<div class="flex flex-col gap-6 p-6">
  <h1 class="text-3xl font-bold tracking-tight">Inntekter / Utgifter</h1>

  <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
    <div class="lg:col-span-2">
      <UtgifterTable bind:this={table} />
    </div>
    <div class="lg:col-span-1">
      <div class="sticky top-6">
        <CreateUtgiftForm onSaved={handleSaved} />
      </div>
    </div>
  </div>
</div>
