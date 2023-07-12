<script setup lang="ts">
import CategoriesSelector, { type categorieVal } from "@/components/supplier/CategoriesSelector.vue"
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
    await getSupplierByID(+route.params.id).then(resp => {
        supplier.value = {
            id: resp.id,
            type: resp.type,
            image: resp.image,
            name: resp.name,
            description: resp.description,
            deliveryTime: resp.deliveryTime,
            deliveryFee: resp.deliveryFee,
            opening: resp.workingHours.opening,
            closing: resp.workingHours.closing
        }
    })

    await getMenusBySupplierID(supplier.value.id).then(resp => {
        let _menus: menuInfo[] = []
        
        resp.menus.forEach(m => {
            let menu: menuInfo = {
                id: m.id,
                name: m.name,
                image: m.image,
                supplierID: m.supplierID,
                categoryID: m.categoryID,
                price: m.price
            }

            _menus.push(menu)
        })

        menus.value = _menus
    })

    await getCategories().then(resp => {
        let _categories: categoryInfo[] = []

        resp.categories.forEach(c => {
            let category: categoryInfo = {
                id: c.id,
                name: c.name
            }

            _categories.push(category)
        })

        categories.value = _categories
    })
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