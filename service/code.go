package service

const (
	SuccessCode       = iota
	ServerError
	CommitDataError
	TokenError
	UsernameRepeated
	LoginError
	EmailRepeated
	EmailSendingError
	CodeError
	EmailUsed
	OldPasswordError
	EmailNotBinding
	RequestRateError
	UnauthorizedEmail
)
