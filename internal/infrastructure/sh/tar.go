package sh

import (
	"context"
	"os/exec"
)

type Tar struct {
}

type TarExtractParams struct {
	ArchivePath string
}

func (t *Tar) Extract(ctx context.Context, params *TarExtractParams) ([]byte, error) {
	return exec.CommandContext(ctx, "tar", "-xf", params.ArchivePath).Output()
}
