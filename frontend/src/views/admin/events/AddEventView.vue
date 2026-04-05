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
import type { CreateEventDTO } from '@/types/DTOs/CreateEvent.dto'
import { useEventStore } from '@/stores/events'
import { useAuthStore } from '@/stores/auth'


const route = useRoute()
const authStore = useAuthStore()
const eventStore = useEventStore()

const confirm = useConfirm();
const toast = useToast();

const currentEventExists = ref<boolean>()

const confirmCreateEvent = (event: any) => {
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
      createEventSubmit()
    },
  });
};

const confirmUpdateEvent = (event: any) => {
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
      updateEventSubmit()
    },
    reject: () => {
      toast.add({ severity: 'danger', summary: 'Rejected', detail: 'You have rejected', life: 3000 });
    }
  });
};

const confirmDeleteEventImage = (event: any, id: string) => {
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
      // eventStore.deleteEventimage(eventStore.currentEvent!.id)
    },
  });
};

onMounted(() => {
  authStore.ensureToken()
  const id = route.params.id as string
  if (id) {
    eventStore.getEventById(id)
    currentEventExists.value = true
  } else {
    eventStore.resetCurrentEvent()
    currentEventExists.value = false
  }
})

function updateEventSubmit() {
  if (eventStore.currentEvent) {
    // eventStore.updateEvent(eventStore.currentEvent)
  }
}
const file = ref(null)
function onFileSelect(event: any) {
  file.value = event.target.files[0]
}
function createEventSubmit() {
  console.log("createEventSubmit")
  eventStore.createEvent(file.value)
}


</script>

<template>
  <div class="page-container">
    <div class="form-container">
      <h2 v-if="!currentEventExists">Create event</h2>
      <h2 v-else>Update event</h2>


      <form @submit.prevent="confirmCreateEvent" class="flex flex-col gap-4">

        <div class="field">
          <label>Title</label>
          <InputText required v-model="eventStore.currentEvent.title" class="w-full" />
        </div>
        <div class="field">
          <label>Description</label>
          <Textarea required v-model="eventStore.currentEvent.description" class="w-full" />
        </div>

        <div class="field">
          <label>Type</label>
          <InputText required v-model="eventStore.currentEvent.type" class="w-full" />
        </div>

        <div class="field">
          <label>Start date</label>
          <DatePicker showTime required v-model="eventStore.currentEvent.start_date" class="w-full" />
        </div>
        <div class="field">
          <label>End date</label>
          <DatePicker showTime required v-model="eventStore.currentEvent.end_date" class="w-full" />
        </div>

        <div class="field">
          <label>Location title</label>
          <Textarea required v-model="eventStore.currentEvent.location_name" class="w-full" />
        </div>
        <div class="field">
          <label>Location city</label>
          <Textarea required v-model="eventStore.currentEvent.location_city" class="w-full" />
        </div>
        <div class="field">
          <label>Location maps link</label>
          <Textarea required v-model="eventStore.currentEvent.location_maps_link" class="w-full" />
        </div>
        <div class="field">
          <label>Image</label>
          <input type="file" @change="onFileSelect" accept="image/*" />
        </div>

        <Button v-if="!currentEventExists" type="submit" label="Add event" severity="secondary"
          @click="confirmCreateEvent($event)" />
        <Button v-else type="submit" label="Add event" severity="secondary" @click="confirmUpdateEvent($event)" />
        <ConfirmPopup></ConfirmPopup>
      </form>
    </div>
  </div>

  <div class="event-images-container">
    <div v-if="eventStore.currentEvent?.poster_path">
      <div class="image-card">
        <div>
          <img style="border-radius: 10px;" :src="`/api/${eventStore.currentEvent.poster_path}`" alt="image">
        </div>
        <div class="image-control-panel">
          <Button severity="danger" @click="confirmDeleteEventImage($event,
            eventStore.currentEvent.id)"><i class="pi pi-trash"></i></Button>
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
