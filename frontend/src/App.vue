<script setup lang="ts"></script>

<template>
  <UApp>
    <div class="app">
      <header class="app-header">
        <h1>Penga</h1>
      </header>

      <main class="app-main">
        <RouterView />
      </main>
    </div>
  </UApp>
</template>

<style scoped>
.app {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.app-header {
  padding: 1rem 2rem;
  background-color: var(--color-bg-surface, #1b1b1f);
  border-bottom: 1px solid var(--color-border-default, #3e3e3e);
}

.app-header h1 {
  margin: 0;
  font-size: 1.5rem;
  font-weight: bold;
}

.app-main {
  flex: 1;
  padding: 2rem;
}
</style>
