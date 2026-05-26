package util

import (
	"context"
)

func ExecuteWithCleanup(parentContext context.Context, onExecute func() error, onCleanup func()) error {
	_ = "STUB: not implemented"
	return nil
}
