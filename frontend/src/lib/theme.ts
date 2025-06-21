// TypeScript型定義ファイル

export interface ColorPalette {
  50: string
  100: string
  200: string
  300: string
  400: string
  500: string
  600: string
  700: string
  800: string
  900: string
}

export interface Theme {
  colors: {
    primary: ColorPalette
    secondary: ColorPalette
    accent: ColorPalette
    warning: ColorPalette
    error: ColorPalette
    gray: ColorPalette
  }
  semantic: {
    background: string
    surface: string
    surfaceVariant: string
    onBackground: string
    onSurface: string
    onSurfaceVariant: string
  }
  text: {
    primary: string
    secondary: string
    tertiary: string
    inverse: string
  }
  border: {
    light: string
    medium: string
    strong: string
  }
  shadow: {
    light: string
    medium: string
    strong: string
  }
}

export type ThemeMode = 'light' | 'dark'

// カラーオブジェクト
export const colors = {
  primary: {
    50: '#eff6ff',
    100: '#dbeafe',
    200: '#bfdbfe',
    300: '#93c5fd',
    400: '#60a5fa',
    500: '#3b82f6',
    600: '#2563eb',
    700: '#1d4ed8',
    800: '#1e40af',
    900: '#1e3a8a',
  },
  secondary: {
    50: '#eef2ff',
    100: '#e0e7ff',
    200: '#c7d2fe',
    300: '#a5b4fc',
    400: '#818cf8',
    500: '#6366f1',
    600: '#4f46e5',
    700: '#4338ca',
    800: '#3730a3',
    900: '#312e81',
  },
  accent: {
    50: '#f0fdf4',
    100: '#dcfce7',
    200: '#bbf7d0',
    300: '#86efac',
    400: '#4ade80',
    500: '#22c55e',
    600: '#16a34a',
    700: '#15803d',
    800: '#166534',
    900: '#14532d',
  },
  warning: {
    50: '#fff7ed',
    100: '#ffedd5',
    200: '#fed7aa',
    300: '#fdba74',
    400: '#fb923c',
    500: '#f97316',
    600: '#ea580c',
    700: '#c2410c',
    800: '#9a3412',
    900: '#7c2d12',
  },
  error: {
    50: '#fef2f2',
    100: '#fee2e2',
    200: '#fecaca',
    300: '#fca5a5',
    400: '#f87171',
    500: '#ef4444',
    600: '#dc2626',
    700: '#b91c1c',
    800: '#991b1b',
    900: '#7f1d1d',
  },
  gray: {
    50: '#f9fafb',
    100: '#f3f4f6',
    200: '#e5e7eb',
    300: '#d1d5db',
    400: '#9ca3af',
    500: '#6b7280',
    600: '#4b5563',
    700: '#374151',
    800: '#1f2937',
    900: '#111827',
  },
} as const satisfies Record<string, ColorPalette>

// CSS変数へのアクセスヘルパー
export const getCSSVariable = (variableName: string): string => {
  if (typeof window !== 'undefined') {
    return getComputedStyle(document.documentElement).getPropertyValue(variableName).trim()
  }
  return ''
}

// テーマユーティリティ関数
export const setThemeMode = (mode: ThemeMode): void => {
  if (typeof document !== 'undefined') {
    if (mode === 'dark') {
      document.documentElement.setAttribute('data-theme', 'dark')
    } else {
      document.documentElement.removeAttribute('data-theme')
    }
    localStorage.setItem('theme-mode', mode)
  }
}

export const getThemeMode = (): ThemeMode => {
  if (typeof window !== 'undefined') {
    const stored = localStorage.getItem('theme-mode') as ThemeMode | null
    if (stored && (stored === 'light' || stored === 'dark')) {
      return stored
    }
    // システムの設定を確認
    return window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light'
  }
  return 'light'
}

export const toggleTheme = (): ThemeMode => {
  const currentMode = getThemeMode()
  const newMode = currentMode === 'light' ? 'dark' : 'light'
  setThemeMode(newMode)
  return newMode
}
