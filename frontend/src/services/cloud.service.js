import { instance } from '@/services/instances.service';

export const getElements = async ( uid ) => {
    try {
        alert("getElements ")
        const response = await instance.get(`/fm/items/${uid}`);
        return response.data;
    } catch (error) {
        return error
    }
}

export const createDirectory = async ( form ) => {
    try {
        const response = await instance.post(`/fm/items/directory`, form);
        return response
    } catch (error) {
        return error.response.data
    }
}

export const createFile = async ( form ) => {
    try {
        const response = await instance.post(`/fm/items/file`, form);
        return response
    } catch (error) {
        return error.response.data
    }
}

