<script setup lang="ts">
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
                <ul class="categories-container" v-for="c in categories">
                    <li class="categorie-wrapper">
                        <input class="categorie-chbx" type="checkbox" :name="c.id" :id="c.id">
                        <label class="categorie-name" :for="c.id">{{ c.name }}</label>
                    </li>
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