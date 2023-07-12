<script setup lang="ts">
import Header from "@/components/Header.vue"
import SearchBar from "@/components/searchbar/SearchBar.vue"
import CategoriesFilter from "@/components/categories/CategoriesFilter.vue"
import Supplier from "@/components/suppliers/Supplier.vue"
import { ref, watchEffect } from "vue";
import { getCategories, getSuppliers } from "@/api/api";
import type { categoryInfo, supplierInfo } from "@/api/models";

const categories = ref([] as categoryInfo[])
const suppliers = ref([] as supplierInfo[])

watchEffect(async () => {
    categories.value = (await getCategories()).categories
    suppliers.value = (await getSuppliers()).suppliers
})
</script>

<template>
    <Header/>
    <main class="main-container">
        <SearchBar placeholder="Search"/>
        <CategoriesFilter :categories="categories"/>
        <section class="content-container">
            <ul class="suppliers-container">
                <Supplier v-for="s in suppliers" :showIf="'all'" :supplier="s"/>
            </ul>
        </section>
    </main>
</template>

<style scoped>
@import url("@/assets/contentLayout.css");
@import url("@/assets/categories.css");
</style>