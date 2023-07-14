<script setup lang="ts">

</script>

<template>
    <header class="auth-header">
        <RouterLink class="nav-back-btn" to="/">
            <!--Icon by Kirill Kazachek - https://www.flaticon.com/authors/kirill-kazachek-->
            <img src="@/assets/icons/red-arrow.png" alt="Back">
        </RouterLink>
        <RouterLink class="auth-logo-wrapper" to="/">
            <!--<img class="logo" src="" alt="">-->
            <h3>Hangry</h3>
        </RouterLink>
    </header>
</template>

<style scoped>
    .auth-header {
        display: flex;
        width: 100%;
        height: var(--h1-size);
        border-bottom: solid var(--sub-border-size) var(--settings-color);
        box-shadow: 0 0 var(--border-size) var(--sub-border-size) var(--settings-color);
    }

    .nav-back-btn {
        cursor: pointer;
        width: var(--h2-size);
        height: fit-content;
        margin: auto var(--h3-size);
    }

    img {
        width: var(--h3-size);
        height: 100%;
        transition: filter var(--tran);
    }

    .auth-logo-wrapper {
        display: grid;
        font-size: var(--h3-size);
        width: 100%
    }

    H3 {
        width: fit-content;
        height: fit-content;
        color: var(--second-color);
        margin: auto;
        padding-right: calc(var(--h3-size) * 3);
    }
</style>