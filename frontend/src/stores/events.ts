import type { Event } from '@/types/types/event'
import axios from 'axios'
import { defineStore } from 'pinia'
import { useToast } from 'primevue'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from './auth'
import type { JsonError } from '@/types/types/json-error'

export const useEventStore = defineStore('events', () => {
  const router = useRouter()
  const toast = useToast()
  const authStore = useAuthStore()

  const events = ref<Event[]>([])

  const loading = ref<boolean>(false)
  const error = ref<any>(null)

  const blankEvent: Event = {
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

  const currentEvent = ref<Event>(structuredClone(blankEvent))

  const resetCurrentEvent = () => (currentEvent.value = structuredClone(blankEvent))

  const createEvent = async (image: File | null) => {
    console.log('createEvent')
    console.log('Current event: ', currentEvent.value)
    loading.value = true
    error.value = null

    try {
      const formData = new FormData()
      formData.append('data', JSON.stringify(currentEvent.value))
      if (image) formData.append('file', image)

      const headers = { Authorization: `Bearer ${authStore.accessToken}` }
      const res = await axios.post('/api/events', formData, { headers })
      events.value.push(res.data)
      router.push({ name: 'admin-events-edit', params: { id: res.data.id } })
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

  const getEventById = async (id: string) => {
    loading.value = true
    error.value = null

    try {
      const res = await axios.get(`/api/events/${id}`)
      currentEvent.value = res.data
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
  const getEvents = async () => {
    loading.value = true
    error.value = null
    try {
      const res = await axios.get('/api/events')
      events.value = res.data
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

  const deleteEvent = async (id: string) => {
    loading.value = true
    error.value = null

    try {
      const headers = { Authorization: `Bearer ${authStore.accessToken}` }
      const res = await axios.delete(`/api/events/${id}`)
      events.value = events.value.filter((e) => e.id != id)
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
    getEventById,
    getEvents,
    deleteEvent,
    createEvent,
    error,
    loading,
    currentEvent,
    events,
    resetCurrentEvent,
  }
})
