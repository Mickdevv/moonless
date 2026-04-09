import type { ContentLinkPlatform } from '../enums/content-link-platform.enum'

export type ContentLink = {
  id: string
  platform: ContentLinkPlatform
  title: string
  description: string
  url: string
  thumbnail_url: string
  published_at: Date
  created_at: Date
  updated_at: Date
  sort_order: number
  is_featured: boolean
}
