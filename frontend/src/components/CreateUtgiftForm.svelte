<script lang="ts">
  import { Button } from "$lib/components/ui/button/index.js";
  import * as Field from "$lib/components/ui/field/index.js";
  import { Input } from "$lib/components/ui/input/index.js";
  import { CreateUtgift } from "../../wailsjs/go/main/App.js";
  import { api } from "../../wailsjs/go/models.js";

  let navn = $state("");
  let beløp = $state(0);
  let dato = $state("");
  let type = $state("");
  let beskrivelse = $state("");
  let fast = $state(false);
  let fastDato = $state("");
  let fastÅrlig = $state(false);
  let eierId = $state(0);

  let submitting = $state(false);
  let error = $state<string | null>(null);
  let success = $state(false);

  async function handleSubmit(e: Event) {
    e.preventDefault();
    submitting = true;
    error = null;
    success = false;

    try {
      const input = {
        Navn: navn,
        Beløp: beløp,
        Dato: new Date(),
        Type: type || undefined,
        Beskrivelse: beskrivelse,
        Fast: fast,
        FastDato: fast && fastDato ? new Date(fastDato) : undefined,
        FastÅrlig: fastÅrlig,
        OwnerID: eierId || undefined,
      };
      await CreateUtgift(input);
      success = true;
      resetForm();
    } catch (e) {
      error = e instanceof Error ? e.message : String(e);
    } finally {
      submitting = false;
    }
  }

  function resetForm() {
    navn = "";
    beløp = 0;
    dato = "";
    type = "";
    beskrivelse = "";
    fast = false;
    fastDato = "";
    fastÅrlig = false;
    eierId = 0;
  }
</script>

<div class="w-full max-w-lg">
  <form onsubmit={handleSubmit}>
    <Field.Group>
      <Field.Set>
        <Field.Legend>Opprett utgift</Field.Legend>
        <Field.Description
          >Fyll inn detaljene for den nye utgiften.</Field.Description
        >
        <Field.Group>
          <Field.Field>
            <Field.Label for="utgift-navn">Navn</Field.Label>
            <Input
              id="utgift-navn"
              bind:value={navn}
              placeholder="f.eks. Leie"
              required
            />
          </Field.Field>
          <Field.Field>
            <Field.Label for="utgift-beløp">Beløp (NOK)</Field.Label>
            <Input
              id="utgift-beløp"
              type="number"
              step="0.01"
              bind:value={beløp}
              placeholder="0"
              required
            />
          </Field.Field>
          <Field.Field>
            <Field.Label for="utgift-dato">Dato</Field.Label>
            <Input id="utgift-dato" type="date" bind:value={dato} required />
          </Field.Field>
          <Field.Field>
            <Field.Label for="utgift-type">Type</Field.Label>
            <select
              id="utgift-type"
              bind:value={type}
              class="flex h-9 w-full rounded-md border border-input bg-transparent px-3 py-1 text-base shadow-sm transition-colors file:border-0 file:bg-transparent file:text-sm file:font-medium file:text-foreground placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:cursor-not-allowed disabled:opacity-50 md:text-sm"
            >
              <option value="">Velg type</option>
              <option value="fast">Fast utgift</option>
              <option value="variabel">Variabel utgift</option>
              <option value="engangs">Engangs utgift</option>
            </select>
          </Field.Field>
          <Field.Field>
            <Field.Label for="utgift-beskrivelse">Beskrivelse</Field.Label>
            <textarea
              id="utgift-beskrivelse"
              bind:value={beskrivelse}
              placeholder="Tilleggsbeskrivelse"
              class="flex min-h-[60px] w-full rounded-md border border-input bg-transparent px-3 py-2 text-base shadow-sm placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:cursor-not-allowed disabled:opacity-50 md:text-sm"
            ></textarea>
          </Field.Field>
        </Field.Group>
      </Field.Set>

      <Field.Separator />

      <Field.Set>
        <Field.Legend>Fast utgift</Field.Legend>
        <Field.Description
          >Angi om dette er en gjentakende utgift.</Field.Description
        >
        <Field.Group>
          <Field.Field orientation="horizontal">
            <input
              type="checkbox"
              id="utgift-fast"
              bind:checked={fast}
              class="peer h-4 w-4 shrink-0 rounded-sm border border-primary shadow focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:cursor-not-allowed disabled:opacity-50 data-[state=checked]:bg-primary data-[state=checked]:text-primary-foreground"
            />
            <Field.Label for="utgift-fast" class="font-normal"
              >Fast utgift</Field.Label
            >
          </Field.Field>

          {#if fast}
            <Field.Field>
              <Field.Label for="utgift-fast-dato">Neste dato</Field.Label>
              <Input id="utgift-fast-dato" type="date" bind:value={fastDato} />
            </Field.Field>
            <Field.Field orientation="horizontal">
              <input
                type="checkbox"
                id="utgift-fast-årlig"
                bind:checked={fastÅrlig}
                class="peer h-4 w-4 shrink-0 rounded-sm border border-primary shadow focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:cursor-not-allowed disabled:opacity-50 data-[state=checked]:bg-primary data-[state=checked]:text-primary-foreground"
              />
              <Field.Label for="utgift-fast-årlig" class="font-normal"
                >Årlig utgift</Field.Label
              >
            </Field.Field>
          {/if}
        </Field.Group>
      </Field.Set>

      <Field.Separator />

      <Field.Set>
        <Field.Group>
          <Field.Field>
            <Field.Label for="utgift-eier-id">Eier ID</Field.Label>
            <Input
              id="utgift-eier-id"
              type="number"
              bind:value={eierId}
              placeholder="Eier ID (valgfritt)"
            />
          </Field.Field>
        </Field.Group>
      </Field.Set>

      {#if error}
        <p class="text-sm text-red-500">{error}</p>
      {/if}
      {#if success}
        <p class="text-sm text-green-500">Utgift opprettet!</p>
      {/if}

      <Field.Field orientation="horizontal">
        <Button type="submit" disabled={submitting}>
          {submitting ? "Lager..." : "Opprett"}
        </Button>
        <Button variant="outline" type="button" on:click={resetForm}
          >Avbryt</Button
        >
      </Field.Field>
    </Field.Group>
  </form>
</div>
