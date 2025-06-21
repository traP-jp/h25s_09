import { useLocalStorage, useToggle } from '@vueuse/core'
import { computed, ref } from 'vue'

import type { ApiError } from '../apis/http-client'
import { defineComposable } from '../store'

// グローバルなローディング状態管理
export const useGlobalLoading = defineComposable('globalLoading', () => {
  const loadingTasks = ref<Set<string>>(new Set())

  const startLoading = (taskId: string) => {
    loadingTasks.value.add(taskId)
  }

  const stopLoading = (taskId: string) => {
    loadingTasks.value.delete(taskId)
  }

  const isLoading = computed(() => loadingTasks.value.size > 0)

  return {
    isLoading,
    startLoading,
    stopLoading,
  }
})

// エラー状態管理
export const useGlobalError = defineComposable('globalError', () => {
  const error = ref<ApiError | null>(null)
  const [isErrorVisible, toggleErrorVisible] = useToggle(false)

  const setError = (newError: ApiError) => {
    error.value = newError
    toggleErrorVisible(true)
  }

  const clearError = () => {
    error.value = null
    toggleErrorVisible(false)
  }

  return {
    error,
    isErrorVisible,
    setError,
    clearError,
    toggleErrorVisible,
  }
})

// ユーザー認証状態管理
export const useAuth = defineComposable('auth', () => {
  const authToken = useLocalStorage<string | null>('auth_token', null)
  const isAuthenticated = computed(() => !!authToken.value)

  const login = (token: string) => {
    authToken.value = token
  }

  const logout = () => {
    authToken.value = null
  }

  return {
    authToken,
    isAuthenticated,
    login,
    logout,
  }
})

// UIの共通状態管理
export const useUI = defineComposable('ui', () => {
  const [isSidebarOpen, toggleSidebar] = useToggle(false)

  const theme = useLocalStorage<'light' | 'dark'>('theme', 'light')

  const toggleTheme = () => {
    theme.value = theme.value === 'light' ? 'dark' : 'light'
  }

  return {
    isSidebarOpen,
    toggleSidebar,
    theme,
    toggleTheme,
  }
})

// フォーム状態管理のヘルパー
export function useFormState<T extends Record<string, unknown>>(initialValues: T) {
  const formData = ref<T>({ ...initialValues })
  const errors = ref<Partial<Record<keyof T, string>>>({})
  const isSubmitting = ref(false)

  const resetForm = () => {
    formData.value = { ...initialValues }
    errors.value = {}
    isSubmitting.value = false
  }

  const setError = (field: keyof T, message: string) => {
    errors.value[field] = message
  }

  const clearError = (field: keyof T) => {
    delete errors.value[field]
  }

  const hasErrors = computed(() => Object.keys(errors.value).length > 0)

  const setSubmitting = (value: boolean) => {
    isSubmitting.value = value
  }

  return {
    formData,
    errors,
    isSubmitting,
    hasErrors,
    resetForm,
    setError,
    clearError,
    setSubmitting,
  }
}

// 無限スクロール対応のページネーション
export function usePagination<T>(
  fetchFn: (page: number, limit: number) => Promise<T[]>,
  limit = 20,
) {
  const items = ref<T[]>([])
  const currentPage = ref(1)
  const isLoading = ref(false)
  const hasMore = ref(true)
  const error = ref<ApiError | null>(null)

  const loadMore = async () => {
    if (isLoading.value || !hasMore.value) return

    isLoading.value = true
    error.value = null

    try {
      const newItems = await fetchFn(currentPage.value, limit)

      if (newItems.length < limit) {
        hasMore.value = false
      }

      ;(items.value as T[]).push(...newItems)
      currentPage.value++
    } catch (err) {
      error.value = err as ApiError
    } finally {
      isLoading.value = false
    }
  }

  const reset = () => {
    items.value = []
    currentPage.value = 1
    hasMore.value = true
    error.value = null
  }

  return {
    items,
    isLoading,
    hasMore,
    error,
    loadMore,
    reset,
  }
}
