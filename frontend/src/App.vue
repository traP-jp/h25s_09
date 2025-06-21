<script setup lang="ts">
import ErrorToast from '@/components/ErrorToast.vue'
import { useTheme } from '@/composables'
import TheFooter from '@/layouts/footer/TheFooter.vue'
import TheHeader from '@/layouts/header/TheHeader.vue'
import TheSidebar from '@/layouts/sidebar/TheSidebar.vue'
import { useBreakpoints } from '@vueuse/core'
import { onMounted } from 'vue'
import { RouterView } from 'vue-router'

const breakpoints = useBreakpoints({
  mobile: 0,
  tablet: 768,
})

const isDesktop = breakpoints.greaterOrEqual('tablet')

// テーマ初期化
useTheme()

onMounted(() => {
  // 初期テーマ設定
  document.body.classList.add('theme-initialized')
})
</script>

<template>
  <div id="app" class="app-container">
    <TheHeader />
    <div class="app-layout">
      <TheSidebar v-if="isDesktop" class="app-sidebar" />
      <main class="app-main">
        <RouterView />
      </main>
    </div>
    <TheFooter v-if="!isDesktop" class="app-footer" />
    <ErrorToast />
  </div>
</template>

<style lang="scss">
// グローバルリセット
* {
  box-sizing: border-box;
}

html {
  height: 100%;
}

body {
  margin: 0;
  padding: 0;
  min-height: 100%;
  background-color: var(--color-background);
  color: var(--color-text-primary);
  font-family:
    -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  line-height: 1.6;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  transition:
    background-color 0.3s ease,
    color 0.3s ease;
}

#app {
  min-height: 100vh;
  width: 100%;
}

.app-container {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  width: 100%;
  background-color: var(--color-background);
}

.app-layout {
  display: flex;
  flex: 1;
  overflow: hidden;
}

.app-sidebar {
  flex-shrink: 0;
}

.app-main {
  flex: 1;
  overflow-y: auto;
  padding: 1rem;
}

.app-footer {
  flex-shrink: 0;
}

@media (max-width: 767px) {
  .app-main {
    padding: 0.5rem;
  }
}

// スクロールバーのスタイリング
::-webkit-scrollbar {
  width: 8px;
  height: 8px;
}

::-webkit-scrollbar-track {
  background: var(--color-surface-variant);
}

::-webkit-scrollbar-thumb {
  background: var(--color-border-medium);
  border-radius: 4px;
}

::-webkit-scrollbar-thumb:hover {
  background: var(--color-border-strong);
}

// フォーカスリングのスタイリング
:focus-visible {
  outline: 2px solid var(--color-primary-500);
  outline-offset: 2px;
}

// 選択テキストのスタイリング
::selection {
  background-color: var(--color-primary-200);
  color: var(--color-primary-900);
}

[data-theme='dark'] ::selection {
  background-color: var(--color-primary-800);
  color: var(--color-primary-100);
}
</style>
