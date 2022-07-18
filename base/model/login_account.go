package model

type AppContext struct {
}

type LoginAccount struct {
	UserId   int64
	TenantId int64
	RoleId   int64
	DeptId   int64
	PostId   int64
	Username string
	RoleKey  string
}
