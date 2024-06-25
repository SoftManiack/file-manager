import { defineStore } from 'pinia'


export const useAuthStore = defineStore({
    id: "",
    state: () => ({
        name: 'Sigismund',
        password: "",
        token: "",
        size: "",
    }),
    actions: {
        
    },
    getters: {
        isAuth(): boolean {
            return (localStorage.getItem("token") != "")
        }
    }

})

