package shared

import (
	"github.com/google/uuid"
)

func GenerateUUIDV4() uuid.UUID {
	return uuid.New()
}

func GenerateUUIDV6() uuid.UUID {
	return uuid.Must(uuid.NewV6())
}

func GenerateUUIDV7() uuid.UUID {
	return uuid.Must(uuid.NewV7())
}
