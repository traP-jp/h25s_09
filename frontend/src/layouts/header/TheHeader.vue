<script setup lang="ts">
import ThemeToggle from '@/components/ThemeToggle.vue'
import UserIcon from '@/components/UserIcon.vue'
import { useUserInfo } from '@/lib/composables'
import { RouterLink } from 'vue-router'

// ユーザー情報取得
const { data: userInfo, isLoading } = useUserInfo()
</script>

<template>
  <header :class="$style.header">
    <div :class="$style.headerContent">
      <div :class="$style.headerLeft">
        <RouterLink to="/" :class="$style.headerTitle">traQ Messages</RouterLink>
      </div>
      <div :class="$style.headerRight">
        <ThemeToggle />
        <div v-if="!isLoading && userInfo" :class="$style.userInfo">
          <RouterLink :to="`/users/${userInfo.traqId}`" :class="$style.userLink">
            <UserIcon :traq-id="userInfo.traqId" size="sm" />
          </RouterLink>
        </div>
        <div v-else-if="isLoading" :class="$style.userPlaceholder">
          <div :class="$style.loadingIcon"></div>
        </div>
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
  text-decoration: none;
  padding: 0.5rem;
  border-radius: 0.375rem;
  transition: background-color 0.2s ease;

  &:hover {
    background-color: var(--color-surface-variant);
  }

  &:focus {
    outline: 2px solid var(--color-primary-500);
    outline-offset: 2px;
  }
}

.headerRight {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.userInfo {
  display: flex;
  align-items: center;
}

.userLink {
  display: flex;
  align-items: center;
  border-radius: 50%;
  transition: transform 0.2s ease;

  &:hover {
    transform: scale(1.05);
  }

  &:focus {
    outline: 2px solid var(--color-primary-500);
    outline-offset: 2px;
  }
}

.userPlaceholder {
  width: 2rem;
  height: 2rem;
  display: flex;
  align-items: center;
  justify-content: center;
}

.loadingIcon {
  width: 1.5rem;
  height: 1.5rem;
  border: 2px solid var(--color-border-light);
  border-top: 2px solid var(--color-primary-600);
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}
</style>
