import { signUp, signIn } from '@/services/auth.service';

const mutations = {
    AUTH_ERROR(state, error) {
        state.error = error;
    },
    SIGN_SUCCESS(state, data){
        localStorage.setItem('uid', data.uid);
        state.email = data.email;
        localStorage.setItem('token', data.token);
    },
    LOGOUT(){
        localStorage.removeItem('token');
        localStorage.removeItem('uid');
    },
}

const actions = {
    async signUp( { commit }, form) {
        const { data, error } = await signUp(form);
        if(!data){
            commit("AUTH_ERROR", error);
        }
    },
    async signIn( { commit },form ) {
        const {data, error}  = await signIn(form);
        if(data){
            commit("SIGN_SUCCESS", data);
        }
        if(error){
            commit("AUTH_ERROR", error);
        }
    },
    async logout( { commit } ) {
        commit('LOGOUT');
    },    
}

const getters = {
    errorAuth: (state) => state.error,
    uidUser: () => localStorage.getItem('uid') || "",
    emailUser: (state) => state.email,
    isAuthorisation: () => localStorage.getItem('token') || "",
}

const state = () => ({
    error: "",
    email:"",
})

export default {
    mutations,
    getters,
    actions,
    state,
}
