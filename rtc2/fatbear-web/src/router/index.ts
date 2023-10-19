import { createRouter, createWebHistory } from 'vue-router'
import LiveRoom from '@/components/LiveRoom.vue'
import Login from '@/components/Login.vue'
import ModelList from '@/components/ModelList.vue'
import ViewRoom from '@/components/ViewRoom.vue'
import Order from '@/components/Order.vue'


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
    },
    {
      path: '/view',
      name: 'view',
      component: ViewRoom
    },
    {
      path: '/order',
      name: 'order',
      component: Order
    },


  ]
})

export default router
