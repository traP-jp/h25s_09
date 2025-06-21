/**
 * 時間表示のフォーマット関数
 */
export function formatRelativeTime(dateString: string): string {
  const date = new Date(dateString)
  const now = new Date()
  const diffInMinutes = (now.getTime() - date.getTime()) / (1000 * 60)

  if (diffInMinutes < 1) {
    return '今'
  } else if (diffInMinutes < 60) {
    return `${Math.floor(diffInMinutes)}分前`
  }

  const diffInHours = diffInMinutes / 60
  if (diffInHours < 24) {
    return date.toLocaleTimeString('ja-JP', {
      hour: '2-digit',
      minute: '2-digit',
    })
  } else if (diffInHours < 24 * 7) {
    return date.toLocaleDateString('ja-JP', {
      month: 'short',
      day: 'numeric',
    })
  } else {
    return date.toLocaleDateString('ja-JP', {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
    })
  }
}

/**
 * 完全な日時文字列を取得
 */
export function formatFullDateTime(dateString: string): string {
  const date = new Date(dateString)
  return date.toLocaleString('ja-JP', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit',
  })
}

/**
 * ファイルサイズをフォーマット
 */
export function formatFileSize(bytes: number): string {
  const sizes = ['B', 'KB', 'MB', 'GB']
  if (bytes === 0) return '0 B'

  const i = Math.floor(Math.log(bytes) / Math.log(1024))
  const size = bytes / Math.pow(1024, i)

  return `${size.toFixed(i === 0 ? 0 : 1)} ${sizes[i]}`
}

/**
 * 画像ファイルかどうかをチェック
 */
export function isImageFile(file: File): boolean {
  return file.type.startsWith('image/')
}

/**
 * 画像ファイルのバリデーション
 */
export function validateImageFile(file: File): { isValid: boolean; error?: string } {
  if (!isImageFile(file)) {
    return { isValid: false, error: '画像ファイルを選択してください' }
  }

  const maxSize = 5 * 1024 * 1024 // 5MB
  if (file.size > maxSize) {
    return {
      isValid: false,
      error: `画像サイズが大きすぎます（${formatFileSize(maxSize)}以下にしてください）`,
    }
  }

  return { isValid: true }
}

/**
 * URLのバリデーション
 */
export function isValidUrl(string: string): boolean {
  try {
    new URL(string)
    return true
  } catch {
    return false
  }
}

/**
 * デバウンス関数
 */
export function debounce<T extends (...args: never[]) => unknown>(
  func: T,
  wait: number,
): (...args: Parameters<T>) => void {
  let timeout: number | undefined

  return (...args: Parameters<T>) => {
    clearTimeout(timeout)
    timeout = window.setTimeout(() => func(...args), wait)
  }
}

/**
 * スロットル関数
 */
export function throttle<T extends (...args: never[]) => unknown>(
  func: T,
  limit: number,
): (...args: Parameters<T>) => void {
  let inThrottle: boolean

  return (...args: Parameters<T>) => {
    if (!inThrottle) {
      func(...args)
      inThrottle = true
      setTimeout(() => (inThrottle = false), limit)
    }
  }
}
