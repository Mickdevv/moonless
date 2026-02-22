<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';
import Button from 'primevue/button';
import Card from 'primevue/card';
import Carousel from 'primevue/carousel';
import { Tag } from 'primevue';
import { useProductStore } from '@/stores/products';

const bandName = 'MOONLESS';
const store = useProductStore();
const router = useRouter();

onMounted(() => {
  store.getProducts();
});

// Carousel responsiveness for different screen sizes
const responsiveOptions = ref([
  { breakpoint: '1400px', numVisible: 4, numScroll: 1 },
  { breakpoint: '1199px', numVisible: 3, numScroll: 1 },
  { breakpoint: '767px', numVisible: 2, numScroll: 1 },
  { breakpoint: '575px', numVisible: 1, numScroll: 1 }
]);

// Severity color for stock
const getSeverity = (stock: number) => {
  if (stock < 5) return 'danger';
  if (stock < 15) return 'warn';
  return 'success';
};

// Navigation
function goToShop() { router.push({ name: 'shop' }); }
function goToMusic() { router.push({ name: 'music' }); }
function goToProduct(slug: string) { router.push({ name: 'product', params: { slug } }); }
function openExternal(url: string) { window.open(url, '_blank', 'noopener,noreferrer'); }
function formatPrice(amount: number) { return `$${amount.toFixed(2)}`; }
</script>

<template>
  <div class="landing-page">

    <!-- HERO -->
    <section class="hero">
      <div class="hero-content">
        <div class="hero-text">
          <p class="hero-kicker">Official Merch</p>
          <h1 class="hero-title">{{ bandName }}</h1>
          <p class="hero-subtitle">New drops, limited runs, and everything you need to rep the band.</p>
          <div class="hero-actions">
            <Button label="Shop Merch" class="p-button-rounded p-button-lg hero-primary-btn" @click="goToShop" />
            <Button label="Listen" class="p-button-text p-button-lg hero-secondary-btn" @click="goToMusic" />
          </div>
        </div>
        <div class="hero-image-wrapper">
          <img src="@/assets/test_image.jpg" alt="Band promo" class="hero-image" />
        </div>
      </div>
    </section>

    <!-- PRODUCT CAROUSEL -->
    <section class="section featured-merch">
      <div class="section-header">
        <h2>Featured Merch</h2>
        <Button label="View All" class="p-button-text p-button-sm" @click="goToShop" />
      </div>

      <div class="carousel-container">
        <!-- Loading / Empty State -->
        <div v-if="store.loading" class="p-6 text-center">Loading products...</div>
        <div v-else-if="store.products.length === 0" class="p-6 text-center">No products available.</div>

        <!-- Carousel -->
        <Carousel v-else :value="store.products" :numVisible="4" :numScroll="1" :responsiveOptions="responsiveOptions"
          :circular="true" :autoplayInterval="5000" class="product-carousel">
          <template #item="{ data }">
            <Card class="carousel-card" @click="goToProduct(data.slug)">
              <div class="image-wrapper">
                <img v-if="data.images?.length" :src="'/api/' + data.images[0].Path" :alt="data.name"
                  class="carousel-image" />
                <Tag v-if="data.stock !== undefined" :value="`${data.stock} in stock`"
                  :severity="getSeverity(data.stock)" class="stock-tag" />
              </div>
              <div class="product-info">
                <h3 class="product-name">{{ data.name }}</h3>
                <p class="product-price">{{ formatPrice(data.price) }}</p>
              </div>
            </Card>
          </template>
        </Carousel>
      </div>
    </section>
  </div>
</template>
