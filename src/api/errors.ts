export const UNEXPECTED = "An unexpected error has occured!"
export const OOPS = "Oops something went wrong!"

const INVALID_CREDENTIALS = "invalid credentials"
export const ACCESS_TOKEN_EXPIRED = "access token expired"
const GET_USER_FAILED = "getting user failed"
const PASSWORD_MISMATCH = "passwords doesn't match"
const PASSWORD_GENERATION_FAILED = "couldn't generate password"
const GET_ALL_ADDRESS_FAILED = "getting all addresses failed"
const DELETING_ADDRESS_FAILED = "deleting address failed"



export enum LoginError {
    invalidCredentials = INVALID_CREDENTIALS
}

export enum SignUpError {
    passwordMismatch = PASSWORD_MISMATCH,
    emailTaken = "email already in use",
    passwordFailed = PASSWORD_GENERATION_FAILED
}

export enum RefreshError {
    invalidCredentials = INVALID_CREDENTIALS,
    refreshTokenExpired = "refresh token expired"
}

export enum ResetError {
    userNotExist = "user doesn't exist with this email",
    passwordMismatch = PASSWORD_MISMATCH,
    getUserFailed = GET_USER_FAILED,
    resetKeyMissing = "bad url, missing reset key",
    invalidCredentials = INVALID_CREDENTIALS,
    passwordFailed = PASSWORD_GENERATION_FAILED,
    resetFailed = "reseting password for user failed"
}

export enum ProfileError {
    invalidCredentials = INVALID_CREDENTIALS,
    getUserFailed = GET_USER_FAILED,
    updatingUserFailed = "updating user failed",
    getAllAddressFailed = GET_ALL_ADDRESS_FAILED,
    deletingAddressFailed = DELETING_ADDRESS_FAILED,
    deletingUserFailed = "deleting user failed"
}

export enum AddressError {
    invalidCredentials = INVALID_CREDENTIALS,
    creatingAddressFailed = "creating address failed",
    getAllAddressFailed = GET_ALL_ADDRESS_FAILED,
    updatingAddressFailed = "updating address failed",
    deletingAddressFailed = DELETING_ADDRESS_FAILED
}

export enum SupplierError {
    getAllSupplierFailed = "getting all suppliers failed",
    getSupplierTypeFailed = "getting supplier type by id failed"
}

export enum MenuError {
    getAllMenuFailed = "getting all menus failed"
}

export enum ItemsMenusError {
    getAllMenuFailed = "getting all menus failed"
}

export enum OrderMenuError {
    creatingOrderMenuFailed = "creating order_menu failed",
    stringConversionFailed = "converting string to integer failed",
    getAllMenuByOrderFailed = "getting all menus by order id failed",
    updatingOrderMenuByOrderFailed = "updating order_menu by order id failed",
    deletingOrderMenuByOrderFailed = "deleting order_menu by order id failed"
}

export enum OrderError {
    creatingOrderFailed = "creating order failed",
    getAllOrderFailed = "getting all orders failed"
}