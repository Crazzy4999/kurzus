<script setup lang="ts">
import AuthHeader from '@/components/AuthHeader.vue';
import { ref } from 'vue';

const errMsg = ref("")
</script>

<template>
    <AuthHeader/>
    <div class="form-wrapper">
        <form class="auth-form">
            <input class="form-auth-input" name="firstName" type="text" placeholder="first name" :required="true">
            <input class="form-auth-input" name="lastName" type="text" placeholder="last name" :required="true">
            <input class="form-auth-input" name="email" type="email" placeholder="email" :required="true">
            <input class="form-auth-input" name="password" type="password" placeholder="password" :required="true">
            <input class="form-auth-input" name="passwordAgain" type="password" placeholder="password again" :required="true">
            <span class="error-msg-box">{{ errMsg }}</span>
            <span class="info-text">Already have an account?</span>
            <RouterLink class="auth-nav-link" to="/login">Log in!</RouterLink>
            <button @click.prevent="">Sign Up</button>
        </form>
    </div>
</template>

<style scoped>
.form-wrapper {
    display: flex;
    height: calc(100% - var(--h1-size) - var(--sub-border-size));
}

.auth-form {
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
    .form-wrapper {
        height: calc(100% - var(--h1-size) - var(--border-size) * 2);
    }

    .auth-form {
        width: 75%;
    }
}
</style>