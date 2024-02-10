package commands

import (
	"github.com/artarts36/quicktool/internal/infrastructure/filesystem"
	"github.com/artarts36/quicktool/internal/presentation/interaction"
)

func CreateList(
	fs filesystem.FileSystem,
) []interaction.Command {
	return []interaction.Command{
		NewUUIDCommand(),
		NewMD5Command(),
		NewJSONPrettyCommand(),
		NewTimestampCommand(),
		NewMkdirsCommand(),
		NewJSONPathCommand(fs),
		NewUserCommand(),
		NewPasswordCommand(),
	}
}
