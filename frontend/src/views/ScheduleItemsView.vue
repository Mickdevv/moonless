<script setup lang="ts">
import 'primeicons/primeicons.css'
import { onMounted } from 'vue';
import { formatDate } from '@/utils/fomatting';
import { useScheduleItemStore } from '@/stores/schedule-items';

const scheduleItemsStore = useScheduleItemStore()

onMounted(() => {
  scheduleItemsStore.getScheduleItems()
})
</script>

<template>
  <section class="events-section">
    <div class="section-header">
      <h1 class="section-title">LIVE DATES</h1>
      <div class="title-accent"></div>
    </div>

    <div v-if="scheduleItemsStore.loading" class="loading-state">
      <span>Loading events&hellip;</span>
    </div>

    <div v-else-if="scheduleItemsStore.scheduleItems?.length" class="events-list">
      <article
        v-for="item in scheduleItemsStore.scheduleItems"
        :key="item.id"
        class="event-card"
      >
        <!-- Poster (hidden on mobile) -->
        <div v-if="item.poster_path" class="event-poster">
          <img :src="`/api/${item.poster_path}`" :alt="item.title" />
        </div>

        <!-- Date column (hidden on mobile, shown inline instead) -->
        <div class="event-date-col">
          <span class="event-date">{{ formatDate(item.start_date) }}</span>
        </div>

        <!-- Main info -->
        <div class="event-info">
          <!-- Date shown here on mobile only -->
          <span class="event-date-inline">{{ formatDate(item.start_date) }}</span>
          <h2 class="event-title">{{ item.title }}</h2>
          <div class="event-venue">
            <a
              v-if="item.location_maps_link"
              :href="item.location_maps_link"
              target="_blank"
              class="venue-link"
            >
              <i class="pi pi-map-marker"></i>
              {{ item.location_name }}<span v-if="item.location_city">, {{ item.location_city }}</span>
            </a>
            <span v-else class="venue-text">
              <i class="pi pi-map-marker"></i>
              {{ item.location_name }}<span v-if="item.location_city">, {{ item.location_city }}</span>
            </span>
          </div>
        </div>

        <div class="event-badge-col">
          <span class="event-badge">LIVE</span>
        </div>
      </article>
    </div>

    <div v-else class="empty-state">
      <p>No upcoming events. Stay tuned!</p>
    </div>
  </section>
</template>

<style scoped>
.events-section {
  width: 100%;
  padding: 4rem 2rem 6rem;
  background: var(--band-black);
  min-height: 100vh;
}

.events-list {
  max-width: 860px;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
}

.event-card {
  display: flex;
  align-items: center;
  gap: 1.25rem;
  padding: 1.25rem 1rem;
  border: 1px solid var(--color-border);
  border-bottom: none;
  background: var(--band-black-card);
  transition: background 0.25s, border-color 0.25s;
}

.event-card:last-child {
  border-bottom: 1px solid var(--color-border);
}

.event-card:hover {
  background: var(--band-black-soft);
  border-color: var(--color-border-hover);
}

/* Poster */
.event-poster {
  flex-shrink: 0;
  width: 68px;
  height: 68px;
  overflow: hidden;
  background: var(--band-black-soft);
}

.event-poster img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

/* Date column — desktop/tablet only */
.event-date-col {
  flex-shrink: 0;
  min-width: 130px;
}

.event-date {
  font-size: 0.75rem;
  font-weight: 700;
  letter-spacing: 0.08em;
  text-transform: uppercase;
  color: var(--band-yellow);
}

/* Date shown inside info on mobile */
.event-date-inline {
  display: none;
  font-size: 0.7rem;
  font-weight: 700;
  letter-spacing: 0.08em;
  text-transform: uppercase;
  color: var(--band-yellow);
  margin-bottom: 0.2rem;
}

/* Info */
.event-info {
  flex: 1;
  min-width: 0;
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
}

.event-title {
  font-size: 1rem;
  font-weight: 700;
  color: var(--band-white);
  letter-spacing: 0.04em;
  line-height: 1.3;
}

.event-venue {
  font-size: 0.82rem;
}

.venue-link {
  color: var(--band-green-light);
  display: inline-flex;
  align-items: center;
  gap: 0.35rem;
  transition: color 0.2s;
}
.venue-link:hover { color: var(--band-yellow); }

.venue-text {
  color: var(--band-white-muted);
  display: inline-flex;
  align-items: center;
  gap: 0.35rem;
}

/* Badge */
.event-badge-col {
  flex-shrink: 0;
}

.event-badge {
  background: var(--band-green);
  color: var(--band-white);
  font-size: 0.6rem;
  font-weight: 700;
  letter-spacing: 0.2em;
  padding: 0.28rem 0.6rem;
  border: 1px solid var(--band-green-light);
  white-space: nowrap;
}

.loading-state,
.empty-state {
  text-align: center;
  color: var(--band-white-muted);
  padding: 5rem 2rem;
  font-size: 1.1rem;
}

/* ── Responsive ── */

/* Tablet */
@media (max-width: 768px) {
  .events-section {
    padding: 3rem 1.5rem 5rem;
  }
}

/* Mobile: hide separate date column and poster, show date inline above title */
@media (max-width: 560px) {
  .events-section {
    padding: 2.5rem 0.75rem 4rem;
  }
  .event-card {
    gap: 0.75rem;
    padding: 1rem 0.75rem;
  }
  .event-date-col {
    display: none;
  }
  .event-poster {
    width: 56px;
    height: 56px;
  }
  .event-date-inline {
    display: block;
  }
  .event-badge {
    font-size: 0.55rem;
    padding: 0.22rem 0.45rem;
    letter-spacing: 0.12em;
  }
}

/* Very small phones: also hide poster */
@media (max-width: 380px) {
  .event-poster {
    display: none;
  }
}
</style>
