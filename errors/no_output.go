package errors

type ErrNoOutput struct {
}

func NewErrNoOutput() *ErrNoOutput {
	return &ErrNoOutput{}
}

func (e ErrNoOutput) Error() string {
	return ""
}

func IsErrNoOutput(err error) bool {
	_, ok := err.(*ErrNoOutput)
	return ok
}
