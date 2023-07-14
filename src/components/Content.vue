<script setup lang="ts">

</script>

<template>
    <main>
        <section>
            <h1>Hangry</h1>
            <p>We all have been there, we are starving haven't eaten anything today in these times we tend to get angry, hence the name Hangry!</p>
            <p>No need to be Hangry find something you like and order now!</p>
            <p>Huge choice with lots of categories!</p>
        </section>
    </main>
</template>

<style scoped>
    section {
        width: 40%;
        margin-inline: auto;
    }

    h1 {
        font-size: var(--h1-size);
        color: var(--second-color);
        margin-block: var(--p-size);
    }
    
    p {
        font-size: var(--h4-size);
        color: var(--second-color);
        margin-block: var(--p-size);
    }

    @media only screen and (hover: none) and (pointer: coarse) {
        section {
            width: 80%
        }
    }
</style>