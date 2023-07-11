<script setup lang="ts">
import { getAddresses, updateAddress } from '@/api/api';
import type { addressInfo } from '@/api/models';
import { ref, watch, type Ref } from 'vue';

const props = defineProps<{
    address: addressInfo
}>()

function setToActive() {
    getAddresses().then(resp => {
        resp.addresses.forEach(address => {
            if(props.address.id !== address.id) updateAddress(address.id, address.userID, false, address.city, address.street, address.houseNumber, address.zipCode, address.floorNumber, address.apartment)
        })
    })
    updateAddress(props.address.id, props.address.userID, true, props.address.city, props.address.street, props.address.houseNumber, props.address.zipCode, props.address.floorNumber, props.address.apartment)
}

let editing = ref(false)
let deleting = ref(false)

function controlScroll(val: boolean) {
    window.scrollTo(0, 0)
    if(val) {
        window.onscroll = () => window.scrollTo(0, 0)
        document.body.classList.add("modal-open")
    } else {
        window.onscroll = () => {}
        document.body.classList.remove("modal-open")
    }
}

watch(editing, (newVal, oldVal) => {
    controlScroll(newVal)
})

watch(deleting, (newVal, oldVal) => {
    controlScroll(newVal)
})
</script>

<template>
    <li class="address-list-item" @click="setToActive()">
        <span class="separator-container">
            <input class="address-radio" type="radio" name="selected-address" :id="address.id+''" :checked="address.isActive">
            <span class="address-container">
                <div class="address-text">{{ address.street }} {{ address.houseNumber }} {{ address.floorNumber }} {{ address.apartment }}</div>
                <div class="address-text">{{ address.zipCode }} {{ address.city }}</div>
            </span>
        </span>
        <span class="separator-container">
            <button class="action-btn" @click.prevent="editing = !editing">
                <!--Icon made by https://www.flaticon.com/authors/freepik-->
                <img class="btn-img" src="@/assets/icons/editing.png" alt="edit">
            </button>
            <button class="action-btn" @click.prevent="deleting = !deleting">
                <!--Icon made by https://www.flaticon.com/authors/freepik-->
                <img class="btn-img" src="@/assets/icons/bin.png" alt="delete">
            </button>
        </span>
        <div v-if="editing" class="block-screen">
            <div class="modal-body">
                <h3 class="title">Edit address</h3>
                <input class="edit-input" type="text" placeholder="City">
                <input class="edit-input" type="text" placeholder="Street">
                <input class="edit-input" type="text" placeholder="House number">
                <input class="edit-input" type="text" placeholder="Zip code">
                <input class="edit-input" type="text" placeholder="Floor number">
                <input class="edit-input" type="text" placeholder="Apartment">
                <span class="btn-container">
                    <button class="modal-btn" @click.prevent="editing = !editing">Cancel</button>
                    <button class="modal-btn" @click.prevent="">Edit</button>
                </span>
            </div>
        </div>
        <div v-if="deleting" class="block-screen">
            <div class="modal-body">
                <h4>Are you sure you want to delete this address?</h4>
                <span class="modal-address-container">
                    <div class="modal-address-text">{{ address.street }} {{ address.houseNumber }} {{ address.floorNumber }} {{ address.apartment }}</div>
                    <div class="modal-address-text">{{ address.zipCode }} {{ address.city }}</div>
                </span>
                <span class="btn-container">
                    <button class="modal-btn" @click.prevent="deleting = !deleting">Cancel</button>
                    <button class="modal-btn" @click.prevent="">Delete</button>
                </span>
            </div>
        </div>
    </li>
</template>

<style scoped>
@import url("@/assets/modal.css");

body.modal-open {
    overflow: hidden;
}

.address-list-item {
    cursor: pointer;
    display: flex;
    justify-content: space-between;
    border: var(--sub-border-size) solid var(--second-color);
    padding-block: var(--sub-p-size);
    margin-block: var(--sub-p-size);
    transition: box-shadow var(--tran);
}

.address-list-item:hover {
    box-shadow: 0 0 var(--tiny-size) 1px var(--second-color);
}

.separator-container {
    display: flex;
}

.address-radio {
    cursor: pointer;
    display: grid;
    place-content: center;
    -webkit-appearance: none;
    appearance: none;
    width: var(--p-size);
    height: var(--p-size);
    border: var(--border-size) solid var(--second-color);
    border-radius: 50%;
    margin: 0;
    margin-left: var(--tiny-size);
    transition: border var(--tran);
}

.address-radio::before {
    content: "";
    border-radius: 50%;
    width: calc(var(--sub-p-size) * 0.8);
    height: calc(var(--sub-p-size) * 0.8);
    background-color: var(--second-color);
    transform: scale(0);
    transition: transform var(--tran);
}

.address-radio:checked::before {
    transform: scale(1);
}

.address-radio:hover, .address-radio:checked {
    border: var(--border-size) solid var(--second-hover);
}

.address-container {
    margin-left: var(--tiny-size);
}

.address-text {
    font-size: var(--sub-p-size);
    color: var(--second-hover);
}

.modal-address-container {
    margin-block: var(--h4-size);
    margin-left: var(--h4-size);
}

.modal-address-text {
    font-size: var(--p-size);
}

.action-btn {
    cursor: pointer;
    width: var(--h4-size);
    height: var(--h4-size);
    margin: auto var(--sub-p-size);
}

.btn-img {
    width: 100%;
}

.title {
    color: var(--second-color);
    margin: var(--sub-p-size) auto;
}

.edit-input {
    font-size: var(--p-size);
    border: var(--sub-border-size) solid var(--second-color);
    background-color: var(--main-color);
    color: var(--second-color);
    width: 50%;
    padding: var(--tiny-size);
    margin: var(--sub-p-size) auto;
    transition: border var(--tran), box-shadow var(--tran);
}

.edit-input:hover {
    border: var(--sub-border-size) solid var(--second-hover);
    box-shadow: 0 0 var(--tiny-size) 1px var(--second-color);
}

.btn-container {
    display: flex;
}

.modal-btn {
    cursor: pointer;
    font-size: var(--p-size);
    border-radius: var(--sub-p-size);
    background-color: var(--second-color);
    color: var(--main-color);
    width: 40%;
    padding: var(--sub-p-size) var(--p-size);
    margin: var(--p-size) auto;
    transition: background-color var(--tran);
}

.modal-btn:hover {
    background-color: var(--second-hover);
}

h4 {
    text-align: center;
    width: 85%;
    margin-inline: auto;
}

@media only screen and (hover: none) and (pointer: coarse) {
    .address-text {
        font-size: calc(var(--p-size) * 0.85);
    }

    .edit-input {
        font-size: var(--h4-size);
        width: 65%;
    }
}
</style>