<script setup lang="ts">
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
import { onMounted } from 'vue';
import DataView from 'primevue/dataview';
import { formatDate, formatDateTime } from '@/utils/fomatting';
import { useScheduleItemStore } from '@/stores/schedule-items';


const scheduleItemsStore = useScheduleItemStore()

onMounted(() => {
  scheduleItemsStore.getScheduleItems()
})

</script>
<template>

  <div v-if="scheduleItemsStore.loading">Loading events...</div>
  <div v-else>
    <div class="card" v-if="scheduleItemsStore.scheduleItems?.length">
      <DataTable :loading="scheduleItemsStore.loading" :value="scheduleItemsStore.scheduleItems"
        tableStyle="min-width: 50rem">
        <Column field="poster_path">
          <template #body="slotProps">
            <img class="poster" :src="`/api/${slotProps.data.poster_path}`" alt="">
          </template>
        </Column>
        <Column field="title" header="Title"></Column>
        <Column field="location_city" header="City"></Column>
        <Column field="start_date" header="Start date">
          <template #body="slotProps">
            <p>{{ formatDateTime(slotProps.data.start_date) }}</p>

          </template>
        </Column>
        <Column field="location_name" header="Location">
          <template #body="slotProps">
            <a v-if="slotProps.data.location_maps_link" target="blank" :href="slotProps.data.location_maps_link"><i
                class="pi pi-map-marker"></i> {{
                  slotProps.data.location_name }}</a>
            <span v-else="slotProps.data.location_maps_link">{{ slotProps.data.location_name }}</span>
          </template>
        </Column>
      </DataTable>
    </div>
    <div v-else>
      <h2>No Events</h2>
    </div>
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

.poster {
  max-width: 10rem;
}
</style>
