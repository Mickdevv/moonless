import type { ContentLinkPlatform } from '../enums/content-link-platform.enum'

export type CreateContentLinkDTO = {
  id: string
  platform: ContentLinkPlatform
  title: string
  description: string
  url: string
  thumbnail_url: string
  published_at: Date
  created_at: Date
  updated_at: Date
}
