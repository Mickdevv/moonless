import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import ProductsView from '@/views/ProductsView.vue'
import AddProductView from '@/views/admin/products/AddProductView.vue'
import UpdateProductView from '@/views/admin/products/UpdateProductView.vue'
import ProductsTableView from '@/views/admin/products/ProductsTableView.vue'
import LoginView from '@/views/LoginView.vue'
import RegisterView from '@/views/RegisterView.vue'
import AddContentLinkView from '@/views/admin/content-links/AddContentLinkView.vue'
import UpdateEventView from '@/views/admin/events/UpdateEventView.vue'
import AddEventView from '@/views/admin/events/AddEventView.vue'
import EventsTableView from '@/views/admin/events/EventsTableView.vue'
import ContentLinksTableView from '@/views/admin/content-links/ContentLinksTableView.vue'
import UpdateContentLinkView from '@/views/admin/content-links/UpdateContentLinkView.vue'

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
    {
      path: '/register',
      name: 'register',
      component: RegisterView,
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView,
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
        {
          path: 'events',
          name: 'admin-events',
          component: EventsTableView, // or AdminProductsView
        },
        {
          path: 'events/add',
          name: 'admin-events-add',
          component: AddEventView,
        },
        {
          path: 'events/:id',
          name: 'admin-events-edit',
          component: UpdateEventView,
          props: true,
        },
        {
          path: 'content-links',
          name: 'admin-content-links',
          component: ContentLinksTableView, // or AdminProductsView
        },
        {
          path: 'content-links/add',
          name: 'admin-content-links-add',
          component: AddContentLinkView,
        },
        {
          path: 'content-links/:id',
          name: 'admin-content-links-edit',
          component: UpdateContentLinkView,
          props: true,
        },
      ],
    },
  ],
})

export default router
