<script setup lang="ts">
import { ref, onMounted } from 'vue';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import ColumnGroup from 'primevue/columngroup';   // optional
import Row from 'primevue/row';
import Badge from 'primevue/badge';

import ConfirmPopup from 'primevue/confirmpopup';


import { useConfirm } from "primevue/useconfirm";
import { useToast } from "primevue/usetoast";
import 'primeicons/primeicons.css'
import { event } from '@primeuix/themes/nora/timeline';
import { useEventStore } from '@/stores/events';

const confirm = useConfirm();

const eventsStore = useEventStore()
onMounted(() => {
  eventsStore.getEvents()
});


const rowClass = (data: any) => {
  return [{ '!bg-primary !text-primary-contrast': data.category === 'Fitness' }];
};
const rowStyle = (data: any) => {
  if (data.quantity === 0) {
    return { fontWeight: 'bold', fontStyle: 'italic' };
  }
};
const stockSeverity = (stock: number) => {
  if (stock == 0) return 'danger';
  else if (stock > 0 && stock < 10) return 'warn';
  else return 'success';
}
const confirmDeleteEvent = (event: any, eventId: string) => {
  confirm.require({
    target: event.currentTarget,
    message: 'Confirm',
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
      eventsStore.deleteEvent(eventId)
    },
  });
};

</script>

<template>

  <div class="card" v-if="eventsStore.events?.length">
    <DataTable :value="eventsStore.events" :rowClass="rowClass" :rowStyle="rowStyle" tableStyle="min-width: 50rem">
      <Column field="name" header="Name"></Column>
      <Column field="category" header="Category"></Column>
      <Column field="active" header="Active">
        <template #body="slotProps">
          <i v-if="slotProps.data.active" class="pi pi-check"></i>
          <i v-else class="pi pi-times"></i>
        </template>
      </Column>
      <Column field="price" header="Price">
        <template #body="slotProps">


        </template>
      </Column>
      <Column field="stock" header="Stock">
        <template #body="slotProps">
          <Badge :value="slotProps.data.stock" :severity="stockSeverity(slotProps.data.stock)" />
        </template>
      </Column>
      <Column field="" header="">
        <template #body="slotProps">
          <router-link :to="`/admin/events/${slotProps.data.id}`" class="editButton
            button"><i class="pi pi-pencil"></i></router-link>
          <button class="deleteButton button" @click="confirmDeleteEvent($event, slotProps.data.id)"><i
              class="pi pi-trash"></i></button>
          <ConfirmPopup></ConfirmPopup>
        </template>
      </Column>
    </DataTable>
  </div>
  <div v-else>
    <h2>No Events</h2>
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
