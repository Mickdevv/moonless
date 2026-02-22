import type { ProductImage } from './product-image'

export type Product = {
  id: string
  name: string
  active: boolean
  description: string
  category: string
  price: number
  stock: number
  images: ProductImage[]
}
