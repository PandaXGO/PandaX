package casbin

type CasbinRule struct {
	Ptype   string `json:"ptype" gorm:"column:ptype"`
	RoleKey string `json:"roleKey" gorm:"column:v0"`
	Path    string `json:"path" gorm:"column:v1"`
	Method  string `json:"method" gorm:"column:v2"`
}
