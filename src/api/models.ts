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

export type productInfo = {
    count: number
    name: string
    price: number
}

export type supplierInfo = {
    id: number
    type: string
    image: string
    name: string
    description: string
    deliveryTime: number
    deliveryFee: number
    opening: string
    closing: string
}

export type categoryInfo = {
    id: number
    name: string
}

export type menuInfo = {
    id: number
    name: string
    image: string
    supplierID: number
    categoryID: number
    price: number
}