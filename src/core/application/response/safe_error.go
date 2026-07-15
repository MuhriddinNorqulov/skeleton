package response

type SafeError struct {
	Code     Code
	Internal error
	Caller   string
}

func NewSafeError(code Code, internal error, caller string) *SafeError {
	return &SafeError{Code: code, Internal: internal, Caller: caller}
}

func (this *SafeError) Error() string {
	return this.Internal.Error()
}
