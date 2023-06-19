import { createRouter, createWebHistory } from 'vue-router'
import HomeView from "@/views/HomeView.vue"
import LoginView from "@/views/LoginView.vue"
import SignUpView from "@/views/SignUpView.vue"
import ProfileView from "@/views/ProfileView.vue"
import SuppliersView from "@/views/SuppliersView.vue"
import CategoriesView from "@/views/CategoriesView.vue"
import SupplierView from "@/views/SupplierView.vue"
import HistoryView from "@/views/HistoryView.vue"
import CartView from "@/views/CartView.vue"

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
      meta: {
        title: 'Hangry - Home page'
      }
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView,
      meta: {
        title: 'Hangry - Login'
      }
    },
    {
      path: '/signup',
      name: 'signup',
      component: SignUpView,
      meta: {
        title: 'Hangry - Sign up'
      }
    },
    {
      path: '/profile',
      name: 'profile',
      component: ProfileView,
      meta: {
        title: 'Hangry - Profile'
      }
    },
    {
      path: '/suppliers',
      name: 'suppliers',
      component: SuppliersView,
      meta: {
        title: 'Hangry - Suppliers'
      }
    },
    {
      path: '/categories',
      name: 'categories',
      component: CategoriesView,
      meta: {
        title: 'Hangry - Categories'
      }
    },
    {
      path: '/supplier',
      name: 'supplier',
      component: SupplierView,
      meta: {
        title: 'Hangry' //might need some other approach if possible
      }
    },
    {
      path: '/history',
      name: 'history',
      component: HistoryView,
      meta: {
        title: 'Hangry - History'
      }
    },
    {
      path: '/cart',
      name: 'cart',
      component: CartView,
      meta: {
        title: 'Hangry - Cart'
      }
    }
  ]
})

export default router
