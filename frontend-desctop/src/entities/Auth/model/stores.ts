import { defineStore } from 'pinia'
import type { ILogin } from './types'

export const useAuthStore = defineStore({
    id: "",
    state: () => ({
        name: 'Sigismund',
        password: "",
        token: "",
        size: "",
    }),
    actions: {
        login( login: ILogin ){
            console.log(login)
            alert(login.login)
        }
    },
    getters: {
        isAuth(): boolean {
            return (localStorage.getItem("token") != "")
        }
    }

})

