<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';
import Button from 'primevue/button';
import Card from 'primevue/card';
import Carousel from 'primevue/carousel';
import { Tag } from 'primevue';
import { useProductStore } from '@/stores/products';
import { formatPrice } from '@/utils/fomattimg';

const bandName = 'MOONLESS';
const productStore = useProductStore();
const router = useRouter();

onMounted(() => {
  productStore.getProducts();
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

</script>

<template>
  <div class="landing-page">
    <!-- PRODUCT CAROUSEL -->

    <div v-if="productStore.loading" class="p-6 text-center">Loading products...</div>
    <div v-else-if="productStore.products.length === 0" class="p-6 text-center">No products available.</div>
    <Carousel v-else :value="productStore.products" :numVisible="3" :numScroll="3"
      :responsiveOptions="responsiveOptions">
      <template #item="{ data }">
        <div class="carousel-card" @click="goToProduct(data.id)">
          <div class="">
            <div class="">
              <img v-if="data.images?.length" :src="'/api/' + data.images[0].Path" :alt="data.name"
                class="carousel-image" />
              <Tag v-if="data.stock !== undefined" :value="`${data.stock} in stock`" :severity="getSeverity(data.stock)"
                class="stock-tag" />
            </div>
          </div>
          <div class="mb-4 font-medium">{{ data.name }}</div>
          <div class="">
            <p class="product-price">{{ formatPrice(data.price) }}</p>
            <span>
              <Button icon="pi pi-heart" severity="secondary" variant="outlined" />
              <Button icon="pi pi-shopping-cart" class="ml-2" />
            </span>
          </div>
        </div>
      </template>
    </Carousel>
  </div>
</template>

<style scoped>
.carousel-card {
  cursor: pointer;
  display: flex;
  flex-direction: column;
  align-items: center;
  height: 100%;
  /* Ensure the card takes the available carousel item height */
}

.carousel-image {
  width: 100%;
  max-height: 200px;
  border-radius: 15px;
  /* or any reasonable value */
  object-fit: cover;
}
</style>
