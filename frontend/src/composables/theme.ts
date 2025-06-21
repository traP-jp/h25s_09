import { getThemeMode, setThemeMode, toggleTheme, type ThemeMode } from '@/lib/theme'
import { onMounted, ref } from 'vue'

export const useTheme = () => {
  const isDark = ref<boolean>(false)
  const currentTheme = ref<ThemeMode>('light')

  const setTheme = (mode: ThemeMode) => {
    currentTheme.value = mode
    isDark.value = mode === 'dark'
    setThemeMode(mode)
  }

  const toggle = () => {
    const newMode = toggleTheme()
    currentTheme.value = newMode
    isDark.value = newMode === 'dark'
  }

  // システムテーマの変更を監視
  const watchSystemTheme = () => {
    if (typeof window !== 'undefined') {
      const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)')

      const handleChange = (e: MediaQueryListEvent) => {
        const stored = localStorage.getItem('theme-mode')
        if (!stored) {
          // ユーザーが手動で設定していない場合のみシステム設定に従う
          setTheme(e.matches ? 'dark' : 'light')
        }
      }

      mediaQuery.addEventListener('change', handleChange)

      return () => {
        mediaQuery.removeEventListener('change', handleChange)
      }
    }
  }

  onMounted(() => {
    const initialTheme = getThemeMode()
    setTheme(initialTheme)

    // システムテーマの監視を開始
    const cleanup = watchSystemTheme()

    // コンポーネントがアンマウントされた時のクリーンアップ
    return cleanup
  })

  return {
    isDark,
    currentTheme,
    setTheme,
    toggle,
  }
}
