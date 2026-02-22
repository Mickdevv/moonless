<script setup>
import { ref, onMounted } from 'vue';
import { useProductStore } from '@/stores/products';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import ColumnGroup from 'primevue/columngroup';   // optional
import Row from 'primevue/row';
import Badge from 'primevue/badge';

import 'primeicons/primeicons.css'

const productsStore = useProductStore()
onMounted(() => {
  productsStore.getProducts()
});


const rowClass = (data) => {
  return [{ '!bg-primary !text-primary-contrast': data.category === 'Fitness' }];
};
const rowStyle = (data) => {
  if (data.quantity === 0) {
    return { fontWeight: 'bold', fontStyle: 'italic' };
  }
};
const stockSeverity = (stock) => {
  if (stock == 0) return 'danger';
  else if (stock > 0 && stock < 10) return 'warn';
  else return 'success';
}

</script>

<template>
  <div class="card" v-if="productsStore.products?.length">
    <DataTable :value="productsStore.products" :rowClass="rowClass" :rowStyle="rowStyle" tableStyle="min-width: 50rem">
      <Column field="name" header="Name"></Column>
      <Column field="category" header="Category"></Column>
      <Column field="active" header="Active">
        <template #body="slotProps">
          <i v-if="slotProps.data.active" class="pi pi-check"></i>
          <i v-else class="pi pi-times"></i>
        </template>
      </Column>
      <Column field="price" header="Price"></Column>
      <Column field="stock" header="Stock">
        <template #body="slotProps">
          <Badge :value="slotProps.data.stock" :severity="stockSeverity(slotProps.data.stock)" />
          {{ slotProps.data.stock }}
        </template>
      </Column>
      <Column field="" header="">
        <template #body="slotProps">
          <router-link :to="`/admin/products/${slotProps.data.id}`" class="editButton
            button"><i class="pi pi-pencil"></i></router-link>
          <router-link class="deleteButton button"><i class="pi
            pi-trash"></i></router-link>
        </template>
      </Column>
    </DataTable>
  </div>
</template>

<style>
.button {
  margin: 0.2rem;
  border-radius: 5px;
}

.editButton {
  color: green;
}

.deleteButton {
  color: red;
}
</style>
