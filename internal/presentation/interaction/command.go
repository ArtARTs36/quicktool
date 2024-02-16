package interaction

import (
	"context"

	"github.com/artarts36/quicktool/internal/domain"
)

type Command interface {
	Definition() *Definition
	Execute(ctx *Context, env *Env) error
}

type Context struct {
	Context context.Context
	User    *domain.User
}

type Definition struct {
	Name        string
	Aliases     []string
	Description string
	Args        []*DefinitionArg
	Opts        []*DefinitionOpt
}

type DefinitionArg struct {
	Name        string
	Description string
	Required    bool
	ValuesEnum  []string
}

func (a *DefinitionArg) IsEnum() bool {
	return len(a.ValuesEnum) > 0
}

type DefinitionOpt struct {
	Name        string
	ShortName   string
	Description string
}
