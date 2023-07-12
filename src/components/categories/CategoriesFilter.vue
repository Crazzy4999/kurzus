<script setup lang="ts">
import type { categoryInfo } from '@/api/models';
import CategorieCheckbox from '@/components/categories/CategorieCheckbox.vue';
import { ref } from 'vue';

defineProps<{
    categories: categoryInfo[]
}>()

let isOpened = ref(false)
</script>

<template>
    <div class="filter-container">    
            <span class="categories-header">Filters</span>
            <div class="categories-container-wrapper" :opened="isOpened">
                <ul class="categories-container">
                    <CategorieCheckbox v-for="c in categories" :categorie="c"/>
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