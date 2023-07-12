<script setup lang="ts">
import type { menuInfo, itemInfo } from "@/api/models"
import { getItemsByMenuID } from '@/api/api'
import ItemModal from "@/components/cart/ItemModal.vue"
import { ref, watchEffect } from 'vue';

const props = defineProps<{
    categoryID: number
    menu: menuInfo
}>()

const addingToCart = ref(false)
const items = ref([] as itemInfo[])

watchEffect(async () => {
    await getItemsByMenuID(props.menu.id).then(resp => {
        resp.items.forEach(i => {
            let item: itemInfo = {
                id: i.id,
                ingredient: i.ingredient
            }

            items.value.push(item)
        })
    })
})

function displayItems(): string {
    let itemsList: string = ""
    items.value.forEach((i, index) => {
        if(index < items.value.length - 1) itemsList += i.ingredient + ", "
        else itemsList += i.ingredient
    })
    return itemsList
}
</script>

<template>
    <div v-if="menu.categoryID === categoryID" class="product-container">
        <div class="product-img-wrapper">
            <img class="product-img" :src="menu.image" :alt="menu.name" @click="addingToCart = true">
        </div>
        <span class="product-infos-container" @click="addingToCart = true">
            <h3 class="product-name">{{ menu.name }}</h3>
            <p class="product-description">{{ displayItems() }}</p>
            <span class="product-price">{{ menu.price }} Ft</span>
        </span>
        <ItemModal :addingToCart="addingToCart" :menu="menu" :items="displayItems()" @close="(val) => addingToCart = val"/>
    </div>
</template>

<style scoped>
@import url("@/assets/modal.css");

.product-container {
    cursor: pointer;
    display: flex;
    border: solid var(--sub-border-size) var(--second-color);
    border-radius: var(--border-size);
    box-shadow: 0 0 var(--border-size) calc(var(--sub-border-size) * 0.1) var(--second-color);
    height: fit-content;
    margin-block: var(--p-size);
    padding: var(--sub-p-size);
    transition: border var(--tran), box-shadow var(--tran);
}

.product-container:hover {
    border: solid var(--sub-border-size) var(--second-hover);
    box-shadow: 0 0 var(--tiny-size) calc(var(--sub-border-size) * 0.1) var(--second-hover);
}

.product-img {
    width: var(--h1-size);
    height: var(--h1-size);
}

.product-infos-container {
    display: flex;
    flex-direction: column;
    width: 100%;
    margin-left: var(--p-size);
}

.product-name {
    font-size: var(--h3-size);
    color: var(--second-hover);
}

.product-description {
    font-size: var(--p-size);
    color: var(--settings-dark)
}

.product-price {
    font-size: var(--p-size);
    color: var(--second-color);
}

@media only screen and (hover: none) and (pointer: coarse) {
    .product-name {
        font-size: var(--h4-size)
    }
}
</style>