package infrastructure

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Orm struct {
	Db *gorm.DB
}

func NewORM(env *Env, logger *zap.Logger) *Orm {
	orm := &Orm{}

	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", env.DbHost, env.DbUser, env.DbPassword, env.DbSid, env.DbPort)

	DB, err := gorm.Open(postgres.Open(connString))
	if err != nil {
		logger.Fatal("failed to connect to database", zap.Error(err))
	}

	//err = DB.AutoMigrate()
	//if err != nil {
	//	logger.Fatal("Error migrating", zap.Error(err))
	//}

	orm.Db = DB

	return orm
}
