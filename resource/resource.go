package resource

import (
	"context"
)

type Options map[string]string
type Arguments map[string]string

func (a Arguments) Get(key string) string {
	return a[key]
}

type Resource interface {
	Status() error
	Exec(ctx context.Context, command string, arguments Arguments) error
	DumpStorage() ([]byte, error)
}

func GetExecID(ctx context.Context) string {
	execID, ok := ctx.Value("exec_id").(string)
	if !ok {
		panic("exec_id is not passed in the context")
	}
	return execID
}

func SetExecID(ctx context.Context, execID string) context.Context {
	return context.WithValue(ctx, "exec_id", execID)
}
