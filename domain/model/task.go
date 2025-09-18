package model

import (
	"errors"
)

type Status string

const (
	StatusNew      Status = "new"
	StatusComplete Status = "complete"
	StatusPending  Status = "pending"
)

func ValidateStatus(status string) error {
	if Status(status) != StatusNew && Status(status) != StatusPending && Status(status) != StatusComplete {
		return errors.New("invalid status")
	}

	return nil
}

// Task ...
type Task struct {
	Id          int
	Status      Status
	Description string
}
