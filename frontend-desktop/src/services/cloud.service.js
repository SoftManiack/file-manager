import { instance } from '@/services/instances.service';

export const getElements = async ( uid ) => {
    try {
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

export const uploadFile = async ( file ) => {
    try {
        const response = await instance.post(`/fm/items/file`, file);
        return response
    } catch (error) {
        return error.response.data
    }
}
