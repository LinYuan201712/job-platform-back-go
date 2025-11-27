package database

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"job-platform-go/internal/config"
)

// AutoMigrate 对应 Flyway 的启动自动迁移功能
func AutoMigrate() {
	// 1. 拼接 DSN
	cfg := config.GlobalConfig.Datasource
	dsn := fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s?charset=%s&multiStatements=true",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.Charset,
	)

	// 2. 创建迁移实例
	// "file://db/migrations" 对应你的资源文件夹路径
	m, err := migrate.New("file://resource/db/migration", dsn)
	if err != nil {
		log.Fatalf("❌ 迁移初始化失败: %v", err)
	}

	// 3. 执行 Up()，这就相当于 Flyway 的 migrate 命令
	// 它会自动判断当前数据库版本，把还没执行的脚本（V2, V3...）全部执行一遍
	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("✅ 数据库已是最新版本，无需迁移")
		} else {
			log.Fatalf("❌ 数据库迁移失败: %v", err)
		}
	} else {
		log.Println("✅ 数据库迁移成功，已更新到最新版本")
	}
}
