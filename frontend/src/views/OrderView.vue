<script setup lang="ts">
import AuthHeader from '@/components/AuthHeader.vue';
import Product from '@/components/order/Product.vue';
import AddressCard from '@/components/order/AddressCard.vue'
import { ref, watchEffect } from 'vue';
import PaymentCard, { type paymentInfo } from '@/components/order/PaymentCard.vue';
import type { addressInfo, supplierInfo } from '@/api/models';
import { useAuthStore, useCartStore } from '@/store';
import { addAddress, addOrderMenu, getAddresses, getOrders, getSupplierByID, makeOrder } from '@/api/api';
import router from '@/router';
import { orderStatus } from '@/api/util';

const useAuth = useAuthStore()
const useCart = useCartStore()
const supplier = ref({} as supplierInfo)
const products = ref(useCart.products)
const addresses = ref([] as addressInfo[])
const activeAddress = ref({} as addressInfo)

watchEffect(async () => {
    if(useCart.products[0] !== undefined) supplier.value = await getSupplierByID(useCart.products[0].supplierID)
    addresses.value = (await getAddresses()).addresses
    addresses.value.forEach(address => {
        if(address.isActive) activeAddress.value = address
    })
})

const addingAddress = ref(false)
const selectingAddress = ref(false)

const city = ref("")
const street = ref("")
const houseNumber = ref("")
const zipCode = ref("")
const floorNumber = ref("")
const apartment = ref("")

function filter(e: KeyboardEvent) {
    if(!/\d/.test(e.key)) return e.preventDefault()
}

let errorMsg = ref("")

function addNewAddress() {
    if(city.value !== "" && street.value !== "" && houseNumber.value !== "" && zipCode.value !== "") {
        addAddress(false, city.value, street.value, houseNumber.value, zipCode.value, floorNumber.value, apartment.value).then(() => addingAddress.value = false)
    } else errorMsg.value = "city, street, house number and zipcode mustn't be empty!"
}

function flip() {
    if(selectingAddress.value) selectingAddress.value = !selectingAddress.value
    if(!addingAddress.value) addingAddress.value = !addingAddress.value
}

function selecting() {
    if(addingAddress.value) addingAddress.value = !addingAddress.value
    if(!selectingAddress.value) selectingAddress.value = !selectingAddress.value
}

let paymentInfos: paymentInfo[] = [
    { text: "Credit/debit card", img: "@/assets/Otp_bank_Logo.svg.png" },
    { text: "OTP", img: "@/assets/Otp_bank_Logo.svg.png" }
]

const paymentMethod = ref({} as paymentInfo)
const note = ref("")

async function makeNewOrder() {
    const userAddresses = (await getAddresses()).addresses
    let activeAddress = {} as addressInfo
    userAddresses.forEach(address => {
        if(address.isActive) activeAddress = address
    })
    await makeOrder(useAuth.id, activeAddress.id, useCart.products[0].supplierID, note.value, new Date(Date.now()+supplier.value.deliveryTime*60000+(1000*60*(-(new Date()).getTimezoneOffset()))).toISOString().replace('T',' ').replace('Z',''))
    const orders = (await getOrders()).orders
    let orderID = -1
    orders.forEach(order => {
        if(order.statusID === orderStatus.CREATED) orderID = order.id
    })
    useCart.products.forEach(product => {
        addOrderMenu(orderID, product.menuID, product.count)
    })
    useCart.products.length = 0
    router.push("/history")
}
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
                        <AddressCard v-if="!(addingAddress || selectingAddress)" :address="activeAddress" :checked="true"/>
                        <div v-if="addingAddress" class="action-container">
                            <input class="address-input" type="text" placeholder="city" v-model="city">
                            <input class="address-input" type="text" placeholder="street" v-model="street">
                            <input class="address-input" type="text" placeholder="house number" v-model="houseNumber" @keypress="(e) => filter(e)">
                            <input class="address-input" type="text" placeholder="zip code" v-model="zipCode" @keypress="(e) => filter(e)">
                            <input class="address-input" type="text" placeholder="floor number" v-model="floorNumber" @keypress="(e) => filter(e)">
                            <input class="address-input" type="text" placeholder="apartment" v-model="apartment" @keypress="(e) => filter(e)">
                            <span class="error-msg">{{ errorMsg }}</span>
                            <button class="add-address-btn" @click.prevent="addNewAddress()">Add</button>
                        </div>
                        <div v-if="selectingAddress" class="action-container">
                            <ul class="address-list">
                                <AddressCard v-for="a in addresses" :address="a"/>
                            </ul>
                        </div>
                        <div class="action-btns-container">
                            <button class="action-btn" @click.prevent="flip()" :show="!addingAddress">Add address</button>
                            <button class="action-btn" @click.prevent="selecting()" :show="!addingAddress">Select address</button>
                        </div>
                    </div>
                    <div class="card">
                        <div class="info">
                            <div class="number-wrapper">
                                <span class="number">2</span>
                            </div>
                            <span class="title">Payment</span>
                        </div>
                        <PaymentCard v-for="p in paymentInfos" :info="p" @select="(val) => paymentMethod = val"/>
                    </div>
                    <div class="card">
                        <div class="info">
                            <div class="number-wrapper">
                                <span class="number">3</span>
                            </div>
                            <span class="title">Note to driver</span>
                        </div>
                        <textarea class="note" name="note" id="note" placeholder="Leave a note to your delivery driver" maxlength="512" v-model="note"></textarea>
                        <div class="note-char-count">512/{{ note.length }}</div>
                    </div>
                </span>
                <span class="order-flex-wrapper">
                    <div class="order-from-container">
                        <span class="order-title">Order from: {{ supplier.name }}</span>
                    </div>
                    <Product v-for="p in  products" :product="p"/>
                    <div class="order-price-container">
                        <span class="order-price-label">Subtotal: </span>
                        <span class="order-price">{{ useCart.total }} Ft</span>
                    </div>
                    <div class="order-price-container">
                        <span class="order-price-label">Delivery fee: </span>
                        <span class="order-price">{{ supplier.deliveryFee === undefined ? "Free" : supplier.deliveryFee + "Ft" }}</span>
                    </div>
                    <div class="order-price-container">
                        <span class="total-price-label">Total: </span>
                        <span class="total-price">{{ useCart.total + (supplier.deliveryFee === undefined ? 0 : supplier.deliveryFee) }} Ft</span>
                    </div>
                </span>
            </div>
            <!--INSERT PAYMENT PROCESSOR HERE XD-->
            <button class="order" @click.prevent="makeNewOrder()" :blocked="Object.keys(paymentMethod).length === 0" :disabled="Object.keys(paymentMethod).length === 0">Order</button>
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

.note {
    position: relative;
    resize: none;
    outline: none !important;
    border: var(--sub-border-size) solid var(--second-color);
    width: calc(100% - var(--sub-border-size) * 4 - var(--tiny-size) * 2);
    height: calc(var(--h1-size) * 2);
    padding: var(--tiny-size);
    margin-block: var(--sub-p-size);
}

.note-char-count {
    color: var(--second-color)
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

.order[blocked=true] {
    background-color: var(--settings-color);   
}

.order:hover {
    background-color: var(--second-hover)
}

.order[blocked=true]:hover {
    background-color: var(--settings-dark);   
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