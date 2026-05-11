<script setup lang="ts">
import { onMounted } from 'vue'
import { useRoute } from 'vue-router'
import InputText from 'primevue/inputtext'
import DatePicker from 'primevue/datepicker'
import InputNumber from 'primevue/inputnumber'
import Checkbox from 'primevue/checkbox'
import Button from 'primevue/button'
import Textarea from 'primevue/textarea'
import ConfirmPopup from 'primevue/confirmpopup';


import { useConfirm } from "primevue/useconfirm";
import { useToast } from "primevue/usetoast";
import 'primeicons/primeicons.css'
import { ref } from 'vue'
import { useAuthStore } from '@/stores/auth'
import { useScheduleItemStore } from '@/stores/schedule-items'


const route = useRoute()
const authStore = useAuthStore()
const scheduleItemStore = useScheduleItemStore()

const confirm = useConfirm();
const toast = useToast();

const currentScheduleItemExists = ref<boolean>()

const confirmCreateScheduleItem = (event: any) => {
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
      createScheduleItemSubmit()
    },
  });
};

const confirmUpdateScheduleItem = (event: any) => {
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
      updateScheduleItemSubmit()
    },
    reject: () => {
      toast.add({ severity: 'danger', summary: 'Rejected', detail: 'You have rejected', life: 3000 });
    }
  });
};

const confirmDeleteScheduleItemImage = (event: any, id: string) => {
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
      // eventStore.deleteEventimage(eventStore.currentScheduleItem!.id)
    },
  });
};

onMounted(async () => {
  await authStore.ensureToken()
  const id = route.params.id as string
  if (id) {
    scheduleItemStore.getScheduleItemById(id)
    currentScheduleItemExists.value = true
  } else {
    scheduleItemStore.resetCurrentScheduleItem()
    currentScheduleItemExists.value = false
  }
})

function updateScheduleItemSubmit() {
  if (scheduleItemStore.currentScheduleItem) {
    // eventStore.updateEvent(eventStore.currentScheduleItem)
  }
}
const file = ref(null)
function onFileSelect(event: any) {
  file.value = event.target.files[0]
}
function createScheduleItemSubmit() {
  console.log("createEventSubmit")
  scheduleItemStore.createScheduleItem(file.value)
}


</script>

<template>
  <div class="page-container">
    <div class="form-container">
      <h2 v-if="!currentScheduleItemExists">Create event</h2>
      <h2 v-else>Update event</h2>


      <form @submit.prevent="confirmCreateScheduleItem" class="flex flex-col gap-4">

        <div class="field">
          <label>Title</label>
          <InputText required v-model="scheduleItemStore.currentScheduleItem.title" class="w-full" />
        </div>
        <div class="field">
          <label>Description</label>
          <Textarea v-model="scheduleItemStore.currentScheduleItem.description" class="w-full" />
        </div>

        <div class="field">
          <label>Type</label>
          <InputText required v-model="scheduleItemStore.currentScheduleItem.type" class="w-full" />
        </div>

        <div class="field">
          <label>Start date</label>
          <DatePicker showTime required v-model="scheduleItemStore.currentScheduleItem.start_date" class="w-full" />
        </div>
        <div class="field">
          <label>End date</label>
          <DatePicker showTime required v-model="scheduleItemStore.currentScheduleItem.end_date" class="w-full" />
        </div>

        <div class="field">
          <label>Location title</label>
          <InputText required v-model="scheduleItemStore.currentScheduleItem.location_name" class="w-full" />
        </div>
        <div class="field">
          <label>Location city</label>
          <InputText required v-model="scheduleItemStore.currentScheduleItem.location_city" class="w-full" />
        </div>
        <div class="field">
          <label>Location maps link</label>
          <InputText v-model="scheduleItemStore.currentScheduleItem.location_maps_link" class="w-full" />
        </div>
        <div class="field">
          <label>Image</label>
          <input type="file" @change="onFileSelect" accept="image/*" />
        </div>

        <Button v-if="!currentScheduleItemExists" type="submit" label="Add event" severity="secondary"
          @click="confirmCreateScheduleItem($event)" />
        <Button v-else type="submit" label="Add event" severity="secondary"
          @click="confirmUpdateScheduleItem($event)" />
        <ConfirmPopup></ConfirmPopup>
      </form>
    </div>
  </div>

  <div class="event-images-container">
    <div v-if="scheduleItemStore.currentScheduleItem?.poster_path">
      <div class="image-card">
        <div>
          <img style="border-radius: 10px;" :src="`/api/${scheduleItemStore.currentScheduleItem.poster_path}`"
            alt="image">
        </div>
        <div class="image-control-panel">
          <Button severity="danger" @click="confirmDeleteScheduleItemImage($event,
            scheduleItemStore.currentScheduleItem.id)"><i class="pi pi-trash"></i></Button>
          <ConfirmPopup></ConfirmPopup>
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

.event-images-container {
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
