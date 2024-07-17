package client

type CommandDispatchingError struct {
	Message   string
	Location  string
	Details   []string
	ErrorCode string
}

func (e CommandDispatchingError) Error() string {
	return e.Message
}
