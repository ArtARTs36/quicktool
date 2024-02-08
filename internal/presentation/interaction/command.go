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
}

type DefinitionArg struct {
	Name        string
	Description string
	Required    bool
	ValuesEnum  []string
}
