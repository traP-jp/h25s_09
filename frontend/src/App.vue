<script setup lang="ts">
import ErrorToast from '@/components/ErrorToast.vue'
import PageLoading from '@/components/PageLoading.vue'
import PostModal from '@/components/PostModal.vue'
import { useTheme } from '@/composables'
import TheFooter from '@/layouts/footer/TheFooter.vue'
import TheHeader from '@/layouts/header/TheHeader.vue'
import TheSidebar from '@/layouts/sidebar/TheSidebar.vue'
import { useBreakpoints } from '@vueuse/core'
import { onMounted } from 'vue'
import { RouterView } from 'vue-router'
import { useCreateAchievement } from './lib/composables'
import { randomBoolean } from './lib/utils'

const breakpoints = useBreakpoints({
  mobile: 0,
  compactSidebar: 900, // コンパクトサイドバー
  fullSidebar: 1200, // フルサイドバー
})

const showFullSidebar = breakpoints.greaterOrEqual('fullSidebar')
const showCompactSidebar = breakpoints.between('compactSidebar', 'fullSidebar')
const showSidebar = breakpoints.greaterOrEqual('compactSidebar')

onMounted(async () => {
  const isCursorLoading = randomBoolean(0.2)
  if (isCursorLoading) {
    document.body.classList.add('cursor-loading')
    await useCreateAchievement().mutateAsync('読込中')
  }
})

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
      <TheSidebar
        v-if="showSidebar"
        class="app-sidebar"
        :class="{
          'app-sidebar--full': showFullSidebar,
          'app-sidebar--compact': showCompactSidebar,
        }"
      />
      <main
        class="app-main"
        :class="{
          'app-main--full-sidebar': showFullSidebar,
          'app-main--compact-sidebar': showCompactSidebar,
        }"
      >
        <RouterView />
      </main>
    </div>
    <TheFooter v-if="!showSidebar" class="app-footer" />
    <ErrorToast />
    <PostModal />
    <PageLoading />
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
  position: relative;
}

.app-sidebar {
  position: fixed;
  top: 4rem; /* ヘッダーの高さ(4rem) */
  left: 0;
  bottom: 0;
  z-index: 10;
  background-color: var(--color-background);
  border-right: 1px solid var(--color-border-light);
  overflow-y: auto;
}

.app-sidebar--full {
  width: 240px; /* フルサイドバーの固定幅 */
}

.app-sidebar--compact {
  width: 100px; /* コンパクトサイドバーの固定幅を狭める */
}

.app-main {
  flex: 1;
  overflow-y: auto;
  padding: 1rem;
  padding-top: calc(4rem + 1rem); /* ヘッダーの高さ(4rem) + 通常のpadding(1rem) */
  min-height: calc(100vh - 4rem); /* ヘッダーの高さ(4rem)を除く */
}

.app-main--full-sidebar {
  margin-left: 300px; /* フルサイドバーの幅分マージン */
}

.app-main--compact-sidebar {
  margin-left: 100px; /* コンパクトサイドバーの幅分マージン */
}

.app-footer {
  flex-shrink: 0;
}

@media (max-width: 899px) {
  .app-main {
    margin-left: 0; /* サイドバー非表示時はマージンなし */
    padding: 0.5rem;
    padding-top: calc(4rem + 0.5rem); /* ヘッダーの高さ(4rem) + 通常のpadding(0.5rem) */
  }
}

// スクロールバーのスタイリング（オーバーレイ表示でUIに影響しない）
* {
  scrollbar-width: thin; /* Firefox用 */
  scrollbar-color: var(--color-border-medium) transparent; /* Firefox用 */
}

/* WebKit系ブラウザ用のオーバーレイスクロールバー */
::-webkit-scrollbar {
  width: 8px;
  height: 8px;
  background: transparent; /* 背景を透明に */
}

::-webkit-scrollbar-track {
  background: transparent; /* トラックも透明に */
}

::-webkit-scrollbar-thumb {
  background: var(--color-border-medium);
  border-radius: 4px;
  background-clip: padding-box; /* パディング内でのみ表示 */
}

::-webkit-scrollbar-thumb:hover {
  background: var(--color-border-strong);
}

/* コンテナ要素にオーバーレイスクロールを適用 */
.app-container,
.app-main,
.app-sidebar {
  overflow: overlay; /* オーバーレイスクロールを使用 */
}

/* 代替案：スクロールバーを非表示にする場合 */
@supports not (overflow: overlay) {
  .app-container,
  .app-main,
  .app-sidebar {
    scrollbar-width: none; /* Firefox */
    -ms-overflow-style: none; /* IE/Edge */
  }

  .app-container::-webkit-scrollbar,
  .app-main::-webkit-scrollbar,
  .app-sidebar::-webkit-scrollbar {
    display: none; /* WebKit */
  }
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

body.cursor-loading {
  cursor: progress;
}
</style>
