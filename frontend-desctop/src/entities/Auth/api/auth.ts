import axios, { type AxiosResponse } from 'axios';
export type { ILogin } from '@/entities/Auth/model/index';

const responseBody = <T>(response: AxiosResponse<T>) => response.data;

const request = {
  get: <T>(url: string) => axios.get<T>(url).then(responseBody),
  post: <T>(url: string, body: {}) =>
    axios.post<T>(url, body).then(responseBody),
};

const authApi = {
    login: (login: ILogin) => request.post<any>('s', login)
}

export default authApi