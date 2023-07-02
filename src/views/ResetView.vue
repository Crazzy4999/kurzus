<script setup lang="ts">
import AuthHeader from "@/components/AuthHeader.vue"
</script>

<template>
    <AuthHeader/>
    <main>
        <form>
            <input type="password" placeholder="password">
            <input type="password" placeholder="password again">
            <div class="error-msg">TEST ERROR MESSAGE THAT INTENTIONALLY OVERFLOWS TO TEST IT</div>
            <button @click.prevent="">Reset Password</button>
        </form>
    </main>
</template>

<style scoped>
    main {
        display: grid;
        width: 100%;
        height: calc(100% - var(--h1-size) * 1.25 - var(--sub-border-size) * 2);
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

    .error-msg {
        text-align: center;
        font-size: var(--sub-p-size);
        color: var(--error-color);
        width: 75%;
        margin: var(--p-size) auto;
    }

    button {
        cursor: pointer;
        font-size: var(--p-size);
        border-radius: var(--p-size);
        background-color: var(--second-color);
        color: var(--main-color);
        width: 45%;
        padding: var(--p-size);
        margin: var(--p-size) auto;
        transition: background-color var(--tran);
    }

    button:hover {
        background-color: var(--second-hover);
    }

    @media only screen and (hover: none) and (pointer: coarse) {
        form {
            width: 75%;
        }

        .error-msg {
            font-size: var(--p-size);
        }

        button {
            width: 65%;
        }
    }
</style>