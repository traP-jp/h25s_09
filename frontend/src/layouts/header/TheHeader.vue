<script setup lang="ts">
import ThemeToggle from '@/components/ThemeToggle.vue'
</script>

<template>
  <header :class="$style.header">
    <div :class="$style.headerContent">
      <div :class="$style.headerLeft">
        <h1 :class="$style.headerTitle">TraQ Messages</h1>
      </div>
      <div :class="$style.headerRight">
        <ThemeToggle />
      </div>
    </div>
  </header>
</template>

<style lang="scss" module>
.header {
  position: sticky;
  top: 0;
  z-index: 50;
  width: 100%;
  background-color: var(--color-surface);
  border-bottom: 1px solid var(--color-border-light);
  backdrop-filter: blur(8px);
  box-shadow: 0 2px 4px var(--color-shadow-light);
}

.headerContent {
  display: flex;
  align-items: center;
  justify-content: space-between;
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 1rem;
  height: 4rem;
}

.headerLeft {
  display: flex;
  align-items: center;
}

.headerTitle {
  margin: 0;
  font-size: 1.5rem;
  font-weight: 600;
  color: var(--color-text-primary);
}

.headerRight {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}
</style>
