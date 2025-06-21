<script lang="ts" setup>
import { Icon } from '@iconify/vue'
import { onMounted, onUnmounted, ref } from 'vue'
import ReplyForm from './ReplyForm.vue'

interface Props {
  /** 返信先のメッセージID */
  messageId: string
  /** 返信先のメッセージ内容（プレビュー用） */
  messageContent: string
  /** 返信先の投稿者 */
  messageAuthor: string
  /** モーダルの表示状態 */
  isOpen: boolean
}

const props = defineProps<Props>()

const emit = defineEmits<{
  close: []
  success: []
  error: [error: Error]
}>()

// モーダル要素の参照
const modalRef = ref<HTMLElement>()

// モーダルを閉じる処理
const handleClose = () => {
  emit('close')
}

// 背景クリックでモーダルを閉じる
const handleBackdropClick = (event: MouseEvent) => {
  if (event.target === modalRef.value) {
    handleClose()
  }
}

// Escキーでモーダルを閉じる
const handleKeydown = (event: KeyboardEvent) => {
  if (event.key === 'Escape') {
    handleClose()
  }
}

// リプライ成功時の処理
const handleReplySuccess = () => {
  emit('success')
  handleClose()
}

// リプライエラー時の処理
const handleReplyError = (error: Error) => {
  emit('error', error)
}

// キーボードイベントの設定
onMounted(() => {
  if (props.isOpen) {
    document.addEventListener('keydown', handleKeydown)
    // フォーカストラップ
    const firstFocusable = modalRef.value?.querySelector(
      'button, [href], input, select, textarea, [tabindex]:not([tabindex="-1"])',
    ) as HTMLElement
    firstFocusable?.focus()
  }
})

onUnmounted(() => {
  document.removeEventListener('keydown', handleKeydown)
})
</script>

<template>
  <Teleport to="body">
    <div
      v-if="isOpen"
      ref="modalRef"
      :class="$style.modalOverlay"
      @click="handleBackdropClick"
      role="dialog"
      aria-modal="true"
      aria-labelledby="reply-modal-title"
    >
      <div :class="$style.modalContent">
        <header :class="$style.modalHeader">
          <h2 id="reply-modal-title" :class="$style.modalTitle">
            <Icon icon="mdi:reply" :class="$style.replyIcon" />
            返信を書く
          </h2>
          <button :class="$style.closeButton" @click="handleClose" aria-label="モーダルを閉じる">
            <Icon icon="mdi:close" />
          </button>
        </header>

        <div :class="$style.originalMessage">
          <p :class="$style.originalMessageLabel">返信先:</p>
          <div :class="$style.originalMessageContent">
            <span :class="$style.originalMessageAuthor">@{{ messageAuthor }}</span>
            <span :class="$style.originalMessageText">{{ messageContent }}</span>
          </div>
        </div>

        <div :class="$style.replyFormContainer">
          <ReplyForm
            :replies-to="messageId"
            :auto-focus="true"
            @success="handleReplySuccess"
            @error="handleReplyError"
            @cancel="handleClose"
          />
        </div>
      </div>
    </div>
  </Teleport>
</template>

<style lang="scss" module>
.modalOverlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
  padding: 1rem;
}

.modalContent {
  background-color: var(--color-surface);
  border-radius: 0.75rem;
  box-shadow:
    0 20px 25px -5px rgba(0, 0, 0, 0.1),
    0 10px 10px -5px rgba(0, 0, 0, 0.04);
  max-width: 600px;
  width: 100%;
  max-height: 80vh;
  overflow-y: auto;
  border: 1px solid var(--color-border-light);
}

.modalHeader {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 1.5rem 1.5rem 0 1.5rem;
  border-bottom: 1px solid var(--color-border-light);
  margin-bottom: 1rem;
}

.modalTitle {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin: 0;
  font-size: 1.25rem;
  font-weight: 600;
  color: var(--color-text-primary);
}

.replyIcon {
  font-size: 1.375rem;
  color: var(--color-primary-600);
}

.closeButton {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 2rem;
  height: 2rem;
  border: none;
  background: none;
  color: var(--color-text-secondary);
  cursor: pointer;
  border-radius: 0.375rem;
  transition: all 0.2s ease;

  &:hover {
    background-color: var(--color-surface-variant);
    color: var(--color-text-primary);
  }

  &:focus {
    outline: 2px solid var(--color-primary-500);
    outline-offset: 2px;
  }
}

.originalMessage {
  padding: 0 1.5rem;
  margin-bottom: 1rem;
}

.originalMessageLabel {
  margin: 0 0 0.5rem 0;
  font-size: 0.875rem;
  font-weight: 600;
  color: var(--color-text-secondary);
}

.originalMessageContent {
  background-color: var(--color-surface-variant);
  border: 1px solid var(--color-border-light);
  border-radius: 0.5rem;
  padding: 0.75rem;
  border-left: 3px solid var(--color-primary-300);
}

.originalMessageAuthor {
  display: block;
  font-size: 0.875rem;
  font-weight: 600;
  color: var(--color-primary-600);
  margin-bottom: 0.25rem;
}

.originalMessageText {
  display: block;
  font-size: 0.875rem;
  color: var(--color-text-primary);
  line-height: 1.4;
  white-space: pre-wrap;
  word-wrap: break-word;
  // 長いテキストは省略
  overflow: hidden;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  line-clamp: 3;
  -webkit-box-orient: vertical;
}

.replyFormContainer {
  padding: 0 1.5rem 1.5rem 1.5rem;
}

/* レスポンシブ対応 */
@media (max-width: 768px) {
  .modalOverlay {
    padding: 0.5rem;
    align-items: flex-start;
    padding-top: 2rem;
  }

  .modalContent {
    max-height: 90vh;
    border-radius: 0.5rem;
  }

  .modalHeader {
    padding: 1rem 1rem 0 1rem;
  }

  .modalTitle {
    font-size: 1.125rem;
  }

  .originalMessage,
  .replyFormContainer {
    padding-left: 1rem;
    padding-right: 1rem;
  }

  .replyFormContainer {
    padding-bottom: 1rem;
  }
}
</style>
