package client

import (
	"github.com/go-redis/redis/v8"
	"github.com/loveyu233/go-utils/public"
	"github.com/minio/minio-go/v7"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
)

var (
	PgSqlClient func(cfg ...public.UseConfig) *gorm.DB

	MySqlClient func(cfg ...public.UseConfig) *gorm.DB

	EsClient *elastic.Client

	RedisClient *redis.Client

	MinioClient *minio.Client
)
