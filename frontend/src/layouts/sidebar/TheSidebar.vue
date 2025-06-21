<script lang="ts" setup>

import { Icon } from "@iconify/vue";
const tools = [
  { name: 'タイムライン', path: '/timeline', icon: 'mdi:home' },
  { name: 'プロフィール', path: '/messages', icon: 'mdi:account' },
  { name: '実績', path: '/achievements', icon: 'mdi:achievement' }
];


</script>

<template>
  <header>
    <nav :class="$style.sidebar">
      <div v-for="tool in tools" :key="tool.name" :class="$style.sidebar__item">
        <router-link :class="$style.sidebar__link" :to="tool.path">
          <icon :icon="tool.icon" height="40px" width="40px" :class="$style.sidebar__icon"/>
          <span>{{ tool.name }}</span>
        </router-link>
      </div>

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
.sidebar__item {  
  cursor: pointer;
  margin:10px;
  border-radius: 30px;
  &:hover {
    background-color: #f0f0f0;
  }
}
.sidebar__link {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 20px;
  color: inherit;
  font-weight: bold;
  font-size: 150%;
  text-decoration: none;
}
</style>
