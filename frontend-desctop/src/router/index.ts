import { createWebHistory, createRouter } from 'vue-router'

import Login from '../pages/Login.vue'
import Main from '../pages/Main.vue'

import { authGuard } from '@/router/authGuard'

//import Login from '@/pages/Login.vue'

const routes = [
  { path: '/login', component: Login,  name: 'login', meta: { requiresAuth: false, layout: "empty" } },
  { path: '/', component: Main, name: 'main',  meta: { requiresAuth: false, layout: "main" }  }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach(async (to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    try {
      await authGuard(to, from)
      next()
    } catch (error) {
      next('/login')
    }
  } else {  
    next()
  }
})

export default router