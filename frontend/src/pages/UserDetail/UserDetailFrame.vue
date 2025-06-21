<script lang="ts" setup>
import UserIcon from '@/components/UserIcon.vue'
import { computed } from 'vue'
import { RouterLink, useRoute } from 'vue-router'

const route = useRoute()

// URLパラメータからユーザーIDを取得
const userId = computed(() => route.params.traqId as string)

// タブの定義
const tabs = computed(() => [
  { id: 'messages', label: 'メッセージ', path: `/users/${userId.value}/messages` },
  { id: 'achievements', label: '実績', path: `/users/${userId.value}/achievements` },
])

// 現在のタブを判定
const currentTab = computed(() => {
  const path = route.path
  if (path.includes('/achievements')) return 'achievements'
  return 'messages'
})
</script>

<template>
  <div :class="$style.userDetailFrame">
    <!-- ユーザーヘッダー -->
    <header :class="$style.userHeader">
      <div :class="$style.userInfo">
        <UserIcon :traq-id="userId" size="lg" />
        <div :class="$style.userDetails">
          <h1 :class="$style.userName">@{{ userId }}</h1>
        </div>
      </div>
    </header>

    <!-- タブナビゲーション -->
    <nav :class="$style.tabNavigation">
      <div :class="$style.tabList">
        <RouterLink
          v-for="tab in tabs"
          :key="tab.id"
          :to="tab.path"
          :class="[$style.tabButton, { [$style.active]: currentTab === tab.id }]"
        >
          {{ tab.label }}
        </RouterLink>
      </div>
    </nav>

    <!-- コンテンツエリア -->
    <main :class="$style.content">
      <router-view />
    </main>
  </div>
</template>

<style lang="scss" module>
.userDetailFrame {
  min-height: 100vh;
  background-color: var(--color-background);
}

.userHeader {
  background-color: var(--color-surface);
  border-bottom: 1px solid var(--color-border-light);
  padding: 2rem 1rem;
}

.userInfo {
  display: flex;
  align-items: center;
  gap: 1rem;
  max-width: 1200px;
  margin: 0 auto;
}

.userDetails {
  flex: 1;
}

.userName {
  margin: 0 0 0.25rem 0;
  font-size: 1.5rem;
  font-weight: 700;
  color: var(--color-text-primary);
}

.tabNavigation {
  background-color: var(--color-surface);
  border-bottom: 1px solid var(--color-border-light);
  position: sticky;
  top: 4rem; /* TheHeaderの高さ分 */
  z-index: 10;
}

.tabList {
  display: flex;
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 1rem;
}

.tabButton {
  padding: 1rem 1.5rem;
  font-size: 0.875rem;
  font-weight: 500;
  color: var(--color-text-secondary);
  border-bottom: 2px solid transparent;
  transition: all 0.2s ease;
  position: relative;
  text-decoration: none;
  display: block;

  &:hover {
    color: var(--color-text-primary);
    background-color: var(--color-surface-variant);
  }

  &.active {
    color: var(--color-primary-600);
    border-bottom-color: var(--color-primary-600);

    &:hover {
      background-color: var(--color-primary-50);
    }

    [data-theme='dark'] & {
      color: var(--color-primary-400);
      border-bottom-color: var(--color-primary-400);

      &:hover {
        background-color: var(--color-primary-900);
      }
    }
  }
}

.content {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0;
  min-height: calc(100vh - 12rem); /* ヘッダーとタブナビゲーションの高さを引く */
}
</style>
