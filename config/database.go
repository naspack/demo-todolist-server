package config

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/naspack/demo-todolist-server/models"
)

var DB *gorm.DB

func InitDB() {
	// 获取数据库路径
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "todo.db"
	}

	// 连接数据库
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 自动迁移数据库表结构
	if err := db.AutoMigrate(&models.Todo{}); err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}

	DB = db
	log.Println("数据库初始化成功")
}