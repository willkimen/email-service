package emailerrors

import "errors"

var (
	ErrTemporaryFailure = errors.New("temporary failure")
	ErrPermanentFailure = errors.New("permanent failure")
)
