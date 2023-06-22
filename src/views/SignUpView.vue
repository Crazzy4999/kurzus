<script setup lang="ts">
import AuthHeader from '@/components/AuthHeader.vue';
import AuthForm, { type inputVal } from '@/components/form/AuthForm.vue';

let inputs: inputVal[] = [
    { name: "firstName", type: "text", placeholder: "first name", required: true },
    { name: "lastName", type: "text", placeholder: "last name", required: true },
    { name: "email", type: "email", placeholder: "email", required: true },
    { name: "password", type: "password", placeholder: "password", required: true },
    { name: "passwordAgain", type: "password", placeholder: "password again", required: true }
]
</script>

<template>
    <AuthHeader/>
    <AuthForm :inputs="inputs" infoText="Already have an account?" btnText="Sign Up" alternateText="Log in!" alternateUrl="/login"/>
</template>

<style scoped>

</style>