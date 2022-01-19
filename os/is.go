package os

import (
	"runtime"
)

// IsMac returns true when macOS.
func IsMac() bool {
	return runtime.GOOS == "darwin"
}

// IsWindows returns true when Windows.
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// IsLinux returns true when Linux.
func IsLinux() bool {
	return runtime.GOOS == "linux"
}
