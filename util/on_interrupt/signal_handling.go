//go:build (darwin || freebsd || netbsd || openbsd) && !plan9 && !windows && !linux
// +build darwin freebsd netbsd openbsd
// +build !plan9
// +build !windows
// +build !linux

package on_interrupt

func OnInterrupt(fn func(), onExitFunc func()) {
	_ = "STUB: not implemented"
	// deal with control+c,etc
	return
}

// controlling terminal close, daemon not exit

// syscall.SIGHUP,
// this causes windows to fail

// syscall.SIGQUIT, // Quit from keyboard, "kill -3"
