import type { tokenPairResponse, userResponse } from "@/api/responses"
import { defineStore } from "pinia"

export const useAuthStore = defineStore("authentication", {
    state: () => ({
        accessToken: "",
        refreshToken: "",

        id: -1,
        firstName: "",
        lastName: "",
        email: ""
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
        }
    }
})



export const useCartStore = defineStore("cart", {
    state: () => ({
        
    }),
    actions: {

    }
})