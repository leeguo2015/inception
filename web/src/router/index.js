import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import publish from '../views/blog/publish.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/about',
      name: 'about',
      component: () => import('../views/AboutView.vue')
    },
    {
      path:'/login',
      name: 'login',
      component: LoginView
    },
    {
      path:'/blog_add',
      name: 'blog_add',
      component: publish
    },    {
      path:'/search',
      name: 'search',
      component: publish
    }
  ]
})

export default router
