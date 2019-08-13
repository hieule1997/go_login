package utils

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func GenUUID() string {
	id := uuid.Must(uuid.NewV4())
	return fmt.Sprintf("%v", id)
}