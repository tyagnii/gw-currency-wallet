package db

import "errors"

var ErrorUserNotFound = errors.New("user not found")
var ErrorInsufficientBalance = errors.New("insufficient balance")
var ErrorAuthenticationFailed = errors.New("authentication failed")
