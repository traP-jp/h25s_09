import { useLoadingStore } from '@/stores/loading'
import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: '/timeline',
    },
    {
      path: '/timeline',
      name: 'timeline',
      component: () => import('@/pages/Timeline/TimelinePage.vue'),
    },
    {
      path: '/messages/:id',
      name: 'message-detail',
      props: true,
      component: () => import('@/pages/MessageDetail/MessageDetailPage.vue'),
    },
    {
      path: '/users/:traqId',
      component: () => import('@/pages/UserDetail/UserDetailFrame.vue'),
      children: [
        {
          path: '',
          redirect: (to) => `/users/${to.params.traqId}/messages`,
        },
        {
          path: 'messages',
          name: 'user-messages',
          component: () => import('@/pages/UserDetail/pages/Messages/UserMessagesPage.vue'),
        },
        {
          path: 'achievements',
          name: 'user-achievements',
          component: () => import('@/pages/UserDetail/pages/Achievements/UserAchievementsPage.vue'),
        },
      ],
    },
  ],
})

// 任意のページ遷移に3秒の遅延を追加
router.beforeEach((to, from, next) => {
  return
  // ローディング状態を開始
  const loadingStore = useLoadingStore()
  loadingStore.setPageLoading(true)

  // すぐにページ遷移を実行
  next()
})

// ページ遷移完了後のフック
router.afterEach(() => {
  const loadingStore = useLoadingStore()

  // 3秒間ローディングを表示
  setTimeout(() => {
    loadingStore.setPageLoading(false)
  }, 3000)
})

export default router
