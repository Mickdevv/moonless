import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import ProductsView from '@/views/ProductsView.vue'
import AddProductView from '@/views/admin/products/AddProductView.vue'
import UpdateProductView from '@/views/admin/products/UpdateProductView.vue'
import ProductsTableView from '@/views/admin/products/ProductsTableView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/about',
      name: 'about',
      component: () => import('../views/AboutView.vue'),
    },
    {
      path: '/shop',
      name: 'shop',
      component: ProductsView,
    },

    // Admin parent
    {
      path: '/admin',
      children: [
        {
          path: 'products',
          name: 'admin-products',
          component: ProductsTableView, // or AdminProductsView
        },
        {
          path: 'products/add',
          name: 'admin-products-add',
          component: AddProductView,
        },
        {
          path: 'products/:id',
          name: 'admin-products-edit',
          component: UpdateProductView,
          props: true,
        },
      ],
    },
  ],
})

export default router
