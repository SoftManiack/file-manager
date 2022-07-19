import { getElements, createDirectory, createFile, uploadFile } from '@/services/cloud.service';
import { loadingStatuses } from "@/store/statusesLoadingConst"

const mutations = {
    GET_ELEMENTS(state, payload ) {
        state.lastRootUid = state.rootUid;
        state.rootUid = payload.uid;
        if(payload.data != null){
            state.elements = payload.data.Directory;
            state.elements.push(...payload.data.File);
            console.log(state.elements)
        }else{
            state.elements = [{uid: "", uidUsers: "", rootUid: "", date_create: "", date_update: "", name: "", isFavorite: false, size: 0, type: "File", countElement: 0}]
        }
    },
    SET_ITEMS(state, data){
        state.elements.push(data);
    },
    ELEMENTS_ERROR(state, error) {
        state.error = error;
    },
}

const actions = {
    async getElements( { commit }, uid) {
        commit("UPDATE_STATUS", loadingStatuses.Loading);
        const data = await getElements(uid);
        if(data || data == null){
            commit("UPDATE_STATUS", loadingStatuses.Ready);
            commit("GET_ELEMENTS", { data, uid } );
        }else{
            commit("UPDATE_STATUS", loadingStatuses.Error);
            commit("ELEMENTS_ERROR", "Нет данных");
        }
    },
    async createDirectory( { commit }, form){
        const { data, error } = await createDirectory(form)
        if(data){
            commit("SET_ITEMS", data);
        }
        if(error){
            commit("ELEMENTS_ERROR", error);
        }
    },
    async createFile( { commit }, form){
        const { data, error } = await createFile(form)
        if(data){
            commit("SET_ITEMS", data);
        }
        if(error){
            commit("ELEMENTS_ERROR", error);
        }
    },
    async rename( form ){
        console.log(form)
        //const { data, error } = await rename(form)
    },
    async uploadFile( { commit }, file){
        const { data, error } = await uploadFile(file)
        if(data){
            commit("SET_ITEMS", data);
        }
        if(error){
            commit("ELEMENTS_ERROR", error);
        }
    },
}

const getters = {
    elements: (state) => state.elements,
    errorElement: (state) => state.error,
    rootUid: (state) => state.rootUid,
    lastRootUid: (state) => state.lastRootUid,
}

const state = () => ({
    error: "",
    elements: [{uid: "", uidUsers: "", rootUid: "", date_create: "", date_update: "", name: "", isFavorite: false, size: 0, type: "File", countElement: 0}],
    rootUid: "",
    lastRootUid: "",
})

export default {
    mutations,
    getters,
    actions,
    state,
}



/*
    Вход

    1) получить id пользовтаеля
    2) установить id текущей диреткории оно равно id пользовтаеля
    

    Зашел первый раз
    1) Получить root


*/