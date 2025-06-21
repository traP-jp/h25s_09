import { setupWorker } from 'msw/browser'
import { handlers } from './handlers'

// Service Worker を設定
export const worker = setupWorker(...handlers)
