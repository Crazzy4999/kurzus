<script setup lang="ts">
import { signUp } from '@/api/api';
import AuthHeader from '@/components/AuthHeader.vue';
import router from '@/router';
import { ref } from 'vue';

const firstName = ref("")
const lastName = ref("")
const email = ref("")
const password = ref("")
const passwordAgain = ref("")

const errMsg = ref("")

function signUpUser(firstName: string, lastName: string, email: string, password: string, passwordAgain: string) {
    if(firstName !== "" && lastName !== "" && email !== "" && password !== "" && passwordAgain !== "") {
        signUp(firstName, lastName, email, password, passwordAgain).then(resp => router.push("/")).catch(err => errMsg.value = err)
    } else errMsg.value = "Please fill out all the fields!"
}
</script>

<template>
    <AuthHeader/>
    <div class="form-wrapper">
        <form class="auth-form">
            <input class="form-auth-input" name="firstName" type="text" placeholder="first name" :required="true" v-model="firstName">
            <input class="form-auth-input" name="lastName" type="text" placeholder="last name" :required="true" v-model="lastName">
            <input class="form-auth-input" name="email" type="email" placeholder="email" :required="true" v-model="email">
            <input class="form-auth-input" name="password" type="password" placeholder="password" :required="true" v-model="password">
            <input class="form-auth-input" name="passwordAgain" type="password" placeholder="password again" :required="true" v-model="passwordAgain">
            <span class="error-msg-box">{{ errMsg }}</span>
            <span class="info-text">Already have an account?</span>
            <RouterLink class="auth-nav-link" to="/login">Log in!</RouterLink>
            <button @click.prevent="signUpUser(firstName, lastName, email, password, passwordAgain)">Sign Up</button>
        </form>
    </div>
</template>

<style scoped>
@import url("@/assets/authForm.css");
</style>