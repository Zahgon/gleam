//go:build plan9
// +build plan9

package on_interrupt

func OnInterrupt(fn func(), onExitFunc func()) { _ = "STUB: not implemented"; return }
