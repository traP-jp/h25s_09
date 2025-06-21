<script lang="ts" setup>
import type { Message } from '@/lib/apis/generated'
import { useMessages } from '@/lib/composables'
import { computed } from 'vue'
import MessageItem from './MessageItem.vue'

interface Props {
  /** 表示するメッセージ配列（外部から渡される場合） */
  messages?: Message[]
  /** ローディング状態（外部から渡される場合） */
  loading?: boolean
  /** エラー状態（外部から渡される場合） */
  error?: Error | null
}

const props = defineProps<Props>()

// 内部でメッセージを取得する場合
const {
  data: internalMessages,
  isLoading: internalLoading,
  error: internalError,
  refetch,
} = useMessages()

// 表示用のデータを決定
const displayMessages = computed(() => props.messages || internalMessages.value || [])
const isLoading = computed(() => props.loading ?? internalLoading.value)
const displayError = computed(() => props.error ?? internalError.value)

// エラー時の再試行
const handleRetry = () => {
  refetch()
}
</script>

<template>
  <div :class="$style.messageListContainer">
    <!-- エラー表示 -->
    <div v-if="displayError" :class="$style.error" role="alert">
      <div :class="$style.errorContent">
        <div :class="$style.errorIcon" aria-hidden="true">⚠️</div>
        <div :class="$style.errorMessage">
          <strong>メッセージの取得に失敗しました</strong>
          <p>{{ displayError.message }}</p>
        </div>
        <button
          :class="$style.retryButton"
          @click="handleRetry"
          aria-label="メッセージの再読み込み"
        >
          再試行
        </button>
      </div>
    </div>

    <!-- ローディング表示 -->
    <div v-else-if="isLoading" :class="$style.loading" role="status" aria-live="polite">
      <div :class="$style.loadingSpinner" aria-hidden="true"></div>
      <p>読み込み中...</p>
    </div>

    <!-- メッセージ一覧 -->
    <div v-else-if="displayMessages.length > 0" :class="$style.messageList">
      <MessageItem v-for="message in displayMessages" :key="message.id" :message="message" />
    </div>

    <!-- 空状態 -->
    <div v-else :class="$style.empty" role="status">
      <div :class="$style.emptyIcon" aria-hidden="true">💬</div>
      <p>まだメッセージがありません</p>
      <p :class="$style.emptySubtext">最初のメッセージを投稿してみませんか？</p>
    </div>
  </div>
</template>

<style lang="scss" module>
.messageListContainer {
  width: 100%;
}

.error {
  background-color: var(--color-error-50);
  border: 1px solid var(--color-error-200);
  border-radius: 0.5rem;
  padding: 1rem;
  margin-bottom: 1rem;

  [data-theme='dark'] & {
    background-color: var(--color-error-900);
    border-color: var(--color-error-800);
  }
}

.errorContent {
  display: flex;
  align-items: flex-start;
  gap: 0.75rem;
}

.errorIcon {
  font-size: 1.5rem;
  flex-shrink: 0;
}

.errorMessage {
  flex: 1;
  color: var(--color-error-700);

  [data-theme='dark'] & {
    color: var(--color-error-200);
  }

  strong {
    display: block;
    margin-bottom: 0.25rem;
    font-size: 0.875rem;
  }

  p {
    margin: 0;
    font-size: 0.8125rem;
    line-height: 1.4;
  }
}

.retryButton {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 0.5rem 0.75rem;
  border-radius: 0.375rem;
  font-weight: 500;
  font-size: 0.75rem;
  line-height: 1.25rem;
  border: 1px solid transparent;
  cursor: pointer;
  transition: all 0.15s ease-in-out;
  background-color: var(--color-error-600);
  color: var(--color-text-inverse);
  flex-shrink: 0;

  &:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }

  &:hover:not(:disabled) {
    background-color: var(--color-error-700);
  }

  &:focus {
    outline: 2px solid var(--color-error-300);
    outline-offset: 2px;
  }
}

.loading {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 3rem 2rem;
  color: var(--color-text-secondary);
  text-align: center;
}

.loadingSpinner {
  width: 2rem;
  height: 2rem;
  border: 2px solid var(--color-border-light);
  border-top: 2px solid var(--color-primary-600);
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin-bottom: 1rem;
}

@keyframes spin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}

.messageList {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  padding: 0.5rem;
  border-radius: 0.5rem;
  border: 1px solid var(--color-border-light);
  background-color: var(--color-background);
  box-shadow: 0 1px 3px var(--color-shadow-light);
}

.empty {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 4rem 2rem;
  text-align: center;
  color: var(--color-text-secondary);
  background-color: var(--color-surface);
  border-radius: 0.5rem;
  border: 1px solid var(--color-border-light);
}

.emptyIcon {
  font-size: 3rem;
  margin-bottom: 1rem;
  opacity: 0.5;
}

.emptySubtext {
  font-size: 0.875rem;
  margin-top: 0.5rem;
  color: var(--color-text-tertiary);
}
</style>
