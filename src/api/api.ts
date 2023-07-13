import { useAuthStore } from "@/store"
import { ACCESS_TOKEN_EXPIRED, AddressError, CategoriesError, ItemsMenusError, LoginError, MenuError, OOPS, OrderError, OrderMenuError, ProfileError, RefreshError, ResetError, SignUpError, SupplierError, UNEXPECTED } from "./errors"
import type { userResponse, tokenPairResponse, addressesCollectionResponse, supplierCollectionResponse, menuCollectionResponse, orderCollectionResponse, supplierResponse, itemCollectionResponse, categoriesCollectionResponse } from "./responses"
import router from "@/router"
import type { addressCollectionInfo, addressInfo, categoriesCollectionInfo, itemCollectionInfo, menuCollectionInfo, orderCollectionInfo, orderMenuCollectionInfo, supplierCollectionInfo, supplierInfo, tokenPairInfo, userInfo } from "./models"

const root = "http://localhost:3000"
let refreshTries = 0

async function apiFetch(url: string, init?: RequestInit | undefined) {
    return await fetch(root+url, init)
}

async function GET(url: string, isProtected: boolean) {
    return await apiFetch(url, {
        method: 'get',
        headers: isProtected ? {
            Authorization: 'Bearer ' + useAuthStore().accessToken,
        } : {},
    }).catch(err => {
        console.log("GET", err, refreshTries)
        if(err === ACCESS_TOKEN_EXPIRED) refresh().then(resp => {
            useAuthStore().setTokens(resp.accessToken, resp.refreshToken)
            return GET(url, isProtected)
        })
        else if(2 < refreshTries) {
            refreshTries = 0
            router.push("/")
        }
        else refreshTries++
    })
}

async function GET_REFRESH(url: string, isProtected: boolean) {
    return await apiFetch(url, {
        method: 'get',
        headers: isProtected ? {
            Authorization: 'Bearer ' + useAuthStore().refreshToken,
        } : {},
    }).catch(err => {
        console.log("GET_REFRESH", err, refreshTries)
        if(err === ACCESS_TOKEN_EXPIRED) refresh().then(resp => {
            useAuthStore().setTokens(resp.accessToken, resp.refreshToken)
            return GET_REFRESH(url, isProtected)
        })
        else if(2 < refreshTries) {
            refreshTries = 0
            router.push("/")
        }
        else refreshTries++
    })
}

async function POST(url: string, isProtected: boolean, data: any) {
    return await apiFetch(url, {
        method: 'post',
        headers: isProtected ? {
            Authorization: 'Bearer ' + useAuthStore().accessToken,
        } : {},
        body: JSON.stringify(data)
    }).catch(err => {
        console.log("POST", err, refreshTries)
        if(err === ACCESS_TOKEN_EXPIRED) refresh().then(resp => {
            useAuthStore().setTokens(resp.accessToken, resp.refreshToken)
            return POST(url, isProtected, data)
        })
        else if(2 < refreshTries) {
            refreshTries = 0
            router.push("/")
        }
        else refreshTries++
    })
}

async function PUT(url: string, isProtected: boolean, data: any) {
    return await apiFetch(url, {
        method: 'put',
        headers: isProtected ? {
            Authorization: 'Bearer ' + useAuthStore().accessToken,
        } : {},
        body: JSON.stringify(data)
    }).catch(err => {
        console.log("PUT", err, refreshTries)
        if(err === ACCESS_TOKEN_EXPIRED) refresh().then(resp => {
            useAuthStore().setTokens(resp.accessToken, resp.refreshToken)
            return PUT(url, isProtected, data)
        })
        else if(2 < refreshTries) {
            refreshTries = 0
            router.push("/")
        }
        else refreshTries++
    })
}

async function DELETE(url: string, isProtected: boolean, data: any) {
    return await apiFetch(url, {
        method: 'delete',
        headers: isProtected ? {
            Authorization: 'Bearer ' + useAuthStore().accessToken,
        } : {},
        body: JSON.stringify(data)
    }).catch(err => {
        console.log("DELETE", err, refreshTries)
        if(err === ACCESS_TOKEN_EXPIRED) refresh().then(resp => {
            useAuthStore().setTokens(resp.accessToken, resp.refreshToken)
            return DELETE(url, isProtected, data)
        })
        else if(2 < refreshTries) {
            refreshTries = 0
            router.push("/")
        }
        else refreshTries++
    })
}





export async function login(email: string, password: string) {
    return POST("/login", false, { email, password }).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case LoginError.invalidCredentials: throw "Invalid credentials!"
                default: throw UNEXPECTED
            }
        }
        return response.json() as Promise<tokenPairInfo>
    })
}

export async function signUp(firstName: string, lastName: string, email: string, password: string, passwordAgain: string) {
    return POST("/signup", false, { firstName, lastName, email, password, passwordAgain }).then(async (response: any) => {
            if(!response.ok) {
                let err = (await response.text()).replace("\n", "")
                switch(err) {
                    case SignUpError.passwordMismatch: throw "The passwords doesn't match!"
                    case SignUpError.emailTaken: throw "Email is already in use!"
                    case SignUpError.passwordFailed: throw OOPS
                    default: throw UNEXPECTED
                }
            }
        })
}

export async function getResetKey(email: string) {
    return POST("/reset", false, { email }).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case ResetError.userNotExist: throw "There is no user registered with this email!"
                default: throw UNEXPECTED
            }
        }
    })
}

export async function resetPassword(password: string, passwordAgain: string) {
    return PUT("/reset?reset_key=" + router.currentRoute.value.query.reset_key, false, { password, passwordAgain }).then(async (response: any) => {
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
    })
}

export async function refresh() {
    return GET_REFRESH("/refresh", true).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case RefreshError.invalidCredentials: throw "Invalid credentials!"
                case RefreshError.refreshTokenExpired: throw RefreshError.refreshTokenExpired
                default: throw UNEXPECTED
            }
        }
        return response.json() as Promise<tokenPairInfo>
    })
}



export async function getProfile() {
    return await GET("/profile", true).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case ProfileError.invalidCredentials: throw "Invalid credentials!"
                case ProfileError.getUserFailed: throw OOPS
                default: throw UNEXPECTED
            }
        }
        return response.json() as Promise<userInfo>
    })
}

export async function updateProfile(firstName: string, lastName: string) {
    return PUT("/profile", true, { firstName, lastName }).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case ProfileError.invalidCredentials: throw "Invalid credentials!"
                case ProfileError.getUserFailed: throw OOPS
                case ProfileError.updatingUserFailed: throw OOPS
                default: throw UNEXPECTED
            }
        }
    })
}

export async function deleteProfile() {
    return DELETE("/profile", true, {}).then(async (response: any) => {
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
    })
}



export async function addAddress(isActive: boolean, city: string, street: string, houseNumber: string, zipCode: string, floorNumber: string, apartment: string) {
    return POST("/address", true, { isActive, city, street, houseNumber, zipCode, floorNumber, apartment }).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case AddressError.invalidCredentials: throw "Invalid credentials!"
                case AddressError.creatingAddressFailed: throw OOPS
                default: throw UNEXPECTED
            }
        }
    })
}

export async function getAddressByID(id: number) {
    return GET("/address/" + id, true).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case AddressError.invalidCredentials: throw "Invalid credentials!"
                case AddressError.getAllAddressFailed: throw OOPS
                default: throw UNEXPECTED
            }
        }
        return response.json() as Promise<addressInfo>
    })
}

export async function getAddresses() {
    return GET("/addresses", true).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case AddressError.invalidCredentials: throw "Invalid credentials!"
                case AddressError.getAllAddressFailed: throw OOPS
                default: throw UNEXPECTED
            }
        }
        return response.json() as Promise<addressCollectionInfo>
    })
}

export async function updateAddress(id: number, userID: number, isActive: boolean, city: string, street: string, houseNumber: string, zipCode: string, floorNumber: string, apartment: string) {
    return PUT("/address", true, { id, userID, isActive, city, street, houseNumber, zipCode, floorNumber, apartment }).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case AddressError.updatingAddressFailed: throw OOPS
                default: throw UNEXPECTED
            }
        }
    })
}

export async function deleteAddress(id: number) {
    return DELETE("/address", true, { id }).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case AddressError.deletingAddressFailed: throw OOPS
                default: throw UNEXPECTED
            }
        }
    })
}



export async function getSupplierByID(id: number) {
    return GET("/supplier/" + id, true).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case SupplierError.getAllSupplierFailed: throw OOPS
                case SupplierError.getSupplierTypeFailed: throw OOPS
                default: throw UNEXPECTED
            }
        }
        return response.json() as Promise<supplierInfo>
    })
}

export async function getSuppliers() {
    return GET("/suppliers", true).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case SupplierError.getAllSupplierFailed: throw OOPS
                case SupplierError.getSupplierTypeFailed: throw OOPS
                default: throw UNEXPECTED
            }
        }
        return response.json() as Promise<supplierCollectionInfo>
    })
}



export async function getCategories() {
    return GET("/categories", true).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case CategoriesError.getAllCategoriesFailed: throw OOPS
                default: throw UNEXPECTED
            }
        }
        return response.json() as Promise<categoriesCollectionInfo>
    })
}



export async function getMenusBySupplierID(id: number) {
    return GET("/menus/" + id, true).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case MenuError.getAllMenuFailed: throw OOPS
                default: throw UNEXPECTED
            }
        }
        return response.json() as Promise<menuCollectionInfo>
    })
}



export async function getItemsByMenuID(id: number) {
    return GET("/itemsmenus/" + id, true).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case ItemsMenusError.getAllMenuFailed: throw OOPS
                default: throw UNEXPECTED
            }
        }
        return response.json() as Promise<itemCollectionInfo>
    })
}



export async function addOrderMenu(orderID: number, menuID: number, quantity: number) {
    return POST("/ordermenu", true, { orderID, menuID, quantity }).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case OrderMenuError.creatingOrderMenuFailed: throw OOPS
                default: throw UNEXPECTED
            }
        }
    })
}

export async function GetMenusByOrderID(orderID: number) {
    return GET("/ordermenu/" + orderID, true).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case OrderMenuError.stringConversionFailed: throw OOPS
                case OrderMenuError.getAllMenuByOrderFailed: throw OOPS
                default: throw UNEXPECTED
            }
        }
        return response.json() as Promise<menuCollectionInfo>
    })
}

export async function GetOrderMenusByOrderID(orderID: number) {
    return GET("/ordermenus/" + orderID, true).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case OrderMenuError.stringConversionFailed: throw OOPS
                case OrderMenuError.getAllOrderMenuByOrderIDFailed: throw OOPS
                default: throw UNEXPECTED
            }
        }
        return response.json() as Promise<orderMenuCollectionInfo>
    })
}

export async function updateOrderMenu(quantity: number) {
    return PUT("/ordermenu", true, { quantity }).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case OrderMenuError.updatingOrderMenuByOrderFailed: throw OOPS
                default: throw UNEXPECTED
            }
        }
    })
}

export async function deleteOrderMenu(orderID: number) {
    return DELETE("/ordermenu", true, { orderID }).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case OrderMenuError.deletingOrderMenuByOrderFailed: throw OOPS
                default: throw UNEXPECTED
            }
        }
    })
}



export async function makeOrder(userID: number, addressID: number, supplierID: number, note: string, estimatedDelivery: string) {
    return POST("/order", true, { userID, addressID, supplierID, note, estimatedDelivery }).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case OrderError.creatingOrderFailed: throw OOPS
                default: throw UNEXPECTED
            }
        }
    })
}

export async function getOrders() {
    return GET("/orders", true).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case OrderError.getAllOrderFailed: throw OOPS
                default: throw UNEXPECTED
            }
        }
        return response.json() as Promise<orderCollectionInfo>
    })
}