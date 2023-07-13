import { useAuthStore } from "@/store"
import { AddressError, CategoriesError, ItemsMenusError, LoginError, MenuError, OOPS, OrderError, OrderMenuError, ProfileError, RefreshError, ResetError, SignUpError, SupplierError, ACCESS_TOKEN_EXPIRED, REFRESH_TOKEN_EXPIRED } from "./errors"
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
                case LoginError.invalidCredentials: throw Error("Invalid credentials!")
                default: throw Error(ACCESS_TOKEN_EXPIRED)
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
                    case SignUpError.passwordMismatch: throw Error("The passwords doesn't match!")
                    case SignUpError.emailTaken: throw Error("Email is already in use!")
                    case SignUpError.passwordFailed: throw Error(OOPS)
                    default: throw Error(ACCESS_TOKEN_EXPIRED)
                }
            }
        })
}

export async function getResetKey(email: string) {
    return POST("/reset", false, { email }).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case ResetError.userNotExist: throw Error("There is no user registered with this email!")
                default: throw Error(ACCESS_TOKEN_EXPIRED)
            }
        }
    })
}

export async function resetPassword(password: string, passwordAgain: string) {
    return PUT("/reset?reset_key=" + router.currentRoute.value.query.reset_key, false, { password, passwordAgain }).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case ResetError.passwordMismatch: throw Error("The passwords doesn't match!")
                case ResetError.getUserFailed: throw Error(OOPS)
                case ResetError.resetKeyMissing: throw Error("Reset key is missing! Request a new reset!")
                case ResetError.invalidCredentials: throw Error("Invalid credentials!")
                case ResetError.passwordFailed: throw Error(OOPS)
                case ResetError.resetFailed: throw Error("Reseting password failed!")
                default: throw Error(ACCESS_TOKEN_EXPIRED)
            }
        }
    })
}

export async function refresh(): Promise<tokenPairInfo> {
    //@ts-ignore
    return GET_REFRESH("/refresh", true).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case RefreshError.invalidCredentials: throw Error("Invalid credentials!")
                case RefreshError.refreshTokenExpired: throw Error(RefreshError.refreshTokenExpired)
                default: throw Error(REFRESH_TOKEN_EXPIRED)
            }
        }
        return response.json() as Promise<tokenPairInfo>
    }).catch((err: Error) => {
        if(err.message === REFRESH_TOKEN_EXPIRED) refresh().then(resp => {
            useAuthStore().setTokens(resp.accessToken, resp.refreshToken)
            return refresh()
        })
        else if(2 < refreshTries) {
            refreshTries = 0
            router.push("/login")
        }
        else refreshTries++
    })
}



export async function getProfile() {
    //@ts-ignore
    return await GET("/profile", true).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case ProfileError.invalidCredentials: throw Error("Invalid credentials!")
                case ProfileError.getUserFailed: throw Error(OOPS)
                default: throw Error(ACCESS_TOKEN_EXPIRED)
            }
        }
        return response.json() as Promise<userInfo>
    })
}

export async function updateProfile(firstName: string, lastName: string) {
    //@ts-ignore
    return PUT("/profile", true, { firstName, lastName }).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case ProfileError.invalidCredentials: throw Error("Invalid credentials!")
                case ProfileError.getUserFailed: throw Error(OOPS)
                case ProfileError.updatingUserFailed: throw Error(OOPS)
                default: throw Error(ACCESS_TOKEN_EXPIRED)
            }
        }
    })
}

export async function deleteProfile() {
    //@ts-ignore
    return DELETE("/profile", true, {}).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case ProfileError.invalidCredentials: throw Error("Invalid credentials!")
                case ProfileError.getAllAddressFailed: throw Error(OOPS)
                case ProfileError.deletingAddressFailed: throw Error(OOPS)
                case ProfileError.deletingUserFailed: throw Error(OOPS)
                default: throw Error(ACCESS_TOKEN_EXPIRED)
            }
        }
    })
}



export async function addAddress(isActive: boolean, city: string, street: string, houseNumber: string, zipCode: string, floorNumber: string, apartment: string) {
    //@ts-ignore
    return POST("/address", true, { isActive, city, street, houseNumber, zipCode, floorNumber, apartment }).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case AddressError.invalidCredentials: throw Error("Invalid credentials!")
                case AddressError.creatingAddressFailed: throw Error(OOPS)
                default: throw Error(ACCESS_TOKEN_EXPIRED)
            }
        }
    })
}

export async function getAddressByID(id: number) {
    //@ts-ignore
    return GET("/address/" + id, true).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case AddressError.invalidCredentials: throw Error("Invalid credentials!")
                case AddressError.getAllAddressFailed: throw Error(OOPS)
                default: throw Error(ACCESS_TOKEN_EXPIRED)
            }
        }
        return response.json() as Promise<addressInfo>
    })
}

export async function getAddresses(): Promise<addressCollectionInfo> {
    //@ts-ignore
    return GET("/addresses", true).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case AddressError.invalidCredentials: throw Error("Invalid credentials!")
                case AddressError.getAllAddressFailed: throw Error(OOPS)
                default: throw Error(ACCESS_TOKEN_EXPIRED)
            }
        }
        return response.json() as Promise<addressCollectionInfo>
    }).catch((err: Error) => {
        if(err.message === ACCESS_TOKEN_EXPIRED) refresh().then(resp => {
            useAuthStore().setTokens(resp.accessToken, resp.refreshToken)
            return getAddresses()
        })
        else if(2 < refreshTries) {
            refreshTries = 0
            router.push("/")
        }
        else refreshTries++
    })
}

export async function updateAddress(id: number, userID: number, isActive: boolean, city: string, street: string, houseNumber: string, zipCode: string, floorNumber: string, apartment: string) {
    //@ts-ignore
    return await PUT("/address", true, { id, userID, isActive, city, street, houseNumber, zipCode, floorNumber, apartment }).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case AddressError.updatingAddressFailed: throw Error(OOPS)
                default: throw Error(ACCESS_TOKEN_EXPIRED)
            }
        }
    }).catch((err: Error) => {
        if(err.message === ACCESS_TOKEN_EXPIRED) refresh().then(resp => {
            useAuthStore().setTokens(resp.accessToken, resp.refreshToken)
            return updateAddress(id, userID, isActive, city, street, houseNumber, zipCode, floorNumber, apartment)
        })
        else if(2 < refreshTries) {
            refreshTries = 0
            router.push("/")
        }
        else refreshTries++
    })
}

export async function deleteAddress(id: number) {
    //@ts-ignore
    return DELETE("/address", true, { id }).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case AddressError.deletingAddressFailed: throw Error(OOPS)
                default: throw Error(ACCESS_TOKEN_EXPIRED)
            }
        }
    }).catch((err: Error) => {
        if(err.message === ACCESS_TOKEN_EXPIRED) refresh().then(resp => {
            useAuthStore().setTokens(resp.accessToken, resp.refreshToken)
            return deleteAddress(id)
        })
        else if(2 < refreshTries) {
            refreshTries = 0
            router.push("/")
        }
        else refreshTries++
    })
}



export async function getSupplierByID(id: number): Promise<supplierInfo> {
    return GET("/supplier/" + id, true).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case SupplierError.getAllSupplierFailed: throw Error(OOPS)
                case SupplierError.getSupplierTypeFailed: throw Error(OOPS)
                default: throw Error(ACCESS_TOKEN_EXPIRED)
            }
        }
        return response.json() as Promise<supplierInfo>
    })
}

export async function getSuppliers(): Promise<supplierCollectionInfo> {
    return GET("/suppliers", true).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case SupplierError.getAllSupplierFailed: throw Error(OOPS)
                case SupplierError.getSupplierTypeFailed: throw Error(OOPS)
                default: throw Error(ACCESS_TOKEN_EXPIRED)
            }
        }
        return response.json() as Promise<supplierCollectionInfo>
    })
}



export async function getCategories(): Promise<categoriesCollectionInfo> {
    return GET("/categories", true).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case CategoriesError.getAllCategoriesFailed: throw Error(OOPS)
                default: throw Error(ACCESS_TOKEN_EXPIRED)
            }
        }
        return response.json() as Promise<categoriesCollectionInfo>
    })
}



export async function getMenusBySupplierID(id: number): Promise<menuCollectionInfo> {
    return GET("/menus/" + id, true).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case MenuError.getAllMenuFailed: throw Error(OOPS)
                default: throw Error(ACCESS_TOKEN_EXPIRED)
            }
        }
        return response.json() as Promise<menuCollectionInfo>
    })
}



export async function getItemsByMenuID(id: number): Promise<itemCollectionInfo> {
    return GET("/itemsmenus/" + id, true).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case ItemsMenusError.getAllMenuFailed: throw Error(OOPS)
                default: throw Error(ACCESS_TOKEN_EXPIRED)
            }
        }
        return response.json() as Promise<itemCollectionInfo>
    })
}



export async function addOrderMenu(orderID: number, menuID: number, quantity: number) {
    //@ts-ignore
    return POST("/ordermenu", true, { orderID, menuID, quantity }).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case OrderMenuError.creatingOrderMenuFailed: throw Error(OOPS)
                default: throw Error(ACCESS_TOKEN_EXPIRED)
            }
        }
    }).catch((err: Error) => {
        if(err.message === ACCESS_TOKEN_EXPIRED) refresh().then(resp => {
            useAuthStore().setTokens(resp.accessToken, resp.refreshToken)
            return addOrderMenu(orderID, menuID, quantity)
        })
        else if(2 < refreshTries) {
            refreshTries = 0
            router.push("/")
        }
        else refreshTries++
    })
}

export async function GetMenusByOrderID(orderID: number): Promise<menuCollectionInfo> {
    //@ts-ignore
    return GET("/ordermenu/" + orderID, true).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case OrderMenuError.stringConversionFailed: throw Error(OOPS)
                case OrderMenuError.getAllMenuByOrderFailed: throw Error(OOPS)
                default: throw Error(ACCESS_TOKEN_EXPIRED)
            }
        }
        return response.json() as Promise<menuCollectionInfo>
    }).catch((err: Error) => {
        if(err.message === ACCESS_TOKEN_EXPIRED) refresh().then(resp => {
            useAuthStore().setTokens(resp.accessToken, resp.refreshToken)
            return GetMenusByOrderID(orderID)
        })
        else if(2 < refreshTries) {
            refreshTries = 0
            router.push("/")
        }
        else refreshTries++
    })
}

export async function GetOrderMenusByOrderID(orderID: number): Promise<orderMenuCollectionInfo> {
    //@ts-ignore
    return GET("/ordermenus/" + orderID, true).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case OrderMenuError.stringConversionFailed: throw Error(OOPS)
                case OrderMenuError.getAllOrderMenuByOrderIDFailed: throw Error(OOPS)
                default: throw Error(ACCESS_TOKEN_EXPIRED)
            }
        }
        return response.json() as Promise<orderMenuCollectionInfo>
    }).catch((err: Error) => {
        if(err.message === ACCESS_TOKEN_EXPIRED) refresh().then(resp => {
            useAuthStore().setTokens(resp.accessToken, resp.refreshToken)
            return GetOrderMenusByOrderID(orderID)
        })
        else if(2 < refreshTries) {
            refreshTries = 0
            router.push("/")
        }
        else refreshTries++
    })
}

export async function updateOrderMenu(quantity: number) {
    //@ts-ignore
    return PUT("/ordermenu", true, { quantity }).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case OrderMenuError.updatingOrderMenuByOrderFailed: throw Error(OOPS)
                default: throw Error(ACCESS_TOKEN_EXPIRED)
            }
        }
    }).catch((err: Error) => {
        if(err.message === ACCESS_TOKEN_EXPIRED) refresh().then(resp => {
            useAuthStore().setTokens(resp.accessToken, resp.refreshToken)
            return updateOrderMenu(quantity)
        })
        else if(2 < refreshTries) {
            refreshTries = 0
            router.push("/")
        }
        else refreshTries++
    })
}

export async function deleteOrderMenu(orderID: number) {
    //@ts-ignore
    return DELETE("/ordermenu", true, { orderID }).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case OrderMenuError.deletingOrderMenuByOrderFailed: throw Error(OOPS)
                default: throw Error(ACCESS_TOKEN_EXPIRED)
            }
        }
    }).catch((err: Error) => {
        if(err.message === ACCESS_TOKEN_EXPIRED) refresh().then(resp => {
            useAuthStore().setTokens(resp.accessToken, resp.refreshToken)
            return deleteOrderMenu(orderID)
        })
        else if(2 < refreshTries) {
            refreshTries = 0
            router.push("/")
        }
        else refreshTries++
    })
}



export async function makeOrder(userID: number, addressID: number, supplierID: number, note: string, estimatedDelivery: string) {
    //@ts-ignore
    return POST("/order", true, { userID, addressID, supplierID, note, estimatedDelivery }).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case OrderError.creatingOrderFailed: throw Error(OOPS)
                default: throw Error(ACCESS_TOKEN_EXPIRED)
            }
        }
    }).catch((err: Error) => {
        if(err.message === ACCESS_TOKEN_EXPIRED) refresh().then(resp => {
            useAuthStore().setTokens(resp.accessToken, resp.refreshToken)
            return makeOrder(userID, addressID, supplierID, note, estimatedDelivery)
        })
        else if(2 < refreshTries) {
            refreshTries = 0
            router.push("/")
        }
        else refreshTries++
    })
}

export async function getOrders(): Promise<orderCollectionInfo> {
    //@ts-ignore
    return GET("/orders", true).then(async (response: any) => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case OrderError.getAllOrderFailed: throw Error(OOPS)
                default: throw Error(ACCESS_TOKEN_EXPIRED)
            }
        }
        return response.json() as Promise<orderCollectionInfo>
    }).catch((err: Error) => {
        if(err.message === ACCESS_TOKEN_EXPIRED) refresh().then(resp => {
            useAuthStore().setTokens(resp.accessToken, resp.refreshToken)
            return getOrders()
        })
        else if(2 < refreshTries) {
            refreshTries = 0
            router.push("/")
        }
        else refreshTries++
    })
}