<script lang="ts" setup>
import { useUserAchievements } from '@/lib/composables'
import { Icon } from '@iconify/vue'
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import AchievementItem from './components/AchievementItem.vue'

const route = useRoute()

// URLパラメータからユーザーIDを取得
const userId = computed(() => route.params.traqId as string)

// ユーザー実績取得
const { data: achievements, isLoading, error } = useUserAchievements(userId)

// 実績の統計情報
const achievementStats = computed(() => {
  if (!achievements.value) return null

  return {
    total: achievements.value.length,
    recent: achievements.value.filter((achievement) => {
      const achievedDate = new Date(achievement.achievedAt)
      const weekAgo = new Date()
      weekAgo.setDate(weekAgo.getDate() - 7)
      return achievedDate > weekAgo
    }).length,
  }
})
</script>

<template>
  <div :class="$style.userAchievementsPage">
    <!-- ローディング状態 -->
    <div v-if="isLoading" :class="$style.loading">
      <div :class="$style.spinner"></div>
      <p>実績を読み込み中...</p>
    </div>

    <!-- エラー状態 -->
    <div v-else-if="error" :class="$style.error">
      <Icon icon="mdi:alert-circle" :class="$style.errorIcon" />
      <h2>実績の取得に失敗しました</h2>
      <p>しばらく待ってから再度お試しください。</p>
    </div>

    <!-- 実績一覧 -->
    <div v-else-if="achievements && achievements.length > 0" :class="$style.content">
      <!-- 統計情報 -->
      <div v-if="achievementStats" :class="$style.stats">
        <div :class="$style.statCard">
          <div :class="$style.statNumber">{{ achievementStats.total }}</div>
          <div :class="$style.statLabel">総実績数</div>
        </div>
        <div :class="$style.statCard">
          <div :class="$style.statNumber">{{ achievementStats.recent }}</div>
          <div :class="$style.statLabel">今週の実績</div>
        </div>
      </div>

      <!-- 実績リスト -->
      <div :class="$style.achievementsList">
        <h2 :class="$style.sectionTitle">達成した実績</h2>
        <div :class="$style.grid">
          <AchievementItem
            v-for="achievement in achievements"
            :key="achievement.id"
            :achievement="achievement"
          />
        </div>
      </div>
    </div>

    <!-- 空の状態 -->
    <div v-else :class="$style.empty">
      <Icon icon="mdi:trophy-outline" :class="$style.emptyIcon" />
      <h2>実績がありません</h2>
      <p>@{{ userId }}さんはまだ実績を達成していません。</p>
    </div>
  </div>
</template>

<style lang="scss" module>
.userAchievementsPage {
  min-height: 100%;
  padding: 1.5rem;
}

.loading,
.error,
.empty {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 4rem 2rem;
  text-align: center;
  color: var(--color-text-secondary);
}

.loading {
  gap: 1rem;
}

.error,
.empty {
  gap: 0.5rem;
}

.spinner {
  width: 2rem;
  height: 2rem;
  border: 2px solid var(--color-border-light);
  border-top: 2px solid var(--color-primary-500);
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}

.errorIcon,
.emptyIcon {
  font-size: 3rem;
  margin-bottom: 1rem;
}

.error h2,
.empty h2 {
  margin: 0 0 0.5rem 0;
  font-size: 1.25rem;
  font-weight: 600;
  color: var(--color-text-primary);
}

.error p,
.empty p {
  margin: 0;
  color: var(--color-text-secondary);
}

.content {
  display: flex;
  flex-direction: column;
  gap: 2rem;
}

.stats {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
  gap: 1rem;
}

.statCard {
  padding: 1.5rem;
  background-color: var(--color-surface);
  border: 1px solid var(--color-border-light);
  border-radius: 0.75rem;
  text-align: center;
  box-shadow: 0 1px 3px var(--color-shadow-light);
}

.statNumber {
  font-size: 2rem;
  font-weight: 700;
  color: var(--color-primary-600);
  margin-bottom: 0.25rem;

  [data-theme='dark'] & {
    color: var(--color-primary-400);
  }
}

.statLabel {
  font-size: 0.875rem;
  color: var(--color-text-secondary);
  font-weight: 500;
}

.achievementsList {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.sectionTitle {
  margin: 0;
  font-size: 1.25rem;
  font-weight: 600;
  color: var(--color-text-primary);
}

.grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 1rem;
}

/* レスポンシブ対応 */
@media (max-width: 768px) {
  .userAchievementsPage {
    padding: 1rem;
  }

  .grid {
    grid-template-columns: 1fr;
  }

  .stats {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 480px) {
  .stats {
    grid-template-columns: 1fr;
  }
}
</style>
