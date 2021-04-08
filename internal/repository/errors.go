package repository

import "errors"

var (
	ErrUserNotFound            = errors.New("user doesn't exists")
	ErrVerificationCodeInvalid = errors.New("verification code is invalid")
)
