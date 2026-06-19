package logics

import "errors"

var ErrTaskNotFound = errors.New("Task is not found")
var ErrTaskAlreadyExist = errors.New("Task is already exists")
