<script setup lang="ts">
import Header from "@/components/Header.vue"
import SearchBar from "@/components/searchbar/SearchBar.vue"
import CategoriesFilter, { type categorie } from "@/components/categories/CategoriesFilter.vue"
import Supplier, { type supplierVal } from "@/components/suppliers/Supplier.vue"

let categories: categorie[] = [

]

let suppliers: supplierVal[] = [

]
</script>

<template>
    <Header/>
    <main class="main-container">
        <SearchBar placeholder="I want to eat..."/>
        <CategoriesFilter :categories="categories"/>
        <section class="content-container">
            <ul class="suppliers-container">
                <Supplier v-for="s in suppliers" :name="s.name" :description="s.description" :img="s.img" :delivery-time="s.deliveryTime" :delivery-fee="s.deliveryFee"/>
            </ul>
        </section>
    </main>
</template>

<style scoped>
@import url("@/assets/contentLayout.css");
@import url("@/assets/categories.css");
</style>