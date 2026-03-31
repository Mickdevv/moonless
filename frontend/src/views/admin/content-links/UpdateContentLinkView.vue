<script setup lang="ts">
import { onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useContentLinksStore } from '@/stores/content-links'
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
import type { CreateContentLinkDTO } from '@/types/DTOs/CreateContentLink.dto'


const route = useRoute()
const contentLinkStore = useContentLinksStore()

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
      updateContentLinkSubmit()
    },
  });
};
const confirmDeleteContentLinkImage = (event: any, imageId: string) => {
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
      // ContentLinkStore.deleteContentLinkimage(ContentLinkStore.currentContentLink!.id, imageId)
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
      toast.add({ severity: 'danger', summary: 'Rejected', detail: 'You have rejected', life: 3000 });
    }
  });
};

onMounted(() => {
  const id = route.params.id as string
  if (id) {
    contentLinkStore.getContentLinkById(id)
  }
})

function updateContentLinkSubmit() {
  if (contentLinkStore.currentContentLink) {
    contentLinkStore.updateContentLink(contentLinkStore.currentContentLink)
  }
}
const file = ref(null)
function onFileSelect(event: any) {
  file.value = event.target.files[0]
}
const contentLinkImage = ref<CreateContentLinkDTO>()
function ContentLinkImageSubmit() {
  if (file.value && contentLinkStore.currentContentLink) {
    // ContentLinkStore.createContentLinkImage(file.value, ContentLinkStore.currentContentLink?.id)
  }
}


</script>

<template>
  <div class="page-container">
    <div v-if="contentLinkStore.currentContentLink && contentLinkStore.currentContentLink!.id" class="form-container">
      <h2>Edit content link</h2>


      <form @submit.prevent="confirm1" class="flex flex-col gap-4">

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
          <InputText required v-model="contentLinkStore.currentContentLink!.platform" class="w-full" />
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

        <Button type="submit" label="Add contentLink" severity="secondary" @click="confirm1($event)" />
        <ConfirmPopup></ConfirmPopup>
      </form>

    </div>

    <div v-else>
      Loading Content link...
    </div>

  </div>

  <div class="ContentLink-images-container">
    <h2>ContentLink Images</h2>
    <div>
      <input type="file" @change="onFileSelect" accept="image/*" />
      <Button label="Upload" @click="ContentLinkImageSubmit()" />
    </div>
    <div v-if="contentLinkStore.currentContentLink?.thumbnail_url">
      <div class="image-card">
        <div>
          <img style="border-radius: 10px;" :src="`/api/${contentLinkStore.currentContentLink.thumbnail_url}`"
            alt="image">
        </div>
        <div class="image-control-panel">
          <Button severity="danger" @click="confirmDeleteContentLinkImage($event,
            contentLinkStore.currentContentLink.id)"><i class="pi pi-trash"></i></Button>
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
