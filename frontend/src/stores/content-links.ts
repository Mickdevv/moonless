import { ref } from 'vue'
import { defineStore } from 'pinia'
import { useRouter } from 'vue-router'
import { useToast } from 'primevue'
import { useAuthStore } from './auth'
import type { ContentLink } from '@/types/types/content-link'
import type { CreateContentLinkDTO } from '@/types/DTOs/CreateContentLink.dto'
import axios from 'axios'

export const useContentLinksStore = defineStore('content-links', () => {
  const router = useRouter()
  const toast = useToast()

  const blankContentLink: ContentLink = {
    id: '',
    platform: '',
    title: '',
    description: '',
    url: '',
    thumbnail_url: '',
    is_featured: false,
    sort_order: 0,
    published_at: new Date(),
    created_at: new Date(),
    updated_at: new Date(),
  }
  const currentContentLink = ref<ContentLink>(structuredClone(blankContentLink))
  const contentLinks = ref<ContentLink[]>([])
  const error = ref<string | null>(null)
  const loading = ref<boolean>(false)

  const authStore = useAuthStore()

  const resetCurrentContentLink = () => {
    currentContentLink.value = structuredClone(blankContentLink)
  }

  const createContentLink = async (contentLink: CreateContentLinkDTO, image?: File) => {
    error.value = null
    loading.value = false

    try {
      const formData = new FormData()
      formData.append('data', JSON.stringify(contentLink))
      if (image) {
        formData.append('file', image)
      }

      const headers = { Authorization: `Bearer ${authStore.accessToken}` }
      const res = await axios.post('/api/content-links', formData, { headers })
      contentLinks.value.push(res.data)
      toast.add({
        severity: 'success',
        summary: 'Confirmed',
        detail: 'Content link created',
        life: 3000,
      })
      router.push({ name: 'admin-content-links-edit', params: { id: res.data.id } })
    } catch (err: any) {
      error.value = err
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

  const getContentLinkById = async (id: string) => {
    loading.value = false
    error.value = null

    try {
      const res = await axios.get(`/api/content-links/${id}`)
      currentContentLink.value = res.data
    } catch (err: any) {
      error.value = err
      toast.add({
        severity: 'danger',
        life: 3000,
        summary: 'error',
        detail: error.value,
      })
    } finally {
      loading.value = false
    }
  }

  const getContentLinks = async () => {
    loading.value = false
    error.value = null

    try {
      const res = await axios.get('/api/content-links')
      contentLinks.value = res.data
    } catch (err: any) {
      error.value = err
      toast.add({
        severity: 'danger',
        life: 3000,
        summary: 'error',
        detail: error.value,
      })
    } finally {
      loading.value = false
    }
  }

  const updateContentLink = async (contentLink: ContentLink) => {
    console.log('TODO: Update content link')
  }

  const deleteContentLink = async (id: string) => {
    loading.value = false
    error.value = null

    try {
      authStore.ensureToken()
      const headers = { Authorization: `Bearer ${authStore.accessToken}` }
      const res = await axios.delete(`/api/content-links/${id}`, { headers })
      contentLinks.value = contentLinks.value.filter((cl) => cl.id != id)
    } catch (err: any) {
      error.value = err
      toast.add({
        severity: 'danger',
        life: 3000,
        summary: 'error',
        detail: error.value,
      })
    } finally {
      loading.value = false
    }
  }
  return {
    resetCurrentContentLink,
    currentContentLink,
    contentLinks,
    createContentLink,
    getContentLinks,
    getContentLinkById,
    updateContentLink,
    deleteContentLink,
    error,
    loading,
  }
})
