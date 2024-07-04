import { defineStore } from 'pinia'
import type { ILogin, IUser } from './types'
import authApi from '../api/auth'

export const useAuthStore = defineStore('useAuthStore',{
    state: () => {
        return {
            user: {} as IUser,
            error: "" as string,
        }
    },
    actions: {
        async login( login: ILogin ){
            
            authApi.login(login)
                .then(() => console.log('then'))
                .catch(() => console.log(12))

        }
    },
    getters: {
        isAuth(): boolean {
            return (localStorage.getItem("token") != "")
        },
        getAuthUser(state){
            return state.user
        },
        getAuthError(state){
            return state.error
        }
    }

})

