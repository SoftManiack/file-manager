import { type AxiosResponse } from 'axios';
import instance from '@/api/instance'

import { ILogin, IUser} from '@/model/Auth'

const responseBody = <T>(response: AxiosResponse<T>) => response.data;

const authApi = {
  login: (data: ILogin) => instance.post<any>('auth/sign-in', JSON.stringify(data)).then(responseBody).catch( (err) => {
      return Promise.reject(err)
    }
  )
}

export default authApi