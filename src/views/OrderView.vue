<script setup lang="ts">
import AuthHeader from '@/components/AuthHeader.vue';
import Product, { type productInfo } from '@/components/order/Product.vue';
import AddressCard from '@/components/order/AddressCard.vue'
import { ref } from 'vue';
import PaymentCard, { type paymentInfo } from '@/components/order/PaymentCard.vue';

export type addressInfo = {
    city: string,
    street: string,
    houseNumber: string,
    zipCode: string,
    floorNumber: string,
    apartment: string
}

let products: productInfo[] = [
    { count: 1, name: "PIZZA PIZZA PIZZA PIZZA", price: 6570 },
    { count: 10, name: "PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA", price: 6570 },
    { count: 1, name: "PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA", price: 6570 },
    { count: 25, name: "PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA PIZZA", price: 6570 }
]

let name = ref("Test étterem")
let total = ref(6750)
let deliveryFee = ref(500)

let addingAddress = ref(false)
let selectingAddress = ref(false)


let city = ref("")
let street = ref("")
let houseNumber = ref("")
let zipCode = ref("")
let floorNumber = ref("")
let apartment = ref("")

function filter(e: KeyboardEvent) {
    if(!/\d/.test(e.key)) return e.preventDefault()
}

let address: addressInfo = {
    city: city.value,
    street: street.value,
    houseNumber: houseNumber.value,
    zipCode: zipCode.value,
    floorNumber: floorNumber.value,
    apartment: apartment.value
}

let errorMsg = ref("")

function addAddress() {
    if(selectingAddress.value) selectingAddress.value = !selectingAddress.value
    if(address.city !== "" || address.street !== "" ||
       address.houseNumber !== "" || address.zipCode !== "") {
           addingAddress.value = !addingAddress.value
    } else errorMsg.value = "Please fill out the following fields: city, street, house number, zip code."
}

function flip() {
    if(selectingAddress.value) selectingAddress.value = !selectingAddress.value
    if(!addingAddress.value) addingAddress.value = !addingAddress.value
}



let addresses: addressInfo[] = [
    { city: "Budapest", street: "Szamos utca", houseNumber: "8", zipCode: "1039", floorNumber: "2", apartment: ""},
    { city: "Budapest", street: "Szamos utca", houseNumber: "9", zipCode: "1039", floorNumber: "2", apartment: ""},
    { city: "Budapest", street: "Szamos utca", houseNumber: "10", zipCode: "1039", floorNumber: "2", apartment: ""},
    { city: "Budapest", street: "Szamos utca", houseNumber: "11", zipCode: "1039", floorNumber: "2", apartment: ""},
]

function selecting() {
    if(addingAddress.value) addingAddress.value = !addingAddress.value
    if(!selectingAddress.value) selectingAddress.value = !selectingAddress.value
}

let paymentInfos: paymentInfo[] = [
    { text: "OTP", img: "@/assets/Otp_bank_Logo.svg.png" },
    { text: "OTP", img: "@/assets/Otp_bank_Logo.svg.png" },
    { text: "OTP", img: "@/assets/Otp_bank_Logo.svg.png" },
    { text: "OTP", img: "@/assets/Otp_bank_Logo.svg.png" }
]
</script>

<template>
    <AuthHeader/>
    <div class="main-wrapper">
        <form>
            <div class="ordering-wrapper">
                <span class="flex-wrapper">
                    <div class="card">
                        <div class="info">
                            <div class="number-wrapper">
                                <span class="number">1</span>
                            </div>
                            <span class="title">Delivery details</span>
                        </div>
                        <div class="card-text">Delivery address</div>
                        <AddressCard v-if="!(addingAddress || selectingAddress)" :address="addresses[0]" :checked="true"/>
                        <div v-if="addingAddress" class="action-container">
                            <input class="address-input" type="text" placeholder="city" v-model="city">
                            <input class="address-input" type="text" placeholder="street" v-model="street">
                            <input class="address-input" type="text" placeholder="house number" v-model="houseNumber" @keypress="(e) => filter(e)">
                            <input class="address-input" type="text" placeholder="zip code" v-model="zipCode" @keypress="(e) => filter(e)">
                            <input class="address-input" type="text" placeholder="floor number" v-model="floorNumber" @keypress="(e) => filter(e)">
                            <input class="address-input" type="text" placeholder="apartment" v-model="apartment" @keypress="(e) => filter(e)">
                            <span class="error-msg">{{ errorMsg }}</span>
                            <button class="add-address-btn" @click.prevent="addAddress">Add</button>
                        </div>
                        <div v-if="selectingAddress" class="action-container">
                            <ul class="address-list">
                                <AddressCard v-for="a in addresses" :address="a"/>
                            </ul>
                        </div>
                        <div class="action-btns-container">
                            <button class="action-btn" @click.prevent="flip" :show="!addingAddress">Add address</button>
                            <button class="action-btn" @click.prevent="selecting" :show="!addingAddress">Select address</button>
                        </div>
                    </div>
                    <div class="card">
                        <div class="info">
                            <div class="number-wrapper">
                                <span class="number">2</span>
                            </div>
                            <span class="title">Payment</span>
                        </div>
                        <PaymentCard v-for="p in paymentInfos" :info="p"/>
                    </div>
                </span>
                <span class="order-flex-wrapper">
                    <div class="order-from-container">
                        <span class="order-title">Order from: {{ name }}</span>
                    </div>
                    <Product v-for="p in  products" :product="p"/>
                    <div class="order-price-container">
                        <span class="order-price-label">Subtotal: </span>
                        <span class="order-price">{{ total }} Ft</span>
                    </div>
                    <div class="order-price-container">
                        <span class="order-price-label">Delivery fee: </span>
                        <span class="order-price">{{ deliveryFee === undefined ? "Free" : deliveryFee + "Ft" }}</span>
                    </div>
                    <div class="order-price-container">
                        <span class="total-price-label">Total: </span>
                        <span class="total-price">{{ total + (deliveryFee === undefined ? 0 : deliveryFee) }} Ft</span>
                    </div>
                </span>
            </div>
            <button class="order">Order</button>
        </form>
    </div>
</template>

<style scoped>
.main-wrapper {
    display: grid;
    width: 100%;
    height: calc(100% - var(--h1-size) * 1.25 - var(--sub-border-size) * 2);
}

form {
    width: 45%;
    margin: auto;
}

.ordering-wrapper {
    display: flex;
    gap: var(--h3-size);
    width: 100%;
}

.flex-wrapper {
    display: flex;
    flex-direction: column;
    width: 100%;
}

.card {
    box-shadow: 0 0 var(--tiny-size) 1px var(--settings-color);
    width: calc(100% - 2 * var(--sub-p-size));
    height: fit-content;
    padding: var(--sub-p-size);
    margin-block: var(--sub-p-size);
}

.info {
    display: flex;
    font-size: var(--h4-size);
}

.number-wrapper {
    display: grid;
    background-color: var(--second-color);
    width: var(--h3-size);
    height: var(--h3-size);
}

.number {
    font-size: var(--p-size);
    color: var(--main-color);
    margin: auto;
}

.title {
    color: var(--second-color);
    line-height: var(--p-size);
    margin-block: auto;
    margin-left: var(--sub-p-size);
}

.card-text {
    font-size: var(--p-size);
    color: var(--settings-dark);
    margin-block: var(--sub-p-size);
}

.action-container {
    display: flex;
    flex-direction: column;
}

.address-input {
    font-size: var(--p-size);
    border: var(--sub-border-size) solid var(--second-color);
    color: var(--second-hover);
    padding: var(--tiny-size);
    margin: var(--tiny-size);
    transition: box-shadow var(--tran);
}

.address-input:hover, .address-input:focus {
    box-shadow: 0 0 var(--tiny-size) 1px var(--second-color);
}

.error-msg {
    text-align: center;
    font-size: var(--p-size);
    color: var(--error-color);
    margin: var(--sub-p-size) 0;
}

.add-address-btn {
    cursor: pointer;
    font-size: var(--p-size);
    border-radius: var(--sub-p-size);
    background-color: var(--second-color);
    color: var(--main-color);
    width: 35%;
    padding: var(--tiny-size);
    margin: var(--p-size) auto;
    transition: background-color var(--tran);
}

.add-address-btn:hover {
    background-color: var(--second-hover);
}

.address-list {
    margin: var(--p-size);
}

.action-btns-container {
    display: flex;
    justify-content: space-between;
}

.action-btn {
    cursor: pointer;
    font-size: var(--p-size);
    color: var(--second-color);
    transition: color var(--tran);
}

.action-btn:hover {
    color: var(--second-hover);
}

.order-flex-wrapper {
    display: flex;
    flex-direction: column;
    width: 50%;
    margin-bottom: var(--h4-size);
}

.order-from-container {
    display: flex;
    justify-content: space-between;
    font-weight: bold;
    margin-block: var(--h4-size);
}

.order-title {
    font-size: var(--p-size);
}

.order-price-container {
    display: flex;
    justify-content: space-between;
    font-size: var(--sub-p-size);
    margin-block: calc(var(--tiny-size) * 0.5);
}

.order-price-label {
    color: var(--settings-dark);
}

.order-price {
    color: var(--settings-dark);
}

.total-price-label {
    font-weight: bold;
}

.total-price {
    font-weight: bold;
}

.order {
    display: block;
    cursor: pointer;
    font-size: var(--h4-size);
    background-color: var(--second-color);
    border-radius: var(--sub-p-size);
    color: var(--main-color);
    padding: var(--tiny-size) var(--h1-size);
    margin: var(--p-size) auto;
    transition: background-color var(--tran);
}

.order:hover {
    background-color: var(--second-hover)
}

@media only screen and (hover: none) and (pointer: coarse) {
    form {
        width: 75%;
    }

    .ordering-wrapper {
        flex-direction: column;
        gap: 0;
    }

    .card-text {
        font-size: var(--p-size);
    }

    .order-flex-wrapper {
        width: 100%;
    }

    .order-title {
        font-size: var(--h4-size);
    }

    .order-price-container {
        font-size: var(--p-size);
    }
}
</style>