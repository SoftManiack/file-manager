import Vue from 'vue'
import VueRouter from 'vue-router'

import Registration from '@/views/Registration.vue'
import Login from '@/views/Login.vue'
import Main from '@/views/main/Main.vue'
import Recent from '@/views/main/Recent.vue'
import Trash from '@/views/main/Trash.vue'

Vue.use(VueRouter);

const routes = [
    {
        path: '/registration',
        name: 'Registration',
        component: Registration,
        meta: { layout: "EmptyLayout", requiresAuth: false }
    },
    {
        path: '/login',
        name: 'Login',
        component: Login,
        meta: { layout: "EmptyLayout", requiresAuth: false }
    },
    {
        path: '/fm/:uid',
        name: 'Main',
        component: Main,
        meta: { layout: "MainLayout", requiresAuth: false }
    },
    {
        path: '/fm/recent',
        name: 'Recent',
        component: Recent,
        meta: { layout: "MainLayout", requiresAuth: false }
    },
    {
        path: '/fm/trash',
        name: 'Trash',
        component: Trash,
        meta: { layout: "MainLayout", requiresAuth: false }
    },
]

const router = new VueRouter({
    mode: 'history',
    routes,
})

export default router;