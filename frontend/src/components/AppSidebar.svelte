<script lang="ts">
  import * as Sidebar from "$lib/components/ui/sidebar/index.js";
  import { BanknoteArrowDownIcon, HouseIcon } from "@lucide/svelte";

  // Menu items.
  const items = [
    {
      title: "Hjem",
      url: "/",
      icon: HouseIcon,
    },
    {
      title: "Inntekter/Utgifter",
      url: "/inntekter-utgifter",
      icon: BanknoteArrowDownIcon,
    },
  ];
</script>

<Sidebar.Root>
  <Sidebar.Content>
    <Sidebar.Group>
      <Sidebar.GroupLabel>Penga</Sidebar.GroupLabel>
      <Sidebar.GroupContent>
        <Sidebar.Menu>
          {#each items as item (item.title)}
            <Sidebar.MenuItem>
              <Sidebar.MenuButton>
                {#snippet child({ props }: { props: Record<string, unknown> })}
                  <a href={item.url} {...props}>
                    <item.icon />
                    <span>{item.title}</span>
                  </a>
                {/snippet}
              </Sidebar.MenuButton>
            </Sidebar.MenuItem>
          {/each}
        </Sidebar.Menu>
      </Sidebar.GroupContent>
    </Sidebar.Group>
  </Sidebar.Content>
</Sidebar.Root>
