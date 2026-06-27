<script lang="ts">
	import type { HTMLInputAttributes } from "svelte/elements";
	import { cn, type WithElementRef } from "$lib/utils.js";

	type Props = WithElementRef<Omit<HTMLInputAttributes, "type">> & {
		unit?: string;
	};

	let {
		ref = $bindable(null),
		value = $bindable(),
		unit = "kr",
		class: className,
		"data-slot": dataSlot = "unit-input",
		...restProps
	}: Props = $props();

	const formatter = new Intl.NumberFormat("nb-NO", {
		maximumFractionDigits: 2,
	});

	let focused = $state(false);

	function formatValue(v: unknown): string {
		if (v === null || v === undefined || v === "") return "";
		const num = typeof v === "number" ? v : Number(v);
		return Number.isNaN(num) ? String(v) : formatter.format(num);
	}

	let display = $derived(
		focused ? (value ?? "") : formatValue(value)
	);

	function handleInput(e: Event) {
		const raw = (e.currentTarget as HTMLInputElement).value;
		const normalized = raw.replace(/\s/g, "").replace(",", ".");
		const num = Number(normalized);
		value = normalized === "" || Number.isNaN(num) ? undefined : num;
	}
</script>

<div class="relative w-full">
	<input
		bind:this={ref}
		data-slot={dataSlot}
		class={cn(
			"bg-input/20 dark:bg-input/30 border-input focus-visible:border-ring focus-visible:ring-ring/30 aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive dark:aria-invalid:border-destructive/50 h-7 rounded-md border pl-2 pr-9 py-0.5 text-sm text-right transition-colors focus-visible:ring-2 aria-invalid:ring-2 md:text-xs/relaxed placeholder:text-muted-foreground w-full min-w-0 outline-none disabled:pointer-events-none disabled:cursor-not-allowed disabled:opacity-50",
			className
		)}
		type="text"
		inputmode="decimal"
		value={display}
		oninput={handleInput}
		onfocus={() => (focused = true)}
		onblur={() => (focused = false)}
		{...restProps}
	/>
	<span
		class="text-muted-foreground pointer-events-none absolute inset-y-0 right-0 flex items-center pr-3 text-sm"
	>
		{unit}
	</span>
</div>
