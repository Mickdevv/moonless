<script setup>
import 'primeicons/primeicons.css'
defineProps(['product'])
</script>

<template>
  <div class="product-card">
    <div class="product-image-wrapper">
      <img
        v-if="product.images?.length"
        :src="'/api/' + product.images[0].path"
        :alt="product.name"
        class="product-image"
      />
      <div v-else class="product-image-placeholder">
        <i class="pi pi-image"></i>
      </div>
    </div>
    <div class="product-info">
      <h3 class="product-name">{{ product.name }}</h3>
      <p class="product-description">{{ product.description }}</p>
      <div class="product-footer">
        <span v-if="product.price" class="product-price">
          €{{ (product.price / 100).toFixed(2) }}
        </span>
        <div class="product-actions">
          <button class="product-btn btn-cart" title="Add to cart">
            <i class="pi pi-shopping-cart"></i>
          </button>
          <button class="product-btn btn-view" title="View details">
            <i class="pi pi-eye"></i>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.product-card {
  background: var(--band-black-card);
  border: 1px solid var(--color-border);
  overflow: hidden;
  display: flex;
  flex-direction: column;
  transition: border-color 0.3s, transform 0.3s, box-shadow 0.3s;
}

.product-card:hover {
  border-color: var(--band-yellow);
  transform: translateY(-5px);
  box-shadow: 0 12px 36px rgba(0, 0, 0, 0.5);
}

.product-image-wrapper {
  width: 100%;
  height: 220px;
  overflow: hidden;
  background: var(--band-black-soft);
  flex-shrink: 0;
}

.product-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.4s ease;
}

.product-card:hover .product-image {
  transform: scale(1.06);
}

.product-image-placeholder {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--band-white-muted);
  font-size: 3rem;
}

.product-info {
  padding: 1.25rem;
  display: flex;
  flex-direction: column;
  flex: 1;
  gap: 0.5rem;
}

.product-name {
  font-size: 0.95rem;
  font-weight: 700;
  letter-spacing: 0.1em;
  color: var(--band-white);
  text-transform: uppercase;
}

.product-description {
  font-size: 0.85rem;
  color: var(--band-white-muted);
  flex: 1;
  line-height: 1.5;
}

.product-footer {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-top: 1rem;
  padding-top: 1rem;
  border-top: 1px solid var(--color-border);
}

.product-price {
  font-size: 1.2rem;
  font-weight: 700;
  color: var(--band-yellow);
}

.product-actions {
  display: flex;
  gap: 0.5rem;
}

.product-btn {
  width: 36px;
  height: 36px;
  display: flex;
  align-items: center;
  justify-content: center;
  border: none;
  cursor: pointer;
  font-size: 0.9rem;
  transition: all 0.2s ease;
}

.btn-cart {
  background: var(--band-yellow);
  color: var(--band-black);
}
.btn-cart:hover {
  background: var(--band-yellow-dark);
}

.btn-view {
  background: transparent;
  border: 1px solid var(--band-white-muted);
  color: var(--band-white-muted);
}
.btn-view:hover {
  border-color: var(--band-white);
  color: var(--band-white);
}
</style>
