package util

import (
	"fmt"
	"github.com/pkg/errors"
)

func Err(msg string, a ...any) error {
	err := errors.WithStack(errors.New(fmt.Sprint(msg, a)))
	Log().LogError(err)
	return err
}

func ErrStack(e error, a ...any) error {
	err := errors.WithStack(errors.WithMessage(e, fmt.Sprint(a)))
	Log().LogError(err)
	return err
}
