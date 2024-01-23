package model

import "time"

type AppContext struct {
}

type LoginAccount struct {
	UserId         int64
	TenantId       int64
	OrganizationId int64
	RoleId         int64
	DeptId         int64
	PostId         int64
	Username       string
	RoleKey        string
}

type OauthAccount struct {
	ID        uint       `json:"id" gorm:"autoIncrement;primaryKey"`
	Name      string     `json:"name" gorm:"size:100;not null;unique"`
	Password  string     `json:"-" gorm:"size:256;"`
	Email     string     `json:"email" gorm:"size:256;"`
	Avatar    string     `json:"avatar" gorm:"size:256;"`
	AuthInfos []AuthInfo `json:"authInfos" gorm:"foreignKey:UserId;references:ID"`
}

type AuthInfo struct {
	ID           uint      `json:"id" gorm:"autoIncrement;primaryKey"`
	UserId       uint      `json:"userId" gorm:"size:256"`
	Url          string    `json:"url" gorm:"size:256"`
	AuthType     string    `json:"authType" gorm:"size:256"`
	AuthId       string    `json:"authId" gorm:"size:256"`
	AccessToken  string    `json:"-" gorm:"size:256"`
	RefreshToken string    `json:"-" gorm:"size:256"`
	Expiry       time.Time `json:"-"`

	BaseModel
}
