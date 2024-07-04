import axios, { AxiosError } from 'axios';
import { type AxiosInstance } from 'axios'

const instance: AxiosInstance = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL,
})

instance.interceptors.request.use((config) => {

  console.log('res interseptors')
  if (localStorage.getItem("token")) {
    config.headers.Authorization = `Bearer ${localStorage.getItem("token")}`;
  }
  return config;
});

instance.interceptors.response.use(
    (res) => res,
    (error: AxiosError) => {
      const { data, status } = error.response!;
      switch (status) {
        case 400:
          console.error(data);
        case 401:
          console.error('unauthorised');
  
        case 404:
          console.error('/not-found');
  
        case 500:
          console.error('/server-error');
      }
      return error
    }
  );

  export default instance