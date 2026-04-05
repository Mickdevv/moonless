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
    { label: 'Content', command: () => router.push({ name: 'content' }) },
    { label: 'Events', command: () => router.push({ name: 'events' }) },
    { label: 'Shop', command: () => router.push({ name: 'shop' }) },
    { label: 'Contact', command: () => router.push({ name: 'contact' }) },

    ...(authStore.accessToken ? [
      { label: 'Add link', command: () => router.push({ name: 'admin-content-links-add' }) },
      { label: 'Links table', command: () => router.push({ name: 'admin-content-links' }) },
      { label: 'Add event', command: () => router.push({ name: 'admin-events-add' }) },
      { label: 'Events table', command: () => router.push({ name: 'admin-events' }) },
      { label: 'Add product', command: () => router.push({ name: 'admin-products-add' }) },
      { label: 'Products table', command: () => router.push({ name: 'admin-products' }) },
      // { label: 'Login', command: () => router.push({ name: 'login' }) },
      // { label: 'Register', command: () => router.push({ name: 'register' }) },
    ] : [])
  ])
});

</script>

<template>
  <div class="layout">
    <Menubar :model="items" class="centered-menubar" />
    <div class="page-content">
      <RouterView />
      <Toast />
    </div>
  </div>
</template>

<style scoped>
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
  /* optional: limit total width */
  margin-top: 2rem;
  /* spacing from top */
  justify-content: center;
  /* center items inside menubar */
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
