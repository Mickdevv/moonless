<script setup lang="ts">
import Menubar from 'primevue/menubar';
import { computed, ref } from 'vue';
import { useRouter } from 'vue-router';
import Toast from 'primevue/toast';
import { useToast } from 'primevue/usetoast';
import { useAuthStore } from './stores/auth';

const router = useRouter();
const toast = useToast();
const authStore = useAuthStore()

const items = computed(() => {
  return ([
    { label: 'Home', command: () => router.push({ name: 'home' }) },
    { label: 'Discography', command: () => router.push({ name: 'discography' }) },
    { label: 'Schedule items', command: () => router.push({ name: 'schedule-items' }) },
    { label: 'Contact', command: () => router.push({ name: 'contact' }) },

    ...(authStore.accessToken ? [
      { label: 'Shop', command: () => router.push({ name: 'shop' }) },
      // { label: 'Add link', command: () => router.push({ name: 'admin-content-links-add' }) },
      { label: 'Links', command: () => router.push({ name: 'admin-content-links' }) },
      // { label: 'Add schedule item', command: () => router.push({ name: 'admin-schedule-items-add' }) },
      { label: 'Events', command: () => router.push({ name: 'admin-schedule-items' }) },
      // { label: 'Add product', command: () => router.push({ name: 'admin-products-add' }) },
      { label: 'Products', command: () => router.push({ name: 'admin-products' }) },
      { label: 'Log out', command: () => router.push({ name: 'logout' }) },
      // { label: 'Login', command: () => router.push({ name: 'login' }) },
      // { label: 'Register', command: () => router.push({ name: 'register' }) },
    ] : [])
  ])
});

</script>

<template>
  <div class="layout">
    <Menubar :model="items" class="centered-menubar">
      <template #start><img class="navbar-logo-image" :src="'/api/static/images/misc/moonless_logo.webp'" /></template>
      <template #end></template>
    </Menubar>
    <div class="page-content">
      <RouterView />
      <Toast />
    </div>
  </div>
</template>

<style scoped>
.navbar-logo-image {
  max-width: 4rem;
  border-radius: 2rem;
}

:deep(.p-menubar-end) {
  margin: 0;
  margin-left: 0;
}

/* Layout wrapper ensures vertical stacking */
.layout {
  display: flex;
  flex-direction: column;
  /* stack items vertically */
  align-items: center;
  /* center everything horizontally */
  width: 100%;
  margin: 0 auto;
}

/* Menubar width and centering */
.centered-menubar {
  width: 100%;
  margin-top: 2rem;
  display: flex;
  justify-content: space-between;
}

/* Page content wrapper */
.page-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  /* center content horizontally */
  width: 100%;
  margin-top: 2rem;
  /* spacing below menubar */
}
</style>
