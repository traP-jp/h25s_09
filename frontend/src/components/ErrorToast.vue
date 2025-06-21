<script lang="ts" setup>
import { useGlobalError } from '@/lib/composables'

const { error, isErrorVisible, clearError } = useGlobalError()
</script>

<template>
  <Teleport to="body">
    <Transition name="error-toast">
      <div v-if="isErrorVisible && error" :class="$style.errorToast" @click="clearError">
        <div :class="$style.errorContent">
          <div :class="$style.errorIcon">⚠️</div>
          <div :class="$style.errorMessage">
            <strong>エラーが発生しました</strong>
            <p>{{ error.message }}</p>
            <small v-if="error.status">Status: {{ error.status }}</small>
          </div>
          <button :class="$style.closeButton" @click="clearError">×</button>
        </div>
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
  background: #f8d7da;
  border: 1px solid #f5c6cb;
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  cursor: pointer;
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
}

.errorMessage {
  flex: 1;
  color: #721c24;

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
    color: #856464;
    font-size: 0.75rem;
  }
}

.closeButton {
  background: none;
  border: none;
  color: #721c24;
  font-size: 1.5rem;
  line-height: 1;
  cursor: pointer;
  padding: 0;
  width: 24px;
  height: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;

  &:hover {
    background: rgba(114, 28, 36, 0.1);
    border-radius: 50%;
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
