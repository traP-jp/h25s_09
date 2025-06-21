<script lang="ts" setup>
import { useCreateMessage, useFormState } from '@/lib/composables'
import { validateImageFile } from '@/lib/utils/format'
import { computed, ref, watch } from 'vue'

interface Props {
  /** 返信先のメッセージID */
  repliesTo?: string
  /** プレースホルダーテキスト */
  placeholder?: string
  /** 送信ボタンのテキスト */
  submitText?: string
}

const props = withDefaults(defineProps<Props>(), {
  placeholder: '今何してる？',
  submitText: '投稿',
})

const emit = defineEmits<{
  success: []
  error: [error: Error]
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
  const fileInput = document.getElementById('image-upload') as HTMLInputElement
  if (fileInput) {
    fileInput.value = ''
  }
}

// メッセージ投稿処理
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
  <form :class="$style.messageForm" @submit.prevent="handleSubmit">
    <div :class="$style.formGroup">
      <label for="message-content" class="sr-only">メッセージ内容</label>
      <textarea
        id="message-content"
        v-model="formData.content"
        :class="$style.textarea"
        :placeholder="placeholder"
        rows="3"
        :disabled="isSubmitting"
        :aria-describedby="isSubmitting ? 'submit-status' : undefined"
      />
    </div>

    <!-- 画像アップロード -->
    <div :class="$style.imageUpload">
      <input
        id="image-upload"
        type="file"
        accept="image/*"
        :class="$style.fileInput"
        :disabled="isSubmitting"
        @change="handleImageChange"
      />
      <label for="image-upload" :class="$style.imageUploadLabel"> 📷 画像を選択 </label>

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
          ✕
        </button>
      </div>
    </div>

    <div :class="$style.formActions">
      <button
        type="submit"
        :class="$style.submitButton"
        :disabled="isSubmitting || !isFormValid"
        :aria-describedby="isSubmitting ? 'submit-status' : undefined"
      >
        {{ isSubmitting ? '投稿中...' : submitText }}
      </button>
      <div v-if="isSubmitting" id="submit-status" class="sr-only">投稿処理中です</div>
    </div>
  </form>
</template>

<style lang="scss" module>
.messageForm {
  background-color: var(--color-surface);
  border-radius: 0.5rem;
  border: 1px solid var(--color-border-light);
  box-shadow: 0 4px 6px var(--color-shadow-medium);
  padding: 1.5rem;
  margin-bottom: 1.5rem;
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
  font-family: inherit;

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
    opacity: 0.7;
  }
}

.imageUpload {
  margin-bottom: 1rem;
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

  &:focus-within {
    outline: 2px solid var(--color-primary-500);
    outline-offset: 2px;
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
  display: block;
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

  &:focus {
    outline: 2px solid var(--color-primary-300);
    outline-offset: 2px;
  }

  &:active:not(:disabled) {
    background-color: var(--color-primary-800);
    border-color: var(--color-primary-800);
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
</style>
