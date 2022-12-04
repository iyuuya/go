package exitcode

// ExitCode is exit code
type ExitCode int

const (
	// ExitOK is succeeded
	ExitOK ExitCode = 0
	// ExitError is failed
	ExitError ExitCode = 1
)
