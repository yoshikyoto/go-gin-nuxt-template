package database

import (
	"api-server/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connector はデータベースに接続するためのstruct
type Connector struct {
	env *config.Env
}

// NewConnector は Connector のコンストラクタ
func NewConnector() *Connector {
	return &Connector{
		env: config.NewEnv(),
	}
}

// Connect はDBに接続します
func (c Connector) Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.env.GetDBUser(),
		c.env.GetDBPassword(),
		c.env.GetDBHost(),
		c.env.GetDBName(),
	)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
