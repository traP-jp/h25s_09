<script setup lang="ts">
import { Icon } from '@iconify/vue'
import { userService } from '@/lib/apis/services.ts'
import { onMounted, ref } from 'vue'
import { computed } from 'vue'
import { useRoute } from 'vue-router'

// ユーザーIDの取得
const userId = ref<string | null>(null)
onMounted(async () => {
  userId.value = (await userService.getUserInfo()).traqId
})
const navItems = computed(() => [
  { to: '/timeline', icon: 'mdi:home', label: 'ホーム' },
  { to: `/users/${userId.value}/achievements`, icon: 'mdi:trophy', label: 'アチーブメント' },
  { to: `/users/${userId.value}/messages`, icon: 'mdi:account', label: 'プロフィール' },
])
const route = useRoute()
const isCurrent = computed(() => {
  return route.path === '/timeline' || route.path.startsWith('/users/')
})
</script>

<template>
  <footer :class="$style.bottomBar">
    <router-link
      v-for="item in navItems"
      :key="item.to"
      :to="item.to"
      :class="[$style.tabItem, { [$style.current]: isCurrent }]"
      active-class="is-active"
    >
      <Icon :icon="item.icon" :class="$style.icon" />
      <span :class="$style.label">{{ item.label }}</span>
    </router-link>
  </footer>
</template>

<style lang="scss" module>
/*
 * 全てのスタイルをこのコンポーネント内にまとめて記述します。
 */

/* 親コンテナのスタイル */
.bottomBar {
  --colorBackground: #000000;
  --colorPrimary: #f9f9fb;
  --colorActive: #7591e4;
  --colorBorder: #000300;
  --shadow: 0 -2px 10px rgba(0, 0, 0, 0.1);
  display: flex;
  justify-content: space-around;
  align-items: center;
  background-color: var(--colorBackground);
  padding: 10px 0;
  position: fixed;
  bottom: 0;
  width: 100%;
  box-shadow: var(--shadow);
  z-index: 1000;
}
.tabItem {
  display: flex;
  flex-direction: column;
  align-items: center;
  text-decoration: none;
  color: var(--colorPrimary);
  font-size: 12px;
  padding: 5px;
  transition:
    color 0.3s ease,
    transform 0.3s ease;
  cursor: pointer;
  border-radius: 10px;
  &:hover {
    color: var(--colorActive);
    transform: scale(2);
    box-shadow: var(--shadow);
    transition:
      color 0.3s ease,
      transform 0.3s ease,
      box-shadow 0.3s ease;
  }
  &.isActive {
    color: var(--colorActive);
    transform: scale(1.1);
    font-weight: bold;
    box-shadow: var(--shadow);
    transition:
      color 0.3s ease,
      transform 0.3s ease,
      box-shadow 0.3s ease;
  }
}
.icon {
  font-size: 24px;
}

.label {
  font-size: 10px;
  margin-top: 2px;
}
.current {
  color: var(--colorPrimary);
  font-weight: bold;
}
</style>
