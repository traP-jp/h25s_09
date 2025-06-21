<script lang="ts" setup>
const tools = [
  { name: 'timeline', path: '/timeline' },
  { name: 'messageId', path: '/messageId' },
  { name: 'messages', path: '/messages' },
  { name: 'achievements', path: '/achievements' }
];


</script>

<template>
  <header>
    <nav :class="$style.sidebar">
      <div v-for="tool in tools" :key="tool.name" :class="$style.sidebar__item">
        <router-link :class="$style.sidebar__link" :to="tool.path">{{ tool.name }}</router-link>
      </div>

    </nav>
  </header>
</template>

<style lang="scss" module>
.sidebar {
  display: flex;
  flex-direction: column;
  width: 200px;
}
.sidebar__item {
  padding: 10px;
  text-decoration: none;
  cursor: pointer;
  &:hover {
    background-color: #f0f0f0;
  }
}
.sidebar__link {
  text-decoration: none;
  color: inherit;
  font-weight: bold;
}
</style>
