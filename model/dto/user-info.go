package dto

type UserInfo struct {
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Digest   string `json:"digest"`
}

type PutV struct {
	ID           uint   `json:"id"  binding:"required"`
	Verification string `json:"verification"`
}
