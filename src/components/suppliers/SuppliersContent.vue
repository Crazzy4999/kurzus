<script setup lang="ts">
import SearchBar from "@/components/searchbar/SearchBar.vue"
import Supplier, { type supplierVal} from "@/components/suppliers/Supplier.vue"

let suppliers: supplierVal[] = [

]
</script>

<template>
    <main class="main-container">
        <SearchBar placeholder="I want to eat..."/>

        <nav class="suppliers-toggle-container">
            <label class="supplier-item-wrapper" for="restaurants">
                <input class="supplier-toggle" type="radio" name="supplier" id="restaurants" checked>
                <span class="supplier-item-name">Restaurants</span>
            </label>
            <label class="supplier-item-wrapper" for="cafes">
                <input class="supplier-toggle" type="radio" name="supplier" id="cafes">
                <span class="supplier-item-name">Cafés</span>
            </label>
            <label class="supplier-item-wrapper" for="markets">
                <input class="supplier-toggle" type="radio" name="supplier" id="markets">
                <span class="supplier-item-name">Markets</span>
            </label>
        </nav>
        <section class="content-container">
            <ul class="suppliers-container">
                <Supplier v-for="s in suppliers" :name="s.name" :description="s.description" :img="s.img" :delivery-time="s.deliveryTime" :delivery-fee="s.deliveryFee"/>
            </ul>
        </section>
    </main>
</template>

<style scoped>
@import url("@/assets/contentLayout.css");

.suppliers-toggle-container {
    display: flex;
    height: var(--h1-size);
    border: solid var(--border-size) var(--second-hover);
    border-radius: var(--p-size);
    background-color: var(--main-color);
    width: calc(100% - var(--sub-p-size) * 2 - var(--border-size) * 2);
    margin: var(--p-size) auto;
}

.supplier-item-wrapper {
    display: grid;
    width: 100%;
    height: 100%;
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
    /*height: calc(100% - var(--h1-size)); MAY CHANGE BACK*/
}

@media only screen and (hover: none) and (pointer: coarse) {
    .suppliers-toggle-container {
        z-index: 2;
        position: -webkit-sticky;
        position: sticky;
        border: none;
        border-block: solid var(--border-size) var(--second-hover);
        border-radius: unset;
        width: 100%;
        top: 0;
    }
}
</style>