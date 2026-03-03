import type { LoginDTO } from '@/types/DTOs/Login.dto'
import type { RegisterDTO } from '@/types/DTOs/Register.dto'
import type { JwtPayload } from '@/types/types/JwtPayload'
import type { User } from '@/types/types/user'
import axios from 'axios'
import { jwtDecode } from 'jwt-decode'
import { defineStore } from 'pinia'
import { useToast } from 'primevue'
import { ref } from 'vue'
import { useRouter } from 'vue-router'

export const useAuthStore = defineStore('auth', () => {
  const toast = useToast()
  const router = useRouter()
  const accessToken = ref<string>()
  const refreshToken = ref<string>()
  const accessTokenPayload = ref<JwtPayload>()
  const loading = ref<boolean>(false)
  const error = ref<string | null>(null)

  const refresh = async () => {
    loading.value = true
    error.value = null
    try {
      const res = await axios.get(`/api/refresh/${refreshToken.value}`)
      accessToken.value = res.data.access_token
      refreshToken.value = res.data.refresh_token
      localStorage.setItem('refresh_token', refreshToken.value!)
      localStorage.setItem('access_token', accessToken.value!)
    } catch (err: any) {
      error.value = err
    } finally {
      loading.value = false
    }
  }

  const login = async (credentials: LoginDTO) => {
    loading.value = true
    error.value = null

    try {
      const res = await axios.post('/api/login', credentials)
      accessToken.value = res.data.access_token
      refreshToken.value = res.data.refresh_token

      if (accessToken.value) {
        accessTokenPayload.value = jwtDecode<JwtPayload>(accessToken.value)
        localStorage.setItem('access_token', accessToken.value)
      }
      if (refreshToken.value) {
        localStorage.setItem('refresh_token', refreshToken.value)
      }

      toast.add({
        severity: 'success',
        summary: 'Success',
        detail: 'Login successful',
        life: 3000,
      })
    } catch (err: any) {
      error.value = err
      toast.add({
        severity: 'danger',
        summary: 'Error',
        detail: err,
        life: 3000,
      })
    } finally {
      loading.value = false
    }
  }

  const register = async (registerDetails: RegisterDTO) => {
    loading.value = false
    error.value = null

    try {
      const res = await axios.post('/api/register', registerDetails)
      router.push('/api/login')
      toast.add({
        severity: 'success',
        summary: 'Success',
        detail: 'Registration successful',
        life: 3000,
      })
    } catch (err: any) {
      error.value = err
      toast.add({
        severity: 'danger',
        summary: 'Error',
        detail: err,
        life: 3000,
      })
    } finally {
      loading.value = false
    }
  }
  const accessTokenFromLocalStorage = localStorage.getItem('access_token')
  const refreshTokenFromLocalStorage = localStorage.getItem('refresh_token')

  if (accessTokenFromLocalStorage) {
    accessToken.value = accessTokenFromLocalStorage
    accessTokenPayload.value = jwtDecode<JwtPayload>(accessToken.value)
  }
  if (refreshTokenFromLocalStorage) {
    refreshToken.value = refreshTokenFromLocalStorage
    if (!accessTokenFromLocalStorage) {
      refresh()
    }
  }

  return {
    register,
    login,
    error,
    loading,
    refreshToken,
    accessToken,
    accessTokenPayload,
  }
})
