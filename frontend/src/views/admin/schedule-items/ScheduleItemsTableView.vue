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
import { formatDate } from '@/utils/fomatting';
import { useScheduleItemStore } from '@/stores/schedule-items';

const confirm = useConfirm();

const scheduleItemsStore = useScheduleItemStore()
onMounted(() => {
  scheduleItemsStore.getScheduleItems()
});


const confirmDeleteScheduleItem = (event: any, scheduleItemId: string) => {
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
      scheduleItemsStore.deleteScheduleItem(scheduleItemId)
    },
  });
};

</script>

<template>

  <router-link :to="`/admin/schedule-items/add`"><i class="pi pi-plus"></i> Add schedule item</router-link>
  <div class="card" v-if="scheduleItemsStore.scheduleItems?.length">
    <DataTable :loading="scheduleItemsStore.loading" :value="scheduleItemsStore.scheduleItems"
      tableStyle="min-width: 50rem">
      <Column field="title" header="Title"></Column>
      <Column field="type" header="Type"></Column>
      <Column field="location_city" header="City"></Column>
      <Column field="location_name" header="Location"></Column>
      <Column field="is_featured" header="Featured">
        <template #body="slotProps">
          <i v-if="slotProps.data.active" class="pi pi-check"></i>
          <i v-else class="pi pi-times"></i>
        </template>
      </Column>
      <Column field="active" header="Active">
        <template #body="slotProps">
          <i v-if="slotProps.data.active" class="pi pi-check"></i>
          <i v-else class="pi pi-times"></i>
        </template>
      </Column>
      <Column field="start_date" header="Start date">
        <template #body="slotProps">
          <p>{{ formatDate(slotProps.data.start_date) }}</p>
        </template>
      </Column>
      <Column field="" header="">
        <template #body="slotProps">
          <router-link :to="`/admin/schedule-items/${slotProps.data.id}`" class="editButton
            button"><i class="pi pi-pencil"></i></router-link>
          <button class="deleteButton button" @click="confirmDeleteScheduleItem($event, slotProps.data.id)"><i
              class="pi pi-trash"></i></button>
          <ConfirmPopup></ConfirmPopup>
        </template>
      </Column>
    </DataTable>
  </div>
  <div v-else>
    <h2>No schedule items</h2>
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
