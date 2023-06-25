<script setup lang="ts">
import { ref, type InputHTMLAttributes, type ReservedProps } from "vue"

const emit = defineEmits<{
    (event: "count", count: number): void
}>()

let count = ref(0)

function decrement() {
    if(0 <= count.value - 1) {
        count.value--
    }
    emit("count", count.value)
}

function restrictInput(e: KeyboardEvent) {
    if(!/\d/.test(e.key) || e.key === "Enter") {
        emit("count", count.value)
        return e.preventDefault()
    }
}

function increment() {
    count.value++
    emit("count", count.value)
}
</script>

<template>
    <div class="counter-container">
        <button class="decrement common" @click.prevent="decrement">-</button>
        <span class="count-wrapper common">
            <input class="count" type="text" :value="count" @keypress="(e) => restrictInput(e)">
        </span>
        <button class="increment common" @click.prevent="increment">+</button>
    </div>
</template>

<style scoped>
.counter-container {
    position: relative;
    display: flex;
}

.common {
    font-size: var(--p-size);
    font-weight: 750;
    color: var(--second-color);
    border: calc(var(--sub-border-size) * 1.25) solid var(--settings-color);
    transition: color var(--tran);
}

button {
    position: relative;
    cursor: pointer;
    margin-block: auto;
}

.decrement {
    border-right: none;
    border-top-left-radius: var(--sub-p-size);
    border-bottom-left-radius: var(--sub-p-size);
    padding-inline: var(--sub-p-size) var(--tiny-size);
}

.count-wrapper {
    display: grid;
    width: var(--h3-size);
}

.count {
    text-align: center;
    font-size: var(--p-size);
    font-weight: 750;
    color: var(--second-color);
    max-width: var(--h3-size);
    margin: auto;
}

.increment {
    border-left: none;
    border-top-right-radius: var(--sub-p-size);
    border-bottom-right-radius: var(--sub-p-size);
    padding-inline: var(--tiny-size) var(--sub-p-size);
}

.decrement:hover, .increment:hover {
    color: var(--second-hover);
}
</style>