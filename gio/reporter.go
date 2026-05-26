package gio

import (
	"context"
)

func (runner *gleamRunner) report(ctx context.Context, f func() error) error {
	_ = "STUB: not implemented"
	//TODO use context for mapper, reducer, and gleam execute
	return nil
}
