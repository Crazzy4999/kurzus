<script setup lang="ts">
import AuthHeader from "@/components/AuthHeader.vue"
import Product from "@/components/cart/Product.vue"
import { ref, watch, watchEffect } from "vue"
import router from "@/router/index"
import { useCartStore } from "@/store";
import { getSupplierByID } from "@/api/api";

const useCart = useCartStore()
const deliveryFee = ref(0)

watchEffect(async () => {
    await getSupplierByID(useCart.products[0].supplierID).then(resp => deliveryFee.value = resp.deliveryFee)
})

let totalCost = ref(0)
</script>

<template>
    <AuthHeader/>
    <div class="main-wrapper">
        <form>
            <div class="inner-wrapper">
                <Product v-for="p in useCart.products" :product="p" @total="(t) => totalCost += t"/>
            </div>
            <div v-if="totalCost !== 0" class="prices-container">
                <div class="pricing">Subtotal: {{ totalCost }}</div>
                <div class="pricing">Delivery fee: {{ deliveryFee === undefined ? "free" : deliveryFee }}</div>
                <div class="total">Total: {{ totalCost + (deliveryFee === undefined ? 0 : deliveryFee) }}</div>
            </div>
            <div v-if="totalCost === 0" class="empty-cart">Your cart is empty</div>
            <button :empty="totalCost === 0" :disabled="totalCost === 0" @click.prevent="router.push('/order')">Checkout</button>
        </form>
    </div>
</template>

<style scoped>
.main-wrapper {
    display: flex;
    height: fit-content;
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
    margin-block: var(--h1-size);
    padding-block: var(--h4-size);
}

.inner-wrapper {
    position: relative;
    margin-inline: var(--h3-size);
}

.inner-wrapper::-webkit-scrollbar {
    display: none;
}

.prices-container {
    border-top: solid calc(var(--tiny-size) * 0.5) var(--settings-color);
    padding-top: var(--p-size);
    margin-top: var(--h3-size);
    margin-inline: var(--h3-size);
}

.pricing {
    font-size: var(--p-size);
}

.total {
    font-size: var(--p-size);
    font-weight: bold;
}

.empty-cart {
    text-align: center;
    font-size: var(--h4-size);
    color: var(--second-color);
}

button {
    cursor: pointer;
    font-size: var(--h4-size);
    border-radius: var(--p-size);
    color: var(--main-color);
    padding: var(--sub-p-size) var(--p-size);
    margin: var(--h4-size) auto;
    transition: background-color var(--tran);
}

button[empty=false] {
    background-color: var(--second-color);
}

button[empty=true] {
    background-color: var(--settings-color);
}

button[empty=false]:hover {
    background-color: var(--second-hover);
}

button[empty=true]:hover {
    background-color: var(--settings-dark);
}

@media only screen and (hover: none) and (pointer: coarse) {
    form {
        width: 75%;
    }
}
</style>