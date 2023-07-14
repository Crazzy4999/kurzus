<script setup lang="ts">
import { useAuthStore } from '@/store';

defineProps<{
    toggle: boolean
}>()

const useAuth = useAuthStore()
</script>

<template>
    <aside :shown="toggle">
        <ul>
            <li>
                <div class="dummy-nav-item"></div>
            </li>
            <li v-if="useAuth.email !== ''">
                <RouterLink to="/profile" class="nav-item">Profile</RouterLink>
            </li>
            <li v-if="useAuth.email === ''">
                <RouterLink to="/login" class="nav-item">Login</RouterLink>
            </li>
            <li v-if="useAuth.email === ''">
                <RouterLink to="/signup" class="nav-item">Sign Up</RouterLink>
            </li>
            <li>
                <RouterLink to="/suppliers"  class="nav-item">Suppliers</RouterLink>
            </li>
            <li>
                <RouterLink to="/categories"  class="nav-item">Categories</RouterLink>
            </li>
            <li v-if="useAuth.email !== ''">
                <RouterLink to="/history" class="nav-item">History</RouterLink>
            </li>
            <li v-if="useAuth.email !== ''">
                <button to="/" class="nav-item" @click.prevent="useAuth.signOut()">Sign Out</button>
            </li>
        </ul>
    </aside>
</template>

<style scoped>
@import "@/assets/layout.css";

aside {
    display: none;
}

@media only screen and (hover: none) and (pointer: coarse) {
    aside {
        z-index: 3;
        position: absolute;
        display: flex;
        flex-direction: column;
        background-color: var(--second-hover);
        width: 50%;
        height: 100%;
        top: calc(var(--h1-size) * -2 - var(--sub-p-size));
        right: 0%;
        padding-top: calc(var(--h1-size) * 2 + var(--sub-p-size));
        transition: right var(--tran);
    }

    aside[shown="false"] {
        right: -50%;
    }

    aside[shown="true"] {
        right: 0%;
    }
}
</style>