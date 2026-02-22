<script setup lang="ts">
import Menubar from 'primevue/menubar';
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();

const items = ref([
  { label: 'Home', command: () => router.push({ name: 'home' }) },
  { label: 'Shop', command: () => router.push({ name: 'shop' }) },
  { label: 'Add product', command: () => router.push({ name: 'admin-products-add' }) },
  { label: 'Products table', command: () => router.push({ name: 'admin-products' }) },
]);
</script>

<template>
  <div class="layout">
    <Menubar :model="items" class="centered-menubar" />
    <div class="page-content">
      <RouterView />
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
