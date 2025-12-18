/*
 * @Author: leeguo leeguo2015@163.com
 * @Date: 2023-09-16 23:43:26
 * @LastEditors: leeguo leeguo2015@163.com
 * @LastEditTime: 2023-11-04 14:44:27
 * @FilePath: \inception\web\src\router\index.js
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
import {createRouter, createWebHistory} from 'vue-router'
import BlogDetail from '../components/blog/detailG.vue'
import ListSidebarComponent from "@/views/ListSidebarBlog.vue";
import Login from "@/views/login.vue";

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'home',
            component: ListSidebarComponent,
            children :[
                {
                    path: 'list/:blogType',
                    name: 'list',
                    component: ListSidebarComponent,
                    props: route => ({role: route.params.role, blogType: route.params.blogType}),
                },
            ]
        },
        {
            path: '/blog/detail/:id',
            name: 'detail',
            component: BlogDetail,
            props: route => ({id: route.params.id}),
        },
                {
            path: '/login',
            name: 'login',
            component: Login,

        }
    ]

})

export default router
