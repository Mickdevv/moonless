<script setup lang="ts">
import { onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useContentLinksStore } from '@/stores/content-links'
import InputText from 'primevue/inputtext'
import Checkbox from 'primevue/checkbox'
import Button from 'primevue/button'
import Textarea from 'primevue/textarea'
import ConfirmPopup from 'primevue/confirmpopup';
import DatePicker from 'primevue/datepicker'
import Select from 'primevue/select'

import { useConfirm } from "primevue/useconfirm";
import { useToast } from "primevue/usetoast";
import 'primeicons/primeicons.css'
import { ref } from 'vue'
import { useAuthStore } from '@/stores/auth'
import { ContentLinkPlatform } from '@/types/enums/content-link-platform.enum'


const route = useRoute()
const contentLinkStore = useContentLinksStore()
const authStore = useAuthStore()

const confirm = useConfirm();
const toast = useToast();

const currentContentLinkExists = ref<boolean>()

const file = ref<File>()
function onFileSelect(event: any) {
  file.value = event.target.files[0]
}


const confirmCreate = (event: any) => {
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
      addContentLinkSubmit()
    },
  });
};
const confirmUpdate = (event: any) => {
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
      updateContentLinkSubmit()
    },
  });
};


onMounted(() => {
  authStore.ensureToken()
  const id = route.params.id as string
  if (id) {
    contentLinkStore.getContentLinkById(id)
    currentContentLinkExists.value = true
  } else {
    contentLinkStore.resetCurrentContentLink()
    currentContentLinkExists.value = false
  }
})

function updateContentLinkSubmit() {
  if (contentLinkStore.currentContentLink) {
    contentLinkStore.updateContentLink(contentLinkStore.currentContentLink)
  }
}

function addContentLinkSubmit() {
  if (contentLinkStore.currentContentLink) {
    console.log(contentLinkStore.currentContentLink)
    contentLinkStore.createContentLink(contentLinkStore.currentContentLink, file.value)
  }
}

const contentLinkPlatformOptions = Object.values(ContentLinkPlatform)
  .map((clv) => {
    return {
      label: clv,
      value: clv
    }
  })



</script>

<template>
  <div class="page-container">
    <div v-if="contentLinkStore.loading">
      Loading Content link...
    </div>
    <div v-else-if="!contentLinkStore.loading" class="form-container">
      <h2 v-if="currentContentLinkExists">Edit content link</h2>
      <h2 v-if="!currentContentLinkExists">Add content link</h2>

      <!-- <form @submit.prevent="confirm1" class="flex flex-col gap-4"> -->

      <form @submit.prevent="confirmCreate" class="flex flex-col gap-4">

        <div class="field">
          <label>Title</label>
          <InputText required v-model="contentLinkStore.currentContentLink!.title" class="w-full" />
        </div>
        <div class="field">
          <label>Description</label>
          <Textarea required v-model="contentLinkStore.currentContentLink!.description" class="w-full" />
        </div>

        <div class="field">
          <label>Platform</label>
          <!-- <InputText required v-model="contentLinkStore.currentContentLink!.platform" class="w-full" /> -->
          <Select v-model="contentLinkStore.currentContentLink.platform" :options="contentLinkPlatformOptions"
            optionLabel="label" optionValue="value" placeholder="Select" class="w-full md:w-56" />

        </div>

        <div class="field">
          <label>URL</label>
          <InputText required v-model="contentLinkStore.currentContentLink!.url" class="w-full" />
        </div>

        <div class="field">
          <label>Published at</label>
          <DatePicker required v-model="contentLinkStore.currentContentLink!.published_at" class="w-full" />
        </div>

        <div class="field">
          <label>Image</label>
          <input type="file" @change="onFileSelect" accept="image/*" />
        </div>

        <Button v-if="currentContentLinkExists" type="submit" label="Edit content link" severity="secondary"
          @click="confirmUpdate($event)" />
        <Button v-else type="submit" label="Add content link" severity="secondary" @click="confirmCreate($event)" />
        <ConfirmPopup></ConfirmPopup>
      </form>

    </div>

    <div v-else>An error occurred</div>
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

.contentLink-images-container {
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
