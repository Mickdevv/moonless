<script setup lang="ts">
import ProductCard from '@/components/Product.vue';
import { useProductStore } from '@/stores/products';
import { onMounted } from 'vue';

const productsStore = useProductStore()

onMounted(() => {
  productsStore.getProducts()
})
</script>

<template>
  <section class="shop-section">
    <div class="section-header">
      <h1 class="section-title">MERCH</h1>
      <div class="title-accent"></div>
    </div>

    <div v-if="!productsStore.products?.length" class="empty-state">
      <p>No products available yet. Check back soon.</p>
    </div>

    <div v-else class="products-grid">
      <ProductCard
        v-for="product in productsStore.products"
        :key="product.id"
        :product="product"
      />
    </div>
  </section>
</template>

<style scoped>
.shop-section {
  width: 100%;
  padding: 4rem 2rem 6rem;
  background: var(--band-black);
  min-height: 100vh;
}

.products-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(min(280px, 100%), 1fr));
  gap: 1.5rem;
  max-width: 1200px;
  margin: 0 auto;
}

.empty-state {
  text-align: center;
  color: var(--band-white-muted);
  padding: 5rem 2rem;
  font-size: 1.1rem;
}

/* Tablet */
@media (max-width: 768px) {
  .shop-section {
    padding: 3rem 1.5rem 5rem;
  }
  .products-grid {
    gap: 1.25rem;
  }
}

/* Mobile */
@media (max-width: 480px) {
  .shop-section {
    padding: 2.5rem 1rem 4rem;
  }
  .products-grid {
    grid-template-columns: 1fr;
    gap: 1rem;
  }
}
</style>
