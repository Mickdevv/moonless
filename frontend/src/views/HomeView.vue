<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import Galleria from 'primevue/galleria';

const router = useRouter();

const images = ref([
  {
    src: '/api/static/images/misc/655833948_10243354944988101_8603034280536231668_n.jpg',
    alt: 'Moonless band photo 1',
  },
  {
    src: '/api/static/images/misc/656173332_10243354940187981_7768129687893519335_n.jpg',
    alt: 'Moonless band photo 2',
  },
  {
    src: '/api/static/images/misc/657085083_10243354929467713_3260572721449297061_n.jpg',
    alt: 'Moonless band photo 3',
  },
])

const activeIndex = ref(0)
</script>

<template>
  <!-- Hero -->
  <section class="hero">
    <div
      class="hero-bg"
      :style="images[0] ? { backgroundImage: `url(${images[0].src})` } : {}"
    ></div>
    <div class="hero-overlay"></div>
    <div class="hero-content">
      <p class="hero-eyebrow">Rock &bull; Darkness &bull; Energy</p>
      <h1 class="hero-title">MOONLESS</h1>
      <p class="hero-tagline">Where the night comes alive</p>
      <div class="hero-cta">
        <button class="cta-btn cta-primary" @click="router.push({ name: 'schedule-items' })">
          See Our Dates
        </button>
        <button class="cta-btn cta-secondary" @click="router.push({ name: 'contact' })">
          Contact Us
        </button>
      </div>
    </div>
    <div class="hero-scroll-hint">
      <span>&#8595;</span>
    </div>
  </section>

  <!-- Photo Gallery -->
  <section class="gallery-section">
    <Galleria
      v-model:activeIndex="activeIndex"
      :value="images"
      :numVisible="3"
      containerStyle="width: 100%"
      :showThumbnails="false"
      :showIndicators="true"
      :circular="true"
      :autoPlay="true"
      :transitionInterval="5000"
      :showItemNavigators="true"
    >
      <template #item="slotProps">
        <img
          :src="slotProps.item.src"
          :alt="slotProps.item.alt"
          class="gallery-image"
        />
      </template>
    </Galleria>
  </section>

  <!-- Quick Links -->
  <section class="quick-links">
    <div class="quick-link-card" @click="router.push({ name: 'discography' })">
      <span class="ql-icon">&#9835;</span>
      <h3>Discography</h3>
      <p>Stream our music on Spotify and YouTube</p>
    </div>
    <div class="quick-link-card" @click="router.push({ name: 'schedule-items' })">
      <span class="ql-icon">&#128197;</span>
      <h3>Live Dates</h3>
      <p>Find us on stage near you</p>
    </div>
    <div class="quick-link-card" @click="router.push({ name: 'contact' })">
      <span class="ql-icon">&#9993;</span>
      <h3>Booking</h3>
      <p>Get in touch for shows and press</p>
    </div>
  </section>
</template>

<style scoped>
/* ── Hero ── */
.hero {
  position: relative;
  width: 100%;
  min-height: 92vh;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
}

.hero-bg {
  position: absolute;
  inset: 0;
  background-size: cover;
  background-position: center 30%;
  filter: brightness(0.35);
}

.hero-overlay {
  position: absolute;
  inset: 0;
  background: linear-gradient(
    to bottom,
    rgba(10, 10, 10, 0.05) 0%,
    rgba(10, 10, 10, 0.25) 50%,
    var(--band-black) 100%
  );
}

.hero-content {
  position: relative;
  z-index: 1;
  text-align: center;
  padding: 2rem;
  width: 100%;
  max-width: 960px;
}

.hero-eyebrow {
  font-size: 0.75rem;
  letter-spacing: 0.35em;
  text-transform: uppercase;
  color: var(--band-green-light);
  margin-bottom: 1.25rem;
  font-weight: 600;
}

.hero-title {
  font-size: 9rem;
  font-weight: 900;
  letter-spacing: 0.25em;
  color: var(--band-white);
  line-height: 1;
  text-shadow:
    0 0 60px rgba(245, 197, 24, 0.25),
    0 4px 24px rgba(0, 0, 0, 0.8);
  margin-bottom: 1.25rem;
}

.hero-tagline {
  font-size: 1rem;
  color: var(--band-white-muted);
  letter-spacing: 0.15em;
  text-transform: uppercase;
  margin-bottom: 3rem;
}

.hero-cta {
  display: flex;
  gap: 1rem;
  justify-content: center;
  flex-wrap: wrap;
}

.cta-btn {
  padding: 0.85rem 2.25rem;
  font-size: 0.8rem;
  font-weight: 700;
  letter-spacing: 0.18em;
  text-transform: uppercase;
  border: none;
  cursor: pointer;
  transition: all 0.25s ease;
  white-space: nowrap;
}

.cta-primary {
  background: var(--band-yellow);
  color: var(--band-black);
}
.cta-primary:hover {
  background: var(--band-yellow-dark);
  transform: translateY(-3px);
  box-shadow: 0 8px 24px rgba(245, 197, 24, 0.35);
}

.cta-secondary {
  background: transparent;
  color: var(--band-white);
  border: 2px solid rgba(240, 240, 240, 0.45);
}
.cta-secondary:hover {
  border-color: var(--band-white);
  transform: translateY(-3px);
}

.hero-scroll-hint {
  position: absolute;
  bottom: 2rem;
  left: 50%;
  transform: translateX(-50%);
  color: var(--band-white-muted);
  font-size: 1.25rem;
  animation: bounce 2s infinite;
  z-index: 1;
}

@keyframes bounce {
  0%, 100% { transform: translateX(-50%) translateY(0); }
  50%       { transform: translateX(-50%) translateY(7px); }
}

/* ── Gallery ── */
.gallery-section {
  width: 100%;
  background: var(--band-black);
}

:deep(.gallery-image) {
  width: 100%;
  height: 55vh;
  object-fit: cover;
  display: block;
}

:deep(.p-galleria) {
  background: var(--band-black);
}

:deep(.p-galleria-indicators) {
  background: transparent;
  padding: 0.75rem;
}

:deep(.p-galleria-indicator button) {
  background: var(--band-white-muted) !important;
  width: 7px !important;
  height: 7px !important;
  border-radius: 50% !important;
}

:deep(.p-galleria-indicator.p-highlight button) {
  background: var(--band-yellow) !important;
}

:deep(.p-galleria-prev-button),
:deep(.p-galleria-next-button) {
  color: var(--band-white) !important;
  background: rgba(0, 0, 0, 0.45) !important;
  border-radius: 0 !important;
}

/* ── Quick Links ── */
.quick-links {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  width: 100%;
  border-top: 1px solid var(--color-border);
}

.quick-link-card {
  padding: 2.5rem 1.5rem;
  text-align: center;
  cursor: pointer;
  border-right: 1px solid var(--color-border);
  background: var(--band-black-soft);
  transition: background 0.3s, border-color 0.3s;
}

.quick-link-card:last-child {
  border-right: none;
}

.quick-link-card:hover {
  background: var(--band-black-card);
}

.quick-link-card:hover .ql-icon {
  color: var(--band-yellow);
}

.ql-icon {
  display: block;
  font-size: 1.75rem;
  margin-bottom: 0.75rem;
  color: var(--band-green-light);
  transition: color 0.3s;
}

.quick-link-card h3 {
  font-size: 0.85rem;
  font-weight: 700;
  letter-spacing: 0.15em;
  text-transform: uppercase;
  color: var(--band-white);
  margin-bottom: 0.4rem;
}

.quick-link-card p {
  font-size: 0.8rem;
  color: var(--band-white-muted);
  line-height: 1.5;
}

/* ── Responsive ── */

/* Large desktop — scale down from fixed 9rem */
@media (max-width: 1200px) {
  .hero-title {
    font-size: clamp(5rem, 9vw, 9rem);
  }
}

/* Tablet */
@media (max-width: 768px) {
  .hero-title {
    font-size: 4.5rem;
    letter-spacing: 0.15em;
  }
  .hero-tagline {
    letter-spacing: 0.08em;
    font-size: 0.9rem;
    margin-bottom: 2.5rem;
  }
  .hero-eyebrow {
    letter-spacing: 0.2em;
  }
  :deep(.gallery-image) {
    height: 40vh;
  }
}

/* Mobile */
@media (max-width: 480px) {
  .hero {
    min-height: 88vh;
  }
  .hero-content {
    padding: 1.25rem;
  }
  .hero-title {
    font-size: clamp(2.4rem, 13vw, 3.5rem);
    letter-spacing: 0.06em;
    margin-bottom: 1rem;
  }
  .hero-eyebrow {
    font-size: 0.65rem;
    letter-spacing: 0.15em;
    margin-bottom: 1rem;
  }
  .hero-tagline {
    font-size: 0.8rem;
    letter-spacing: 0.05em;
    margin-bottom: 2rem;
  }
  .cta-btn {
    width: 100%;
    padding: 1rem 1.5rem;
  }
  .hero-cta {
    flex-direction: column;
    gap: 0.75rem;
  }
  :deep(.gallery-image) {
    height: 30vh;
  }

  /* Stack quick links on mobile */
  .quick-links {
    grid-template-columns: 1fr;
  }
  .quick-link-card {
    border-right: none;
    border-bottom: 1px solid var(--color-border);
    padding: 1.75rem 1.5rem;
  }
  .quick-link-card:last-child {
    border-bottom: none;
  }
}

/* Tablet: 2-col quick links */
@media (min-width: 481px) and (max-width: 767px) {
  .quick-links {
    grid-template-columns: repeat(2, 1fr);
  }
  /* Last item spans full width when 3 items in 2-col grid */
  .quick-link-card:last-child {
    border-right: none;
    border-top: 1px solid var(--color-border);
    grid-column: span 2;
  }
}
</style>
