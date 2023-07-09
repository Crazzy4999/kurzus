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
        async setTokens(tokenPair: tokenPairResponse) {
            this.$state.accessToken = tokenPair.accessToken
            this.$state.refreshToken = tokenPair.refreshToken
        },

        async setUser(userData: userResponse) {
            this.$state.id = userData.id
            this.$state.firstName = userData.firstName
            this.$state.lastName = userData.lastName
            this.$state.email = userData.email
        }
    }
})



export const useCartStore = defineStore("cart", {
    state: () => ({
        
    }),
    actions: {

    }
})