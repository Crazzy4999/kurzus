<script setup lang="ts">
import type { supplierInfo } from '@/api/models';

defineProps<{
    supplier: supplierInfo
}>()
</script>

<template>
    <li class="supplier">
        <RouterLink to="/suppliers"> <!--TODO ADD ROUTING TO SUPPLIERS-->
            <div class="supplier-img-container">
                <img class="supplier-img" :src="supplier.image" :alt="supplier.name">
                <span class="delivery-time-wrapper">
                    <span class="delivery-time">{{ supplier.deliveryTime }} min</span>
                </span>
            </div>
            <h2 class="supplier-name">{{ supplier.name }}</h2>
            <p class="supplier-description">{{ supplier.description }}</p>
            <span class="supplier-rates"></span>
            <span class="supplier-delivery-fee">{{ supplier.deliveryFee === 0 ? "Free delivery" : supplier.deliveryFee }}</span>
        </RouterLink>
    </li>
</template>

<style scoped>
.supplier {
    cursor: pointer;
    display: flex;
    flex-grow: 1;
    flex-direction: column;
    max-width: calc(25% - var(--sub-p-size));
    height: fit-content;
    margin: calc(var(--sub-p-size) / 2);
}

.supplier-img-container {
    position: relative;
    width: 100%;
    height: 50%;
}

.supplier-img {
    width: 100%;
    height: 100%;
}

/*FIX THIS on desktop view*/
.delivery-time-wrapper {
    overflow: hidden;
    position: absolute;
    font-size: var(--p-size);
    bottom: calc(var(--sub-p-size) / 4);
    right: 0;
    padding: var(--tiny-size) calc(var(--sub-p-size) * 0.75);
}

.delivery-time-wrapper::after {
    content: "";
    position: absolute;
    width: var(--p-size);
    height: var(--p-size);
    top: 0;
    left: 0;
    rotate: 45deg;
    margin: calc(-1 * var(--p-size) / 2);
    box-shadow: 0 0 0 calc(var(--h1-size) * 2) var(--second-color);
    transition: box-shadow var(--tran);
}

.supplier:hover .delivery-time-wrapper::after {
    box-shadow: 0 0 0 calc(var(--h1-size) * 2) var(--second-hover);
}

.delivery-time {
    z-index: 1;
    position: relative;
    font-size: var(--p-size);
    font-weight: bold;
    color: var(--main-color);
}

.supplier-name {
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    font-size: var(--h4-size);
    color: var(--second-color);
    transition: color var(--tran);
}

.supplier:hover .supplier-name {
    color: var(--second-hover);
}

.supplier-description {
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    font-size: var(--p-size);
    color: var(--second-hover)
}

.supplier-rates {

}

.supplier-delivery-fee {
    font-size: var(--sub-p-size);
    color: var(--settings-dark);
}

@media only screen and (hover: none) and (pointer: coarse) {
    .supplier {
        max-width: calc(100% - var(--sub-p-size));
        margin: var(--sub-p-size) calc(var(--sub-p-size) / 2);
    }

    .delivery-time {
        font-size: var(--h4-size);
    }
}
</style>