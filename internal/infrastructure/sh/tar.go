package sh

import (
	"bytes"
	"context"
	"errors"
	"os/exec"
	"strings"
)

type Tar struct {
}

type TarExtractParams struct {
	ArchivePath string
}

func (t *Tar) Extract(ctx context.Context, params *TarExtractParams) ([]byte, error) {
	var errBuffer bytes.Buffer

	cmd := exec.CommandContext(ctx, "tar", "-xf", params.ArchivePath)
	cmd.Stderr = &errBuffer
	res, err := cmd.Output()
	if err != nil {
		err = errors.New(strings.Join([]string{
			err.Error(),
			errBuffer.String(),
		}, ": "))
	}

	return res, err
}
