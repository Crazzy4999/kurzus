<script setup lang="ts">
import DeliveryPoint, { type deliveryInfo } from "@/components/history/DeliveryPoint.vue"
import { ref, watchEffect } from 'vue'
import ItemCard from "@/components/history/ItemCard.vue"
import type { addressInfo, menuInfo, orderMenuInfo, orderInfo, productInfo, supplierInfo } from "@/api/models";
import { getAddressByID, GetMenusByOrderID, GetOrderMenusByOrderID, getSupplierByID, makeOrder } from "@/api/api";
import router from "@/router";
import { time } from "console";

const props = defineProps<{
    order: orderInfo
}>()

const menus = ref([] as menuInfo[])
const orderMenus = ref([] as orderMenuInfo[])
const products = ref([] as productInfo[])
const supplier = ref({} as supplierInfo)
const address = ref({} as addressInfo)
const deliveryInfos = ref([] as deliveryInfo[]) 

watchEffect(async () => {
    menus.value = (await GetMenusByOrderID(props.order.id)).menus
    orderMenus.value = (await GetOrderMenusByOrderID(props.order.id)).orderMenus
    menus.value.forEach(menu => {
        orderMenus.value.forEach(orderMenu => {
            if(menu.id === orderMenu.menuID) products.value.push({ count: orderMenu.quantity, menuID: menu.id, supplierID: menu.supplierID, name: menu.name, price: menu.price })
        })
    })
    supplier.value = await getSupplierByID(props.order.supplierID)
    address.value = await getAddressByID(props.order.addressID)
    deliveryInfos.value = [
        { description: "Order from:", text: supplier.value.name, URL: "/supplier/" + supplier.value.id },
        { description: "Delivery to:", text: address.value.city + " " + address.value.street + " " + address.value.houseNumber + " " + address.value.zipCode + " " + address.value.floorNumber  + " " + address.value.apartment }
    ]
})

let subTotal = () => {
    let t = 0
    products.value.forEach(product => t += product.price * product.count)
    return t
}

let total = () => {
    return subTotal() + (supplier.value.deliveryFee === undefined ? 0 : supplier.value.deliveryFee)
}

let open = ref(false)

function formatDelivering() {
    const date = new Date(Date.parse(props.order.deliveredAt))
    const timeTillDone = new Date(Date.parse(props.order.estimatedDelivery) - Date.now())
    const months = ['Jan','Feb','Mar','Apr','May','Jun','Jul','Aug','Sep','Oct','Nov','Dec'];
    if(props.order.deliveredAt === "") return "Delivery in progress - " + (timeTillDone.getMinutes() < 5 ? "any minute" : "~" + timeTillDone.getMinutes() + " mins")
    return "Delivered at " + date.getFullYear() + " " + months[date.getMonth()] + ". " + date.getDay() + ". " + date.getHours() + ":" + date.getMinutes()
}

async function reorder() {
    /*await makeOrder(props.order.userID, props.order.addressID, props.order.supplierID, props.order.note, new Date(Date.now()+supplier.value.deliveryTime*60000+(1000*60*(-(new Date()).getTimezoneOffset()))).toISOString().replace('T',' ').replace('Z',''))
    router.push("/")*/
}
</script>

<template>
    <div class="order-container">
        <div class="order-info-container">
            <div class="order-overview">
                <div class="supplier-order-container">
                    <span class="img-wrapper">
                        <img class="supplier-img" :src="supplier.image" :alt="supplier.name">
                    </span>
                    <span class="short-details-container">
                        <div class="supplier-name">{{ supplier.name }}</div>
                        <div class="items-container">
                            <span class="item" v-for="i in products">{{ i.count }} x {{ i.name }}</span>
                        </div>
                    </span>
                </div>
                <span class="reorder-info-container">
                    <span class="reorder-price">{{ total() }} Ft</span>
                    <button class="reorder-btn" @click.prevent="reorder()">Reorder</button>
                </span>
            </div>
            <div class="delivery-date">{{ formatDelivering() }}</div>
        </div>
        <div class="details-container" @click="open = !open">
            <span class="details-text">Order Details</span>
            <span class="details-arrow" :flip="open">&and;</span>
        </div>
        <div class="order-details-container" v-if="open">
            <DeliveryPoint :deliveryInfos="deliveryInfos"/>
            <ItemCard :items="products"/>
            <div class="sub-details-container">
                <div class="pricing-container">
                    <span class="pricing-text">Subtotal</span>
                    <span class="pricing-text">{{ subTotal() }} Ft</span>
                </div>
                <div class="pricing-container">
                    <span class="pricing-text">Delivery Fee</span>
                    <span class="pricing-text">{{ supplier.deliveryFee }} Ft</span>
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
                            <img class="payment-img" src="@/assets/Otp_bank_logo.svg.png" alt=""> <!--order.paymentMethod.name-->
                        </span>
                        <span class="payment-name">{{  }}</span> <!--order.paymentMethod.name-->
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
    overflow: hidden;
    border-radius: var(--sub-p-size);
    border: calc(var(--sub-border-size) * 1.5) solid var(--settings-color);
    width: var(--h2-size);
    height: var(--h2-size);
}

.supplier-img {
    width: 400%;
    height: var(--h2-size);
    transform: translateX(-37.5%);
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

    .supplier-img {
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