package main

import (
	"context"
	"log"

	"github.com/artarts36/quicktool/internal/infrastructure/filesystem"
	"github.com/artarts36/quicktool/internal/infrastructure/repository"
	"github.com/artarts36/quicktool/internal/presentation/commands"
	"github.com/artarts36/quicktool/internal/presentation/interaction"
)

func main() {
	ctx := context.TODO()
	fs := &filesystem.LocalFileSystem{}

	userRepo := &repository.OsUserRepository{}
	user, err := userRepo.GetCurrent(ctx)
	if err != nil {
		log.Printf("unable to get current user: %s", err)
	}

	console := interaction.NewConsole(commands.CreateList(fs))

	console.Run(&interaction.Context{
		Context: ctx,
		User:    user,
	})
}
