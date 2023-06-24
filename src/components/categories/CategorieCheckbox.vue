<script setup lang="ts">
import type { categorie } from '@/components/categories/CategoriesFilter.vue';
import { ref } from 'vue';

defineProps<{
    categorie: categorie
}>()

let checked = ref(false)
</script>

<template>
    <li class="categorie-wrapper" :checked="checked">
        <input class="categorie-chbx" type="checkbox" :name="categorie.id" :id="categorie.id">
        <label class="categorie-name" :for="categorie.id" @click="checked = !checked">{{ categorie.name }}</label>
    </li>
</template>

<style scoped>
.categorie-chbx {
    display: none;
}

.categorie-name {
    cursor: pointer;
    font-size: var(--p-size);
    color: var(--second-color);
    transition: color var(--tran);
}

.categorie-wrapper {
    cursor: pointer;
    display: flex;
    border: solid var(--border-size) var(--second-color);
    border-radius: var(--p-size);
    margin: var(--sub-p-size);
    padding: var(--tiny-size) var(--sub-p-size);
    transition: background-color var(--tran);
}

.categorie-wrapper[checked="true"] {
    background-color: var(--second-color);
}

.categorie-wrapper[checked="true"] .categorie-name {
    color: var(--main-color);
}
</style>