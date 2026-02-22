<script setup lang="ts">
import { onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useProductStore } from '@/stores/products'

import InputText from 'primevue/inputtext'
import InputNumber from 'primevue/inputnumber'
import Checkbox from 'primevue/checkbox'
import Button from 'primevue/button'
import Textarea from 'primevue/textarea'

const route = useRoute()
const productStore = useProductStore()

onMounted(() => {
  const id = route.params.id as string
  if (id) {
    productStore.getProductById(id)
  }
})

function onFormSubmit() {
  if (productStore.currentProduct) {
    productStore.updateProduct(productStore.currentProduct)
  }
}
</script>

<template>
  <div v-if="productStore.currentProduct" class="formContainer">
    <h2>Edit Product</h2>

    <form @submit.prevent="onFormSubmit" class="flex flex-col gap-4">

      <!-- Name -->
      <div class="field">
        <label>Name</label>
        <InputText v-model="productStore.currentProduct.name" class="w-full" />
      </div>

      <!-- Price -->
      <div class="field">
        <label>Price</label>
        <InputNumber v-model="productStore.currentProduct.price" mode="currency" currency="USD" locale="en-US"
          class="w-full" />
      </div>

      <!-- Stock -->
      <div class="field">
        <label>Stock</label>
        <InputNumber v-model="productStore.currentProduct.stock" :min="0" class="w-full" />
      </div>

      <!-- Category -->
      <div class="field">
        <label>Category</label>
        <InputText v-model="productStore.currentProduct.category" class="w-full" />
      </div>

      <!-- Description -->
      <div class="field">
        <label>Description</label>
        <Textarea v-model="productStore.currentProduct.description" rows="4" class="w-full" />
      </div>

      <!-- Active -->
      <div class="field flex items-center gap-2">
        <label>Active</label>
        <Checkbox v-model="productStore.currentProduct.active" :binary="true" />
      </div>

      <Button type="submit" label="Update Product" severity="secondary" />

    </form>
  </div>

  <div v-else>
    Loading product...
  </div>
</template>

<style>
.formContainer {
  min-width: 50%;
}

.field {
  justify-content: space-between;
  display: flex;
  margin: 0.5rem
}
</style>
