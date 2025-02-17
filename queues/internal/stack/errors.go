package stack

import "errors"

var ErrStackIsEmpty = errors.New("stack is empty")

var ErrUnknownPriority = errors.New("unknown priority")
