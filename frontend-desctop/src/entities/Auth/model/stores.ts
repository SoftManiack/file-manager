import { defineStore } from 'pinia'
import type { ILogin, IUser } from './types'

export const useAuthStore = defineStore('useAuthStore',{
    state: () => {
        return {
            user: {} as IUser,
            error: "" as string,
        }
    },
    actions: {
        login( login: ILogin ){
            console.log(login)
            alert(login.login)
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

