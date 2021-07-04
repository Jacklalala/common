package errors

import (
  "github.com/pkg/errors"
)

var (
  New         = errors.New
  Wrap        = errors.Wrap
  Wrapf       = errors.Wrapf
  Errorf      = errors.Errorf
  WithStack   = errors.WithStack
  WithMessage = errors.WithMessage
  Cause       = errors.Cause
)

type Error struct {
  code int
  msg string
}

func (e *Error) Error() string {
  return e.msg
}

func (e *Error) Code() int {
  return e.code
}


