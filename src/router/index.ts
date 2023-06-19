import { createRouter, createWebHistory } from 'vue-router'
import HomeView from "@/views/HomeView.vue"
import SuppliersView from "@/views/SuppliersView.vue"

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
      meta: {
        title: 'Hangry - Home page',
        metaTags: [
          {
            name: 'description',
            content: 'The home page of our example app.'
          },
          {
            property: 'og:description',
            content: 'The home page of our example app.'
          }
        ]
      }
    },
    {
      path: '/',
      name: 'suppliers',
      component: SuppliersView,
      meta: {
        title: 'Hangry - Suppliers'
      }
    },
    {
      path: '/',
      name: 'suppliers',
      component: SuppliersView,
      meta: {
        title: 'Hangry - Suppliers'
      }
    }
  ]
})

export default router
