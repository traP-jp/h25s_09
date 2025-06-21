<script lang="ts" setup>
import {
  useAddReaction,
  useCreateMessage,
  useFormState,
  useMessageDetail,
  useRemoveReaction,
} from '@/lib/composables'
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const route = useRoute()
const router = useRouter()

// URLパラメータからメッセージIDを取得
const messageId = computed(() => route.params.id as string)

// メッセージ詳細取得
const { data: message, isLoading, error } = useMessageDetail(messageId)

// リプライ作成
const createMessageMutation = useCreateMessage()

// リアクション機能
const addReactionMutation = useAddReaction()
const removeReactionMutation = useRemoveReaction()

// リプライフォーム状態管理
const {
  formData: replyForm,
  resetForm,
  setSubmitting,
  isSubmitting,
} = useFormState({
  content: '',
})

// リアクション切り替え処理
const toggleReaction = async () => {
  if (!messageId.value || !message.value) return

  try {
    if (message.value.reactions.myReaction) {
      await removeReactionMutation.mutateAsync(messageId.value)
    } else {
      await addReactionMutation.mutateAsync(messageId.value)
    }
  } catch (error) {
    console.error('Failed to toggle reaction:', error)
  }
}

// リプライ投稿処理
const handleReplySubmit = async () => {
  if (!replyForm.value.content.trim() || !messageId.value) return

  setSubmitting(true)

  try {
    await createMessageMutation.mutateAsync({
      content: replyForm.value.content,
      repliesTo: messageId.value,
    })
    resetForm()
  } catch (error) {
    console.error('Failed to create reply:', error)
  } finally {
    setSubmitting(false)
  }
}

// 戻る
const goBack = () => {
  router.back()
}
</script>

<template>
  <main :class="$style.messageDetail">
    <div :class="$style.container">
      <!-- ヘッダー -->
      <div :class="$style.header">
        <button :class="$style.backButton" @click="goBack">← 戻る</button>
        <h1>メッセージの詳細</h1>
      </div>

      <!-- エラー表示 -->
      <div v-if="error" :class="$style.error">
        <p>メッセージの取得に失敗しました</p>
        <button @click="goBack">戻る</button>
      </div>

      <!-- ローディング表示 -->
      <div v-else-if="isLoading" :class="$style.loading">
        <p>読み込み中...</p>
      </div>

      <!-- メッセージ詳細 -->
      <div v-else-if="message" :class="$style.content">
        <!-- メインメッセージ -->
        <div :class="$style.mainMessage">
          <div :class="$style.messageHeader">
            <div :class="$style.userInfo">
              <strong>{{ message.author || 'Unknown User' }}</strong>
              <span :class="$style.timestamp">{{
                new Date(message.createdAt).toLocaleString()
              }}</span>
            </div>
          </div>
          <div :class="$style.messageContent">
            {{ message.content }}
          </div>

          <!-- リアクション表示（読み取り専用） -->
          <div v-if="message.reactions && message.reactions.count > 0" :class="$style.reactions">
            <div :class="$style.reaction">
              👍 {{ message.reactions.count }}
              <span v-if="message.reactions.myReaction" :class="$style.myReaction"
                >（あなたがリアクション済み）</span
              >
            </div>
          </div>

          <!-- リアクションボタン -->
          <div :class="$style.reactionActions">
            <button
              :class="[$style.reactionButton, { [$style.active]: message.reactions.myReaction }]"
              @click="toggleReaction"
              :disabled="
                addReactionMutation.isPending.value || removeReactionMutation.isPending.value
              "
            >
              👍 {{ message.reactions.myReaction ? 'リアクション済み' : 'リアクション' }}
            </button>
          </div>
        </div>

        <!-- リプライフォーム -->
        <div :class="$style.replyForm">
          <h3>このメッセージにリプライ</h3>
          <form @submit.prevent="handleReplySubmit">
            <div :class="$style.formGroup">
              <textarea
                v-model="replyForm.content"
                :class="$style.textarea"
                placeholder="リプライ内容を入力..."
                rows="3"
                :disabled="isSubmitting"
              />
            </div>
            <div :class="$style.formActions">
              <button
                type="submit"
                :class="$style.submitButton"
                :disabled="isSubmitting || !replyForm.content.trim()"
              >
                {{ isSubmitting ? 'リプライ中...' : 'リプライ' }}
              </button>
            </div>
          </form>
        </div>

        <!-- リプライ一覧（読み取り専用） -->
        <div v-if="message?.replies && message.replies.length > 0" :class="$style.replies">
          <h3>リプライ ({{ message.replies.length }})</h3>
          <div v-for="reply in message.replies" :key="reply.id" :class="$style.reply">
            <div :class="$style.replyHeader">
              <strong>{{ reply.author || 'Unknown User' }}</strong>
              <span :class="$style.timestamp">{{
                new Date(reply.createdAt).toLocaleString()
              }}</span>
            </div>
            <div :class="$style.replyContent">
              {{ reply.content }}
            </div>
          </div>
        </div>
      </div>

      <!-- メッセージが見つからない場合 -->
      <div v-else :class="$style.notFound">
        <p>メッセージが見つかりません</p>
        <button @click="goBack">戻る</button>
      </div>
    </div>
  </main>
</template>

<style lang="scss" module>
.messageDetail {
  padding: 1rem;
  max-width: 800px;
  margin: 0 auto;
}

.container {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.header {
  display: flex;
  align-items: center;
  gap: 1rem;

  h1 {
    margin: 0;
    color: var(--color-text-primary);
  }
}

.backButton {
  background: none;
  border: 1px solid var(--color-border-medium);
  border-radius: 4px;
  padding: 0.5rem 1rem;
  cursor: pointer;
  color: var(--color-primary-600);

  &:hover {
    background: var(--color-surface-variant);
  }
}

.error,
.notFound {
  background: var(--color-error-50);
  color: var(--color-error-700);
  padding: 1.5rem;
  border-radius: 8px;
  text-align: center;

  button {
    background: var(--color-error-600);
    color: var(--color-text-inverse);
    border: none;
    border-radius: 4px;
    padding: 0.5rem 1rem;
    margin-top: 1rem;
    cursor: pointer;

    &:hover {
      background: var(--color-error-700);
    }
  }
}

.loading {
  text-align: center;
  padding: 3rem;
  color: var(--color-text-secondary);
}

.content {
  display: flex;
  flex-direction: column;
  gap: 2rem;
}

.mainMessage {
  background: var(--color-surface);
  border-radius: 8px;
  padding: 1.5rem;
  box-shadow: 0 2px 4px var(--color-shadow-light);
}

.messageHeader {
  margin-bottom: 1rem;
}

.userInfo {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.timestamp {
  color: var(--color-text-secondary);
  font-size: 0.875rem;
}

.messageContent {
  font-size: 1.1rem;
  line-height: 1.6;
  margin-bottom: 1rem;
  white-space: pre-wrap;
}

.reactions {
  display: flex;
  gap: 0.5rem;
  margin-bottom: 1rem;
}

.reaction {
  background: var(--color-surface-variant);
  border: 1px solid var(--color-border-light);
  border-radius: 20px;
  padding: 0.25rem 0.75rem;
  font-size: 0.875rem;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 0.5rem;

  &:hover {
    background: var(--color-surface);
  }
}

.myReaction {
  font-size: 0.75rem;
  color: var(--color-primary-600);
  font-style: italic;
}

.reactionActions {
  display: flex;
  gap: 0.25rem;
  margin-top: 1rem;
}

.emojiButton {
  background: none;
  border: 1px solid #ddd;
  border-radius: 4px;
  padding: 0.25rem 0.5rem;
  font-size: 1.2rem;
  cursor: pointer;

  &:hover {
    background: #f8f9fa;
  }
}

.replies {
  h3 {
    margin-bottom: 1rem;
    color: var(--color-text-primary);
  }
}

.reply {
  background: var(--color-surface-variant);
  border-radius: 8px;
  padding: 1rem;
  margin-bottom: 1rem;
}

.replyHeader {
  display: flex;
  align-items: center;
  gap: 1rem;
  margin-bottom: 0.5rem;
}

.replyContent {
  line-height: 1.5;
  white-space: pre-wrap;
}

.replyForm {
  background: var(--color-surface);
  border-radius: 8px;
  padding: 1.5rem;
  box-shadow: 0 2px 4px var(--color-shadow-light);
}

.formGroup {
  margin-bottom: 1rem;
}

.textarea {
  width: 100%;
  border: 1px solid var(--color-border-medium);
  border-radius: 4px;
  padding: 0.75rem;
  font-size: 1rem;
  resize: vertical;
  background-color: var(--color-surface);
  color: var(--color-text-primary);

  &:focus {
    outline: none;
    border-color: var(--color-primary-500);
    box-shadow: 0 0 0 2px var(--color-primary-200);
  }

  &:disabled {
    background-color: var(--color-surface-variant);
    cursor: not-allowed;
  }
}

.formActions {
  display: flex;
  justify-content: flex-end;
}

.submitButton {
  background: var(--color-primary-600);
  color: var(--color-text-inverse);
  border: none;
  border-radius: 4px;
  padding: 0.5rem 1.5rem;
  font-size: 1rem;
  cursor: pointer;
  transition: background-color 0.2s;

  &:hover:not(:disabled) {
    background: var(--color-primary-700);
  }

  &:disabled {
    background: var(--color-gray-500);
    cursor: not-allowed;
  }
}

.reactionButton {
  background: none;
  border: 1px solid var(--color-border-medium);
  border-radius: 4px;
  padding: 0.25rem 0.5rem;
  font-size: 1.2rem;
  cursor: pointer;

  &:hover {
    background: var(--color-surface-variant);
  }
}
</style>
