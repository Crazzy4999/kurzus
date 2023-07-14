<script setup lang="ts">
import type { categoryInfo } from '@/api/models';
import CategorieCheckbox from '@/components/categories/CategorieCheckbox.vue';
import { ref, watch } from 'vue';

defineProps<{
    categories: categoryInfo[]
}>()

const filters = ref([] as categoryInfo[])

function processFilter(val: categoryInfo) {
    if(val.name.includes("remove")) {
        const categoryName = val.name.replace("remove", "")
        filters.value.forEach((filter, i) => {
            if(filter.name === categoryName) filters.value.splice(i, 1)
        })
    } else filters.value.push(val)
}

const emit = defineEmits<{
    (event: "filters", val: categoryInfo[]): void
}>()

watch(filters.value, () => {
    emit("filters", filters.value)
})

let isOpened = ref(false)
</script>

<template>
    <div class="filter-container">    
        <span class="categories-header">Filters</span>
        <div class="categories-container-wrapper" :opened="isOpened">
            <ul class="categories-container">
                <CategorieCheckbox v-for="c in categories" :categorie="c" @filter="(val) => processFilter(val)"/>
            </ul>
        </div>
        <span class="categories-container-toggle" @click="isOpened = !isOpened">
            <!--Icon by Sabr Studio - https://www.iconfinder.com/sabrstudio-->
            <img class="categories-container-arrow" src="@/assets/icons/4781845_arrows_bottom_chevron_direction_move_icon.svg" alt="dropdown arrow">
        </span>
    </div>
</template>

<style scoped>
@import url("@/assets/categories.css");
</style>