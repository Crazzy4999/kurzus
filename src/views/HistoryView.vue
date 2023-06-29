<script setup lang="ts">
import Header from '@/components/Header.vue';
import OrderCard, { type orderInfo } from "@/components/history/OrderCard.vue"

let orders: orderInfo[] = [
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
]
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