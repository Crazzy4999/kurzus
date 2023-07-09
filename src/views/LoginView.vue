<script setup lang="ts">
import AuthHeader from '@/components/AuthHeader.vue';
import { ref } from 'vue';

const errMsg = ref("")
</script>

<template>
    <AuthHeader/>
    <div class="form-wrapper">
        <form class="auth-form">
            <input class="form-auth-input" name="email" type="email" placeholder="email" :required="true">
            <input class="form-auth-input" name="password" type="password" placeholder="password" :required="true">
            <span class="error-msg-box">{{ errMsg }}</span>
            <span class="info-text">Don't have an account yet?</span>
            <RouterLink class="auth-nav-link" to="/signup">Sign up now!</RouterLink>
            <button @click.prevent="">Login</button>
        </form>
    </div>
</template>

<style scoped>
@import url("@/assets/authForm.css");
</style>