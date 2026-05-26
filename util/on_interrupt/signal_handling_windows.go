//go:build windows
// +build windows

package on_interrupt

func OnInterrupt(fn func(), onExitFunc func()) {
	_ = "STUB: not implemented"
	// deal with control+c,etc
	return
}

// controlling terminal close, daemon not exit

// syscall.SIGHUP,
// syscall.SIGINFO, // this causes windows to fail

// syscall.SIGQUIT, // Quit from keyboard, "kill -3"
