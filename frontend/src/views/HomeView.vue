<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';
import Button from 'primevue/button';
import Card from 'primevue/card';
import Carousel from 'primevue/carousel';
import { Tag } from 'primevue';
import { useProductStore } from '@/stores/products';
import Galleria from 'primevue/galleria';

const bandName = 'MOONLESS';
const productStore = useProductStore();
const router = useRouter();

onMounted(() => {
  productStore.getProducts();
});

const images = ref([{
  itemImageSrc: '/api/static/images/misc/655833948_10243354944988101_8603034280536231668_n.jpg',
  alt: 'Description for Image 1',
  title: 'Title 1'
}, {
  itemImageSrc: '/api/static/images/misc/656173332_10243354940187981_7768129687893519335_n.jpg',
  alt: 'Description for Image 1',
  title: 'Title 1'
}, {
  itemImageSrc: '/api/static/images/misc/657085083_10243354929467713_3260572721449297061_n.jpg',
  alt: 'Description for Image 1',
  title: 'Title 1'
},
])
const activeIndex = ref(0)
const responsiveOptions = ref()

const next = () => {
  activeIndex.value = activeIndex.value === images.value.length - 1 ? images.value.length - 1 : activeIndex.value + 1;
};
const prev = () => {
  activeIndex.value = activeIndex.value === 0 ? 0 : activeIndex.value - 1;
};
</script>

<template>
  <div class="card">

    <Galleria v-model:activeIndex="activeIndex" :value="images" :responsiveOptions="responsiveOptions" :numVisible="5"
      containerStyle="width: 100%" :showThumbnails="false" :showIndicators="true" :circular="true" :autoPlay="true"
      :transitionInterval="5000" :showItemNavigators="true">
      <template #item="slotProps">
        <img :src="slotProps.item.itemImageSrc" :alt="slotProps.item.alt" style="width: 100%" />
      </template>
      <template #thumbnail="slotProps">
        <img class="carousel-thumbnail" :src="slotProps.item.itemImageSrc" :alt="slotProps.item.alt" />
      </template>
    </Galleria>
  </div>
</template>

<style scoped>
.carousel-thumbnail {
  max-width: 5rem;
}

.carousel-card {
  cursor: pointer;
  display: flex;
  flex-direction: column;
  align-items: center;
  height: 100%;
  /* Ensure the card takes the available carousel item height */
}

.carousel-image {
  width: 100%;
  max-height: 200px;
  border-radius: 15px;
  /* or any reasonable value */
  object-fit: cover;
}

.card {
  display: block;
}
</style>
