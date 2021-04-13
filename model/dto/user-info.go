package dto

type UserInfo struct {
	Nickname		string	`json:",omitempty" gorm:"type:varchar(64);index"`
	Avatar			string	`json:",omitempty" gorm:"type:varchar(255);"`
	Digest			string	`json:",omitempty" gorm:"type:varchar(65535);"`
}
