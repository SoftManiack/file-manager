import { useAuthStore } from '@/entities/Auth/model/stores'

import type { NavigationGuard } from 'vue-router';

export const authGuard: NavigationGuard = async (to, from, next) => {
   
  console.log(to)
  console.log(from)
  const authStore = useAuthStore()

  console.log(authStore.isAuth)

  if(to.fullPath == "/" && authStore.isAuth == false ){
    next({ name: 'login' })
  }else {
    next()
  }
};