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
  { to: '/timeline', icon: 'home', label: 'ホーム' },
  { to: `/users/${userId.value}/achievements`, icon: 'trophy', label: 'アチーブメント' },
  { to: `/users/${userId.value}/messages`, icon: 'account-circle', label: 'プロフィール' },
])

const route = useRoute()

// 各ナビゲーションアイテムが現在のルートと一致するかチェック
const isCurrentRoute = (itemPath: string) => {
  return route.path === itemPath
}

// SidebarItemと同じロジックでアイコンパスを生成
const getIconPath = (iconName: string, isActive: boolean) => {
  return isActive ? `material-symbols:${iconName}` : `material-symbols-light:${iconName}-outline`
}
</script>

<template>
  <footer :class="$style.bottomBar">
    <router-link
      v-for="item in navItems"
      :key="item.to"
      :to="item.to"
      :class="[$style.tabItem, { [$style.isActive]: isCurrentRoute(item.to) }]"
    >
      <Icon :icon="getIconPath(item.icon, isCurrentRoute(item.to))" :class="$style.icon" />
      <span :class="$style.label">{{ item.label }}</span>
    </router-link>
  </footer>
</template>

<style lang="scss" module>
.bottomBar {
  --colorBackground: var(--color-background);
  --colorPrimary: var(--color-text-primary);
  --colorActive: var(--color-primary);
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
    transform: scale(1.1);
    box-shadow: var(--shadow);
  }

  &.isActive {
    color: var(--colorActive);
    transform: scale(1.1);
    font-weight: bold;
    box-shadow: var(--shadow);
    background-color: var(--color-shadow-light);
  }
}

.icon {
  font-size: 24px;
}

.label {
  font-size: 10px;
  margin-top: 2px;
}
</style>
