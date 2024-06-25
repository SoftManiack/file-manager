import { createMemoryHistory, createRouter } from 'vue-router'

import { Login } from '@/pages/login'
import { authGuard } from '@/app/router/authGuard'

//import Login from '@/pages/Login.vue'

const routes = [
  { path: '/login', component: Login,  name: 'login', meta: { requiresAuth: false, layout: "empty" } },
  { path: '/', name: 'main',  component: Login, meta: { requiresAuth: true, layout: "main" }}
]

const router = createRouter({
  history: createMemoryHistory(),
  routes,
})

// сделал асинхронным хз )

router.beforeEach( async (to, from, next):Promise<void>  => {
  // ...
  // explicitly return false to cancel the navigation

  if (to.matched.some(record => record.meta.requiresAuth)) {
    authGuard(to, from, next);
  } else {
    next();
  }

})


export default router