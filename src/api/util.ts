import type { categoryInfo } from "./models"

export enum supplierType {
    RESTAURANT = "restaurant",
    SUPERMARKET = "supermarket",
    COFFEE_SHOP = "coffee_shop",
    ALL = "all"
} 

export type query = {
    search: string
    categories: categoryInfo[]
    type: supplierType
}