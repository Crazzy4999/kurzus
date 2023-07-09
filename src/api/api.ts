import { useAuthStore } from "@/store"
import { AddressError, LoginError, OOPS, OrderError, OrderMenuError, ProfileError, ResetError, SignUpError, SupplierError, UNEXPECTED } from "./errors"
import type { userResponse, tokenPairResponse, addressesCollectionResponse, supplierCollectionResponse, menuCollectionResponse, orderCollectionResponse } from "./responses"

const root = "http://localhost:3000"

function apiFetch(url: string, init?: RequestInit | undefined) {
    return fetch(root+url, init)
}

function GET(url: string, isProtected: boolean) {
    return apiFetch(url, {
        method: 'get',
        headers: isProtected ? {
            Authorization: 'Bearer ' + useAuthStore().accessToken,
        } : {},
    })
}

function GET_REFRESH(url: string, isProtected: boolean) {
    return apiFetch(url, {
        method: 'get',
        headers: isProtected ? {
            Authorization: 'Bearer ' + useAuthStore().refreshToken,
        } : {},
    })
}

function POST(url: string, isProtected: boolean, data: any) {
    return apiFetch(url, {
        method: 'post',
        headers: isProtected ? {
            Authorization: 'Bearer ' + useAuthStore().refreshToken,
        } : {},
        body: JSON.stringify(data)
    })
}

function PUT(url: string, isProtected: boolean, data: any) {
    return apiFetch(url, {
        method: 'put',
        headers: isProtected ? {
            Authorization: 'Bearer ' + useAuthStore().accessToken,
        } : {},
        body: JSON.stringify(data)
    })
}

function DELETE(url: string, isProtected: boolean, data: any) {
    return apiFetch(url, {
        method: 'delete',
        headers: isProtected ? {
            Authorization: 'Bearer ' + useAuthStore().accessToken,
        } : {},
        body: JSON.stringify(data)
    })
}





export async function login(email: string, password: string) {
    return POST("/login", false, { email, password }).then(async response => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case LoginError.invalidCredentials: throw "Invalid credentials!"
                default: throw UNEXPECTED
            }
        }
        return response.json() as Promise<tokenPairResponse>
    })
}

export async function signUp(firstName: string, lastName: string, email: string, password: string, passwordAgain: string) {
    return POST("/signup", false, { firstName, lastName, email, password, passwordAgain }).then(async response => {
            if(!response.ok) {
                let err = (await response.text()).replace("\n", "")
                switch(err) {
                    case SignUpError.passwordMismatch: throw "The passwords doesn't match!"
                    case SignUpError.emailTaken: throw "Email is already in use!"
                    case SignUpError.passwordFailed: throw OOPS
                    default: throw UNEXPECTED
                }
            }
            return response.json()
        })
}

export async function getResetKey(email: string) {
    return POST("/reset", false, { email }).then(async response => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case ResetError.userNotExist: throw "There is no user registered with this email!"
                default: throw UNEXPECTED
            }
        }
        return response.json()
    })
}

export async function resetPassword(email: string, password: string, passwordAgain: string) {
    return PUT("/reset", false, { email, password, passwordAgain }).then(async response => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case ResetError.passwordMismatch: throw "The passwords doesn't match!"
                case ResetError.getUserFailed: throw OOPS
                case ResetError.resetKeyMissing: throw "Reset key is missing! Request a new reset!"
                case ResetError.invalidCredentials: throw "Invalid credentials!"
                case ResetError.passwordFailed: throw OOPS
                case ResetError.resetFailed: throw "Reseting password failed!"
                default: throw UNEXPECTED
            }
        }
        return response.json()
    })
}

export async function refresh() {
    return GET_REFRESH("/refresh", true).then(async response => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case LoginError.invalidCredentials: throw "Invalid credentials!"
                default: throw UNEXPECTED
            }
        }
        return response.json() as Promise<tokenPairResponse>
    })
}



export async function getProfile() {
    return GET("/profile", true).then(async response => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case ProfileError.invalidCredentials: throw "Invalid credentials!"
                case ProfileError.getUserFailed: throw OOPS
                default: throw UNEXPECTED
            }
        }
        return response.json() as Promise<userResponse>
    })
}

export async function updateProfile(firstName: string, lastName: string) {
    return PUT("/profile", true, { firstName, lastName }).then(async response => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case ProfileError.invalidCredentials: throw "Invalid credentials!"
                case ProfileError.getUserFailed: throw OOPS
                case ProfileError.updatingUserFailed: throw OOPS
                default: throw UNEXPECTED
            }
        }
        return response.json()
    })
}

export async function deleteProfile() {
    return DELETE("/profile", true, {}).then(async response => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case ProfileError.invalidCredentials: throw "Invalid credentials!"
                case ProfileError.getAllAddressFailed: throw OOPS
                case ProfileError.deletingAddressFailed: throw OOPS
                case ProfileError.deletingUserFailed: throw OOPS
                default: throw UNEXPECTED
            }
        }
        return response.json()
    })
}



export async function addAddress(isActive: boolean, city: string, street: string, houseNumber: string, zipCode: string, floorNumber: string, apartment: string) {
    return POST("/address", true, { isActive, city, street, houseNumber, zipCode, floorNumber, apartment }).then(async response => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case AddressError.invalidCredentials: throw "Invalid credentials!"
                case AddressError.creatingAddressFailed: throw OOPS
                default: throw UNEXPECTED
            }
        }
        return response.json()
    })
}

export async function getAddresses() {
    return GET("/addresses", true).then(async response => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case AddressError.invalidCredentials: throw "Invalid credentials!"
                case AddressError.getAllAddressFailed: throw OOPS
                default: throw UNEXPECTED
            }
        }
        return response.json() as Promise<addressesCollectionResponse>
    })
}

export async function updateAddress(isActive: boolean, city: string, street: string, houseNumber: string, zipCode: string, floorNumber: string, apartment: string) {
    return PUT("/address", true, { isActive, city, street, houseNumber, zipCode, floorNumber, apartment }).then(async response => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case AddressError.updatingAddressFailed: throw OOPS
                default: throw UNEXPECTED
            }
        }
        return response.json()
    })
}

export async function deleteAddress(id: number) {
    return DELETE("/address", true, { id }).then(async response => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case AddressError.deletingAddressFailed: throw OOPS
                default: throw UNEXPECTED
            }
        }
        return response.json()
    })
}



export async function getSuppliers() {
    return GET("/suppliers", true).then(async response => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case SupplierError.getAllSupplierFailed: throw OOPS
                case SupplierError.getSupplierTypeFailed: throw OOPS
                default: throw UNEXPECTED
            }
        }
        return response.json() as Promise<supplierCollectionResponse>
    })
}



export async function addOrderMenu() {
    return POST("/ordermenu", true, {}).then(async response => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case OrderMenuError.creatingOrderMenuFailed: throw OOPS
                default: throw UNEXPECTED
            }
        }
        return response.json()
    })
}

export async function getOrderMenus(orderID: number) {
    return GET("/ordermenus/" + orderID, true).then(async response => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case OrderMenuError.stringConversionFailed: throw OOPS
                case OrderMenuError.getAllMenuByOrderFailed: throw OOPS
                default: throw UNEXPECTED
            }
        }
        return response.json() as Promise<menuCollectionResponse>
    })
}

export async function updateOrderMenu(quantity: number) {
    return PUT("/ordermenu", true, { quantity }).then(async response => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case OrderMenuError.updatingOrderMenuByOrderFailed: throw OOPS
                default: throw UNEXPECTED
            }
        }
        return response.json()
    })
}

export async function deleteOrderMenu(orderID: number) {
    return DELETE("/ordermenu", true, { orderID }).then(async response => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case OrderMenuError.deletingOrderMenuByOrderFailed: throw OOPS
                default: throw UNEXPECTED
            }
        }
        return response.json()
    })
}



export async function makeOrder(userID: number, supplierID: number, note: string) {
    return POST("/order", true, { userID, supplierID, note }).then(async response => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case OrderError.creatingOrderFailed: throw OOPS
                default: throw UNEXPECTED
            }
        }
        return response.json()
    })
}

export async function getOrders() {
    return GET("/orders", true).then(async response => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case OrderError.getAllOrderFailed: throw OOPS
                default: throw UNEXPECTED
            }
        }
        return response.json() as Promise<orderCollectionResponse>
    })
}