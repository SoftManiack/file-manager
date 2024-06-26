import axios, { AxiosError } from 'axios';

axios.defaults.baseURL = import.meta.env.BASE_URL;

axios.interceptors.request.use((config) => {
  if (localStorage.getItem("token")) {
    config.headers.Authorization = `Bearer ${localStorage.getItem("token")}`;
  }
  return config;
});

axios.interceptors.response.use(
    (res) => res,
    (error: AxiosError) => {
      const { data, status, config } = error.response!;
      switch (status) {
        case 400:
          console.error(data);
          break;
  
        case 401:
          console.error('unauthorised');
          break;
  
        case 404:
          console.error('/not-found');
          break;
  
        case 500:
          console.error('/server-error');
          break;
      }
      return Promise.reject(error);
    }
  );