<script setup lang="ts">
import Counter from "@/components/cart/Counter.vue"
import { ref, watch } from "vue";

const props = defineProps<{
    name: string
    price: number
}>()

const emit = defineEmits<{
    (event: "total", val: number): void
}>()

let count = ref(0)

watch(count, (newVal, oldVal) => {
    emit("total", newVal * props.price - oldVal * props.price)
})
</script>

<template>
    <div class="product">
        <span class="infos-container">
            <span class="name">{{ name }}</span>
            <div class="price">{{ price * count }} Ft</div>
        </span>
        <Counter class="counter" @count="(val) => count = val"/>
    </div>
</template>

<style scoped>
.product {
    display: flex;
    flex-direction: column;
    gap: var(--tiny-size);
    border-top: var(--sub-border-size) solid var(--settings-color);
    padding-top: var(--tiny-size);
    margin-top: var(--p-size);
}

.name {
    display: -webkit-box;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
    overflow: hidden;
    text-overflow: ellipsis;
    font-size: var(--p-size);
    word-break: break-all;
    color: var(--second-color);
    width: 65%;
}

.infos-container {
    display: flex;
    justify-content: space-between;
}

.price {
    text-align: right;
    font-size: var(--p-size);
    color: var(--settings-color);
}

.counter {
    margin-left: auto;
}
</style>