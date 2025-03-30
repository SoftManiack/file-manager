
import { defineStore } from 'pinia'
import { ILogin, IUser } from '@/model/Auth'
import authApi from "@/api/auth"

export const useAuthStore = defineStore('user', {
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
                    this.error = null
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