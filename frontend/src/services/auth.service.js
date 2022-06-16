import { instance } from '@/services/instances.service';

export const signUp = async (form) => {
    try {
        const response = await instance.post(`auth/sign-up`, form );
        return response
    } catch (error) {
        return error.response.data
    }
}

export const signIn = async (form) => {
    try {
        const response = await instance.post(`auth/sign-in`, form );
        return response
    } catch (error) {
        return error.response.data
    }
}

