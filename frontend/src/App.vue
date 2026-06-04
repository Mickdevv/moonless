<script setup lang="ts">
import Menubar from 'primevue/menubar';
import { computed } from 'vue';
import { useRouter } from 'vue-router';
import Toast from 'primevue/toast';
import { useAuthStore } from './stores/auth';

const router = useRouter();
const authStore = useAuthStore()

const items = computed(() => {
  return ([
    { label: 'Home', command: () => router.push({ name: 'home' }) },
    { label: 'Discography', command: () => router.push({ name: 'discography' }) },
    { label: 'Live Dates', command: () => router.push({ name: 'schedule-items' }) },
    { label: 'Contact', command: () => router.push({ name: 'contact' }) },

    ...(authStore.accessToken ? [
      { label: 'Shop', command: () => router.push({ name: 'shop' }) },
      {
        label: 'Admin', items: [
          { label: 'Links', command: () => router.push({ name: 'admin-content-links' }) },
          { label: 'Events', command: () => router.push({ name: 'admin-schedule-items' }) },
          { label: 'Products', command: () => router.push({ name: 'admin-products' }) },
        ]
      },
      { label: 'Log out', command: () => router.push({ name: 'logout' }) },
    ] : [])
  ])
});
</script>

<template>
  <div class="layout">
    <div class="nav-wrapper">
      <Menubar :model="items" class="site-nav">
        <template #start>
          <img class="nav-logo" :src="'/api/static/images/misc/moonless_logo.webp'" alt="Moonless" />
        </template>
      </Menubar>
    </div>
    <div class="page-content">
      <RouterView />
      <Toast />
    </div>
  </div>
</template>

<style scoped>
.layout {
  display: flex;
  flex-direction: column;
  width: 100%;
}

/* Wrapper gives the sticky bar a stacking context for the mobile dropdown */
.nav-wrapper {
  position: sticky;
  top: 0;
  z-index: 100;
  width: 100%;
}

.nav-logo {
  width: 2.75rem;
  height: 2.75rem;
  border-radius: 50%;
  object-fit: cover;
  flex-shrink: 0;
}

.site-nav {
  width: 100%;
}

.page-content {
  width: 100%;
}

/* ── PrimeVue Menubar: dark theme ── */
:deep(.p-menubar) {
  background: var(--band-black-soft);
  border: none;
  border-bottom: 1px solid var(--color-border);
  border-radius: 0;
  padding: 0.65rem 2rem;
}

:deep(.p-menuitem-link) {
  padding: 0.55rem 0.9rem;
  border-radius: 0 !important;
}

:deep(.p-menuitem-text),
:deep(.p-menuitem-icon) {
  color: var(--band-white-muted) !important;
  font-weight: 600;
  letter-spacing: 0.05em;
  text-transform: uppercase;
  font-size: 0.78rem;
}

:deep(.p-menuitem:hover > .p-menuitem-link .p-menuitem-text),
:deep(.p-menuitem:hover > .p-menuitem-link .p-menuitem-icon),
:deep(.p-menuitem.p-focus > .p-menuitem-link .p-menuitem-text) {
  color: var(--band-yellow) !important;
}

:deep(.p-menuitem:hover > .p-menuitem-link),
:deep(.p-menuitem.p-focus > .p-menuitem-link) {
  background: transparent !important;
}

:deep(.p-menubar-submenu) {
  background: var(--band-black-soft) !important;
  border: 1px solid var(--color-border) !important;
  border-radius: 0 !important;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.6) !important;
  min-width: 160px;
}

:deep(.p-menubar-submenu .p-menuitem-link) {
  padding: 0.65rem 1.25rem;
}

:deep(.p-submenu-icon) {
  color: var(--band-white-muted) !important;
}

/* ── Mobile hamburger button ── */
:deep(.p-menubar-button) {
  background: transparent !important;
  border: 1px solid var(--color-border) !important;
  color: var(--band-white) !important;
  border-radius: 2px !important;
  width: 2.2rem !important;
  height: 2.2rem !important;
}

:deep(.p-menubar-button:hover) {
  background: var(--band-black-card) !important;
  border-color: var(--band-yellow) !important;
}

:deep(.p-menubar-button .p-icon) {
  color: var(--band-white) !important;
  font-size: 1rem;
}

/* ── Mobile open menu ── */
@media (max-width: 960px) {
  :deep(.p-menubar) {
    padding: 0.65rem 1rem;
    flex-wrap: wrap;
    position: relative;
  }

  :deep(.p-menubar-root-list) {
    width: 100%;
    background: var(--band-black-soft) !important;
    border-top: 1px solid var(--color-border) !important;
    border-bottom: 1px solid var(--color-border) !important;
    border-left: none !important;
    border-right: none !important;
    border-radius: 0 !important;
    box-shadow: 0 8px 20px rgba(0, 0, 0, 0.5) !important;
    padding: 0.25rem 0 !important;
  }

  :deep(.p-menubar-root-list > .p-menuitem > .p-menuitem-link) {
    padding: 0.75rem 1.25rem !important;
    border-bottom: 1px solid var(--color-border);
  }

  :deep(.p-menubar-root-list > .p-menuitem:last-child > .p-menuitem-link) {
    border-bottom: none;
  }

  :deep(.p-menuitem-text),
  :deep(.p-menuitem-icon) {
    font-size: 0.82rem;
  }
}
</style>
