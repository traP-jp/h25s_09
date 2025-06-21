import MessageDetailPage from '@/pages/MessageDetail/MessageDetailPage.vue'
import TimelinePage from '@/pages/Timeline/TimelinePage.vue'
import UserAchievementsPage from '@/pages/UserDetail/pages/Achievements/UserAchievementsPage.vue'
import UserMessagesPage from '@/pages/UserDetail/pages/Messages/UserMessagesPage.vue'
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
      component: TimelinePage,
    },
    {
      path: '/messages/:id',
      name: 'message-detail',
      props: true,
      component: MessageDetailPage,
    },
    {
      path: '/users/:id/messages',
      name: 'user-messages',
      props: true,
      component: UserMessagesPage,
    },
    {
      path: '/users/:id/achievements',
      name: 'user-achievements',
      props: true,
      component: UserAchievementsPage,
    },
  ],
})

export default router
