<script lang="ts" setup>
import MessageItem from '@/components/MessageItem.vue'
import { useCreateMessage, useFormState, useMessages } from '@/lib/composables'
import { ref } from 'vue'

// 画像プレビューURL管理
const imagePreviewUrl = ref<string | null>(null)

// メッセージ一覧取得
const { data: messages, isLoading, error, refetch } = useMessages()

// メッセージ作成
const createMessageMutation = useCreateMessage()

// フォーム状態管理
const { formData, resetForm, setSubmitting, isSubmitting } = useFormState({
  content: '',
  image: null as File | null,
})

// 画像アップロード処理
const handleImageChange = (event: Event) => {
  const input = event.target as HTMLInputElement
  const file = input.files?.[0]
  if (file) {
    formData.value.image = file
    // プレビューURL生成
    if (imagePreviewUrl.value) {
      window.URL.revokeObjectURL(imagePreviewUrl.value)
    }
    imagePreviewUrl.value = window.URL.createObjectURL(file)
  }
}

// 画像削除処理
const removeImage = () => {
  formData.value.image = null
  if (imagePreviewUrl.value) {
    window.URL.revokeObjectURL(imagePreviewUrl.value)
    imagePreviewUrl.value = null
  }
}

// メッセージ投稿処理
const handleSubmit = async () => {
  if (!formData.value.content.trim()) return

  setSubmitting(true)

  try {
    await createMessageMutation.mutateAsync({
      content: formData.value.content,
      image: formData.value.image || undefined,
    })
    resetForm()
    // 画像プレビューもクリーンアップ
    if (imagePreviewUrl.value) {
      window.URL.revokeObjectURL(imagePreviewUrl.value)
      imagePreviewUrl.value = null
    }
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

        <!-- 画像アップロード -->
        <div :class="$style.imageUpload">
          <input
            id="image-upload"
            type="file"
            accept="image/*"
            :class="$style.fileInput"
            @change="handleImageChange"
          />
          <label for="image-upload" :class="$style.imageUploadLabel"> 📷 画像を選択 </label>

          <!-- 選択された画像のプレビュー -->
          <div v-if="formData.image && imagePreviewUrl" :class="$style.imagePreview">
            <img :src="imagePreviewUrl" alt="プレビュー" :class="$style.previewImage" />
            <button type="button" :class="$style.removeImageButton" @click="removeImage">✕</button>
          </div>
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
  background-color: var(--color-surface);
  border-radius: 0.5rem;
  border: 1px solid var(--color-border-light);
  box-shadow: 0 4px 6px var(--color-shadow-medium);
  padding: 1.5rem;
}

.formGroup {
  margin-bottom: 1rem;
}

.textarea {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid var(--color-border-medium);
  border-radius: 0.375rem;
  background-color: var(--color-surface);
  color: var(--color-text-primary);
  font-size: 0.875rem;
  transition: border-color 0.15s ease-in-out;
  resize: vertical;
  min-height: 80px;

  &:focus {
    outline: none;
    border-color: var(--color-primary-500);
    box-shadow: 0 0 0 3px var(--color-primary-200);
  }

  &::placeholder {
    color: var(--color-text-tertiary);
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
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 0.5rem 1rem;
  border-radius: 0.375rem;
  font-weight: 500;
  font-size: 0.875rem;
  line-height: 1.25rem;
  border: 1px solid transparent;
  cursor: pointer;
  transition: all 0.15s ease-in-out;
  background-color: var(--color-primary-600);
  color: var(--color-text-inverse);
  border-color: var(--color-primary-600);

  &:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }

  &:hover:not(:disabled) {
    background-color: var(--color-primary-700);
    border-color: var(--color-primary-700);
  }

  &:active {
    background-color: var(--color-primary-800);
    border-color: var(--color-primary-800);
  }
}

.imageUpload {
  margin-top: 0.75rem;
}

.fileInput {
  display: none;
}

.imageUploadLabel {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 1rem;
  background: var(--color-surface-variant);
  border: 1px solid var(--color-border-medium);
  border-radius: 0.375rem;
  cursor: pointer;
  font-size: 0.875rem;
  color: var(--color-text-secondary);
  transition: all 0.2s ease;

  &:hover {
    background: var(--color-surface);
    border-color: var(--color-border-strong);
  }
}

.imagePreview {
  position: relative;
  margin-top: 0.75rem;
  display: inline-block;
}

.previewImage {
  max-width: 200px;
  max-height: 200px;
  border-radius: 0.5rem;
  border: 1px solid var(--color-border-light);
  object-fit: cover;
}

.removeImageButton {
  position: absolute;
  top: -8px;
  right: -8px;
  width: 24px;
  height: 24px;
  border-radius: 50%;
  background: var(--color-error-600);
  color: white;
  border: none;
  font-size: 0.75rem;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;

  &:hover {
    background: var(--color-error-700);
  }
}

.error {
  background-color: var(--color-error-50);
  color: var(--color-error-700);
  padding: 1rem;
  border-radius: 0.5rem;
  border: 1px solid var(--color-error-200);

  button {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    padding: 0.25rem 0.75rem;
    border-radius: 0.375rem;
    font-weight: 500;
    font-size: 0.875rem;
    line-height: 1.25rem;
    border: 1px solid transparent;
    cursor: pointer;
    transition: all 0.15s ease-in-out;
    background-color: var(--color-error-600);
    color: var(--color-text-inverse);
    margin-top: 0.5rem;

    &:disabled {
      opacity: 0.5;
      cursor: not-allowed;
    }

    &:hover:not(:disabled) {
      background-color: var(--color-error-700);
    }
  }
}

.loading {
  text-align: center;
  padding: 2rem;
  color: var(--color-text-secondary);
}

.messageList {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.empty {
  text-align: center;
  padding: 3rem;
  color: var(--color-text-secondary);
  font-style: italic;
}

:global([data-theme='dark']) .error {
  background-color: var(--color-error-900);
  color: var(--color-error-200);
  border-color: var(--color-error-800);
}
</style>
