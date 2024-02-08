package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/artarts36/quicktool/internal/domain"
)

type OsUserRepository struct {
}

func (r *OsUserRepository) GetCurrent(_ context.Context) (*domain.User, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("unable to get home dir: %s", err.Error())
	}

	return &domain.User{
		Name:    os.Getenv("USER"),
		HomeDir: homeDir,
	}, nil
}
