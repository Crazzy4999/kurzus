import type { addressInfo } from "@/api/models"
import { defineStore } from "pinia"
import { ref } from "vue"

export const useAuthStore = defineStore("authentication", {
    state: () => ({
        accessToken: "",
        refreshToken: "",

        id: -1,
        firstName: "",
        lastName: "",
        email: "",

        addresses: [] as addressInfo[]
    }),
    actions: {
        async setTokens(accessToken: string, refreshToken: string) {
            this.$state.accessToken = accessToken
            this.$state.refreshToken = refreshToken
        },

        async setUser(id: number, firstName: string, lastName: string, email: string) {
            this.$state.id = id
            this.$state.firstName = firstName
            this.$state.lastName = lastName
            this.$state.email = email
        },

        async setAddresses(addresses: addressInfo[]) {
            this.$state.addresses = addresses
        }
    }
})



export const useCartStore = defineStore("cart", {
    state: () => ({
        
    }),
    actions: {

    }
})