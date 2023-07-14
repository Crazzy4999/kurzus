<script setup lang="ts">
export type deliveryInfo = {
    description: string
    text: string
    URL?: string
}

const props = defineProps<{
    deliveryInfos: deliveryInfo[]
}>()
</script>

<template>
    <div class="container">
        <div v-for="d, i in deliveryInfos" class="inner-container">
            <span class="texts-container">
                <span class="img-wrapper">
                    <!--Icon made by https://www.flaticon.com/authors/freepik-->
                    <img src="@/assets/icons/maps-and-flags.png" alt="marker">
                </span>
                <div class="description-text">{{ d.description }}</div>
            </span>
            <div v-if="d.URL === undefined" :class="i < props.deliveryInfos.length-1 ? 'text path' : 'text'">{{ d.text }}</div>
            <RouterLink v-if="d.URL !== undefined" :to="d.URL" :class="i < props.deliveryInfos.length-1 ? 'link-text path' : 'link-text'">{{ d.text }}</RouterLink>
        </div>
    </div>
</template>

<style scoped>
.container {
    display: flex;
    flex-direction: column;
    border-top: var(--sub-border-size) solid var(--settings-color);
    padding-top: var(--sub-p-size);
    margin-block: var(--sub-p-size) var(--tiny-size);
}

.inner-container {
    display: flex;
    flex-direction: column;
    padding-bottom: var(--tiny-size);
}

.img-wrapper {
    width: calc(var(--sub-p-size) * 1.125);
    height: calc(var(--sub-p-size) * 1.125);
}

img {
    background-color: var(--main-color);
    width: 100%;
}

.texts-container {
    display: flex;
    font-size: var(--sub-p-size);
}

.description-text {
    color: var(--settings-color);
    margin-bottom: var(--tiny-size);
    margin-left: var(--tiny-size);
}

.text {
    font-size: calc(var(--sub-p-size) * 1.25);
    width: fit-content;
    padding-left: calc(var(--tiny-size) * 2 + 1px);
    margin-left: var(--tiny-size);
}

.path {
    border-left: var(--border-size) dotted var(--second-color);
    padding-left: calc(var(--tiny-size) * 2 - 1px);
}

.link-text {
    font-size: calc(var(--sub-p-size) * 1.25);
    color: var(--second-color);
    width: fit-content;
    padding-left: calc(var(--tiny-size) * 2 + 1px);
    margin-left: var(--tiny-size);
    transition: color var(--tran);
}

.link-text:hover {
    color: var(--second-hover);
}
</style>