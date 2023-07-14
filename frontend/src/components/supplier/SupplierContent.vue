<script setup lang="ts">
import CategoriesSelector from "@/components/supplier/CategoriesSelector.vue"
import SubCategorie from "./SubCategorie.vue";
import type { categoryInfo, menuInfo, supplierInfo } from "@/api/models";
import { getCategories, getMenusBySupplierID, getSupplierByID } from "@/api/api";
import { ref, watchEffect } from "vue";
import { useRoute } from "vue-router";

const route = useRoute()
const supplier = ref({} as supplierInfo)
const menus = ref([] as menuInfo[])
const categories = ref([] as categoryInfo[])

watchEffect(async () => {
    supplier.value = await getSupplierByID(+route.params.id)
    menus.value = (await getMenusBySupplierID(supplier.value.id)).menus
    const filteredCategories: categoryInfo[] = []
    const fetchedCategories = (await getCategories()).categories.forEach(category => {
        menus.value.forEach(menu => {
            if(category.id === menu.categoryID && !filteredCategories.includes(category)) {
                filteredCategories.push(category)
            }
        })
    })
    categories.value = filteredCategories
})

/*
<CategoriesSelector :categories="selectorCategories"/>
*/
</script>

<template>
    <main class="main-container">
        <section class="content-container">
            <img class="supplier-background-img" :src="supplier.image" :alt="supplier.name">
            <h1 class="supplier-name">{{ supplier.name }}</h1>
            <h2 class="supplier-description">{{ supplier.description }}</h2>
            <div class="supplier-infos">
                <span class="supplier-delivery-text">Delivery cost:</span>
                <span class="supplier-delivery-cost">{{ supplier.deliveryFee }}Ft</span>
            </div>
            <ul class="supplier-categories-container">
                <li v-for="c in categories" class="supplier-categorie">{{ c.name }}</li>
            </ul>

            <SubCategorie :categories="categories" :menus="menus"/>
        </section>
    </main>
</template>

<style scoped>
@import url("@/assets/contentLayout.css");
@import url("@/assets/supplier.css");
</style>