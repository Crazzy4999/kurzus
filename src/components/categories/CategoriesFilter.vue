<script setup lang="ts">
import CategorieCheckbox from '@/components/categories/CategorieCheckbox.vue';

export type categorie = {
    id: string,
    name: string
}

defineProps<{
    categories: categorie[]
}>()
</script>

<template>
    <div class="filter-container">    
            <span class="categories-header">Filters</span>
            <div class="categories-container-wrapper" opened="false">
                <ul class="categories-container">
                    <CategorieCheckbox v-for="c in categories" :categorie="c"/>
                </ul>
            </div>
            <span class="categories-container-toggle">
                <!--Icon by Sabr Studio - https://www.iconfinder.com/sabrstudio-->
                <img class="categories-container-arrow" src="@/assets/icons/4781845_arrows_bottom_chevron_direction_move_icon.svg" alt="dropdown arrow">
            </span>
        </div>
</template>

<style scoped>
@import url("@/assets/categories.css");
</style>