<script lang="ts" setup>
import type { Achievement } from '@/lib/apis/generated'
import { computed } from 'vue'

interface Props {
  /** 実績データ */
  achievement: Achievement
}

const props = defineProps<Props>()

// フォーマット済み達成日時
const formattedAchievedAt = computed(() => {
  const date = new Date(props.achievement.achievedAt)
  return date.toLocaleDateString('ja-JP', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  })
})

// 実績のアイコンを取得（名前に基づいて簡単なアイコンを返す）
const achievementIcon = computed(() => {
  const name = props.achievement.name.toLowerCase()

  if (name.includes('first') || name.includes('初')) return '🥇'
  if (name.includes('message') || name.includes('メッセージ')) return '💬'
  if (name.includes('reaction') || name.includes('リアクション')) return '👍'
  if (name.includes('reply') || name.includes('返信')) return '💭'
  if (name.includes('image') || name.includes('画像')) return '📸'
  if (name.includes('login') || name.includes('ログイン')) return '🔑'
  if (name.includes('streak') || name.includes('連続')) return '🔥'
  if (name.includes('time') || name.includes('時間')) return '⏰'
  if (name.includes('friend') || name.includes('友達')) return '👥'
  if (name.includes('star') || name.includes('スター')) return '⭐'

  return '🏆' // デフォルトアイコン
})
</script>

<template>
  <div :class="$style.achievementItem">
    <div :class="$style.iconContainer">
      <span :class="$style.icon">{{ achievementIcon }}</span>
    </div>

    <div :class="$style.content">
      <h3 :class="$style.name">{{ achievement.name }}</h3>
      <time :class="$style.achievedAt" :datetime="achievement.achievedAt">
        {{ formattedAchievedAt }}に達成
      </time>
    </div>

    <div :class="$style.badge">
      <span :class="$style.badgeText">達成済み</span>
    </div>
  </div>
</template>

<style lang="scss" module>
.achievementItem {
  display: flex;
  align-items: center;
  gap: 1rem;
  padding: 1rem;
  background-color: var(--color-surface);
  border: 1px solid var(--color-border-light);
  border-radius: 0.75rem;
  transition: all 0.2s ease;
  animation: slideInUp 0.3s ease;

  &:hover {
    border-color: var(--color-primary-200);
    box-shadow: 0 4px 12px var(--color-shadow-light);
    transform: translateY(-1px);
  }
}

.iconContainer {
  flex-shrink: 0;
  width: 3rem;
  height: 3rem;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, var(--color-primary-100), var(--color-primary-200));
  border-radius: 50%;

  [data-theme='dark'] & {
    background: linear-gradient(135deg, var(--color-primary-800), var(--color-primary-700));
  }
}

.icon {
  font-size: 1.5rem;
  filter: drop-shadow(0 1px 2px var(--color-shadow-light));
}

.content {
  flex: 1;
  min-width: 0; /* flexの子要素のoverflowを制御 */
}

.name {
  margin: 0 0 0.25rem 0;
  font-size: 1rem;
  font-weight: 600;
  color: var(--color-text-primary);
  line-height: 1.4;

  /* 長い名前の場合の省略表示 */
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.achievedAt {
  display: block;
  font-size: 0.75rem;
  color: var(--color-text-secondary);
  margin: 0;
}

.badge {
  flex-shrink: 0;
  padding: 0.25rem 0.75rem;
  background-color: var(--color-accent-100);
  color: var(--color-accent-800);
  border-radius: 1rem;
  font-size: 0.75rem;
  font-weight: 500;

  [data-theme='dark'] & {
    background-color: var(--color-accent-800);
    color: var(--color-accent-100);
  }
}

.badgeText {
  display: inline-block;
}

/* レスポンシブ対応 */
@media (max-width: 640px) {
  .achievementItem {
    padding: 0.75rem;
    gap: 0.75rem;
  }

  .iconContainer {
    width: 2.5rem;
    height: 2.5rem;
  }

  .icon {
    font-size: 1.25rem;
  }

  .name {
    font-size: 0.875rem;
  }

  .badge {
    padding: 0.25rem 0.5rem;
  }
}
</style>
