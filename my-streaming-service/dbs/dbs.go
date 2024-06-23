package dbs

//引入gorm和postgres驱动
import (
	"fmt"
	"log"
	"my-streaming-service/internal/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// 定义一个全局的DB变量
var DB *gorm.DB

// 初始化数据库连接
func InitDB() {
	var err error
	dsn := "host=localhost user=postgres password=123456 dbname=test port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal(err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(time.Minute * 30)

	fmt.Println("数据库连接成功")

	// 自动迁移数据库表结构

	err = DB.AutoMigrate(&models.User{}, &models.Media{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("数据库表结构自动迁移成功")

}
