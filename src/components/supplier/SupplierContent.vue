<script setup lang="ts">
import CategoriesSelector, { type categorieVal } from "@/components/supplier/CategoriesSelector.vue"
import SubCategorie, { type productVal } from "./SubCategorie.vue";

const props = defineProps<{
    img: string
    name: string
    rating: number
    ratesCount: number
    deliveryCost: number
    categorieKeys: string[]
    items: productVal[]
}>()

let selectorCategories: categorieVal[] = [

]

function formatRatings() {

}

let categories: { subCategorie: categorieVal, products: productVal[] }[] = [
    
]
</script>

<template>
    <main class="main-container">
        <section class="content-container">
            <img class="restaurant-background-img" :src="img" :alt="name">
            <h1 class="restaurant-name">{{ name }}</h1>
            <div class="restaurant-infos">
                <span class="restaurant-rates-container">
                    <!--Icon by Google - https://www.iconfinder.com/iconsets/material-core-->
                    <img class="restaurant-rate-img" src="@/assets/icons/326703_favorite_rate_star_icon.svg" alt="star">
                    <span class="restaurant-rate">{{ rating }}</span>
                    <span class="restaurant-rate-max">/5</span>
                </span>
                <span class="restaurant-rates-count">({{ formatRatings() }}+)</span>
                <span class="restaurant-delivery-text">Delivery cost:</span>
                <span class="restaurant-delivery-cost">{{ deliveryCost }}Ft</span>
            </div>
            <ul class="restaurant-categories-container">
                <li v-for="cKey in categorieKeys" class="restaurant-categorie">{{ cKey }}</li>
            </ul>

            <CategoriesSelector :categories="selectorCategories"/>
            <SubCategorie v-for="c in categories" :categorieId="c.subCategorie.id" :categorieText="c.subCategorie.text" :products="c.products"/>
        </section>
    </main>
</template>

<style scoped>
@import url("@/assets/contentLayout.css");
@import url("@/assets/supplier.css");
</style>