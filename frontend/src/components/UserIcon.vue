<script lang="ts" setup>
interface Props {
  /** ユーザーのtraqID */
  traqId: string
  /** アイコンのサイズ */
  size?: 'sm' | 'md' | 'lg'
  /** クリック可能にするかどうか */
  clickable?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  size: 'md',
  clickable: false,
})

const emit = defineEmits<{
  click: [traqId: string]
}>()

const handleClick = () => {
  if (props.clickable) {
    emit('click', props.traqId)
  }
}
</script>

<template>
  <div
    :class="[$style.userIcon, $style[`size-${size}`], { [$style.clickable]: clickable }]"
    @click="handleClick"
  >
    <img
      :src="`https://q.trap.jp/api/v3/public/icon/${traqId}`"
      :alt="`${traqId}のアイコン`"
      :class="$style.avatar"
    />
    <div :class="$style.fallback">
      {{ traqId.charAt(0).toUpperCase() }}
    </div>
  </div>
</template>

<style lang="scss" module>
.userIcon {
  position: relative;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  background-color: var(--color-surface-variant);
  overflow: hidden;
  flex-shrink: 0;
}

.size-sm {
  width: 2rem;
  height: 2rem;
  font-size: 0.75rem;
}

.size-md {
  width: 2.5rem;
  height: 2.5rem;
  font-size: 0.875rem;
}

.size-lg {
  width: 3rem;
  height: 3rem;
  font-size: 1rem;
}

.clickable {
  cursor: pointer;
  transition: transform 0.2s ease;

  &:hover {
    transform: scale(1.05);
  }

  &:active {
    transform: scale(0.95);
  }
}

.avatar {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
}

.fallback {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
  color: var(--color-text-secondary);
  background-color: var(--color-surface-variant);
  z-index: -1;
}

// 画像が読み込まれない場合のフォールバック
.avatar:not([src]),
.avatar[src=''] {
  display: none;
}
</style>
