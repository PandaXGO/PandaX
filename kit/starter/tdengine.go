package starter

import (
	"database/sql"
	"fmt"
)

type TaosDB struct {
	Username string
	Password string
	Host     string
	Port     int
	Dbname   string
	Config   string //配置
	Db       *sql.DB
}

func (d *TaosDB) InitTdDB() error {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",
		d.Username, d.Password, "http", d.Host, d.Port, d.Dbname)
	if d.Config != "" {
		dsn = fmt.Sprintf("%s&%s", dsn, d.Config)
	}
	var err error
	d.Db, err = sql.Open("taosRestful", dsn)
	return err
}
