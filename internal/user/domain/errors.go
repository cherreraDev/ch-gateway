package domain

import "errors"

var ErrIncorrectPassword error = errors.New("the password provided it's not valid")
