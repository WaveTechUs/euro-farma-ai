import { createRouter, createWebHistory } from 'vue-router'

const allowedRoutes = ['/', '/login', '/forgot-password', '/reset-password']

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  scrollBehavior(to, from, savedPosition) {
    return { top: 0 }
  },
  routes: [
    { path: '/', component: () => import('../views/Login.vue')},
    { path: '/criar-conta', component: () => import('../views/CriarConta.vue')},
    { path: '/dashboard', component: () => import('../views/Dashboard.vue')},
    { path: '/chat', component: () => import('../views/Chat.vue')},
    { path: '/informacoes', component: () => import('../views/Infos.vue')},
  ]
})

export default router
