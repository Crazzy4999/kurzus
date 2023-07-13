<script setup lang="ts">
import { getSuppliers } from "@/api/api";
import type { supplierInfo } from "@/api/models";
import { supplierType, type query } from "@/api/util";
import SearchBar from "@/components/searchbar/SearchBar.vue"
import Supplier from "@/components/suppliers/Supplier.vue"
import { reactive, ref, watchEffect } from "vue";

const showIf = ref({
    search: "",
    categories: [],
    type: supplierType.ALL
} as query)
const suppliers = reactive({ values: [] as supplierInfo[] })

watchEffect(async () => {
    suppliers.values = (await getSuppliers()).suppliers
})
</script>

<template>
    <main class="main-container">
        <SearchBar placeholder="Search"/>

        <nav class="suppliers-toggle-container">
            <label class="supplier-item-wrapper" for="restaurants" @click="showIf.type = supplierType.RESTAURANT">
                <input class="supplier-toggle" type="radio" name="supplier" id="restaurants">
                <span class="supplier-item-name">Restaurants</span>
            </label>
            <label class="supplier-item-wrapper" for="cafes" @click="showIf.type = supplierType.COFFEE_SHOP">
                <input class="supplier-toggle" type="radio" name="supplier" id="cafes">
                <span class="supplier-item-name">Cafés</span>
            </label>
            <label class="supplier-item-wrapper" for="markets" @click="showIf.type = supplierType.SUPERMARKET">
                <input class="supplier-toggle" type="radio" name="supplier" id="markets">
                <span class="supplier-item-name">Markets</span>
            </label>
            <label class="supplier-item-wrapper" for="all" @click="showIf.type = supplierType.ALL">
                <input class="supplier-toggle" type="radio" name="supplier" id="all" checked>
                <span class="supplier-item-name">All</span>
            </label>
        </nav>
        <section class="content-container">
            <ul class="suppliers-container">
                <Supplier v-for="s in suppliers.values" :showIf="showIf" :supplier="s"/>
            </ul>
        </section>
    </main>
</template>

<style scoped>
@import url("@/assets/contentLayout.css");

.suppliers-toggle-container {
    display: flex;
    position: sticky;
    flex-wrap: wrap;
    height: fit-content;
    border: solid var(--border-size) var(--second-hover);
    border-radius: var(--p-size);
    background-color: var(--main-color);
    width: calc(100% - var(--sub-p-size) * 2 - var(--border-size) * 2);
    margin: var(--p-size) auto;
}

.supplier-item-wrapper {
    display: grid;
    flex: 0 0 33.3333%;
    flex-grow: 1;
    width: 100%;
    height: 100%;
    margin-inline: auto;
}

.supplier-toggle {
    display: none;
}

.supplier-item-name {
    cursor: pointer;
    text-align: center;
    font-size: calc(var(--p-size) * 0.9);
    font-weight: bold;
    line-height: var(--h3-size);
    border: solid var(--border-size)  var(--second-hover);
    border-radius: var(--h3-size);
    background-color: var(--main-color);
    color: var(--second-color);
    width: 65%;
    margin: var(--p-size) auto;
    padding-inline: var(--p-size);
    transition: background-color var(--tran), color var(--tran);
}

.supplier-toggle:checked + .supplier-item-name {
    background-color: var(--second-hover);
    color: var(--main-color);
}



.suppliers-container {
    display: flex;
    flex-wrap: wrap;
}

@media only screen and (hover: none) and (pointer: coarse) {
    .suppliers-toggle-container {
        z-index: 2;
        position: -webkit-sticky;
        display: flex;
        position: sticky;
        flex-wrap: wrap;
        border: none;
        border-block: solid var(--border-size) var(--second-hover);
        border-radius: unset;
        width: 100%;
        top: 0;
    }
}
</style>