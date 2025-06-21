<template>
  <RouterLink
    :class="[$style.sidebar__link, { [$style.current]: isCurrent }]"
    :to="`${props.path}`"
    :title="props.name"
  >
    <Icon :icon="`${iconPath}`" height="40px" width="40px" :class="$style.icon" />
    <span :class="$style.text">{{ props.name }}</span>
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
  border-radius: 20px;
  margin: 5px;
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 15px 10px; /* パディングを調整 */
  color: inherit;
  font-weight: normal;
  font-size: 150%;
  text-decoration: none;
  min-width: 0; /* flexアイテムの縮小を許可 */

  &:hover {
    background-color: var(--color-shadow-light);
  }
}

.icon {
  flex-shrink: 0; /* アイコンは縮小しない */
}

.text {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  min-width: 0; /* テキストの縮小を許可 */
}

.current {
  font-weight: bold;
}

/* コンパクトサイドバー用のスタイル（100px幅） */
@container (max-width: 110px) {
  .sidebar__link {
    flex-direction: column;
    justify-content: center;
    align-items: center;
    gap: 2px;
    padding: 8px 1px;
    border-radius: 12px;
    margin: 2px;
  }

  .icon {
    width: 28px !important;
    height: 28px !important;
  }

  .text {
    font-size: 0.5em; /* より小さなフォントサイズ */
    line-height: 1;
    text-align: center;
    white-space: nowrap; /* 改行を防ぐ */
    overflow: hidden;
    text-overflow: ellipsis;
    max-width: 100%;
  }
}
</style>
