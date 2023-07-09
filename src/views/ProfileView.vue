<script setup lang="ts">
import Header from "@/components/Header.vue"
import { reactive, ref, watch } from "vue";
import AddressCard from '@/components/order/AddressCard.vue'
import type { addressInfo } from "./OrderView.vue";
import { useAuthStore } from "@/store";

let addresses: addressInfo[] = [
    { city: "Budapest", street: "Szamos utca", houseNumber: "8", zipCode: "1039", floorNumber: "2", apartment: ""},
    { city: "Budapest", street: "Szamos utca", houseNumber: "9", zipCode: "1039", floorNumber: "2", apartment: ""},
    { city: "Budapest", street: "Szamos utca", houseNumber: "10", zipCode: "1039", floorNumber: "2", apartment: ""},
    { city: "Budapest", street: "Szamos utca", houseNumber: "11", zipCode: "1039", floorNumber: "2", apartment: ""},
]

const useAuth = useAuthStore()

const firstName = ref(useAuth.firstName)
const lastName = ref(useAuth.lastName)

let hasChanged = ref(firstName.value !== useAuth.firstName)

watch(firstName, (n, o) => {
    console.log(firstName.value, useAuth.firstName, hasChanged.value)
})

</script>

<template>
    <Header/>
    <main>
        <form>
            <h2 class="title">Profile settings</h2>
            <h3>Name</h3>
            <input type="text" placeholder="First name" v-model="firstName">
            <input type="text" placeholder="Last name" v-model="lastName">
            <h3>Addresses</h3>
            <p>
                <AddressCard v-for="a in addresses" :address="a"/>
            </p>
            <h3>Password</h3>
            <p>Incase you forgot your password or you want to change it you can do it here.</p>
            <button class="btn" @click.prevent="">New password</button>
            <button class="btn save-btn" @click.prevent="hasChanged = !hasChanged" :changed="hasChanged" :disabled="!hasChanged">Save Changes</button>
            <h3>Delete Profile</h3>
            <p>This action can not be undone, are you sure you want to delete your Hangry profile?</p>
            <button class="delete-btn" @click.prevent="">Delete profile</button>
        </form>
    </main>
</template>

<style scoped>
    main {
        display: grid;
        width: 100%;
        height: calc(100% - var(--h1-size) * 1.25 - var(--sub-border-size) * 2);
    }
    
    form {
        width: 30%;
        box-shadow: 0 0 var(--tiny-size) 1px var(--settings-color);
        padding-block: var(--h2-size);
        margin: auto;
    }

    h2, h3, p {
        color: var(--second-color);
        margin: var(--p-size);
    }

    input {
        display: block;
        text-align: left;
        font-size: var(--h4-size);
        border: solid var(--border-size) var(--second-color);
        border-radius: var(--h3-size);
        color: var(--second-color);
        width: 45%;
        margin: var(--sub-p-size) var(--p-size);
        padding: var(--tiny-size) var(--p-size);
        transition: box-shadow var(--tran);
    }

    input:hover {
        box-shadow: 0 0 var(--border-size) calc(var(--sub-border-size) * 0.25) var(--second-hover);
    }

    button {
        display: block;
        margin: var(--p-size) auto;
    }

    .btn {
        cursor: pointer;
        font-size: var(--p-size);
        border-radius: var(--sub-p-size);
        background-color: var(--second-color);
        color: var(--main-color);
        padding: var(--sub-p-size);
        transition: background-color var(--tran);
    }
    
    .btn:hover {
        background-color: var(--second-hover);
    }

    .save-btn[changed=false] {
        background-color: var(--settings-color);
    }

    .save-btn[changed=false]:hover {
        background-color: var(--settings-dark);
    }

    .save-btn {
        font-size: var(--h4-size);
        margin-block: var(--h2-size);
    }
    
    .delete-btn {
        cursor: pointer;
        font-size: var(--p-size);
        border: var(--sub-border-size) solid var(--second-color);
        border-radius: var(--sub-p-size);
        background-color: var(--main-color);
        color: var(--error-color);
        padding: var(--sub-p-size);
        transition: border var(--tran), background-color var(--tran);
    }
    
    .delete-btn:hover {
        border: var(--sub-border-size) solid var(--second-hover);
        background-color: var(--second-hover);
    }

    @media only screen and (hover: none) and (pointer: coarse) {
        form {
            width: 100%;
        }

        .title {
            font-size: calc(var(--h2-size) * 0.95);
        }

        input {
            width: calc(100% - 2 * var(--sub-border-size) - 4 * var(--p-size));
        }
    }
</style>