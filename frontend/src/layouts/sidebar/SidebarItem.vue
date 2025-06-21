<template>
  <RouterLink :class="$style.sidebar__link" :to="`${props.path}`">
    <Icon :icon="`${props.icon}`" height="40px" width="40px" />
    <span>{{ props.name }}</span>
  </RouterLink>
</template>

<script lang="ts" setup>
import { Icon } from '@iconify/vue'

const props = defineProps({
  name: String,
  path: String,
  icon: String,
})
</script>

<style lang="scss" module>
.sidebar__link {
  cursor: pointer;
  border-radius: 30px;
  margin: 10px;
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 20px;
  color: inherit;
  font-weight: bold;
  font-size: 150%;
  text-decoration: none;
  &:hover {
    background-color: var(--color-shadow-light);
  }
}
</style>
