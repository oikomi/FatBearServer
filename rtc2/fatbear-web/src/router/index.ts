import { createRouter, createWebHistory } from 'vue-router'
import LiveRoom from '@/components/LiveRoom.vue'
import Login from '@/components/Login.vue'
import ModelList from '@/components/ModelList.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/:catchAll(.*)',
      redirect: { name: 'live' },
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },

    {
      path: '/live',
      name: 'live',
      component: LiveRoom
    },
    {
      path: '/model',
      name: 'model',
      component: ModelList
    }
  ]
})

export default router
