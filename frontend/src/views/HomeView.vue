<script setup lang="ts">
import { computed } from 'vue';
import { useRouter } from 'vue-router';
import Button from 'primevue/button';
import Card from 'primevue/card';
import { useProductStore } from '@/stores/products';
import { p } from 'vue-router/dist/index-CUL6Z3eo.mjs';

// Example: you could pull this from a Pinia store later
const bandName = 'MOONLESS';

const store = useProductStore()
store.getProducts()

const router = useRouter();

interface FeaturedItem {
  id: string;
  name: string;
  slug: string;
  price: number;
  image: string;
  tagline: string;
}

interface ListenLink {
  id: string;
  label: string;
  url: string;
}

interface ShowInfo {
  city: string;
  date: string;  // e.g., "March 12, 2026"
  venue: string;
  ticketsUrl?: string;
}

// Static mock data for now (later you can fetch / use Pinia)
const featuredItems = computed<FeaturedItem[]>(() => [
  {
    id: 'tee-1',
    name: 'Logo Tee',
    slug: 'logo-tee',
    price: 25,
    image: "@/assets/test_image.jpg",
    tagline: 'Classic logo on ultra-soft cotton.',
  },
  {
    id: 'hoodie-1',
    name: 'Tour Hoodie',
    slug: 'tour-hoodie',
    price: 45,
    image: '@/assets/test_image.jpg',
    tagline: 'Limited tour dates print. Heavyweight.',
  },
  {
    id: 'vinyl-1',
    name: 'Limited Vinyl',
    slug: 'limited-vinyl',
    price: 30,
    image: '@/assets/merch/test_image.jpg',
    tagline: 'Numbered edition, hand-signed inserts.',
  },
]);

const listenLinks = computed<ListenLink[]>(() => [
  { id: 'spotify', label: 'Spotify', url: 'https://open.spotify.com' },
  { id: 'apple', label: 'Apple Music', url: 'https://music.apple.com' },
  { id: 'youtube', label: 'YouTube', url: 'https://youtube.com' },
  { id: 'bandcamp', label: 'Bandcamp', url: 'https://bandcamp.com' },
]);

const nextShow = computed<ShowInfo | null>(() => ({
  city: 'Your City',
  date: 'April 20, 2026',
  venue: 'Your Venue',
  ticketsUrl: 'https://tickets.example.com',
}));

function goToShop() {
  router.push({ name: 'shop' }); // ensure you have a route named 'shop'
}

function goToMusic() {
  router.push({ name: 'music' }); // or change to a real route you have
}

function goToProduct(slug: string) {
  router.push({ name: 'product', params: { slug } }); // ensure a route named 'product'
}

function formatPrice(amount: number): string {
  return `$${amount.toFixed(2)}`;
}

function openExternal(url: string) {
  window.open(url, '_blank', 'noopener,noreferrer');
}
</script>
<template>
  <div class="landing-page">
    <!-- Hero Section -->
    <section class="hero">
      <div class="hero-content">
        <div class="hero-text">
          <p class="hero-kicker">Official Merch</p>
          <h1 class="hero-title">{{ bandName }}</h1>
          <p class="hero-subtitle">
            New drops, limited runs, and everything you need to rep the band.

          </p>
          <button @click="store.getProducts()">Refresh products</button>

          <p>{{ store.products.length }}</p>
          <p>{{ store.loading }}</p>

          <div class="hero-actions">
            <Button label="Shop Merch" class="p-button-rounded p-button-lg hero-primary-btn" @click="goToShop" />
            <Button label="Listen" class="p-button-text p-button-lg hero-secondary-btn" @click="goToMusic" />
          </div>
        </div>

        <div class="hero-image-wrapper">
          <!-- Replace with your own band image / artwork -->
          <img src="@/assets/test_image.jpg" alt="Band promo" class="hero-image" />
        </div>
      </div>
    </section>

    <!-- Featured Merch Section -->
    <section class="section featured-merch">
      <div class="section-header">
        <h2>Featured Merch</h2>
        <Button label="View All" class="p-button-text p-button-sm" @click="goToShop" />
      </div>

      <div class="featured-grid">
        <Card v-for="item in featuredItems" :key="item.id" class="featured-card" @click="goToProduct(item.slug)">
          <template #header>
            <img :src="item.image" :alt="item.name" class="featured-image" />
          </template>
          <template #title>
            {{ item.name }}
          </template>
          <template #subtitle>
            {{ formatPrice(item.price) }}
          </template>
          <template #content>
            <p class="featured-description">
              {{ item.tagline }}
            </p>
          </template>
        </Card>
      </div>
    </section>

    <!-- Listen Section -->
    <section class="section listen-section">
      <div class="section-header">
        <h2>Listen</h2>
      </div>

      <div class="listen-grid">
        <Button v-for="link in listenLinks" :key="link.id" :label="link.label" icon-pos="left"
          class="p-button-rounded listen-button" @click="openExternal(link.url)">
          <span class="listen-button-label">{{ link.label }}</span>
        </Button>
      </div>
    </section>

    <!-- Upcoming Show Section (optional) -->
    <section v-if="nextShow" class="section show-section">
      <div class="show-card">
        <h2>Next Show</h2>
        <p class="show-city">{{ nextShow.city }}</p>
        <p class="show-date">{{ nextShow.date }}</p>
        <p class="show-venue">{{ nextShow.venue }}</p>
        <Button v-if="nextShow.ticketsUrl" label="Tickets" class="p-button-rounded p-button-sm show-tickets-btn"
          @click="openExternal(nextShow.ticketsUrl)" />
      </div>
    </section>
  </div>
</template>

<style scoped>
.landing-page {
  display: flex;
  flex-direction: column;
  gap: 2rem;
  padding: 1.5rem 1.25rem 3rem;
  max-width: 960px;
  margin: 0 auto;
}

/* HERO SECTION */
.hero {
  display: flex;
  flex-direction: column;
  padding: 1rem 0;
}

.hero-content {
  display: flex;
  flex-direction: column;
  gap: 1.75rem;
}

.hero-text {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.hero-kicker {
  text-transform: uppercase;
  letter-spacing: 0.14em;
  font-size: 0.75rem;
  color: #f97316;
  /* orange-ish accent */
}

.hero-title {
  font-size: 2.1rem;
  line-height: 1.15;
  margin: 0;
}

.hero-subtitle {
  margin: 0;
  font-size: 0.95rem;
  color: #6b7280;
}

.hero-actions {
  display: flex;
  flex-wrap: wrap;
  gap: 0.75rem;
  margin-top: 0.75rem;
}

.hero-primary-btn {
  font-weight: 600;
}

.hero-secondary-btn {
  font-weight: 500;
}

/* Hero image */
.hero-image-wrapper {
  width: 100%;
  border-radius: 1rem;
  overflow: hidden;
  background: #111827;
  display: flex;
  justify-content: center;
  align-items: center;
}

.hero-image {
  width: 100%;
  height: auto;
  display: block;
  object-fit: cover;
}

/* Sections */
.section {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.section-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.section-header h2 {
  font-size: 1.35rem;
  margin: 0;
}

/* Featured merch */
.featured-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 1rem;
}

.featured-card {
  cursor: pointer;
}

.featured-image {
  width: 100%;
  height: 180px;
  object-fit: cover;
}

.featured-description {
  margin: 0;
  font-size: 0.9rem;
  color: #6b7280;
}

/* Listen section */
.listen-grid {
  display: flex;
  flex-wrap: wrap;
  gap: 0.75rem;
}

.listen-button {
  flex: 1 1 calc(50% - 0.75rem);
  justify-content: center;
}

.listen-button-label {
  font-weight: 500;
}

/* Next show section */
.show-section {
  margin-top: 0.5rem;
}

.show-card {
  padding: 1.25rem;
  border-radius: 0.75rem;
  background: #111827;
  color: #f9fafb;
}

.show-card h2 {
  margin: 0 0 0.5rem;
}

.show-city {
  font-size: 1.05rem;
  font-weight: 600;
  margin: 0.25rem 0;
}

.show-date,
.show-venue {
  margin: 0.1rem 0;
  font-size: 0.9rem;
}

.show-tickets-btn {
  margin-top: 0.75rem;
}

/* RESPONSIVE */
/* @media (min-width: 768px) { */
/*   .landing-page { */
/*     padding: 2rem 1.5rem 4rem; */
/*   } */
/**/
/*   .hero-content { */
/*     flex-direction: row; */
/*     align-items: center; */
/*   } */
/**/
/*   .hero-text { */
/*     flex: 1; */
/*   } */
/**/
/*   .hero-image-wrapper { */
/*     flex: 1; */
/*     max-width: 420px; */
/*     margin-left: auto; */
/*   } */
/**/
/*   .featured-grid { */
/*     grid-template-columns: repeat(3, minmax(0, 1fr)); */
/*   } */
/* } */
</style>
