package pg

import (
	"database/sql"
	"fmt"
	"github.com/loveyu233/go-utils/client"
	"github.com/loveyu233/go-utils/public"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var pgClient *gorm.DB

type PgConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	Debug    bool
}

var (
	DefaultPgConfig = PgConfig{
		Host:     "127.0.0.1",
		Port:     5432,
		User:     "pgsql",
		Password: "pgsql",
		Database: "pgsql",
		Debug:    true,
	}
)

func MustInitPgSqlClient(config ...PgConfig) *gorm.DB {
	var (
		err      error
		sqlDB    *sql.DB
		pgConfig PgConfig
	)

	if len(config) == 0 {
		pgConfig = DefaultPgConfig
	} else {
		pgConfig = config[0]
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		pgConfig.Host,
		pgConfig.User,
		pgConfig.Password,
		pgConfig.Database,
		pgConfig.Port,
	)
	if pgClient, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		logrus.Panicf("pgsql 连接失败: %v", err)
	}

	if pgConfig.Debug {
		pgClient = pgClient.Debug()
	}

	logrus.Info("pgsql 连接成功")

	if sqlDB, err = pgClient.DB(); err != nil {
		logrus.Panicf("pgsql 获取数据库失败: %v", err)
	}

	sqlDB.SetConnMaxLifetime(2 * time.Hour)
	sqlDB.SetMaxIdleConns(10)

	client.PgSqlClient = func(cfg ...public.UseConfig) *gorm.DB {
		var dUseConfig public.UseConfig
		if len(cfg) == 0 {
			dUseConfig = public.DefaultUseConfig
		} else {
			dUseConfig = cfg[0]
		}
		tx := pgClient.Session(&gorm.Session{}).WithContext(dUseConfig.Ctx)

		if dUseConfig.Debug {
			tx = tx.Debug()
		}

		return tx
	}

	return pgClient
}
