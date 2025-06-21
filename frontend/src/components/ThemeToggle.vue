<template>
  <button
    :class="$style.themeToggle"
    @click="toggleTheme"
    :title="isDark ? 'ライトモードに切り替え' : 'ダークモードに切り替え'"
    :aria-label="isDark ? 'ライトモードに切り替え' : 'ダークモードに切り替え'"
  >
    <Icon :icon="isDark ? 'mdi:weather-sunny' : 'mdi:weather-night'" :size="20" />
  </button>
</template>

<script setup lang="ts">
import { useTheme } from '@/composables'
import { Icon } from '@iconify/vue'

const { isDark, toggle } = useTheme()

const toggleTheme = () => {
  toggle()
}
</script>

<style module lang="scss">
.themeToggle {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 2.5rem;
  height: 2.5rem;
  border: none;
  border-radius: 0.5rem;
  background-color: var(--color-surface-variant);
  color: var(--color-text-secondary);
  cursor: pointer;
  transition: all 0.2s ease-in-out;

  &:hover {
    background-color: var(--color-primary-100);
    color: var(--color-primary-600);
    transform: scale(1.05);
  }

  &:active {
    transform: scale(0.95);
  }

  &:focus {
    outline: 2px solid var(--color-primary-500);
    outline-offset: 2px;
  }
}

:global([data-theme='dark']) .themeToggle {
  &:hover {
    background-color: var(--color-primary-800);
    color: var(--color-primary-300);
  }
}
</style>
