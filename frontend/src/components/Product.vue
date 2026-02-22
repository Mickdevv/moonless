<script setup>
import { onMounted, ref } from 'vue'
import Card from 'primevue/card';

defineProps(['product'])

</script>

<template>
  <Card class="product-card">
    <template #header>
      <div class="image-wrapper">
        <img v-if="product.images?.length" :src="'/api/' + product.images[0].Path" :alt="product.name"
          class="product-image" />
      </div>
    </template>
    <template #title>{{ product.name }}</template>
    <template #content>
      <p>{{ product.description }}</p>
    </template>
  </Card>
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
}

.product-card {
  overflow: hidden;
  max-width: 25rem;
  padding: 1rem;
  gap: 1rem;
  border: white;
  border-radius: 2rem;
  border-width: 2px;
  display: flex;
  background-color: white;
  color: black;
  justify-content: center;
  max-height: 200px;
}
</style>
