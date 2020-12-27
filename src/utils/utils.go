package utils

import (
	"os"

	"github.com/google/uuid"
)

type Reply int

type RegisterComponentArgs struct {
	ComponentType string
	ID            uuid.UUID
}

// RunningInDocker Returns if the code is running within a Docker container
func RunningInDocker() bool {
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}

	return false
}
