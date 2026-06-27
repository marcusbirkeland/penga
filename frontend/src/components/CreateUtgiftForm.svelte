<script lang="ts">
  import { Button } from "$lib/components/ui/button/index.js";
  import * as Field from "$lib/components/ui/field/index.js";
  import { Input } from "$lib/components/ui/input/index.js";
  import { UnitInput } from "$lib/components/ui/unit-input/index.js";
  import { CreateUtgift, GetEiere } from "../../wailsjs/go/main/App.js";
  import { api, models } from "../../wailsjs/go/models.js";

  function getFirstOfMonth() {
    const now = new Date();
    const first = new Date(now.getFullYear(), now.getMonth(), 1);
    return first.toISOString().split("T")[0];
  }

  let navn = $state("");
  let beløp = $state<number | undefined>(undefined);
  let type = $state("");
  let beskrivelse = $state("");
  let fast = $state(true);
  let fastDato = $state(getFirstOfMonth());
  let fastÅrlig = $state(false);
  let eierId = $state(0);
  let eiere = $state<models.Eier[]>([]);

  let { onSaved = () => {} } = $props();

  async function loadEiere() {
    try {
      eiere = await GetEiere();
      if (eiere.length > 0) {
        eierId = eiere[0].id;
      }
    } catch (e) {
      console.error("Failed to load eiere:", e);
    }
  }
  loadEiere();

  let submitting = $state(false);
  let error = $state<string | null>(null);
  let success = $state(false);

  async function handleSubmit(e: Event) {
    e.preventDefault();
    submitting = true;
    error = null;
    success = false;

    try {
      const input = api.UtgiftInput.createFrom({
        Navn: navn,
        Beløp: beløp ?? 0,
        Dato: new Date().toISOString(),
        Type: type || undefined,
        Beskrivelse: beskrivelse,
        Fast: fast,
        FastDato:
          fast && fastDato ? new Date(fastDato).toISOString() : undefined,
        FastÅrlig: fastÅrlig,
        OwnerID: eierId || undefined,
      });
      await CreateUtgift(input);
      success = true;
      resetForm();
      onSaved();
    } catch (e) {
      error = e instanceof Error ? e.message : String(e);
    } finally {
      submitting = false;
    }
  }

  function resetForm() {
    navn = "";
    beløp = undefined;
    beskrivelse = "";
    fast = true;
    fastDato = getFirstOfMonth();
    fastÅrlig = false;
    eierId = eiere.length > 0 ? eiere[0].id : 0;
  }
</script>

<div class="w-full max-w-lg mx-auto">
  <div class="rounded-lg border bg-card text-card-foreground shadow-sm">
    <div class="flex flex-col space-y-1.5 p-6">
      <h3 class="text-2xl font-semibold leading-none tracking-tight">
        Opprett utgift
      </h3>
    </div>
    <div class="p-6 pt-0">
      <form onsubmit={handleSubmit} class="w-xs">
        <Field.Group>
          <Field.Set class="pt-4">
            <Field.Group>
              <Field.Field>
                <Field.Label for="utgift-eier">Eier</Field.Label>
                <select
                  id="utgift-eier"
                  bind:value={eierId}
                  class="flex h-9 w-full rounded-md border border-input bg-transparent px-3 py-1 text-base shadow-sm transition-colors file:border-0 file:bg-transparent file:text-sm file:font-medium file:text-foreground placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:cursor-not-allowed disabled:opacity-50 md:text-sm"
                >
                  {#each eiere as eier}
                    <option value={eier.id}>{eier.navn}</option>
                  {/each}
                </select>
              </Field.Field>
            </Field.Group>
          </Field.Set>
          <div class="grid gap-4">
            <div class="grid grid-cols-2 gap-4">
              <Field.Field>
                <Field.Label for="utgift-beløp">Beløp</Field.Label>
                <UnitInput
                  id="utgift-beløp"
                  step="0.01"
                  unit="kr"
                  bind:value={beløp}
                  placeholder="0"
                  required
                />
              </Field.Field>
              <Field.Field>
                <Field.Label for="utgift-navn">Navn</Field.Label>
                <Input
                  id="utgift-navn"
                  bind:value={navn}
                  placeholder="f.eks. Leie"
                  required
                />
              </Field.Field>
            </div>
          </div>

          <Field.Set class="pt-4">
            <Field.Group>
              <div class="grid grid-cols-2 items-start gap-4">
                <div class="flex flex-col gap-3">
                  <Field.Field orientation="horizontal">
                    <input
                      type="checkbox"
                      id="utgift-fast"
                      bind:checked={fast}
                      class="peer h-4 w-4 shrink-0 rounded-sm border border-primary shadow focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:cursor-not-allowed disabled:opacity-50 data-[state=checked]:bg-primary data-[state=checked]:text-primary-foreground"
                    />
                    <Field.Label for="utgift-fast" class="font-normal"
                      >Fast</Field.Label
                    >
                  </Field.Field>

                  {#if fast}
                    <Field.Field orientation="horizontal">
                      <input
                        type="checkbox"
                        id="utgift-fast-årlig"
                        bind:checked={fastÅrlig}
                        class="peer h-4 w-4 shrink-0 rounded-sm border border-primary shadow focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:cursor-not-allowed disabled:opacity-50 data-[state=checked]:bg-primary data-[state=checked]:text-primary-foreground"
                      />
                      <Field.Label for="utgift-fast-årlig" class="font-normal"
                        >Årlig</Field.Label
                      >
                    </Field.Field>
                  {/if}
                </div>

                {#if fast}
                  <Field.Field>
                    <Field.Label for="utgift-fast-dato">Neste dato</Field.Label>
                    <Input
                      id="utgift-fast-dato"
                      type="date"
                      bind:value={fastDato}
                    />
                  </Field.Field>
                {/if}
              </div>
            </Field.Group>
          </Field.Set>

          <Field.Separator />

          <Field.Field>
            <Field.Label for="utgift-type">Type</Field.Label>
            <select
              id="utgift-type"
              bind:value={type}
              class="flex h-9 w-full rounded-md border border-input bg-transparent px-3 py-1 text-base shadow-sm transition-colors file:border-0 file:bg-transparent file:text-sm file:font-medium file:text-foreground placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:cursor-not-allowed disabled:opacity-50 md:text-sm"
            >
              <option value="">Velg type</option>
              <option value="fast">Regninger</option>
              <option value="variabel">Mat</option>
              <option value="engangs">Moro</option>
            </select>
          </Field.Field>

          <Field.Field>
            <Field.Label for="utgift-beskrivelse">Beskrivelse</Field.Label>
            <textarea
              id="utgift-beskrivelse"
              bind:value={beskrivelse}
              placeholder="Tilleggsbeskrivelse (valgfritt)"
              rows="2"
              class="flex min-h-15 w-full rounded-md border border-input bg-transparent px-3 py-2 text-base shadow-sm placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:cursor-not-allowed disabled:opacity-50 md:text-sm resize-none"
            ></textarea>
          </Field.Field>

          {#if error}
            <div class="rounded-md bg-destructive/15 p-3">
              <p class="text-sm font-medium text-destructive">{error}</p>
            </div>
          {/if}
          {#if success}
            <div class="rounded-md bg-green-500/15 p-3">
              <p class="text-sm font-medium text-green-600">
                Utgift opprettet!
              </p>
            </div>
          {/if}

          <div class="flex justify-end gap-2 pt-4">
            <Button size="lg" type="submit" disabled={submitting}>
              {submitting ? "Lager..." : "Opprett"}
            </Button>
          </div>
        </Field.Group>
      </form>
    </div>
  </div>
</div>
