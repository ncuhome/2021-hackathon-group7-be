package service

const (
	SuccessCode       = iota
	ErrorServer
	ErrorCommitData
	ErrorToken
	ErrorUsernameRepeated
	ErrorLogin
	ErrorRequestRate
)
