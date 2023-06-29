<script setup lang="ts">
import DeliveryPoint, { type deliveryInfo } from "@/components/history/DeliveryPoint.vue"
import { ref } from 'vue'
import type { itemInfo } from "@/components/history/ItemCard.vue"
import ItemCard from "@/components/history/ItemCard.vue"

export type orderInfo = {
    img: string,
    supplier: string,
    supplierURL: string
    items: itemInfo[],
    deliveryFee: number,
    deliveredAt: string,
    address: string,
    paymentMethod: { img: string, name: string }
}

const props = defineProps<{
    order: orderInfo
}>()

let subTotal = () => {
    let t = 0
    props.order.items.forEach(i => t += i.price)
    return t
}

let total = () => {
    return subTotal() + (props.order.deliveryFee === undefined ? 0 : props.order.deliveryFee)
}

let open = ref(false)

let deliveryInfos: deliveryInfo[] = [
    { description: "Order from:", text: props.order.supplier, URL: props.order.supplierURL },
    { description: "Delivery to:", text: props.order.address }
]
</script>

<template>
    <div class="order-container">
        <div class="order-info-container">
            <div class="order-overview">
                <div class="supplier-order-container">
                    <span class="img-wrapper">
                        <img class="supplier-img" src="@/assets/Otp_bank_logo.svg.png" :alt="order.supplier">
                    </span>
                    <span class="short-details-container">
                        <div class="supplier-name">{{ order.supplier }}</div>
                        <div class="items-container">
                            <span class="item" v-for="i in order.items">{{ i.count }} x {{ i.name }}</span>
                        </div>
                    </span>
                </div>
                <span class="reorder-info-container">
                    <span class="reorder-price">{{ total() }} Ft</span>
                    <button class="reorder-btn">Reorder</button>
                </span>
            </div>
            <div class="delivery-date">Delivered on {{ order.deliveredAt }}</div>
        </div>
        <div class="details-container" @click="open = !open">
            <span class="details-text">Order Details</span>
            <span class="details-arrow" :flip="open">&and;</span>
        </div>
        <div class="order-details-container" v-if="open">
            <DeliveryPoint :deliveryInfos="deliveryInfos"/>
            <ItemCard :items="order.items"/>
            <div class="sub-details-container">
                <div class="pricing-container">
                    <span class="pricing-text">Subtotal</span>
                    <span class="pricing-text">{{ subTotal() }} Ft</span>
                </div>
                <div class="pricing-container">
                    <span class="pricing-text">Delivery Fee</span>
                    <span class="pricing-text">{{ order.deliveryFee }} Ft</span>
                </div>
                <div class="pricing-container">
                    <span class="pricing-text">Total</span>
                    <span class="pricing-text">{{ total() }} Ft</span>
                </div>
            </div>
            <div class="sub-details-container">
                <div class="paid-with-label">Paid with</div>
                <span class="payment-container">
                    <span class="payment-info-container">
                        <span class="payment-img-wrapper">
                            <img class="payment-img" src="@/assets/Otp_bank_logo.svg.png" :alt="order.paymentMethod.name">
                        </span>
                        <span class="payment-name">{{ order.paymentMethod.name }}</span>
                        </span>
                    <span class="payment-price-label">{{ total() }} Ft</span>
                </span>
            </div>
        </div>
    </div>
</template>

<style scoped>
.order-container {
    box-shadow: 0 0 var(--tiny-size) 1px var(--settings-color);
    padding: var(--p-size) var(--sub-p-size);
    margin-block: var(--p-size);
}

.order-info-container {
    display: flex;
    flex-direction: column;
    font-size: var(--sub-p-size);
}

.order-overview {
    display: flex;
    justify-content: space-between;
}

.supplier-order-container {
    display: flex;
    gap: var(--sub-p-size);
}

.img-wrapper {
    width: var(--h2-size);
    height: var(--h2-size);
}

.supplier-img {
    border-radius: var(--sub-p-size);
    border: calc(var(--sub-border-size) * 1.5) solid var(--settings-color);
    width: 100%;
}

.short-details-container {
    display: flex;
    flex-direction: column;
    width: 70%;
}

.supplier-name {
    font-weight: bold;
}

.items-container {
    display: flex;
    flex-direction: column;
    width: 100%;
}

.item {
    position: relative;
    text-overflow: ellipsis;
    overflow: hidden;
    white-space: nowrap;
    width: calc(var(--h1-size) * 2.5);
    color: var(--settings-color);
}

.reorder-info-container {
    display: flex;
    flex-direction: column;
    gap: var(--sub-p-size);
    padding: var(--sub-p-size);
    padding-top: 0;
}

.reorder-price {
    font-size: calc(var(--p-size) * 0.85);
    text-align: center;
    font-weight: bold;
}

.reorder-btn {
    cursor: pointer;
    font-size: var(--sub-p-size);
    border: var(--sub-border-size) solid var(--second-color);
    border-radius: var(--tiny-size);
    color: var(--second-color);
    padding: var(--tiny-size);
    transition: border var(--tran), color var(--tran), background-color var(--tran);
}

.reorder-btn:hover {
    border: var(--sub-border-size) solid var(--second-hover);
    color: var(--main-color);
    background-color: var(--second-hover)
}

.delivery-date {
    font-size: calc(var(--sub-p-size) * 0.75);
    color: var(--settings-dark);
    width: fit-content;
    margin-block: var(--sub-p-size);
}

.details-container {
    cursor: pointer;
    display: flex;
    justify-content: space-between;
    font-size: var(--sub-p-size);
}

.details-text {
    color: var(--second-color)
}

.details-arrow {
    display: block;
    user-select: none;
    text-align: center;
    color: var(--second-color);
    width: var(--p-size);
    height: var(--p-size);
}

.details-arrow[flip=true] {
    transform: rotateZ(-180deg)
}

.order-details-container {
    position: relative;
    display: flex;
    flex-direction: column;
}



.sub-details-container {
    border-top: var(--sub-border-size) solid var(--settings-color);
    padding-block: var(--tiny-size);
}

.pricing-container {
    display: flex;
    justify-content: space-between;
    margin: var(--p-size) var(--border-size);
}

.pricing-text {
    font-size: var(--p-size);
    color: var(--settings-color);
}



.paid-with-label {
    font-weight: bold;
    font-size: var(--p-size);
    margin-left: var(--border-size);
    margin-top: var(--sub-p-size);
}

.payment-container {
    display: flex;
    justify-content: space-between;
    margin-inline: var(--border-size);
}

.payment-info-container {
    display: flex;
    gap: var(--p-size);
    margin-block: var(--sub-p-size);
}

.payment-img-wrapper {
    width: var(--h3-size);
    height: var(--h3-size);
}

.payment-img {
    border: var(--sub-border-size) solid var(--settings-color);
    box-shadow: 0 0 var(--tiny-size) 0.1px var(--settings-color);
    width: 100%;
}

.payment-name {
    font-size: var(--p-size);
    margin-block: auto;
}

.payment-price-label {
    font-size: var(--p-size);
    font-weight: bold;
    color: var(--second-color);
    margin-block: auto;
}

@media only screen and (hover: none) and (pointer: coarse) {
    .order-info-container {
        font-size: var(--p-size);
    }

    .img-wrapper {
        width: var(--h1-size);
        height: var(--h1-size);
    }

    .short-details-container {
        width: 40%;
    }

    .reorder-btn {
        font-size: var(--p-size);
    }

    .delivery-date {
        font-size: var(--sub-p-size)
    }

    .details-container {
        font-size: var(--p-size);
    }
}
</style>