<script setup lang="ts">
import Header from "@/components/Header.vue"
import SearchBar from "@/components/searchbar/SearchBar.vue"
import CategoriesFilter from "@/components/categories/CategoriesFilter.vue"
import Supplier from "@/components/suppliers/Supplier.vue"
import { ref, watch, watchEffect } from "vue";
import { getCategories, getSuppliers } from "@/api/api";
import type { categoryInfo, supplierInfo } from "@/api/models";
import { supplierType, type query } from "@/api/util";

const categories = ref([] as categoryInfo[])
const suppliers = ref([] as supplierInfo[])
const showIf = ref({ search: '', categories: [], type: supplierType.ALL } as query)

watchEffect(async () => {
    categories.value = (await getCategories()).categories
    suppliers.value = (await getSuppliers()).suppliers
})

//watch(showIf.value, () => console.log(showIf.value))
</script>

<template>
    <Header/>
    <main class="main-container">
        <SearchBar placeholder="Search" @search="(val) => showIf.search = val"/>
        <CategoriesFilter :categories="categories" @filters="(val) => showIf.categories = val"/>
        <section class="content-container">
            <ul class="suppliers-container">
                <Supplier v-for="s in suppliers" :showIf="showIf" :supplier="s"/>
            </ul>
        </section>
    </main>
</template>

<style scoped>
@import url("@/assets/contentLayout.css");
@import url("@/assets/categories.css");

.suppliers-container {
    display: flex;
    flex-wrap: wrap;
}
</style>