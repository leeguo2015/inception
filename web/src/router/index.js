/*
 * @Author: leeguo leeguo2015@163.com
 * @Date: 2023-09-16 23:43:26
 * @LastEditors: leeguo leeguo2015@163.com
 * @LastEditTime: 2023-11-04 14:44:27
 * @FilePath: \inception\web\src\router\index.js
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
import {createRouter, createWebHistory} from 'vue-router'
import HomeView from '../views/home.vue'
import LoginView from '../views/login.vue'
import Publish from '../components/blog/Publish.vue'
import BlogDetail from '../components/blog/detail.vue'
import Hall from '../components/Hall.vue'
import App from '@/App.vue'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        // 之前页面。未完善
        {
            path: '/',
            name: 'home',
            redirect: '/visitor/list',
            // children:    
        },
        {
            path: '/:role',
            name: 'role',
            props: true, // 允许通过路由传递 props
            children: [
                {
                    path: 'list',
                    name: 'list',
                    // component: Hall,
                    children: [
                        {
                            path: ':blogType',
                            name: 'listType',
                            component: App,
                            props: true // 允许通过路由传递 props
                        }
                    ]
                },
                {
                    path: 'blog/detail/:id',
                    name: 'detail',
                    component: BlogDetail,
                    props: true // 允许通过路由传递 props
                },
                
            
            ]
        },

        // {
        //   path: '/login',
        //   name: 'login',
        //   component: LoginView
        // },
        // {
        //   path: '/blog_add',
        //   name: 'blog_add',
        //   component: Publish
        // }, {
        //   path: '/search',
        //   name: 'search',
        //   component: Publish
        // },
        // {
        //     path: '/list',
        //     name: 'list',
        //     component: Hall,
        //     children: [
        //         {
        //             path: ':blogType',
        //             name: 'listType',
        //             component: Hall,
        //             props: true // 允许通过路由传递 props
        //         }
        //     ]
        // },
        // {
        //     path: '/blog/detail/:id',
        //     name: 'search',
        //     component: BlogDetail
        // },
    ]
})

export default router
