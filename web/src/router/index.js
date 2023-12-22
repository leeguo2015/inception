/*
 * @Author: leeguo leeguo2015@163.com
 * @Date: 2023-09-16 23:43:26
 * @LastEditors: leeguo leeguo2015@163.com
 * @LastEditTime: 2023-11-04 14:44:27
 * @FilePath: \inception\web\src\router\index.js
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import Publish from '../views/blog/Publish.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path:'/login',
      name: 'login',
      component: LoginView
    },
        {
      path:'/login',
      name: 'login',
      component: LoginView
    },
    {
      path:'/blog_add',
      name: 'blog_add',
      component: Publish
    },    {
      path:'/search',
      name: 'search',
      component: Publish
    }
  ]
})

export default router
