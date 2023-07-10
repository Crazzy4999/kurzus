import { createRouter, createWebHistory } from 'vue-router'
import HomeView from "@/views/HomeView.vue"
import LoginView from "@/views/LoginView.vue"
import ForgotPasswordView from "@/views/ForgotPasswordView.vue"
import SignUpView from "@/views/SignUpView.vue"
import ProfileView from "@/views/ProfileView.vue"
import ResetView from "@/views/ResetView.vue"
import SuppliersView from "@/views/SuppliersView.vue"
import CategoriesView from "@/views/CategoriesView.vue"
import SupplierView from "@/views/SupplierView.vue"
import HistoryView from "@/views/HistoryView.vue"
import CartView from "@/views/CartView.vue"
import OrderView from "@/views/OrderView.vue"

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
      path: '/forgot',
      name: 'forgot',
      component: ForgotPasswordView,
      meta: {
        title: 'Hangry - Forgot password'
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
      path: '/reset',
      name: 'reset',
      component: ResetView,
      meta: {
        title: 'Hangry - Reset Password'
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
    },
    {
      path: '/order',
      name: 'order',
      component: OrderView,
      meta: {
        title: 'Hangry - Order'
      }
    }
  ]
})

export default router
