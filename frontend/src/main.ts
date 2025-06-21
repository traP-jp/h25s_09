import { VueQueryPlugin } from '@tanstack/vue-query'
import { createPinia } from 'pinia'
import { createApp } from 'vue'

import App from '@/App.vue'
import router from '@/router'

// グローバルテーマ変数をインポート
import '@/assets/styles/variables.scss'

// MSWを開発環境で有効にする
async function enableMocking() {
  if (import.meta.env.MODE !== 'development' || import.meta.env.VITE_ENABLE_MSW !== 'true') {
    return
  }

  const { worker } = await import('@/mocks')

  // Service Worker を開始
  return worker.start({
    onUnhandledRequest: 'warn',
  })
}

enableMocking().then(() => {
  const app = createApp(App)

  app.use(createPinia())
  app.use(router)
  app.use(VueQueryPlugin)

  app.mount('#app')
})
