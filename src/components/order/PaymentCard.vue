<script setup lang="ts">
export type paymentInfo = {
    text: string,
    img: string
}

defineProps<{
    info: paymentInfo
}>()
</script>

<template>
    <li class="payment-card-container">
        <div class="payment-info-container">
            <input class="payment-radio" type="radio" name="payment-processor-select" id="">
            <span class="payment-text">{{ info.text }}</span>
        </div>
        <div class="img-wrapper">
            <img class="payment-img" src="@/assets/Otp_bank_Logo.svg.png" :alt="info.text">
        </div>
    </li>
</template>

<style scoped>
.payment-card-container {
    cursor: pointer;
    display: flex;
    justify-content: space-between;
    border: var(--sub-border-size) solid var(--second-color);
    padding: var(--tiny-size);
    transition: box-shadow var(--tran);
    margin-block: var(--sub-p-size);
}

.payment-card-container:hover {
    box-shadow: 0 0 var(--tiny-size) 1px var(--second-color);
}

.payment-info-container {
    display: flex;
}

.payment-radio {
    cursor: pointer;
    display: grid;
    place-content: center;
    -webkit-appearance: none;
    appearance: none;
    width: var(--p-size);
    height: var(--p-size);
    border: var(--border-size) solid var(--second-color);
    border-radius: 50%;
    margin-left: var(--tiny-size);
    margin-block: auto;
    transition: border var(--tran);
}

.payment-radio::before {
    content: "";
    border-radius: 50%;
    width: calc(var(--sub-p-size) * 0.8);
    height: calc(var(--sub-p-size) * 0.8);
    background-color: var(--second-color);
    transform: scale(0);
    transition: transform var(--tran);
}

.payment-radio:checked::before {
    transform: scale(1);
}

.payment-radio:hover, .payment-radio:checked {
    border: var(--border-size) solid var(--second-hover);
}

.payment-text {
    font-size: var(--sub-p-size);
    line-height: var(--h3-size);
    color: var(--second-hover);
    margin-left: var(--tiny-size);
}

.img-wrapper {
    width: var(--h4-size);
    height: 100%;
    border: var(--sub-border-size) solid var(--settings-dark);
    box-shadow: 0 0 var(--border-size) 0.1px var(--settings-dark);
    margin-block: auto;
}

.payment-img {
    width: 100%;
}

@media only screen and (hover: none) and (pointer: coarse) {
    .payment-text {
        font-size: calc(var(--p-size) * 0.8)
    }
}
</style>