<script setup lang="ts">
import Header from "@/components/Header.vue"
import Sidebar from "@/components/Sidebar.vue"
import Content from "@/components/Content.vue"
import Footer from "@/components/Footer.vue"
import { ref } from "vue";

let toggleSidebar = ref(false)
</script>

<template>
    <Header @toggle="(val) => toggleSidebar = val"/>
    <Sidebar :toggle="toggleSidebar"/>
    <Content/>
    <Footer/>
</template>

<style scoped>

</style>