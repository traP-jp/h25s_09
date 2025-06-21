<template>
  <RouterLink
    :class="[$style.sidebar__link, { [$style.current]: isCurrent }]"
    :to="`${props.path}`"
  >
    <Icon :icon="`${iconPath}`" height="40px" width="40px" />
    <span>{{ props.name }}</span>
  </RouterLink>
</template>

<script lang="ts" setup>
import { Icon } from '@iconify/vue'
import { computed } from 'vue'
import { useRoute } from 'vue-router'
const props = defineProps({
  name: String,
  path: String,
  icon: String,
})
const route = useRoute()
const currentPath = computed(() => route.path)
const isCurrent = computed(() => {
  return currentPath.value === props.path
})
const iconPath = computed(() => {
  return isCurrent.value
    ? `material-symbols:${props.icon}`
    : `material-symbols-light:${props.icon}-outline`
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
  font-weight: normal;
  font-size: 150%;
  text-decoration: none;
  &:hover {
    background-color: var(--color-shadow-light);
  }
}

.current {
  font-weight: bold;
}
</style>
