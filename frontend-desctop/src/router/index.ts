import { createMemoryHistory, createRouter } from 'vue-router'

import Login from '../pages/Login.vue'
import Main from '../pages/Main.vue'

import { authGuard } from '@/router/authGuard'

//import Login from '@/pages/Login.vue'

const routes = [
  { path: '/login', component: Login,  name: 'login', meta: { requiresAuth: false, layout: "empty" } },
  { path: '/', component: Main, name: 'main',  meta: { requiresAuth: false, layout: "empty" }  }
]

const router = createRouter({
  history: createMemoryHistory(),
  routes,
})

// сделал асинхронным хз )

router.beforeEach( async (to, from, next):Promise<void> => {
  // ...
  // explicitly return false to cancel the navigation

  if (to.matched.some(record => record.meta.requiresAuth)) {
    authGuard(to, from, next);
  } else {
    next();
  }

})

export default router