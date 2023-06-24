<script setup lang="ts">
import { ref } from "vue"

let count = ref(0)

function decrement() {
    if(0 <= count.value - 1) {
        count.value--
    }
}

function restrictInput(e: KeyboardEvent) {
    if(!/\d/.test(e.key)) return e.preventDefault()
}
</script>

<template>
    <div class="counter-container">
        <button class="decrement common" @click="decrement">-</button>
        <span class="count-wrapper common">
            <input class="count" type="text" :value="count" @keypress="(e) => restrictInput(e as KeyboardEvent)">
        </span>
        <button class="increment common" @click="count++">+</button>
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
    border: calc(var(--sub-border-size) * 1.25) solid var(--settings-color)
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
</style>