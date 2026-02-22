<script setup lang="ts">
import { onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useProductStore } from '@/stores/products'
import InputText from 'primevue/inputtext'
import InputNumber from 'primevue/inputnumber'
import Checkbox from 'primevue/checkbox'
import Button from 'primevue/button'
import Textarea from 'primevue/textarea'
import ConfirmDialog from 'primevue/confirmdialog'
import { useConfirm } from "primevue/useconfirm";
import { useToast } from "primevue/usetoast";
import 'primeicons/primeicons.css'
import Card from 'primevue/card';


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
// const confirmPosition = (position: string) => {
//   confirm.require({
//     group: 'positioned',
//     message: 'Are you sure you want to proceed?',
//     header: 'Confirmation',
//     icon: 'pi pi-info-circle',
//     position: position,
//     rejectProps: {
//       label: 'Cancel',
//       severity: 'secondary',
//       text: true
//     },
//     acceptProps: {
//       label: 'Confirm',
//       text: true
//     },
//     accept: () => {
//       toast.add({ severity: 'info', summary: 'Confirmed', detail: 'Request submitted', life: 3000 });
//     },
//     reject: () => {
//       toast.add({ severity: 'error', summary: 'Rejected', detail: 'Process incomplete', life: 3000 });
//     }
//   });
// };
function deleteImage() {

}
</script>

<template>
  <div class="page-container">
    <div v-if="productStore.currentProduct" class="form-container">
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
          <InputNumber v-model="productStore.currentProduct.price" mode="currency" currency="EUR" locale="en-UK"
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

    <div v-if="productStore.currentProduct?.images.length" class="product-images-container">
      <h2>Product Images</h2>
      <div class="image-card" v-for="image in productStore.currentProduct?.images">
        <div>
          <img style="border-radius: 10px;" :src="`/api/${image.path}`" :alt="image.id">
        </div>
        <div class="image-control-panel">
          <Button :v-if="!image.is_primary" severity="info">Make primary</Button>
          <Button severity="danger"><i class="pi pi-trash"></i></Button>
          <!-- <Button @click="confirmPosition('top')" severity="danger"><i class="pi pi-trash"></i></Button> -->
        </div>
      </div>
    </div>
  </div>
</template>

<style>
.image-control-panel {
  padding: 0.2rem;
  display: flex;
  justify-content: space-between;
}

.page-container {
  min-width: 50%;
  max-width: 800px;
}

.image-card {
  border: 2px solid white;
  border-radius: 10px;
}

Button {
  margin-top: 0;
}

.product-images-container {
  display: flex;
  gap: 1rem;
  justify-content: space-around;
  flex-wrap: wrap;
}

.form-container {
  min-width: 50%;
}

.field {
  justify-content: space-between;
  display: flex;
  margin: 0.5rem;
}

h2 {
  width: 100%;
  padding: 2rem;
}
</style>
