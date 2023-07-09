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
@import url("@/assets/authForm.css");
</style>