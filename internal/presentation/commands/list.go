package commands

import (
	"github.com/artarts36/quicktool/internal/infrastructure/filesystem"
	"github.com/artarts36/quicktool/internal/infrastructure/sh"
	"github.com/artarts36/quicktool/internal/presentation/interaction"
)

func CreateList(
	fs filesystem.FileSystem,
) []interaction.Command {
	return []interaction.Command{
		NewUUID(),
		NewMD5(),
		NewJSONPretty(),
		NewTimestamp(),
		NewMkdirs(),
		NewJSONPath(fs),
		NewUser(),
		NewPassword(),
		NewBase64(fs),
		NewRange(),
		NewGpg(),
		NewUntar(&sh.Tar{}),
	}
}
