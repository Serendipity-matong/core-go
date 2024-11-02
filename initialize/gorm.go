package initialize

import (
	"MiMengCore/config"
	"MiMengCore/global"
	"MiMengCore/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func InitMySQL() {
	addr, port, username, password, database := config.DB_ADDR, config.DB_PORT, config.DB_USER, config.DB_PASSWORD, config.DB_DATABASE
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, addr, port, database)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,  // Slow SQL threshold
			LogLevel:      logger.Error, // Log level
			Colorful:      true,         // Disable color
		},
	)
	var err error
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MySQL database successfully")

	// 获取 migrator
	migrator := global.DB.Migrator()

	// 检查 content 表是否存在
	if !migrator.HasTable(&model.Content{}) {
		// 创建 content 表
		err = migrator.CreateTable(&model.Content{})
		if err != nil {
			panic(err)
		}
		fmt.Println("Content table created")
	} else {
		fmt.Println("Content table already exists")
	}

	fmt.Println()
}
