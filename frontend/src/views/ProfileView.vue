<script setup lang="ts">
import Header from "@/components/Header.vue"
import { nextTick, ref, watch, watchEffect } from "vue";
import AddressCard from '@/components/order/AddressCard.vue'
import { useAuthStore } from "@/store";
import { addAddress, deleteProfile, getAddresses, getResetKey, updateProfile } from "@/api/api";
import router from "@/router";

const useAuth = useAuthStore()

const addingAddress = ref(false)
const city = ref("")
const street = ref("")
const zipCode = ref("")
const houseNumber = ref("")
const floorNumber = ref("")
const apartment = ref("")
const errorMsg = ref("")

function filter(e: KeyboardEvent) {
    if(!/\d/.test(e.key)) return e.preventDefault()
}

async function addNewAddress() {
    if(city.value !== "" && street.value !== "" && houseNumber.value !== "" && zipCode.value !== "") {
        await addAddress(false, city.value, street.value, houseNumber.value, zipCode.value, floorNumber.value, apartment.value)
        router.push("/suppliers")
    } else errorMsg.value = "city, street, house number and zipcode mustn't be empty!"
}

async function getUserAddresses() {
    useAuth.setAddresses((await getAddresses()).addresses)
}

watchEffect(async () => {
    await getUserAddresses()
})

//UPDATE ADDRESS LIST AFTER NEW ADDRESS WAS ADDED



const emailSent = ref("")

function getResetEmail() {
    if(useAuth.email !== "") getResetKey(useAuth.email).then(() => emailSent.value = "Reset email succesfully sent, check your email!").catch(err => emailSent.value = err)
    else emailSent.value = "You must be signed in!"
}



const firstName = ref(useAuth.firstName)
const lastName = ref(useAuth.lastName)

function hasChanged(): boolean {
    return firstName.value !== useAuth.firstName || lastName.value !== useAuth.lastName
}

function updateUser(firstName: string, lastName: string) {
    if(hasChanged()) updateProfile(firstName, lastName).then(() => {
        useAuth.firstName = firstName
        useAuth.lastName = lastName
        router.push("/")
    })
}



const showDeleteModal = ref(false)
const deleteErr = ref("")

function deleteUser() {
    deleteProfile().then(() => {
        useAuth.setUser(-1, "", "", "")
        router.push("/")
    }).catch(err => deleteErr.value = err)
}
</script>

<template>
    <Header/>
    <main>
        <form>
            <h2 class="title">Profile settings</h2>
            <h3>Name</h3>
            <input class="profile-input" type="text" placeholder="First name" v-model="firstName">
            <input class="profile-input" type="text" placeholder="Last name" v-model="lastName">
            <h3>Addresses</h3>
            <p>
                <AddressCard v-for="a in useAuth.addresses" :address="a"/>
                <button class="btn" @click.prevent="addingAddress = !addingAddress">New address</button>
            </p>
            <h3>Password</h3>
            <p>Incase you forgot your password or you want to change it you can do it here.</p>
            <button class="btn" @click.prevent="getResetEmail()">New password</button>
            <p class="info-msg">{{ emailSent }}</p>
            <button class="btn save-btn" @click.prevent="updateUser(firstName, lastName)" :changed="hasChanged()" :disabled="!hasChanged()">Save Changes</button>
            <h3>Delete Profile</h3>
            <p>This action can not be undone, are you sure you want to delete your Hangry profile?</p>
            <button class="delete-btn" @click.prevent="showDeleteModal = !showDeleteModal">Delete profile</button>
            <p class="info-msg">{{ deleteErr }}</p>

            <div class="block-screen" v-if="showDeleteModal">
                <div class="modal-body">
                    <h3 class="modal-header">Delete user profile</h3>
                    <p class="modal-text">Are you sure you want to delete your user profile?</p>
                    <span class="modal-flex">
                        <button class="btn" @click.prevent="showDeleteModal = !showDeleteModal">Cancel</button>
                        <button class="btn" @click.prevent="deleteUser()">Delete profile</button>
                    </span>
                </div>
            </div>

            <div class="block-screen" v-if="addingAddress">
                <div class="modal-body">
                    <h3 class="modal-header">Add new address</h3>
                    <input class="address-input" type="text" placeholder="city" v-model="city">
                    <input class="address-input" type="text" placeholder="street" v-model="street">
                    <input class="address-input" type="text" placeholder="house number" v-model="houseNumber" @keypress="(e) => filter(e)">
                    <input class="address-input" type="text" placeholder="zip code" v-model="zipCode" @keypress="(e) => filter(e)">
                    <input class="address-input" type="text" placeholder="floor number" v-model="floorNumber" @keypress="(e) => filter(e)">
                    <input class="address-input" type="text" placeholder="apartment" v-model="apartment" @keypress="(e) => filter(e)">
                    <span class="error-msg">{{ errorMsg }}</span>
                    <span class="modal-flex">
                        <button class="btn" @click.prevent="addingAddress = !addingAddress">Cancel</button>
                        <button class="btn" @click.prevent="addNewAddress()">Add address</button>
                    </span>
                </div>
            </div>
        </form>
    </main>
</template>

<style scoped>
    @import url("@/assets/modal.css");

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

    .profile-input {
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

    .address-input {
        display: block;
        text-align: left;
        font-size: var(--h4-size);
        border: solid var(--border-size) var(--second-color);
        color: var(--second-color);
        width: 75%;
        margin: var(--sub-p-size) auto;
        padding: var(--tiny-size) var(--p-size);
        transition: box-shadow var(--tran);
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

    .info-msg {
        text-align: center;
        margin: var(--h3-size) auto;
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

    .error-msg {
        text-align: center;
        font-size: var(--p-size);
        color: var(--error-color);
        margin: var(--sub-p-size) 0;
    }

    @media only screen and (hover: none) and (pointer: coarse) {
        form {
            width: 100%;
        }

        .title {
            font-size: calc(var(--h2-size) * 0.95);
        }

        .profile-input {
            width: calc(100% - 2 * var(--sub-border-size) - 4 * var(--p-size));
        }
    }
</style>