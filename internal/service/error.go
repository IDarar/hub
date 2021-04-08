package service

import "errors"

var (
	ErrUserNotFound = errors.New("user doesn't exists")

	ErrUnknownCallbackType     = errors.New("unknown callback type")
	ErrVerificationCodeInvalid = errors.New("verification code is invalid")
)
