<script lang="ts" setup>
import { useGlobalError } from '@/lib/composables'
import { onMounted, ref } from 'vue'

const { error, isErrorVisible, clearError } = useGlobalError()

// 自動消失のタイマー
const autoHideTimer = ref<number | null>(null)

// 自動消失機能（5秒後）
const startAutoHide = () => {
  if (autoHideTimer.value) {
    clearTimeout(autoHideTimer.value)
  }
  autoHideTimer.value = window.setTimeout(() => {
    clearError()
  }, 5000)
}

// エラーが表示されたら自動消失タイマーを開始
const handleErrorShow = () => {
  if (isErrorVisible.value) {
    startAutoHide()
  }
}

// クリックで手動で消す
const handleManualClose = () => {
  if (autoHideTimer.value) {
    clearTimeout(autoHideTimer.value)
  }
  clearError()
}

onMounted(() => {
  // エラーが表示されたら自動消失タイマーを開始
  if (isErrorVisible.value) {
    startAutoHide()
  }
})
</script>

<template>
  <Teleport to="body">
    <Transition name="error-toast" @after-enter="handleErrorShow">
      <div
        v-if="isErrorVisible && error"
        :class="$style.errorToast"
        role="alert"
        aria-live="assertive"
        @click="handleManualClose"
      >
        <div :class="$style.errorContent">
          <div :class="$style.errorIcon" aria-hidden="true">⚠️</div>
          <div :class="$style.errorMessage">
            <strong>エラーが発生しました</strong>
            <p>{{ error.message }}</p>
            <small v-if="error.status">Status: {{ error.status }}</small>
          </div>
          <button
            :class="$style.closeButton"
            @click.stop="handleManualClose"
            aria-label="エラーメッセージを閉じる"
          >
            ✕
          </button>
        </div>
        <div :class="$style.progressBar"></div>
      </div>
    </Transition>
  </Teleport>
</template>

<style lang="scss" module>
.errorToast {
  position: fixed;
  top: 1rem;
  right: 1rem;
  z-index: 9999;
  max-width: 400px;
  min-width: 300px;
  background: var(--color-error-50);
  border: 1px solid var(--color-error-200);
  border-radius: 8px;
  box-shadow: 0 4px 12px var(--color-shadow-strong);
  cursor: pointer;
  overflow: hidden;

  [data-theme='dark'] & {
    background: var(--color-error-900);
    border-color: var(--color-error-800);
  }
}

.errorContent {
  display: flex;
  align-items: flex-start;
  padding: 1rem;
  gap: 0.75rem;
}

.errorIcon {
  font-size: 1.5rem;
  flex-shrink: 0;
  display: flex;
  align-items: center;
  justify-content: center;
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
  }

  p {
    margin: 0;
    font-size: 0.875rem;
    line-height: 1.4;
  }

  small {
    display: block;
    margin-top: 0.25rem;
    color: var(--color-error-600);
    font-size: 0.75rem;

    [data-theme='dark'] & {
      color: var(--color-error-400);
    }
  }
}

.closeButton {
  background: none;
  border: none;
  color: var(--color-error-700);
  font-size: 1.2rem;
  line-height: 1;
  cursor: pointer;
  padding: 0;
  width: 24px;
  height: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  border-radius: 50%;
  transition: background-color 0.2s ease;

  [data-theme='dark'] & {
    color: var(--color-error-200);
  }

  &:hover {
    background: var(--color-error-100);

    [data-theme='dark'] & {
      background: var(--color-error-800);
    }
  }

  &:focus {
    outline: 2px solid var(--color-error-300);
    outline-offset: 2px;
  }
}

.progressBar {
  height: 3px;
  background: var(--color-error-600);
  animation: progress 5s linear forwards;

  [data-theme='dark'] & {
    background: var(--color-error-400);
  }
}

@keyframes progress {
  from {
    width: 100%;
  }
  to {
    width: 0%;
  }
}

// トランジション
.error-toast-enter-active,
.error-toast-leave-active {
  transition: all 0.3s ease;
}

.error-toast-enter-from {
  opacity: 0;
  transform: translateX(100%);
}

.error-toast-leave-to {
  opacity: 0;
  transform: translateX(100%);
}
</style>
