<script setup lang="ts">
import { getOrders } from '@/api/api';
import type { orderInfo } from '@/api/models';
import Header from '@/components/Header.vue';
import OrderCard from "@/components/history/OrderCard.vue"
import { ref, watchEffect } from 'vue';

/*let orders: orderInfo[] = [
{
        img: "@/assets/Otp_bank_logo.svg.png",
        supplier: "Pizza King",
        supplierURL: "/",
        items: [
            { count: 1, name: "Popeye Pizza Popeye Pizza Popeye Pizza", price: 3500 },
            { count: 25, name: "Popeye Pizza", price: 3500 },
            { count: 10, name: "Popeye Pizza", price: 7500 }
        ],
        deliveryFee: 799,
        deliveredAt: "Apr. 14. 18:35",
        address: "1039 Budapest Szamos utca 8",
        paymentMethod: { img: "@/assets/Otp_bank_logo.svg.png", name: "OTP" }
    },
    {
        img: "@/assets/Otp_bank_logo.svg.png",
        supplier: "Pizza King",
        supplierURL: "/",
        items: [
            { count: 1, name: "Popeye Pizza", price: 3500 },
            { count: 25, name: "Popeye Pizza", price: 3500 },
            { count: 10, name: "Popeye Pizza", price: 7500 }
        ],
        deliveryFee: 799,
        deliveredAt: "Apr. 14. 18:35",
        address: "1039 Budapest Szamos utca 8",
        paymentMethod: { img: "@/assets/Otp_bank_logo.svg.png", name: "OTP" }
    },
    {
        img: "@/assets/Otp_bank_logo.svg.png",
        supplier: "Pizza King",
        supplierURL: "/",
        items: [
            { count: 1, name: "Popeye Pizza", price: 3500 },
            { count: 25, name: "Popeye Pizza", price: 3500 },
            { count: 10, name: "Popeye Pizza", price: 7500 }
        ],
        deliveryFee: 799,
        deliveredAt: "Apr. 14. 18:35",
        address: "1039 Budapest Szamos utca 8",
        paymentMethod: { img: "@/assets/Otp_bank_logo.svg.png", name: "OTP" }
    },
    {
        img: "@/assets/Otp_bank_logo.svg.png",
        supplier: "Pizza King",
        supplierURL: "/",
        items: [
            { count: 1, name: "Popeye Pizza", price: 3500 },
            { count: 25, name: "Popeye Pizza", price: 3500 },
            { count: 10, name: "Popeye Pizza", price: 70000 }
        ],
        deliveryFee: 799,
        deliveredAt: "Apr. 14. 18:35",
        address: "1039 Budapest Szamos utca 8",
        paymentMethod: { img: "@/assets/Otp_bank_logo.svg.png", name: "OTP" }
    },
    {
        img: "@/assets/Otp_bank_logo.svg.png",
        supplier: "Pizza King",
        supplierURL: "/",
        items: [
            { count: 1, name: "Pizza", price: 3500 },
            { count: 25, name: "Pizza", price: 3500 },
            { count: 10, name: "Pizza", price: 7500 }
        ],
        deliveryFee: 799,
        deliveredAt: "Apr. 14. 18:35",
        address: "1039 Budapest Szamos utca 8",
        paymentMethod: { img: "@/assets/Otp_bank_logo.svg.png", name: "OTP" }
    }
]*/

const orders = ref([] as orderInfo[])

watchEffect(async () => {
    orders.value = (await getOrders()).orders
})
</script>

<template>
    <Header/>
    <main>
        <section>
            <OrderCard v-for="o in orders" :order="o"/>
        </section>
    </main>
</template>

<style scoped>
main {
    position: relative;
    display: grid;
    width: 100%;
    height: 100%;
}

section {
    position: relative;
    width: 30%;
    margin: auto;
}

@media only screen and (hover: none) and (pointer: coarse) {
    section {
        width: 100%;
    }
}
</style>