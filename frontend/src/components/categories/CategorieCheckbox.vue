<script setup lang="ts">
import type { categoryInfo } from '@/api/models';
import { ref } from 'vue';

const props = defineProps<{
    categorie: categoryInfo
}>()

const emit = defineEmits<{
    (event: "filter", val: categoryInfo): void
}>()

const checked = ref(false)

function toggle() {
    checked.value = !checked.value
    if(checked.value) emit("filter", props.categorie)
    else emit("filter", {name: "remove"+props.categorie.name} as categoryInfo)
}
</script>

<template>
    <li class="categorie-wrapper" :checked="checked">
        <input class="categorie-chbx" type="checkbox" :name="categorie.id+''" :id="categorie.id+''">
        <label class="categorie-name" :for="categorie.id+''" @click="toggle()">{{ categorie.name }}</label>
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