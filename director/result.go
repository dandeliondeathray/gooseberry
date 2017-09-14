package director

type result int

const (
	// Success means that a step finished successfully
	success result = iota
	// Failure means that a step encountered a problem, e.g. compile error, test failure.
	failure result = iota
	// Error means that a step could not be executed as planned, e.g. out of disk space, network
	// error.
	error result = iota
	// Timeout means that the step did not complete within the expected time frame.
	timeout result = iota
)

// Result encapsulates success or failure.
type Result struct {
	result result
}

// SuccessResult creates a success result.
func SuccessResult() Result {
	return Result{success}
}

// FailureResult creates a failure result.
func FailureResult() Result {
	return Result{failure}
}

// ErrorResult creates an error.
func ErrorResult() Result {
	return Result{error}
}

// TimeoutResult creates a timeout result.
func TimeoutResult() Result {
	return Result{timeout}
}
