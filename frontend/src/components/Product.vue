<script setup>
import { onMounted, ref } from 'vue'
import Card from 'primevue/card';

import 'primeicons/primeicons.css'
import Button from 'primevue/button'

defineProps(['product'])

</script>

<template>
  <div class="product-card">
    <div class="image-wrapper">
      <img v-if="product.images?.length" :src="'/api/' + product.images[0].path" :alt="product.name"
        class="product-image" />
    </div>
    <p>{{ product.description }}</p>
    <div class="product-card-toolbar">

      <Button severity="danger"><i class="pi pi-shopping-cart"></i></Button>
      <Button severity="danger"><i class="pi pi-eye"></i></Button>
    </div>
  </div>
</template>

<style>
.product-image {
  max-width: 25rem;
}

/* This creates the clipped image area */
.image-wrapper {
  width: 100%;
  height: 200px;
  /* define visible height */
  overflow: hidden;
  /* clip excess */
  padding: 0.75rem;
  /* gap from card edges */
  box-sizing: border-box;
  display: block;
}

.product-card {
  overflow: hidden;
  max-width: 25rem;
  padding: 1rem;
  gap: 1rem;
  border: 2px solid white;
  border-radius: 2rem;
  border-width: 2px;
  display: flex;
  color: black;

  justify-content: center;
  max-height: 200px;
}
</style>
