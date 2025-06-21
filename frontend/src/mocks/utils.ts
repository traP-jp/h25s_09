/**
 * MSW用のユーティリティ関数とヘルパー
 */

// レスポンスの遅延を追加する関数
export const delay = (ms: number = 1000) => new Promise((resolve) => setTimeout(resolve, ms))

// ランダムな遅延を追加する関数
export const randomDelay = (min: number = 500, max: number = 2000) => {
  const ms = Math.floor(Math.random() * (max - min + 1)) + min
  return delay(ms)
}

// ランダムエラーを発生させる関数（テスト用）
export const maybeError = (probability: number = 0.1) => {
  return Math.random() < probability
}

// ログ出力用の関数
export const mswLog = (type: 'info' | 'warn' | 'error', message: string, data?: unknown) => {
  const prefix = '[MSW]'
  const timestamp = new Date().toISOString()

  switch (type) {
    case 'info':
      console.info(`${prefix} ${timestamp} INFO: ${message}`, data || '')
      break
    case 'warn':
      console.warn(`${prefix} ${timestamp} WARN: ${message}`, data || '')
      break
    case 'error':
      console.error(`${prefix} ${timestamp} ERROR: ${message}`, data || '')
      break
  }
}

// UUID生成関数
export const generateUUID = () => {
  return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function (c) {
    const r = (Math.random() * 16) | 0
    const v = c === 'x' ? r : (r & 0x3) | 0x8
    return v.toString(16)
  })
}
