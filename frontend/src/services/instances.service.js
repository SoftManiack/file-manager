import axios from "axios";

export const instance = axios.create({
    baseURL: "http://localhost:8003/v1/",
});

instance.interceptors.request.use( config  => {
    const token = localStorage.getItem("token");
    config.headers.Authorization = token ? `Bearer ${token}` : "";
    return config;
});

instance.interceptors.response.use( response => {
    return response;
}, error => {
    if (401 == error.response.status){
        window.location = '/login';
    }
})