<script setup lang="ts">
import type { productInfo } from '@/api/models';

defineProps<{
    product: productInfo
}>()
</script>

<template>
    <div class="main-container">
        <div class="info-flex">
            <span class="number">{{ product.count }}</span>
            <span class="times">x</span>
            <span class="name">{{ product.name }}</span>
        </div>
        <div class="price">{{ product.price }} Ft</div>
    </div>
</template>

<style scoped>
.main-container {
    display: flex;
    flex-direction: column;
    margin-block: var(--tiny-size);
}

.info-flex {
    display: flex;
    font-size: var(--sub-p-size);
}

.number {
    font-weight: bold;
}

.times {
    font-weight: bold;
    margin-inline: var(--tiny-size);
}

.name {
    color: var(--settings-dark);
    word-wrap: break-word;
}

.price {
    color: var(--second-color);
    font-size: var(--sub-p-size);
    margin-block: var(--tiny-size);
}

@media only screen and (hover: none) and (pointer: coarse) {
    .info-flex {
        font-size: var(--p-size);
    }
    
    .price {
        font-size: var(--p-size);
    }
}
</style>