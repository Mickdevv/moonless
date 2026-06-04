<script setup lang="ts">
import YouTube from 'vue3-youtube';
import { computed, onMounted } from 'vue';
import { useContentLinksStore } from '@/stores/content-links';

const contentLinkStore = useContentLinksStore()

onMounted(() => {
  contentLinkStore.getContentLinks()
})

const spotifyLinks = computed(() =>
  contentLinkStore.contentLinks.filter(cl => cl.platform.toLowerCase() === 'spotify')
)
const youtubeLinks = computed(() =>
  contentLinkStore.contentLinks.filter(cl => cl.platform.toLowerCase() === 'youtube')
)
</script>

<template>
  <section class="discography-section">
    <div class="section-header">
      <h1 class="section-title">DISCOGRAPHY</h1>
      <div class="title-accent"></div>
    </div>

    <!-- Spotify -->
    <div v-if="spotifyLinks.length" class="platform-block">
      <div class="platform-heading">
        <div class="platform-dot spotify-dot"></div>
        <span>Spotify</span>
      </div>
      <div class="spotify-list">
        <iframe
          v-for="link in spotifyLinks"
          :key="link.id"
          class="spotify-embed"
          :src="link.url.replace('/track', '/embed/track')"
          width="100%"
          height="352"
          frameBorder="0"
          allowfullscreen="false"
          allow="autoplay; clipboard-write; encrypted-media; fullscreen; picture-in-picture"
          loading="lazy"
        ></iframe>
      </div>
    </div>

    <!-- YouTube -->
    <div v-if="youtubeLinks.length" class="platform-block">
      <div class="platform-heading">
        <div class="platform-dot youtube-dot"></div>
        <span>YouTube</span>
      </div>
      <div class="youtube-grid">
        <!--
          vue3-youtube appends "px" to its width/height props, so we must NOT
          pass percentage values. Instead we size it purely via CSS using the
          padding-bottom aspect-ratio trick and :deep() overrides.
        -->
        <div v-for="video in youtubeLinks" :key="video.id" class="youtube-wrapper">
          <YouTube :src="video.url" />
        </div>
      </div>
    </div>

    <div
      v-if="!spotifyLinks.length && !youtubeLinks.length && !contentLinkStore.loading"
      class="empty-state"
    >
      <p>No content available yet. Check back soon.</p>
    </div>
  </section>
</template>

<style scoped>
.discography-section {
  width: 100%;
  padding: 4rem 2rem 6rem;
  background: var(--band-black);
  min-height: 100vh;
}

.platform-block {
  max-width: 900px;
  margin: 0 auto 4rem;
}

.platform-heading {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  font-size: 0.75rem;
  font-weight: 700;
  letter-spacing: 0.25em;
  text-transform: uppercase;
  color: var(--band-white-muted);
  margin-bottom: 1.25rem;
  padding-bottom: 0.65rem;
  border-bottom: 1px solid var(--color-border);
}

.platform-dot {
  width: 9px;
  height: 9px;
  border-radius: 50%;
  flex-shrink: 0;
}

.spotify-dot  { background: #1db954; }
.youtube-dot  { background: #ff0000; }

/* ── Spotify ── */
.spotify-list {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.spotify-embed {
  display: block;
  border: none;
  border-radius: 12px;
  /* Keep the standard 352px height on all viewports — the Spotify player
     needs this height to render correctly. Let it scroll on small screens. */
  width: 100%;
}

/* ── YouTube ── */
.youtube-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(min(320px, 100%), 1fr));
  gap: 1.25rem;
}

/*
  Responsive YouTube embed technique:
  - Set padding-bottom: 56.25% on the wrapper to create a 16:9 box
  - Position the YouTube component (and the iframe it injects) absolutely
    to fill the box
  - :deep() targets both the component's outer div and the iframe itself
*/
.youtube-wrapper {
  position: relative;
  width: 100%;
  padding-bottom: 56.25%; /* 16:9 */
  height: 0;
  overflow: hidden;
  background: var(--band-black-soft);
}

.youtube-wrapper :deep(div),
.youtube-wrapper :deep(iframe) {
  position: absolute !important;
  top: 0 !important;
  left: 0 !important;
  width: 100% !important;
  height: 100% !important;
}

/* ── Empty / loading ── */
.empty-state {
  text-align: center;
  color: var(--band-white-muted);
  padding: 5rem 2rem;
  font-size: 1.1rem;
}

/* ── Responsive ── */
@media (max-width: 768px) {
  .discography-section {
    padding: 3rem 1.5rem 5rem;
  }
  .platform-block {
    margin-bottom: 3rem;
  }
}

@media (max-width: 480px) {
  .discography-section {
    padding: 2.5rem 1rem 4rem;
  }
  .youtube-grid {
    grid-template-columns: 1fr;
  }
  .platform-heading {
    letter-spacing: 0.1em;
  }
}
</style>
