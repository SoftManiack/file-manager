import Vue from 'vue'
import VueRouter from 'vue-router'
import store from '@/store/index'

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
        path: '/v1/fm/:uid',
        name: 'Main',
        component: Main,
        meta: { layout: "MainLayout", requiresAuth: true, keepAlive: true }
    },
    {
        path: '/v1/fm/recent',
        name: 'Recent',
        component: Recent,
        meta: { layout: "MainLayout", requiresAuth: true, keepAlive: true }
    },
    {
        path: '/v1/fm/trash',
        name: 'Trash',
        component: Trash,
        meta: { layout: "MainLayout", requiresAuth: true, keepAlive: true }
    },
]

const router = new VueRouter({
    mode: 'history',
    routes,
})

router.beforeEach((to, from, next) => {
    if(to.matched.some( record => record.name == "Main")){
        store.dispatch('getElements', to.path.split('fm/')[1]).then(next())
    }
    if(to.matched.some(record => record.meta.requiresAuth)) {
        if(localStorage.getItem('token')) {
            next()
            return
        }
        next('/Login') 
    }else {
        next() 
    }
})

export default router;