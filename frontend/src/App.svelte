<script lang="ts">
  import { Router } from "sv-router";
  import "sv-router/generated";
  import * as Sidebar from "$lib/components/ui/sidebar/index.js";
  import AppSidebar from "./components/AppSidebar.svelte";
</script>

<Sidebar.Provider>
  <AppSidebar />
  <main>
    <Router />
  </main>
</Sidebar.Provider>
