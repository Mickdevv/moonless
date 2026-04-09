import type { ScheduleItem } from '@/types/types/schedule-item'
import axios from 'axios'
import { defineStore } from 'pinia'
import { useToast } from 'primevue'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from './auth'
import type { JsonError } from '@/types/types/json-error'

export const useScheduleItemStore = defineStore('scheduleItems', () => {
  const router = useRouter()
  const toast = useToast()
  const authStore = useAuthStore()

  const scheduleItems = ref<ScheduleItem[]>([])

  const loading = ref<boolean>(false)
  const error = ref<any>(null)

  const blankScheduleItem: ScheduleItem = {
    id: '',
    title: '',
    type: '',
    start_date: new Date(),
    end_date: new Date(),
    description: '',
    is_featured: false,
    poster_path: '',
    location_city: '',
    location_maps_link: '',
    location_name: '',
    created_at: new Date(),
    updated_at: new Date(),
  }

  const currentScheduleItem = ref<ScheduleItem>(structuredClone(blankScheduleItem))

  const resetCurrentScheduleItem = () =>
    (currentScheduleItem.value = structuredClone(blankScheduleItem))

  const createScheduleItem = async (image: File | null) => {
    console.log('createscheduleItem')
    console.log('Current scheduleItem: ', currentScheduleItem.value)
    loading.value = true
    error.value = null

    try {
      const formData = new FormData()
      formData.append('data', JSON.stringify(currentScheduleItem.value))
      if (image) formData.append('file', image)

      const headers = { Authorization: `Bearer ${authStore.accessToken}` }
      const res = await axios.post('/api/schedule-items', formData, { headers })
      scheduleItems.value.push(res.data)
      router.push({ name: 'admin-schedule-items-edit', params: { id: res.data.id } })
    } catch (err: any) {
      error.value = err.message
      toast.add({
        severity: 'danger',
        summary: 'Error',
        detail: error.value,
        life: 3000,
      })
    } finally {
      loading.value = false
    }
  }

  const getScheduleItemById = async (id: string) => {
    loading.value = true
    error.value = null

    console.log(id)
    try {
      const res = await axios.get(`/api/schedule-items/${id}`)
      currentScheduleItem.value = {
        ...res.data,
        end_date: new Date(res.data.end_date),
        start_date: new Date(res.data.start_date),
        created_at: new Date(res.data.created_at),
        updated_at: new Date(res.data.updated_at),
      }
    } catch (err: any) {
      error.value = err.message
      toast.add({
        severity: 'danger',
        summary: 'Error',
        detail: error.value,
        life: 3000,
      })
    } finally {
      loading.value = false
    }
  }
  const getScheduleItems = async () => {
    loading.value = true
    error.value = null
    try {
      const res = await axios.get('/api/schedule-items')
      scheduleItems.value = res.data.map((e: ScheduleItem) => {
        return {
          ...e,
          end_date: new Date(e.end_date),
          start_date: new Date(e.start_date),
          created_at: new Date(e.created_at),
          updated_at: new Date(e.updated_at),
        }
      })
    } catch (err: any) {
      error.value = err.message
      toast.add({
        severity: 'danger',
        summary: 'Error',
        detail: error.value,
        life: 3000,
      })
    } finally {
      loading.value = false
    }
  }

  const deleteScheduleItem = async (id: string) => {
    loading.value = true
    error.value = null

    try {
      const headers = { Authorization: `Bearer ${authStore.accessToken}` }
      const res = await axios.delete(`/api/schedule-items/${id}`)
      scheduleItems.value = scheduleItems.value.filter((e: ScheduleItem) => e.id != id)
    } catch (err: any) {
      error.value = err.message
      toast.add({
        severity: 'danger',
        summary: 'Error',
        detail: error.value,
        life: 3000,
      })
    } finally {
      loading.value = false
    }
  }

  return {
    getScheduleItemById,
    getScheduleItems,
    deleteScheduleItem,
    createScheduleItem,
    error,
    loading,
    currentScheduleItem,
    scheduleItems,
    resetCurrentScheduleItem,
  }
})
