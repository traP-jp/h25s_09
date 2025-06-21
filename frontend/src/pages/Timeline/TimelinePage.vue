<script lang="ts" setup>
import MessageItem from '@/components/MessageItem.vue'
import { useCreateMessage, useFormState, useMessages } from '@/lib/composables'

// メッセージ一覧取得
const { data: messages, isLoading, error, refetch } = useMessages()

// メッセージ作成
const createMessageMutation = useCreateMessage()

// フォーム状態管理
const { formData, resetForm, setSubmitting, isSubmitting } = useFormState({
  content: '',
})

// メッセージ投稿処理
const handleSubmit = async () => {
  if (!formData.value.content.trim()) return

  setSubmitting(true)

  try {
    await createMessageMutation.mutateAsync(formData.value.content)
    resetForm()
  } catch (error) {
    console.error('Failed to create message:', error)
  } finally {
    setSubmitting(false)
  }
}
</script>

<template>
  <main :class="$style.timeline">
    <div :class="$style.container">
      <!-- メッセージ投稿フォーム -->
      <form :class="$style.messageForm" @submit.prevent="handleSubmit">
        <div :class="$style.formGroup">
          <textarea
            v-model="formData.content"
            :class="$style.textarea"
            placeholder="今何してる？"
            rows="3"
            :disabled="isSubmitting"
          />
        </div>
        <div :class="$style.formActions">
          <button
            type="submit"
            :class="$style.submitButton"
            :disabled="isSubmitting || !formData.content.trim()"
          >
            {{ isSubmitting ? '投稿中...' : '投稿' }}
          </button>
        </div>
      </form>

      <!-- エラー表示 -->
      <div v-if="error" :class="$style.error">
        <p>メッセージの取得に失敗しました</p>
        <button @click="() => refetch()">再試行</button>
      </div>

      <!-- ローディング表示 -->
      <div v-if="isLoading" :class="$style.loading">
        <p>読み込み中...</p>
      </div>

      <!-- メッセージ一覧 -->
      <div v-else-if="messages && messages.length > 0" :class="$style.messageList">
        <MessageItem v-for="message in messages" :key="message.id" :message="message" />
      </div>

      <!-- 空状態 -->
      <div v-else :class="$style.empty">
        <p>まだメッセージがありません</p>
      </div>
    </div>
  </main>
</template>

<style lang="scss" module>
.timeline {
  padding: 1rem;
  max-width: 800px;
  margin: 0 auto;
}

.container {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.messageForm {
  background: white;
  border-radius: 8px;
  padding: 1.5rem;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.formGroup {
  margin-bottom: 1rem;
}

.textarea {
  width: 100%;
  border: 1px solid #ddd;
  border-radius: 4px;
  padding: 0.75rem;
  font-size: 1rem;
  resize: vertical;
  min-height: 80px;

  &:focus {
    outline: none;
    border-color: #007bff;
    box-shadow: 0 0 0 2px rgba(0, 123, 255, 0.25);
  }

  &:disabled {
    background-color: #f8f9fa;
    cursor: not-allowed;
  }
}

.formActions {
  display: flex;
  justify-content: flex-end;
}

.submitButton {
  background: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  padding: 0.5rem 1.5rem;
  font-size: 1rem;
  cursor: pointer;
  transition: background-color 0.2s;

  &:hover:not(:disabled) {
    background: #0056b3;
  }

  &:disabled {
    background: #6c757d;
    cursor: not-allowed;
  }
}

.error {
  background: #f8d7da;
  color: #721c24;
  padding: 1rem;
  border-radius: 4px;
  border: 1px solid #f5c6cb;

  button {
    background: #dc3545;
    color: white;
    border: none;
    border-radius: 4px;
    padding: 0.25rem 0.75rem;
    margin-top: 0.5rem;
    cursor: pointer;

    &:hover {
      background: #c82333;
    }
  }
}

.loading {
  text-align: center;
  padding: 2rem;
  color: #6c757d;
}

.messageList {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.empty {
  text-align: center;
  padding: 3rem;
  color: #6c757d;
  font-style: italic;
}
</style>
