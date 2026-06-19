package http

import (
	"encoding/json"
	"errors"
	"time"
)

type TaskDTO struct {
	Title       string
	Description string
}

type completeTaskDTO struct {
	Complete bool
}

func (t TaskDTO) ValidateForCreate() error {
	if t.Title == "" {
		return errors.New("Title is empty")
	}

	if t.Description == "" {
		return errors.New("Description is empty")
	}

	return nil
}

type ErrorDto struct {
	Message string
	Time    time.Time
}

func (e ErrorDto) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ")

	if err != nil {
		panic(err)
	}

	return string(b)
}
