<script lang="ts" setup>
import MessageForm from '@/components/MessageForm.vue'
import { useGlobalError } from '@/lib/composables'
import { Icon } from '@iconify/vue'
import { useBreakpoints } from '@vueuse/core'
import { computed, onMounted, ref } from 'vue'

// ToggleEvent型の定義
interface ToggleEvent extends Event {
  newState: 'open' | 'closed'
  oldState: 'open' | 'closed'
}

const { setError } = useGlobalError()

// ブレークポイント設定（App.vueと同じ設定）
const breakpoints = useBreakpoints({
  mobile: 0,
  compactSidebar: 900,
  fullSidebar: 1200,
})

// フッターが表示されるかどうかを判定
const showFooter = computed(() => !breakpoints.greaterOrEqual('compactSidebar').value)

const popoverRef = ref<HTMLElement | null>(null)
const messageFormRef = ref<InstanceType<typeof MessageForm> | null>(null)

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
      const toggleEvent = event as ToggleEvent
      const isOpen = toggleEvent.newState === 'open'
      if (isOpen) {
        // モーダルが開いた時にキーボードイベントリスナーを追加
        document.addEventListener('keydown', handleKeydown)
        // 背景のスクロールを無効化
        document.body.style.overflow = 'hidden'
        // 背景クリック用のイベントリスナーを追加
        setTimeout(() => {
          document.addEventListener('click', handleOutsideClick)
        }, 0)
        // MessageFormのテキストエリアにフォーカス
        setTimeout(() => {
          messageFormRef.value?.focusTextarea()
        }, 100)
      } else {
        // モーダルが閉じた時にキーボードイベントリスナーを削除
        document.removeEventListener('keydown', handleKeydown)
        // 背景のスクロールを有効化
        document.body.style.overflow = ''
        // 背景クリック用のイベントリスナーを削除
        document.removeEventListener('click', handleOutsideClick)
      }
    })
  }
})

// 外側クリックでモーダルを閉じる
const handleOutsideClick = (event: Event) => {
  if (popoverRef.value && !popoverRef.value.contains(event.target as Node)) {
    closeModal()
  }
}
</script>

<template>
  <!-- 投稿ボタン -->
  <button
    :class="[$style.postButton, { [$style.postButtonWithFooter]: showFooter }]"
    @click="openModal"
    aria-label="新しい投稿を作成"
    title="新しい投稿を作成"
  >
    <Icon icon="mdi:feather" :class="$style.buttonIcon" />
  </button>

  <!-- 投稿モーダル -->
  <div
    ref="popoverRef"
    popover="manual"
    :class="$style.postModal"
    role="dialog"
    aria-labelledby="modal-title"
    aria-modal="true"
  >
    <div :class="$style.modalHeader">
      <h2 id="modal-title" :class="$style.modalTitle">新しい投稿</h2>
      <button :class="$style.closeButton" @click="closeModal" aria-label="モーダルを閉じる">
        <Icon icon="mdi:close" />
      </button>
    </div>

    <MessageForm
      ref="messageFormRef"
      :class="$style.modalContent"
      @success="handlePostSuccess"
      @error="handlePostError"
    />
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

// フッターがある時のボタン位置調整
.postButtonWithFooter {
  bottom: 5.5rem; // フッターの高さ（約3.5rem）+ マージン（約2rem）
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
    /* バックドロップのクリックを有効化 */
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
  padding: 1.5rem 1.5rem 0.5rem;
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
  margin-bottom: 0;
  border: transparent;
}

// モバイル対応
@media (max-width: 640px) {
  .postButton {
    bottom: 1rem;
    right: 1rem;
    width: 3rem;
    height: 3rem;
  }

  // モバイルでフッターがある時の位置調整
  .postButtonWithFooter {
    bottom: 4.5rem; // フッターの高さ（約3.5rem）+ マージン（約1rem）
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
