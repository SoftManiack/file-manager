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
            
            console.log(login)
            authApi.login(login)
                .then((data) => {
                    this.user = data.data
                    console.log(this.user)
                })
                .catch((err) => {
                    this.error = err
                })
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

