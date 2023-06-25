<script setup lang="ts">
import AuthHeader from "@/components/AuthHeader.vue"
import Product from "@/components/cart/Product.vue"
import { ref } from "vue";

defineProps<{
    deliveryFee: number
}>()

let totalCost = ref(0)
</script>

<template>
    <AuthHeader/>
    <div class="main-wrapper">
        <form>
            <div class="inner-wrapper">
                <Product name="Pizza Pizza Pizza Pizza Pizza Pizza Pizza Pizza Pizza Pizza Pizza Pizza Pizza Pizza" :price="3000" @total="(t) => totalCost += t"/>
                <Product name="PizzaPizzaPizzaPizzaPizzaPizzaPizzaPizzaPizzaPizzaPizzaPizzaPizzaPizzaPizzaPizzaPizza" :price="3000" @total="(t) => totalCost += t"/>
                <Product name="Pizza Pizza Pizza Pizza" :price="3000" @total="(t) => totalCost += t"/>
                <Product name="PizzaPizzaPizzaPizzaPizza" :price="3000" @total="(t) => totalCost += t"/>
                <Product name="Pizza Pizza Pizza Pizza Pizza Pizza Pizza Pizza Pizza Pizza Pizza Pizza Pizza Pizza" :price="3000" @total="(t) => totalCost += t"/>
                <Product name="PizzaPizzaPizzaPizzaPizzaPizzaPizzaPizzaPizzaPizzaPizzaPizzaPizzaPizzaPizzaPizzaPizza" :price="3000" @total="(t) => totalCost += t"/>
                <Product name="Pizza Pizza Pizza Pizza" :price="3000" @total="(t) => totalCost += t"/>
                <Product name="PizzaPizzaPizzaPizzaPizza" :price="3000" @total="(t) => totalCost += t"/>
                <div class="prices-container">
                    <div class="pricing">Total: {{ totalCost }}</div>
                    <div class="pricing">Delivery fee: {{ deliveryFee }}</div>
                    <div class="total">Total: {{ totalCost + deliveryFee }}</div>
                </div>
            </div>
            <button @click.prevent="">Checkout</button>
        </form>
    </div>
</template>

<style scoped>
.main-wrapper {
    display: flex;
    height: calc(100% - var(--h1-size) - var(--sub-border-size));
}

form {
    display: flex;
    flex-direction: column;
    border: solid var(--sub-border-size) var(--settings-color);
    border-radius: var(--h3-size);
    box-shadow: 0 0 var(--border-size) calc(var(--sub-border-size) * 0.5) var(--settings-color);
    width: 25%;
    height: min-content;
    margin: auto;
    padding-block: var(--h4-size);
}

.inner-wrapper {
    position: relative;
    overflow-y: scroll;
    max-height: calc(var(--h1-size) * 6);
    margin-inline: var(--h3-size);
}

.inner-wrapper::-webkit-scrollbar {
    display: none;
}

.prices-container {
    border-top: solid calc(var(--tiny-size) * 0.5) var(--settings-color);
    padding-top: var(--p-size);
    margin-top: var(--h3-size);
}

.pricing {
    font-size: var(--p-size);
}

.total {
    font-size: var(--p-size);
    font-weight: bold;
}

button {
    cursor: pointer;
    font-size: var(--h4-size);
    border-radius: var(--p-size);
    background-color: var(--second-color);
    color: var(--main-color);
    padding: var(--sub-p-size) var(--p-size);
    margin: var(--h4-size) auto;
    transition: background-color var(--tran);
}

button:hover {
    background-color: var(--second-hover);
}

@media only screen and (hover: none) and (pointer: coarse) {
    .main-wrapper {
        height: calc(100% - var(--h1-size) - var(--border-size) * 2);
    }

    form {
        width: 75%;
    }
}
</style>