<script setup lang="ts">
import { ref } from 'vue';

export type inputVal = { name: string, type: string, placeholder: string, required: boolean }

const props = defineProps<{
    action?: string
    method?: string
    inputs: inputVal[]
    infoText: string
    btnText: string
    alternateText?: string
    alternateUrl?: string
}>()

let errorMsg = ref("")
let alternateUrlString = ref(props.alternateUrl as string)
</script>

<template>
    <div>
        <form :action="action" :method="method">
            <input v-for="i in props.inputs" class="form-auth-input" :name="i.name" :type="i.type" :placeholder="i.placeholder" :required="i.required">
            <span class="error-msg-box">{{ errorMsg }}</span>
            <span class="info-text">{{ infoText }}</span>
            <RouterLink v-if="alternateText !== undefined && alternateUrl !== undefined" class="auth-nav-link" :to="alternateUrlString">{{ alternateText }}</RouterLink>
            <button>{{ btnText }}</button>
        </form>
    </div>
</template>

<style scoped>
div {
    display: flex;
    height: calc(100% - var(--h1-size) - var(--sub-border-size));
}

form {
    display: flex;
    flex-direction: column;
    border: solid var(--sub-border-size) var(--settings-color);
    border-radius: var(--h3-size);
    box-shadow: 0 0 var(--border-size) calc(var(--sub-border-size) * 0.5) var(--settings-color);
    width: 25%;
    height: min-content;
    margin: auto;
    padding-block: var(--h4-size);
}

input {
    font-size: var(--h4-size);
    border: solid var(--border-size) var(--second-color);
    border-radius: var(--h3-size);
    color: var(--second-color);
    width: calc(90% - var(--border-size) * 2 - var(--p-size) * 2 - var(--h4-size) * 2);
    margin: var(--sub-p-size) calc(5% + var(--h4-size));
    padding: var(--tiny-size) var(--p-size);
    transition: box-shadow var(--tran);
}

input:hover {
    box-shadow: 0 0 var(--border-size) calc(var(--sub-border-size) * 0.25) var(--second-hover);
}

.error-msg-box {
    text-align: center;
    font-size: var(--p-size);
    color: var(--error-color);
    width: 75%;
    margin: var(--border-size) auto;
}

.info-text {
    text-align: center;
    font-size: var(--p-size);
    color: var(--settings-dark);
    width: 70%;
    margin: var(--sub-p-size) auto;
}

.auth-nav-link {
    text-align: center;
    font-size: var(--p-size);
    font-weight: bold;
    color: var(--second-color);
    width: 70%;
    margin-block: 0 var(--p-size);
    margin-inline: auto;
}

.auth-nav-link:hover {
    color: var(--second-hover);
}

button {
    cursor: pointer;
    font-size: var(--h4-size);
    border: solid var(--border-size) var(--second-color);
    border-radius: var(--h3-size);
    color: var(--second-color);
    width: fit-content;
    margin: var(--sub-p-size) auto;
    padding: var(--tiny-size) var(--h4-size);
    transition: border var(--tran), background-color var(--tran), color var(--tran), box-shadow var(--tran);
}

button:hover {
    border: solid var(--border-size) var(--second-hover);
    background-color: var(--second-hover);
    color: var(--main-color);
    box-shadow: 0 0 var(--border-size) calc(var(--border-size) * 0.25) var(--second-hover);
}

@media only screen and (hover: none) and (pointer: coarse) {
    div {
        height: calc(100% - var(--h1-size) - var(--border-size) * 2);
    }

    form {
        width: 75%;
    }
}
</style>