<script setup lang="ts">
import Product from '@/components/Product.vue';
import { useProductStore } from '@/stores/products';
import YouTube from 'vue3-youtube';
import { onMounted } from 'vue';
import { useContentLinksStore } from '@/stores/content-links';

const contentLinkStore = useContentLinksStore()

onMounted(() => {
  contentLinkStore.getContentLinks()
})

</script>
<template>
  <h2>Spotify</h2>
  <iframe class="content_link" v-for="spotifyLink in contentLinkStore.contentLinks.filter((cl) =>
    cl.platform.toLowerCase() == 'spotify')" :key="spotifyLink.id" data-testid="embed-iframe" style="border-radius:12px"
    :src="spotifyLink.url.replace('/track', '/embed/track')" width="100%" height="352" frameBorder="0"
    allowfullscreen="false" allow="autoplay; clipboard-write; encrypted-media; fullscreen; picture-in-picture"
    loading="lazy"></iframe>

  <h2>Youtube</h2>
  <YouTube class="content_link" v-for="youtubeVideo in contentLinkStore.contentLinks.filter((cl) =>
    cl.platform.toLowerCase() == 'youtube')" :key="youtubeVideo.id" :src="youtubeVideo.url" />

</template>

<style>
.content_link {
  display: flex;
  margin-bottom: 1rem;
}

h2 {
  margin: 1rem 1rem 0rem 5rem;
}
</style>
