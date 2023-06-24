<script setup lang="ts">
export type categorieVal = {
    id: string,
    text: string
}

defineProps<{
    categories: categorieVal[]
}>()
</script>

<template>
    <div class="sub-categories-container">
        <span class="categories-arrow-wrapper" left>
            <img class="sub-categories-arrow sub-categories-larrow" src="@/assets/icons/4781845_arrows_bottom_chevron_direction_move_icon.svg" alt="scroll left">
        </span>
        <ul class="restaurant-sub-categories">
            <li v-for="c in categories" class="sub-categorie-item">
                <a class="sub-categorie-link selected" :href="'#'+c.id">{{ c.text }}</a>
            </li>
        </ul>
        <span class="categories-arrow-wrapper" right>
            <img class="sub-categories-arrow sub-categories-rarrow" src="@/assets/icons/4781845_arrows_bottom_chevron_direction_move_icon.svg" alt="scroll right">
        </span>
    </div>
</template>

<style scoped>
@import url("@/assets/supplier.css");
</style>