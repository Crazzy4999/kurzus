<script setup lang="ts">
import CategoriesSelector, { type categorieVal } from "@/components/supplier/CategoriesSelector.vue"
import SubCategorie, { type productVal } from "./SubCategorie.vue";
import type { supplierInfo } from "@/api/models";
import { getSupplierByID } from "@/api/api";
import { ref, watch, watchEffect } from "vue";
import { useRoute } from "vue-router";

const route = useRoute()
const supplier = ref({} as supplierInfo)

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
})

let selectorCategories: categorieVal[] = [

]

let categories: { subCategorie: categorieVal, products: productVal[] }[] = [
    
]

/*
<CategoriesSelector :categories="selectorCategories"/>
<SubCategorie v-for="c in categories" :categorieId="c.subCategorie.id" :categorieText="c.subCategorie.text" :products="c.products"/>
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

            
        </section>
    </main>
</template>

<style scoped>
@import url("@/assets/contentLayout.css");
@import url("@/assets/supplier.css");
</style>