<script setup lang="ts">
import AuthHeader from '@/components/AuthHeader.vue';
import AuthForm, { type inputVal } from '@/components/form/AuthForm.vue';

let inputs: inputVal[] = [
    { name: "email", type: "email", placeholder: "email", required: true },
    { name: "password", type: "password", placeholder: "password", required: true }
]
</script>

<template>
    <AuthHeader/>
    <AuthForm :inputs="inputs" infoText="Don't have an account yet?" btnText="Login" alternateText="Sign up now!" alternateUrl="/signup"/>
</template>

<style scoped>

</style>