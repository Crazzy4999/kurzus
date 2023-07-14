<script setup lang="ts">
import { ref, watch, watchEffect } from 'vue';

defineProps<{
    placeholder: string
}>()

const emit = defineEmits<{
    (event: "search", val: string): void
}>()

const search = ref("")

watchEffect(() => {
    emit("search", search.value.toLowerCase())
})
</script>

<template>
    <span class="searchbar-container">
        <span class="searchbar-btn">
            <!--Icon by Ionicons - https://www.iconfinder.com/iconsets/ionicons-->
            <img class="searchbar-img" src="@/assets/icons/211817_search_strong_icon.svg" alt="search">
        </span>
        <input class="searchbar" type="text" :placeholder="placeholder" v-model="search" @click="emit('search', search)">
    </span>
</template>

<style scoped>
.searchbar-container {
    display: flex;
    flex-direction: row-reverse;
    width: calc(100% - var(--sub-p-size) * 2);
    margin: var(--p-size) auto;
}

.searchbar {
    position: relative;
    font-size: var(--h4-size);
    border: solid var(--border-size) var(--second-color);
    color: var(--second-color);
    width: 100%;
    padding: var(--tiny-size) var(--sub-p-size);
    transition: border var(--tran);
}

.searchbar-btn {
    cursor: pointer;
    display: grid;
    background-color: var(--second-color);
    width: calc(var(--h4-size) * 2 + var(--border-size) * 2 + var(--sub-p-size));
    height: 100%;
    transition: background-color var(--tran);
}

.searchbar-btn:hover {
    background-color: var(--second-hover);
}

.searchbar-btn:hover + .searchbar {
    border: solid var(--border-size) var(--second-hover);
}

.searchbar-img {
    width: 100%;
    height: 100%;
}
</style>