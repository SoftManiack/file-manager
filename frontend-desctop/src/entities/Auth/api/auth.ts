import { type AxiosResponse } from 'axios';
import instance from '@/shared/api/instance'

export type { ILogin } from '@/entities/Auth/model/index';

const responseBody = <T>(response: AxiosResponse<T>) => response.data;

const authApi = {
  login: (data: ILogin) => instance.post<any>('auth/sign-in', data).then(responseBody).catch( (err) => err )
}

export default authApi