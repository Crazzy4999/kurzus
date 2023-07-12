import type { addressInfo, productInfo } from "@/api/models"
import { defineStore } from "pinia"

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
        products: [] as productInfo[]
    }),
    actions: {
        async addToCart(count: number, supplierID: number, name: string, price: number) {
            let alreadyInCart = false;
            this.products.forEach(p => {
                if(p.name === name) alreadyInCart = true
            })
            if(1 <= this.$state.products.length && this.$state.products[this.$state.products.length-1].supplierID !== supplierID) this.$state.products.length = 0
            if(!alreadyInCart) this.$state.products.push({ count: count, supplierID: supplierID, name: name, price: price })
        }
    }
})