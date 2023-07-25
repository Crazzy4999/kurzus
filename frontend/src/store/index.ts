import type { addressInfo, productInfo } from "@/api/models"
import router from "@/router"
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
        },

        signOut() {
            this.setUser(-1, "", "", "")
            router.push("/")
        }
    }
})



export const useCartStore = defineStore("cart", {
    state: () => ({
        products: [] as productInfo[],
        total: 0
    }),
    actions: {
        async addToCart(count: number, menuID: number, supplierID: number, name: string, price: number) {
            let alreadyInCart = false;
            this.products.forEach(p => {
                if(p.name === name) alreadyInCart = true
            })
            if(1 <= this.$state.products.length && this.$state.products[this.$state.products.length-1].supplierID !== supplierID) {
                this.$state.products.length = 0
                this.$state.total = 0
            }
            if(!alreadyInCart) {
                this.$state.products.push({ count: count, menuID: menuID, supplierID: supplierID, name: name, price: price })
                this.$state.total += price * count
            } else {
                this.$state.products.forEach((p, index) => {
                    if(p.menuID === menuID) {
                        this.$state.products[index].count += count
                        this.$state.total += price * count
                    }
                })
            }
        }
    }
})