<script lang="ts" setup>
import MessageItem from '@/components/MessageItem.vue'
import { useUserMessages } from '@/lib/composables'
import { computed } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()

// URLパラメータからユーザーIDを取得
const userId = computed(() => route.params.traqId as string)

// ユーザーメッセージ取得
const { data: messages, isLoading, error } = useUserMessages(userId)
</script>

<template>
  <div :class="$style.userMessagesPage">
    <!-- ローディング状態 -->
    <div v-if="isLoading" :class="$style.loading">
      <div :class="$style.spinner"></div>
      <p>メッセージを読み込み中...</p>
    </div>

    <!-- エラー状態 -->
    <div v-else-if="error" :class="$style.error">
      <div :class="$style.errorIcon">⚠️</div>
      <h2>メッセージの取得に失敗しました</h2>
      <p>しばらく待ってから再度お試しください。</p>
    </div>

    <!-- メッセージ一覧 -->
    <div v-else-if="messages && messages.length > 0" :class="$style.messagesList">
      <MessageItem v-for="message in messages" :key="message.id" :message="message" />
    </div>

    <!-- 空の状態 -->
    <div v-else :class="$style.empty">
      <div :class="$style.emptyIcon">📝</div>
      <h2>メッセージがありません</h2>
      <p>@{{ userId }}さんはまだメッセージを投稿していません。</p>
    </div>
  </div>
</template>

<style lang="scss" module>
.userMessagesPage {
  min-height: 100%;
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

.messagesList {
  background-color: var(--color-surface);
  border-radius: 0.5rem;
  overflow: hidden;
  box-shadow: 0 1px 3px var(--color-shadow-light);

  // 最後の要素のボーダーを削除
  & > :last-child {
    border-bottom: none;
  }
}
</style>
