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
  const refreshTokenExpires = ref<Date>()
  const accessTokenPayload = ref<JwtPayload>()
  const loading = ref<boolean>(false)
  const error = ref<string | null>(null)

  const refresh = async () => {
    console.log('REFRESH')
    loading.value = true
    error.value = null
    try {
      const res = await axios.get(`/api/refresh/${refreshToken.value}`)
      accessToken.value = res.data.access_token
      refreshToken.value = res.data.refresh_token
      refreshTokenExpires.value = res.data.refresh_token_expires
      localStorage.setItem('access_token', accessToken.value!)
      localStorage.setItem('refresh_token', refreshToken.value!)
      localStorage.setItem('refresh_token_expires', refreshTokenExpires.value!.toISOString())

      console.warn(res)
      console.log(accessToken.value, refreshToken.value, refreshTokenExpires.value)

      return true
    } catch (err: any) {
      error.value = err
      console.error('Error when refreshing token')
      console.error(err)
      await logout()
      return false
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
      refreshTokenExpires.value = new Date(res.data.refresh_token_expires)

      console.log(res.data)

      if (accessToken.value) {
        accessTokenPayload.value = jwtDecode<JwtPayload>(accessToken.value)
        localStorage.setItem('access_token', accessToken.value)
      }
      if (refreshTokenExpires.value) {
        localStorage.setItem('refresh_token_expires', String(refreshTokenExpires.value))
      }
      if (refreshToken.value) {
        localStorage.setItem('refresh_token', refreshToken.value)
      }
      await router.push(`/`)

      toast.add({
        severity: 'success',
        summary: 'Success',
        detail: 'Login successful',
        life: 3000,
      })
    } catch (err: any) {
      error.value = err
      logout()
      toast.add({
        severity: 'danger',
        summary: 'Error',
        detail: `Please log in again: ${err}`,
        life: 3000,
      })
    } finally {
      loading.value = false
    }
  }

  const register = async (registerDetails: RegisterDTO) => {
    loading.value = true
    error.value = null

    try {
      const res = await axios.post('/api/register', registerDetails)
      router.push('/login')
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

  const logout = async () => {
    console.log('LOGOUT')
    refreshToken.value = undefined
    accessToken.value = undefined
    localStorage.removeItem('access_token')
    localStorage.removeItem('refresh_token')
    localStorage.removeItem('refresh_token_expires')
    await router.push(`/login`)
  }

  const ensureToken = async (attempts = 0, redirect = true): Promise<boolean> => {
    if (!accessToken.value && !accessTokenPayload.value && !refreshToken.value) {
      if (redirect) {
        await logout()
      }
      return false
    }

    const isAccessExpired = accessTokenPayload.value
      ? accessTokenPayload.value.exp * 1000 < Date.now()
      : false
    const isRefreshExpired = isExpiredDate(refreshTokenExpires.value)

    if (isAccessExpired) {
      console.log('Access expired')
      if (refreshToken.value && !isRefreshExpired) {
        if (attempts > 2) {
          if (redirect) {
            await logout()
          }
          console.log('Attempt count exceeded: ' + attempts)
          return false
        }
        if (await refresh()) {
          console.log('REFRESH SUCCESSFUL')
          return true
        }
        // TODO: Could be dangerous to have a recursive call like this
        attempts += 1
        return await ensureToken(attempts)
      } else {
        console.log('Catch-all from refresh ', refreshToken.value, isRefreshExpired)
        if (redirect) {
          await logout()
        }
        return false
      }
    }
    return true
  }

  const accessTokenFromLocalStorage = localStorage.getItem('access_token')
  const refreshTokenFromLocalStorage = localStorage.getItem('refresh_token')
  const refreshTokenExpiresFromLocalStorage = localStorage.getItem('refresh_token_expires')

  if (accessTokenFromLocalStorage) {
    accessToken.value = accessTokenFromLocalStorage
    accessTokenPayload.value = jwtDecode<JwtPayload>(accessToken.value)
  }
  if (refreshTokenFromLocalStorage) {
    refreshToken.value = refreshTokenFromLocalStorage
  }
  if (refreshTokenExpiresFromLocalStorage) {
    refreshTokenExpires.value = new Date(refreshTokenExpiresFromLocalStorage)
  }

  const isExpiredUnix = (time: number | undefined) => {
    return !time || time * 1000 < Date.now()
  }
  const isExpiredDate = (time: Date | undefined) => {
    if (!time) return true

    const date = time instanceof Date ? time : new Date(time)
    return date.getTime() < Date.now()
  }

  if (refreshToken.value && isExpiredDate(refreshTokenExpires.value)) {
    ensureToken()
  }

  return {
    logout,
    register,
    login,
    error,
    loading,
    refreshToken,
    refreshTokenExpires,
    accessToken,
    accessTokenPayload,
    ensureToken,
  }
})
