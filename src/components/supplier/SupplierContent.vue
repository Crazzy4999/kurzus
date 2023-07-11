<script setup lang="ts">
import CategoriesSelector, { type categorieVal } from "@/components/supplier/CategoriesSelector.vue"
import SubCategorie, { type productVal } from "./SubCategorie.vue";
import type { menuInfo, supplierInfo } from "@/api/models";
import { getMenusBySupplierID, getSupplierByID } from "@/api/api";
import { ref, watchEffect } from "vue";
import { useRoute } from "vue-router";

const route = useRoute()
const supplier = ref({} as supplierInfo)
const menus = ref([] as menuInfo[])

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
})

let selectorCategories: categorieVal[] = [

]

let categories: { subCategorie: categorieVal, products: productVal[] }[] = [
    
]

/*
<CategoriesSelector :categories="selectorCategories"/>

*/
</script>

<template>
    <main class="main-container">
        <section class="content-container">
            <img class="supplier-background-img" :src="supplier.image" :alt="supplier.name">
            <h1 class="supplier-name">{{ supplier.name }}</h1>
            <div class="supplier-infos">
                <span class="supplier-delivery-text">Delivery cost:</span>
                <span class="supplier-delivery-cost">{{ supplier.deliveryFee }}Ft</span>
            </div>
            <ul class="supplier-categories-container">
                <li v-for="c in categories" class="supplier-categorie">{{ c }}</li>
            </ul>

            <SubCategorie v-for="c in categories" :categorieId="c.subCategorie.id" :categorieText="c.subCategorie.text" :products="c.products"/>
        </section>
    </main>
</template>

<style scoped>
@import url("@/assets/contentLayout.css");
@import url("@/assets/supplier.css");
</style>