<script setup lang="ts">
import { getOrders } from '@/api/api';
import type { orderInfo } from '@/api/models';
import { orderStatus } from '@/api/util';
import Header from '@/components/Header.vue';
import OrderCard from "@/components/history/OrderCard.vue"
import { ref, watchEffect } from 'vue';

const orders = ref([] as orderInfo[])

watchEffect(async () => {
    orders.value = (await getOrders()).orders
})
</script>

<template>
    <Header/>
    <main>
        <section>
            <h2>Inprogess orders</h2>
            <span v-for="o in orders">
                <OrderCard v-if="o.statusID === orderStatus.CREATED || o.statusID === orderStatus.DELIVERING" :order="o"/>
            </span>
            <h2>Previous orders</h2>
            <span v-for="o in orders">
                <OrderCard v-if="o.statusID === orderStatus.DONE" :order="o"/>
            </span>
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
    margin-inline: auto;
}

h2 {
    border-bottom: var(--border-size) solid var(--second-color);
    font-size: var(--h3-size);
    color: var(--second-color);
    margin-block: var(--p-size);
}

@media only screen and (hover: none) and (pointer: coarse) {
    section {
        width: 100%;
    }

    h2 {
        padding-inline: var(--sub-p-size);
    }
}
</style>