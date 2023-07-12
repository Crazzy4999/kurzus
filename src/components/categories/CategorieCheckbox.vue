<script setup lang="ts">
import type { categoryInfo } from '@/api/models';
import { ref } from 'vue';

defineProps<{
    categorie: categoryInfo
}>()

let checked = ref(false)
</script>

<template>
    <li class="categorie-wrapper" @click="checked = !checked" :checked="checked">
        <input class="categorie-chbx" type="checkbox" :name="categorie.id+''" :id="categorie.id+''">
        <label class="categorie-name" :for="categorie.id+''" @click="checked = !checked">{{ categorie.name }}</label>
    </li>
</template>

<style scoped>
.categorie-chbx {
    display: none;
}

.categorie-name {
    user-select: none;
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