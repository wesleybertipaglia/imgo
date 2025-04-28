package utils

import "errors"

var ErrInvalidWriter = errors.New("invalid writer: expected *os.File")
