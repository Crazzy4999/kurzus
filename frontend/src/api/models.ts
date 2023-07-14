export type tokenPairInfo = {
    accessToken: string
    refreshToken: string
}

export type userInfo = {
    id: number
    firstName: string
    lastName: string
    email: string
}

export type addressInfo = {
    id: number
    userID: number
    isActive: boolean
    city: string
    street: string
    houseNumber: string
    zipCode: string
    floorNumber: string
    apartment: string
}

export type addressCollectionInfo = {
    addresses: addressInfo[]
}

type workingHours = {
    opening: string
    closing: string
}

export type supplierInfo = {
    id: number
    type: string
    image: string
    name: string
    description: string
    deliveryTime: number
    deliveryFee: number
    workingHours: workingHours
}

export type supplierCollectionInfo = {
    suppliers: supplierInfo[]
}

export type itemInfo = {
    id: number
    ingredient: string
}

export type cartItemInfo = {
    count: number
    name: string
    price: number
}

export type itemCollectionInfo = {
    items: itemInfo[]
}

export type categoryInfo = {
    id: number
    name: string
}

export type categoriesCollectionInfo = {
    categories: categoryInfo[]
}

export type productInfo = {
    count: number
    menuID: number
    supplierID: number
    name: string
    price: number
}

export type menuInfo = {
    id: number
    name: string
    image: string
    supplierID: number
    categoryID: number
    price: number
}

export type menuCollectionInfo = {
    menus: menuInfo[]
}

export type orderMenuInfo = {
    orderID: number
    menuID: number
    quantity: number
}

export type orderMenuCollectionInfo = {
    orderMenus: orderMenuInfo[]
}

export type orderInfo = {
    id: number
    userID: number
    addressID: number
    supplierID: number
    driverID: number
    statusID: number
    note: string
    createdAt: string
    estimatedDelivery: string
    deliveredAt: string
}

export type orderCollectionInfo = {
    orders: orderInfo[]
}