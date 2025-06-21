<script lang="ts" setup>
import MessageForm from '@/components/MessageForm.vue'
import { useGlobalError } from '@/lib/composables'
import { Icon } from '@iconify/vue'
import { onMounted, ref } from 'vue'

const { setError } = useGlobalError()

const popoverRef = ref<HTMLElement | null>(null)

// メッセージ投稿成功時の処理
const handlePostSuccess = () => {
  console.log('Message posted successfully')
  // モーダルを閉じる
  if (popoverRef.value) {
    popoverRef.value.hidePopover()
  }
}

// メッセージ投稿エラー時の処理
const handlePostError = (error: Error) => {
  setError({
    message: error.message,
    status: 500,
  })
}

// モーダルを開く
const openModal = () => {
  if (popoverRef.value) {
    popoverRef.value.showPopover()
  }
}

// モーダルを閉じる
const closeModal = () => {
  if (popoverRef.value) {
    popoverRef.value.hidePopover()
  }
}

// バックドロップクリックを防ぐ
const handleBackdropClick = (event: Event) => {
  // モーダル内のクリックは無視
  if (event.target === popoverRef.value) {
    closeModal()
  }
}

// Escキーでモーダルを閉じる
const handleKeydown = (event: KeyboardEvent) => {
  if (event.key === 'Escape') {
    closeModal()
  }
}

// コンポーネントマウント時にイベントリスナーを設定
onMounted(() => {
  if (popoverRef.value) {
    // ポップオーバーが開かれた時のイベント
    popoverRef.value.addEventListener('toggle', (event) => {
      const isOpen = (event as any).newState === 'open'
      if (isOpen) {
        // モーダルが開いた時にキーボードイベントリスナーを追加
        document.addEventListener('keydown', handleKeydown)
        // 背景のスクロールを無効化
        document.body.style.overflow = 'hidden'
      } else {
        // モーダルが閉じた時にキーボードイベントリスナーを削除
        document.removeEventListener('keydown', handleKeydown)
        // 背景のスクロールを有効化
        document.body.style.overflow = ''
      }
    })
  }
})
</script>

<template>
  <!-- 投稿ボタン -->
  <button
    :class="$style.postButton"
    @click="openModal"
    aria-label="新しい投稿を作成"
    title="新しい投稿を作成"
  >
    <Icon icon="mdi:plus" :class="$style.buttonIcon" />
  </button>

  <!-- 投稿モーダル -->
  <div
    ref="popoverRef"
    popover="manual"
    :class="$style.postModal"
    role="dialog"
    aria-labelledby="modal-title"
    aria-modal="true"
    @click="handleBackdropClick"
  >
    <div :class="$style.modalHeader" @click.stop>
      <h2 id="modal-title" :class="$style.modalTitle">新しい投稿</h2>
      <button :class="$style.closeButton" @click="closeModal" aria-label="モーダルを閉じる">
        <Icon icon="mdi:close" />
      </button>
    </div>

    <div :class="$style.modalContent" @click.stop>
      <MessageForm @success="handlePostSuccess" @error="handlePostError" />
    </div>
  </div>
</template>

<style lang="scss" module>
.postButton {
  position: fixed;
  bottom: 2rem;
  right: 2rem;
  width: 3.5rem;
  height: 3.5rem;
  background-color: var(--color-primary-600);
  color: var(--color-text-inverse);
  border: none;
  border-radius: 50%;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.3s ease;
  z-index: 1000;

  &:hover {
    background-color: var(--color-primary-700);
    transform: scale(1.05);
    box-shadow: 0 6px 16px rgba(0, 0, 0, 0.2);
  }

  &:active {
    transform: scale(0.95);
  }

  &:focus {
    outline: 2px solid var(--color-primary-300);
    outline-offset: 2px;
  }
}

.buttonIcon {
  font-size: 1.5rem;
}

.postModal {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 90vw;
  max-width: 600px;
  max-height: 80vh;
  background-color: var(--color-surface);
  border: 1px solid var(--color-border-light);
  border-radius: 0.75rem;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.15);
  z-index: 1001;
  padding: 0;
  margin: 0;

  &::backdrop {
    background-color: rgba(0, 0, 0, 0.5);
    backdrop-filter: blur(4px);
    /* バックドロップクリックを無効化 */
    pointer-events: auto;
  }

  /* ポップオーバーが開いている時のスタイル */
  &:popover-open {
    animation: modalFadeIn 0.1s ease-out;
    /* モーダル全体でクリックイベントを受け取る */
    pointer-events: auto;
  }

  /* モーダルコンテンツ内のクリックイベントを保持 */
  & > * {
    pointer-events: auto;
  }

  // アニメーション
  @keyframes modalFadeIn {
    from {
      opacity: 0;
      transform: translate(-50%, -50%) scale(0.95);
    }
    to {
      opacity: 1;
      transform: translate(-50%, -50%) scale(1);
    }
  }
}

.modalHeader {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 1.5rem 1.5rem 0;
  border-bottom: 1px solid var(--color-border-light);
  margin-bottom: 1rem;
}

.modalTitle {
  font-size: 1.25rem;
  font-weight: 600;
  color: var(--color-text-primary);
  margin: 0;
}

.closeButton {
  width: 2rem;
  height: 2rem;
  border: none;
  background: none;
  border-radius: 0.375rem;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--color-text-secondary);
  transition: all 0.2s ease;

  &:hover {
    background-color: var(--color-surface-variant);
    color: var(--color-text-primary);
  }

  &:focus {
    outline: 2px solid var(--color-primary-300);
    outline-offset: 2px;
  }
}

.modalContent {
  padding: 0 1.5rem 1.5rem;
  overflow-y: auto;
  max-height: calc(80vh - 4rem);

  // MessageFormのmargin-bottomを調整
  :global(.messageForm) {
    margin-bottom: 0;
  }
}

// モバイル対応
@media (max-width: 640px) {
  .postButton {
    bottom: 1rem;
    right: 1rem;
    width: 3rem;
    height: 3rem;
  }

  .buttonIcon {
    font-size: 1.25rem;
  }

  .postModal {
    width: 95vw;
    max-height: 90vh;
  }

  .modalHeader {
    padding: 1rem 1rem 0;
  }

  .modalContent {
    padding: 0 1rem 1rem;
    max-height: calc(90vh - 3rem);
  }
}
</style>
