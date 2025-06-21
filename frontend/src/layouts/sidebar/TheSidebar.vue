<script lang="ts" setup>
import { userService } from '@/lib/apis/services.ts'
import { onMounted, ref } from 'vue'
import SidebarItem from '@/layouts/sidebar/SidebarItem.vue'

const userId = ref<string | null>(null)
onMounted(async () => {
  userId.value = (await userService.getUserInfo()).traqId
})
</script>

<template>
  <header>
    <nav :class="$style.sidebar">
      <SidebarItem name="ホーム" path="/timeline" icon="home" />
      <SidebarItem name="プロフィール" :path="`/user/${userId}/messages`" icon="account-circle" />
      <SidebarItem name="実績" path="/achievements" icon="trophy" />
    </nav>
  </header>
</template>

<style lang="scss" module>
.sidebar {
  display: flex;
  flex-direction: column;
  width: 400px;
  padding: 30px;
}
</style>
