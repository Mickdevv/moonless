<script setup lang="ts">
import Product from '@/components/Product.vue';
import { useProductStore } from '@/stores/products';
import { onMounted } from 'vue';

const productsStore = useProductStore()

onMounted(() => {
  productsStore.getProducts()
})

</script>
<template>
  <div class="products-list-container">
    <div class="products-list">
      <Product v-for="product in productsStore.products" :key="product.id" :product="product" />
    </div>
  </div>
</template>
<style>
.products-list-container {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  justify-content: center;
}

.products-list {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-evenly;
}
</style>
