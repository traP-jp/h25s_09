<script lang="ts" setup>
import { useCreateMessage, useFormState } from '@/lib/composables'
import { validateImageFile } from '@/lib/utils/format'
import { Icon } from '@iconify/vue'
import { computed, ref, watch } from 'vue'

interface Props {
  /** 返信先のメッセージID */
  repliesTo: string
  /** プレースホルダーテキスト */
  placeholder?: string
  /** 自動フォーカスを有効にするか */
  autoFocus?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  placeholder: '返信を入力してください...',
  autoFocus: false,
})

const emit = defineEmits<{
  success: []
  error: [error: Error]
  cancel: []
}>()

// 画像プレビューURL管理
const imagePreviewUrl = ref<string | null>(null)

// メッセージ作成
const createMessageMutation = useCreateMessage()

// フォーム状態管理
const { formData, resetForm, setSubmitting, isSubmitting } = useFormState({
  content: '',
  image: null as File | null,
})

// フォームの有効性チェック
const isFormValid = computed(() => {
  return formData.value.content.trim().length > 0
})

// 画像のクリーンアップ
const cleanupImagePreview = () => {
  if (imagePreviewUrl.value) {
    window.URL.revokeObjectURL(imagePreviewUrl.value)
    imagePreviewUrl.value = null
  }
}

// 画像アップロード処理
const handleImageChange = (event: Event) => {
  const input = event.target as HTMLInputElement
  const file = input.files?.[0]

  if (file) {
    // 画像のバリデーション
    const validation = validateImageFile(file)
    if (!validation.isValid) {
      emit('error', new Error(validation.error))
      return
    }

    formData.value.image = file

    // 古いプレビューURLをクリーンアップ
    cleanupImagePreview()

    // 新しいプレビューURL生成
    imagePreviewUrl.value = window.URL.createObjectURL(file)
  }
}

// 画像削除処理
const removeImage = () => {
  formData.value.image = null
  cleanupImagePreview()

  // file inputもリセット
  const fileInput = document.getElementById(
    `reply-image-upload-${props.repliesTo}`,
  ) as HTMLInputElement
  if (fileInput) {
    fileInput.value = ''
  }
}

// リプライ投稿処理
const handleSubmit = async () => {
  if (!isFormValid.value) return

  setSubmitting(true)

  try {
    await createMessageMutation.mutateAsync({
      content: formData.value.content,
      image: formData.value.image || undefined,
      repliesTo: props.repliesTo,
    })

    resetForm()
    cleanupImagePreview()
    emit('success')
  } catch (error) {
    emit('error', error as Error)
  } finally {
    setSubmitting(false)
  }
}

// キャンセル処理
const handleCancel = () => {
  resetForm()
  cleanupImagePreview()
  emit('cancel')
}

// Escキーでキャンセル
const handleKeydown = (event: KeyboardEvent) => {
  if (event.key === 'Escape') {
    handleCancel()
  }
}

// コンポーネントがアンマウントされる際のクリーンアップ
watch(
  () => formData.value.image,
  (newImage, oldImage) => {
    if (oldImage && !newImage) {
      cleanupImagePreview()
    }
  },
)
</script>

<template>
  <form :class="$style.replyForm" @submit.prevent="handleSubmit" @keydown="handleKeydown">
    <div :class="$style.formGroup">
      <label :for="`reply-content-${repliesTo}`" class="sr-only">返信内容</label>
      <textarea
        :id="`reply-content-${repliesTo}`"
        v-model="formData.content"
        :class="$style.textarea"
        :placeholder="placeholder"
        rows="3"
        :disabled="isSubmitting"
        :autofocus="autoFocus"
        :aria-describedby="isSubmitting ? 'submit-status' : undefined"
      />
    </div>

    <!-- 画像アップロード -->
    <div :class="$style.imageUpload">
      <input
        :id="`reply-image-upload-${repliesTo}`"
        type="file"
        accept="image/*"
        :class="$style.fileInput"
        :disabled="isSubmitting"
        @change="handleImageChange"
      />
      <label :for="`reply-image-upload-${repliesTo}`" :class="$style.imageUploadLabel">
        <Icon icon="mdi:camera" />
        画像を選択
      </label>

      <!-- 選択された画像のプレビュー -->
      <div v-if="formData.image && imagePreviewUrl" :class="$style.imagePreview">
        <img :src="imagePreviewUrl" alt="プレビュー" :class="$style.previewImage" loading="lazy" />
        <button
          type="button"
          :class="$style.removeImageButton"
          :disabled="isSubmitting"
          @click="removeImage"
          aria-label="画像を削除"
        >
          <Icon icon="mdi:close" />
        </button>
      </div>
    </div>

    <div :class="$style.formActions">
      <button
        type="button"
        :class="$style.cancelButton"
        :disabled="isSubmitting"
        @click="handleCancel"
      >
        キャンセル
      </button>
      <button
        type="submit"
        :class="$style.submitButton"
        :disabled="isSubmitting || !isFormValid"
        :aria-describedby="isSubmitting ? 'submit-status' : undefined"
      >
        <span :class="$style.submitButtonContent">
          {{ isSubmitting ? '返信中...' : '返信する' }}
          <Icon v-if="!isSubmitting" icon="mdi:send" :class="$style.submitButtonIcon" />
          <Icon v-else icon="mdi:loading" :class="[$style.submitButtonIcon, $style.spinning]" />
        </span>
      </button>
      <div v-if="isSubmitting" id="submit-status" class="sr-only">返信処理中です</div>
    </div>
  </form>
</template>

<style lang="scss" module>
.replyForm {
  background-color: var(--color-surface);
  border-radius: 0.5rem;
  border: 1px solid var(--color-border-light);
  padding: 1rem;
}

.formGroup {
  margin-bottom: 0.75rem;
}

.textarea {
  width: 100%;
  padding: 0.625rem;
  border: 1px solid var(--color-border-medium);
  border-radius: 0.375rem;
  background-color: var(--color-surface);
  color: var(--color-text-primary);
  font-size: 0.875rem;
  transition: border-color 0.15s ease-in-out;
  resize: vertical;
  min-height: 60px;
  font-family: inherit;

  &:focus {
    outline: none;
    border-color: var(--color-primary-500);
    box-shadow: 0 0 0 2px var(--color-primary-200);
  }

  &::placeholder {
    color: var(--color-text-tertiary);
  }

  &:disabled {
    background-color: var(--color-surface-variant);
    cursor: not-allowed;
    opacity: 0.7;
  }
}

.imageUpload {
  margin-bottom: 0.75rem;
}

.fileInput {
  position: absolute;
  width: 1px;
  height: 1px;
  padding: 0;
  margin: -1px;
  overflow: hidden;
  clip: rect(0, 0, 0, 0);
  white-space: nowrap;
  border: 0;
}

.imageUploadLabel {
  display: inline-flex;
  align-items: center;
  gap: 0.375rem;
  padding: 0.375rem 0.75rem;
  background: var(--color-surface-variant);
  border: 1px solid var(--color-border-medium);
  border-radius: 0.375rem;
  cursor: pointer;
  font-size: 0.8rem;
  color: var(--color-text-secondary);
  transition: all 0.2s ease;

  &:hover {
    background: var(--color-surface);
    border-color: var(--color-border-strong);
  }

  &:focus-within {
    outline: 2px solid var(--color-primary-500);
    outline-offset: 2px;
  }
}

.imagePreview {
  position: relative;
  margin-top: 0.5rem;
  display: inline-block;
}

.previewImage {
  max-width: 150px;
  max-height: 150px;
  border-radius: 0.375rem;
  border: 1px solid var(--color-border-light);
  object-fit: cover;
  display: block;
}

.removeImageButton {
  position: absolute;
  top: -6px;
  right: -6px;
  width: 20px;
  height: 20px;
  border-radius: 50%;
  background: var(--color-error-600);
  color: white;
  border: none;
  font-size: 0.7rem;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: background-color 0.2s ease;

  &:hover {
    background: var(--color-error-700);
  }

  &:focus {
    outline: 2px solid var(--color-error-300);
    outline-offset: 2px;
  }

  &:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }
}

.formActions {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  gap: 0.5rem;
}

.cancelButton {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 0.5rem 0.875rem;
  border-radius: 0.375rem;
  font-weight: 500;
  font-size: 0.875rem;
  line-height: 1.25rem;
  cursor: pointer;
  transition: all 0.15s ease-in-out;
  background-color: var(--color-surface-variant);
  color: var(--color-text-secondary);
  border: 1px solid var(--color-border-medium);

  &:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }

  &:hover:not(:disabled) {
    background-color: var(--color-surface);
    border-color: var(--color-border-strong);
    color: var(--color-text-primary);
  }

  &:focus {
    outline: 2px solid var(--color-primary-300);
    outline-offset: 2px;
  }
}

.submitButton {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 0.5rem 0.875rem;
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

  &:focus {
    outline: 2px solid var(--color-primary-300);
    outline-offset: 2px;
  }

  &:active:not(:disabled) {
    background-color: var(--color-primary-800);
    border-color: var(--color-primary-800);
  }
}

.submitButtonContent {
  display: flex;
  align-items: center;
  gap: 0.375rem;
}

.submitButtonIcon {
  font-size: 0.875rem;
  flex-shrink: 0;
}

.spinning {
  animation: spin 1s linear infinite;
}

@keyframes spin {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}

// スクリーンリーダー専用のスタイル
.sr-only {
  position: absolute;
  width: 1px;
  height: 1px;
  padding: 0;
  margin: -1px;
  overflow: hidden;
  clip: rect(0, 0, 0, 0);
  white-space: nowrap;
  border: 0;
}

/* レスポンシブ対応 */
@media (max-width: 768px) {
  .replyForm {
    padding: 0.75rem;
  }

  .textarea {
    font-size: 0.8rem;
    padding: 0.5rem;
  }

  .formActions {
    flex-direction: column-reverse;
    align-items: stretch;
    gap: 0.375rem;
  }

  .cancelButton,
  .submitButton {
    width: 100%;
    justify-content: center;
  }

  .imageUploadLabel {
    font-size: 0.75rem;
    padding: 0.3rem 0.6rem;
  }

  .previewImage {
    max-width: 120px;
    max-height: 120px;
  }
}
</style>
