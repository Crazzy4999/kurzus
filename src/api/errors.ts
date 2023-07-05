export enum LoginError {
    invalidCredentials = "invalid credentials"
}

export enum SignUpError {
    passwordMismatch = "passwords doesn't match",
    emailTaken = "email already in use",
    passwordFailed = "couldn't generate password"
}

export enum ResetError {
    userNotExist = "user doesn't exist with this email",
    passwordMismatch = "passwords doesn't match",
    getUserFailed = "getting user failed",
    resetKeyMissing = "bad url, missing reset key",
    invalidCredentials = "invalid credentials",
    passwordFailed = "couldn't generate password",
    resetFailed = "reseting password for user failed"
}