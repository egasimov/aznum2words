package aznum2words

import "errors"

var (
	ErrInvalidArgument      = errors.New("an invalid argument provided")
	ErrTooBigNumber         = errors.New("max supported number digits is 66")
	ErrTooBigScale          = errors.New("max supported scale for real numbers is 15")
	ErrUnexpectedBehaviour  = errors.New("an unexpected behaviour occurred")
	ErrUnsupportedOperation = errors.New("an operation is not supported")
)
