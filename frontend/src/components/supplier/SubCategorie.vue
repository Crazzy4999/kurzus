<script setup lang="ts">
import type { categoryInfo, menuInfo } from '@/api/models'
import SupplierProduct from '@/components/supplier/SupplierProduct.vue';

defineProps<{
    categories: categoryInfo[]
    menus: menuInfo[]
}>()
</script>

<template>
    <div class="sub-categorie-container" v-for="c in categories">
        <h2 class="sub-categorie-header" :id="c.id+''">{{ c.name }}</h2>
        <div v-for="m in menus" class="wrapper-for-v-for">
            <SupplierProduct :categoryID="c.id" :menu="m"/>
        </div>
    </div>
</template>

<style scoped>
@import url("@/assets/supplier.css");

.sub-categorie-container {

}

.sub-categorie-header {
    font-size: calc(var(--h4-size) * 1.25);
    color: var(--second-color);
    scroll-margin-top: calc(var(--h1-size) + var(--border-size) * 2 + var(--h4-size));
}

.sub-categorie-header::before {
    content: "";
    z-index: 0;
    position: relative;
    display: block;
    height: var(--border-size);
    background-color: var(--settings-color);
    opacity: 0.5;
    margin-top: var(--h4-size);
}

@media only screen and (hover: none) and (pointer: coarse) {
    .sub-categorie-header {
        scroll-margin-top: var(--h2-size);
    }
}
</style>