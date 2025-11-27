package main

import (
	"fmt"
	"job-platform-go/internal/config"
	"job-platform-go/internal/controller"
	"job-platform-go/internal/middleware"
	"job-platform-go/pkg/database"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	//1.初始化配置
	config.InitConfig()
	//2.连接数据库
	database.InitDB()

	database.AutoMigrate()
	//3.初始化路由
	r := gin.Default()
	// 初始化 Controller

	authCtrl := controller.NewAuthController()
	companyCtrl := controller.NewCompanyController()
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", authCtrl.Login)
		authGroup.POST("/register", authCtrl.Register)
	}
	hrGroup := r.Group("/hr", middleware.JWTAuth())

	{
		hrGroup.GET("/company/profile", companyCtrl.GetProfile)
		hrGroup.PUT("/company/profile", companyCtrl.UpdateProfile)
	}

	r.GET("/ping", func(c *gin.Context) {
		sqlDB, err := database.DB.DB()
		if err != nil {
			log.Fatalf("获取数据库实例失败：%v", err)
			c.JSON(500, gin.H{"message": "Database error"})
			return
		}
		if err := sqlDB.Ping(); err != nil {
			log.Fatalf("❌ 数据库 Ping 失败: %v", err)
			c.JSON(500, gin.H{"message": "Database ping failed"})
			return
		}
		c.JSON(200, gin.H{
			"message":   "pong",
			"db_status": "connected",
			"app_name":  "Job Platform Go",
		})
		fmt.Println("✅ 成功连接到 MySQL 数据库！")
	})
	//启动服务
	port := config.GlobalConfig.Server.Port
	fmt.Printf("Server is running on port %s\n", port)
	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}
