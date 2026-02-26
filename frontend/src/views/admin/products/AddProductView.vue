<script setup lang="ts">
import { onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useProductStore } from '@/stores/products'
import InputText from 'primevue/inputtext'
import InputNumber from 'primevue/inputnumber'
import Checkbox from 'primevue/checkbox'
import Button from 'primevue/button'
import Textarea from 'primevue/textarea'
import ConfirmPopup from 'primevue/confirmpopup';


import { useConfirm } from "primevue/useconfirm";
import { useToast } from "primevue/usetoast";
import 'primeicons/primeicons.css'
import { ref } from 'vue'
import type { CreateProductImageDto } from '@/types/DTOs/CreateProductImage.dto'


const route = useRoute()
const productStore = useProductStore()

const confirm = useConfirm();
const toast = useToast();


const confirm1 = (event: any) => {
  confirm.require({
    target: event.currentTarget,
    message: 'Are you sure you want to proceed?',
    icon: 'pi pi-exclamation-triangle',
    rejectProps: {
      label: 'Cancel',
      severity: 'secondary',
      outlined: true
    },
    acceptProps: {
      label: 'Save'
    },
    accept: () => {
      addProductSubmit()
    },
  });
};

const confirm2 = (event: any) => {
  confirm.require({
    target: event.currentTarget,
    message: 'Do you want to delete this record?',
    icon: 'pi pi-info-circle',
    rejectProps: {
      label: 'Cancel',
      severity: 'secondary',
      outlined: true
    },
    acceptProps: {
      label: 'Delete',
      severity: 'danger'
    },
    accept: () => {
      toast.add({ severity: 'info', summary: 'Confirmed', detail: 'Record deleted', life: 3000 });
    },
    reject: () => {
      toast.add({ severity: 'error', summary: 'Rejected', detail: 'You have rejected', life: 3000 });
    }
  });
};

onMounted(() => {
  productStore.resetCurrentProduct()
})

function addProductSubmit() {
  if (productStore.currentProduct) {
    productStore.createProduct(productStore.currentProduct)
  }
}



</script>

<template>
  <div class="page-container">
    <div v-if="productStore.currentProduct && productStore.currentProduct!.id == ''" class="form-container">
      <h2>Edit Product</h2>

      <!-- <form @submit.prevent="confirm1" class="flex flex-col gap-4"> -->

      <form @submit.prevent="confirm1" class="flex flex-col gap-4">

        <!-- Name -->
        <div class="field">
          <label>Name</label>
          <InputText v-model="productStore.currentProduct!.name" class="w-full" />
        </div>

        <!-- Price -->
        <div class="field">
          <label>Price</label>
          <InputNumber v-model="productStore.currentProduct!.price" mode="currency" currency="EUR" locale="en-UK"
            class="w-full" />
        </div>

        <!-- Stock -->
        <div class="field">
          <label>Stock</label>
          <InputNumber v-model="productStore.currentProduct!.stock" :min="0" class="w-full" />
        </div>

        <!-- Category -->
        <div class="field">
          <label>Category</label>
          <InputText v-model="productStore.currentProduct!.category" class="w-full" />
        </div>

        <!-- Description -->
        <div class="field">
          <label>Description</label>
          <Textarea v-model="productStore.currentProduct!.description" rows="4" class="w-full" />
        </div>

        <!-- Active -->
        <div class="field flex items-center gap-2">
          <label>Active</label>
          <Checkbox v-model="productStore.currentProduct!.active" :binary="true" />
        </div>

        <Button type="submit" label="Add Product" severity="secondary" @click="confirm1($event)" />
        <ConfirmPopup></ConfirmPopup>
      </form>

    </div>

    <div v-else>
      Loading product...
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
  padding-bottom: 2rem;
}
</style>
