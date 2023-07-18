package data

import (
	"github.com/xushuhui/kratos-microservice-layout/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewDb, NewsRedis)

type Data struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewDb(conf *conf.Data_Database, logger log.Logger) *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       conf.Source, // DSN data source name
		DefaultStringSize:         256,         // string 类型字段的默认长度
		DisableDatetimePrecision:  true,        // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,        // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,        // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,       // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		log.NewHelper(logger).Fatalf("failed opening connection to db: %v", err)
	}

	return db
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	db := NewDb(c.Database, logger)
	rdb := NewsRedis(c.Redis, logger)
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		rdb.Close()
	}
	return &Data{
		db:  db,
		rdb: rdb,
	}, cleanup, nil
}

func NewsRedis(c *conf.Data_Redis, logger log.Logger) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       int(c.Db),
	})
}
