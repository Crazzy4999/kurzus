<script setup lang="ts">
import type { menuInfo } from '@/api/models';
import Counter from "@/components/cart/Counter.vue"

defineProps<{
    addingToCart: boolean
    menu: menuInfo
    items: string
}>()
</script>

<template>
    <div class="block-screen" v-if="addingToCart">
        <div class="item-modal">
            <div class="close-btn"></div>
            <div class="img-wrapper">
                <img class="image" :src="menu.image" :alt="menu.name">
            </div>
            <div class="container">
                <div class="info-container">
                    <h2 class="name">{{ menu.name }}</h2>
                    <p class="price">{{ menu.price }} Ft</p>
                    <p class="items">{{ items }}</p>
                </div>
                <span class="cart-flex">
                    <span class="counter-wrapper">
                        <Counter/>
                    </span>
                    <button class="btn">Add to cart</button>
                </span>
            </div>
        </div>
    </div>
</template>

<style scoped>
@import url("@/assets/modal.css");

.item-modal {
    overflow: hidden;
    border-radius: var(--sub-p-size);
    background-color: var(--main-color);
    width: 25%;
    height: 75%;
    margin: auto;
    box-shadow: 0 0 var(--tiny-size) 1px var(--settings-color);
}

.close-btn {
    position: absolute;
}

.img-wrapper {
    position: relative;
    width: 100%;
}

.image {
    width: 100%;
    height: calc(var(--h1-size) * 3);
}

.container {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    width: 92.5%;
    height: calc(100% - var(--h1-size) * 3.5);
    margin-inline: auto;
}

.name {
    font-size: var(--h4-size);
    color: var(--second-color);
    margin-block: var(--p-size) var(--sub-p-size);
}

.price {
    font-weight: bold;
    color: var(--second-hover);
    margin-block: var(--tiny-size);
}

.items {
    color: var(--settings-color);
}

.cart-flex {
    display: flex;
    justify-content: space-around;
}

.counter-wrapper {
    margin-block: auto;
}

.btn {
    cursor: pointer;
    font-size: var(--p-size);
    border-radius: var(--tiny-size);
    background-color: var(--second-color);
    color: var(--main-color);
    padding: var(--sub-p-size);
    transition: background-color var(--tran);
}

.btn:hover {
    background-color: var(--second-hover);
}

@media only screen and (hover: none) and (pointer: coarse) {
    .item-modal {
        width: 75%;
    }
}
</style>