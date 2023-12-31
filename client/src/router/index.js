import { createRouter, createWebHistory } from 'vue-router'
import PasswordView from '@/views/PasswordView.vue'
import DashboardView from '@/views/DashboardView.vue'

const routes = [
  {
    path: '/',
    name: 'password',
    component: PasswordView
  },
  {
    path: '/dashboard',
    name: 'dashboard',
    component: DashboardView
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
