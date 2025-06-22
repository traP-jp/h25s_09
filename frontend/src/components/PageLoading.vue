<template>
  <div v-if="isVisible" class="page-loading-overlay">
    <div class="loading-spinner">
      <div class="spinner"></div>
      <p class="loading-text">ページを読み込み中...</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useLoadingStore } from '@/stores/loading'
import { computed } from 'vue'

const loadingStore = useLoadingStore()

const isVisible = computed(() => loadingStore.isPageLoading)
</script>

<style scoped>
.page-loading-overlay {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: var(--color-background);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 10;
}

.loading-spinner {
  text-align: center;
}

.spinner {
  width: 50px;
  height: 50px;
  border: 4px solid var(--color-border-light);
  border-top: 4px solid var(--color-primary-500);
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin: 0 auto 16px;
}

@keyframes spin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}

.loading-text {
  margin: 0;
  font-size: 16px;
  color: var(--color-text-secondary);
}

/* ダークモード用のスタイルは削除（CSS変数で対応） */
</style>
