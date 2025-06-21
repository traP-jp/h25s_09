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
      path: '/users/:id/messages',
      name: 'user-messages',
      props: true,
      component: () => import('@/pages/UserDetail/pages/Messages/UserMessagesPage.vue'),
    },
    {
      path: '/users/:id/achievements',
      name: 'user-achievements',
      props: true,
      component: () => import('@/pages/UserDetail/pages/Achievements/UserAchievementsPage.vue'),
    },
  ],
})

export default router
