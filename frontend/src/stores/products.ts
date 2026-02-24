import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import type { Product } from '@/types/types/product'
import axios from 'axios'
import type { ProductImage } from '@/types/types/product-image'
import { useToast } from 'primevue/usetoast'

export const useProductStore = defineStore('products', () => {
  const toast = useToast()
  const products = ref<Product[]>([])
  const currentProduct = ref<Product>()
  const error = ref<string | null>('')
  const loading = ref<boolean>(false)

  const createProductImage = async (image: File, productId: string, productImage: ProductImage) => {
    loading.value = true
    error.value = null
    try {
      const res = await axios.post('/api/product-images', productImage)
      const product = products.value.find((p) => p.id == productId)
      if (product) {
        product.images.push(res.data)
      }
    } catch (err: any) {
      error.value = err.message || 'Failed to create product image'
    } finally {
      loading.value = false
    }
  }

  const deleteProductimage = async (productId: string, productImageId: string) => {
    loading.value = true
    error.value = null
    try {
      const res = await axios.delete(`/api/product-images/${productImageId}`)
      const product = products.value.find((p) => p.id == productId)
      if (product) {
        product.images = product.images.filter((i) => i.id != productImageId)
      }
    } catch (err: any) {
      error.value = err.message || 'Failed to delete product image'
    } finally {
      loading.value = false
    }
  }

  const getProducts = async () => {
    loading.value = true
    error.value = null
    try {
      const res = await axios.get('/api/products')
      products.value = res.data.products
    } catch (err: any) {
      error.value = err.message || 'Failed to fetch products'
    } finally {
      loading.value = false
    }
  }

  const getProductById = async (id: string) => {
    try {
      loading.value = true
      const res = await axios.get(`/api/products/${id}`)
      currentProduct.value = res.data.data
    } catch (err: any) {
      error.value = err.message || 'Failed to fetch product'
    } finally {
      loading.value = false
    }
  }

  const createProduct = async (product: Product) => {
    try {
      loading.value = true
      const res = await axios.post('/api/products', product)
      products.value.push(res.data.product_images)
    } catch (err: any) {
      error.value = err.message
    } finally {
      loading.value = false
    }
  }
  const removeProduct = async (id: string) => {
    try {
      loading.value = true
      const res = await axios.delete(`/api/products/${id}`)
      products.value = products.value.filter((p) => p.id != res.data.id)
    } catch (err: any) {
      error.value = err.message
    } finally {
      loading.value = false
    }
  }

  const updateProduct = async (product: Product) => {
    try {
      loading.value = true
      const res = await axios.put(`/api/products/${product.id}`, product)

      products.value = products.value.filter((p) => p.id != res.data.id)
      products.value.push(res.data)
      toast.add({
        severity: 'success',
        summary: 'Confirmed',
        detail: 'You have accepted21',
        life: 3000,
      })
    } catch (err: any) {
      error.value = err.message
    } finally {
      loading.value = false
    }
  }
  return {
    error,
    loading,
    products,
    createProduct,
    getProductById,
    currentProduct,
    removeProduct,
    updateProduct,
    getProducts,
    deleteProductimage,
    createProductImage,
  }
})
