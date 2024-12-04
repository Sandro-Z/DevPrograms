package service

import (
	"errors"
)

var (
	ErrShutdown = errors.New("service is shut down")
)
