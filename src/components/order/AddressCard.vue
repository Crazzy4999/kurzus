<script setup lang="ts">
import type { addressInfo } from '@/views/OrderView.vue';

defineProps<{
    address: addressInfo
    checked?: boolean
}>()
</script>

<template>
    <li class="address-list-item">
        <input class="address-radio" type="radio" name="selected-address" id="" :checked="checked">
        <span class="address-container">
            <div class="address-text">{{ address.street }} {{ address.houseNumber }} {{ address.floorNumber }} {{ address.apartment }}</div>
            <div class="address-text">{{ address.zipCode }} {{ address.city }}</div>
        </span>
    </li>
</template>

<style scoped>
.address-list-item {
    cursor: pointer;
    display: flex;
    border: var(--sub-border-size) solid var(--second-color);
    padding-block: var(--sub-p-size);
    margin-block: var(--sub-p-size);
    transition: box-shadow var(--tran);
}

.address-list-item:hover {
    box-shadow: 0 0 var(--tiny-size) 1px var(--second-color);
}

.address-radio {
    cursor: pointer;
    display: grid;
    place-content: center;
    -webkit-appearance: none;
    appearance: none;
    width: var(--p-size);
    height: var(--p-size);
    border: var(--border-size) solid var(--second-color);
    border-radius: 50%;
    margin: 0;
    margin-left: var(--tiny-size);
    transition: border var(--tran);
}

.address-radio::before {
    content: "";
    border-radius: 50%;
    width: calc(var(--sub-p-size) * 0.8);
    height: calc(var(--sub-p-size) * 0.8);
    background-color: var(--second-color);
    transform: scale(0);
    transition: transform var(--tran);
}

.address-radio:checked::before {
    transform: scale(1);
}

.address-radio:hover, .address-radio:checked {
    border: var(--border-size) solid var(--second-hover);
}

.address-container {
    margin-left: var(--tiny-size);
}

.address-text {
    font-size: var(--sub-p-size);
    color: var(--second-hover);
}

@media only screen and (hover: none) and (pointer: coarse) {
    .address-text {
        font-size: calc(var(--p-size) * 0.85);
    }
}
</style>