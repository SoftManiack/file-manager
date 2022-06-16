import Vue from 'vue'
import Vuex from 'vuex'
import cloud from "@/store/cloud.store"
import auth from "@/store/auth.store"
import { loadingStatuses } from "@/store/statusesLoadingConst"

Vue.use(Vuex);

export default new Vuex.Store({
    getters: {
        loadingStatus: (state) => state.loadingStatus,
    },
    state: () => ({
        loadingStatus: loadingStatuses.Empty,
    }),
    mutations: {
        UPDATE_STATUS(state, data){
            state.loadingStatus = data;
        },
    },
    modules: {
        auth,
        cloud,
    }
})

