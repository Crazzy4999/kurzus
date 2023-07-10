<script setup lang="ts">
import { getProfile, login } from '@/api/api';
import AuthHeader from '@/components/AuthHeader.vue';
import router from '@/router';
import { useAuthStore } from '@/store';
import { ref } from 'vue';

const useAuth = useAuthStore()

const email = ref("")
const password = ref("")

const errMsg = ref("")

function loginUser(email: string, password: string) {
    if(email !== "" && password !== "") {
        login(email, password).then(resp => {
            useAuth.setTokens(resp.accessToken, resp.refreshToken)
            getProfile().then(resp => useAuth.setUser(resp.id, resp.firstName, resp.lastName, resp.email))
            router.push("/")
        }).catch(err => errMsg.value = err)
    } else errMsg.value = "Please fill out all the fields!"
    
}
</script>

<template>
    <AuthHeader/>
    <div class="form-wrapper">
        <form class="auth-form">
            <input class="form-auth-input" name="email" type="email" placeholder="email" :required="true" v-model="email">
            <input class="form-auth-input" name="password" type="password" placeholder="password" :required="true" v-model="password">
            <span class="error-msg-box">{{ errMsg }}</span>
            <span class="info-text">Forgot your password?</span>
            <RouterLink class="auth-nav-link" to="/forgot">Reset password!</RouterLink>
            <span class="info-text">Don't have an account yet?</span>
            <RouterLink class="auth-nav-link" to="/signup">Sign up now!</RouterLink>
            <button @click.prevent="loginUser(email, password)">Login</button>
        </form>
    </div>
</template>

<style scoped>
@import url("@/assets/authForm.css");
</style>