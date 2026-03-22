// // api.ts
// import axios from 'axios'
// import { useAuthStore } from '@/stores/auth'
//
// export const api = axios.create({
//   baseURL: '/api',
//   withCredentials: true, // important for cookies
// })
//
// api.interceptors.request.use((config) => {
//   const auth = useAuthStore()
//   if (auth.accessToken) {
//     config.headers.Authorization = `Bearer ${auth.accessToken}`
//   }
//   return config
// })
//
// api.interceptors.response.use(
//   (res) => res,
//   async (error) => {
//     const auth = useAuthStore()
//
//     if (error.response?.status === 401) {
//       await auth.refresh()
//       error.config.headers.Authorization = `Bearer ${auth.accessToken}`
//       return api.request(error.config)
//     }
//
//     return Promise.reject(error)
//   },
// )
