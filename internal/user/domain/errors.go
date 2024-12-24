package domain

import "errors"

var ErrIncorrectPassword error = errors.New("the password provided it's not valid")
var ErrUserNotFound error = errors.New("user not found")
