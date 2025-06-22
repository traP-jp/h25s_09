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
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(255, 255, 255, 0.9);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 9999;
}

.loading-spinner {
  text-align: center;
}

.spinner {
  width: 50px;
  height: 50px;
  border: 4px solid #f3f3f3;
  border-top: 4px solid #3498db;
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
  color: #666;
}

/* ダークモード対応 */
@media (prefers-color-scheme: dark) {
  .page-loading-overlay {
    background-color: rgba(0, 0, 0, 0.9);
  }

  .spinner {
    border: 4px solid #333;
    border-top: 4px solid #3498db;
  }

  .loading-text {
    color: #ccc;
  }
}
</style>
