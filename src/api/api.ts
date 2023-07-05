import { LoginError, ResetError, SignUpError } from "./errors"
import type { tokenPairResponse } from "./responses"

const root = "http://localhost:3000"

function apiFetch(url: string, init?: RequestInit | undefined) {
    return fetch(root+url, init)
}

function GET(url: string, data: any) {
    return apiFetch(url, {
        method: 'get',
        body: JSON.stringify(data)
    })
}

function POST(url: string, data: any) {
    return apiFetch(url, {
        method: 'post',
        body: JSON.stringify(data)
    })
}

function PUT(url: string, data: any) {
    return apiFetch(url, {
        method: 'put',
        body: JSON.stringify(data)
    })
}

function DELETE(url: string, data: any) {
    return apiFetch(url, {
        method: 'delete',
        body: JSON.stringify(data)
    })
}



export async function login(email: string, password: string) {
    return POST("/login", { email, password }).then(async response => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case LoginError.invalidCredentials: throw "Invalid credentials!"
                default: throw "An unexpected error has occured!"
            }
        }
        return response.json() as Promise<tokenPairResponse>
    })
}

export async function signUp(firstName: string, lastName: string, email: string, password: string, passwordAgain: string) {
    return POST("/signup", { firstName, lastName, email, password, passwordAgain }).then(async response => {
            if(!response.ok) {
                let err = (await response.text()).replace("\n", "")
                switch(err) {
                    case SignUpError.passwordMismatch: throw "The passwords doesn't match!"
                    case SignUpError.emailTaken: throw "Email is already in use!"
                    case SignUpError.passwordFailed: throw "Oops, something went wrong!"
                    default: throw "An unexpected error has occured!"
                }
            }
            return response.json()
        })
}

export async function getResetKey(email: string) {
    return POST("/reset", { email }).then(async response => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case ResetError.userNotExist: throw "There is no user registered with this email!"
                default: throw "An unexpected error has occured!"
            }
        }
        return response.json()
    })
}

export async function resetPassword(email: string, password: string, passwordAgain: string) {
    return PUT("/reset", { email, password, passwordAgain }).then(async response => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case ResetError.passwordMismatch: throw "The passwords doesn't match!"
                case ResetError.getUserFailed: throw "Oops something went wrong!"
                case ResetError.resetKeyMissing: throw "Reset key is missing! Request a new reset!"
                case ResetError.invalidCredentials: throw "Invalid credentials!"
                case ResetError.passwordFailed: throw "Oops something went wrong!"
                case ResetError.resetFailed: throw "Reseting password failed!"
                default: throw "An unexpected error has occured!"
            }
        }
        return response.json()
    })
}

export async function refresh() {
    return GET("/refresh", {}).then(async response => {
        if(!response.ok) {
            let err = (await response.text()).replace("\n", "")
            switch(err) {
                case LoginError.invalidCredentials: throw "Invalid credentials!"
                default: throw "An unexpected error has occured!"
            }
        }
        return response.json()
    })
}