<script setup lang="ts">
import type { itemInfo } from '@/api/models';

defineProps<{
    items: itemInfo[]
}>()
</script>

<template>
    <div class="main-container">
        <div v-for="i in items" class="container">
            <span class="inner-container">
                <span class="text common">{{ i.count }}</span>
                <span class="times common">x</span>
                <span class="text common">{{ i.name }}</span>
            </span>
            <span class="price common">{{ i.price }} Ft</span>
        </div>
    </div>
</template>

<style scoped>
.main-container {
    display: flex;
    flex-direction: column;
    border-top: var(--sub-border-size) solid var(--settings-color);
    padding-block: var(--tiny-size);
}

.container {
    display: flex;
    justify-content: space-between;
    margin-inline: var(--border-size);
}

.inner-container {
    display: flex;
    width: 70%;
}

.common {
    font-size: calc(var(--sub-p-size) * 1.25);
    margin-block: calc(var(--tiny-size) * 1.25);
}

.text {
    font-weight: bold;
}

.times {
    font-weight: bold;
    margin-inline: var(--tiny-size);
}

.price {
    font-weight: 450;
}

@media only screen and (hover: none) and (pointer: coarse) {
    
}
</style>