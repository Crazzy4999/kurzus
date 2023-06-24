<script setup lang="ts">
export type productVal = {
    img: string
    name: string
    ingredients: string[]
    price: number
}

defineProps<{
    categorieId: string
    categorieText: string
    products: productVal[]
}>()

function foramtIngredients(ingredients: string[]): string {
    return ""
}

function formatPrice(price: number): string {
    return ""
}
</script>

<template>
    <div class="sub-categorie-container">
        <h2 class="sub-categorie-header" :id="categorieId">{{ categorieText }}</h2>
        <div v-for="p in products" class="product-container">
            <img class="product-img" :src="p.img" :alt="p.name">
            <span class="product-infos-container">
                <h3 class="product-name">{{ p.name }}</h3>
                <p class="product-description">{{ foramtIngredients(p.ingredients) }}</p>
                <span class="product-price">{{ formatPrice(p.price) }} Ft</span>
            </span>
        </div>
    </div>
</template>

<style scoped>

</style>