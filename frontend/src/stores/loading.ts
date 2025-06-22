import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useLoadingStore = defineStore('loading', () => {
  const isPageLoading = ref(false)

  const setPageLoading = (loading: boolean) => {
    isPageLoading.value = loading
  }

  return {
    isPageLoading,
    setPageLoading,
  }
})
