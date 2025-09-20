package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"nft-capture-game/internal/config"
	"nft-capture-game/internal/database"
	"nft-capture-game/internal/handlers"
	"nft-capture-game/internal/services"
	"nft-capture-game/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	// 初始化日志
	logger.Init()

	// 加载配置
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// 连接数据库
	db, err := database.Connect(cfg.Database)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 运行数据库迁移
	if err := database.Migrate(db); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// 初始化服务
	userService := services.NewUserService(db)
	nftService := services.NewNFTService(db)
	captureService := services.NewCaptureService(db, userService, nftService)

	// 初始化处理器
	userHandler := handlers.NewUserHandler(userService)
	captureHandler := handlers.NewCaptureHandler(captureService)
	nftHandler := handlers.NewNFTHandler(nftService)

	// 设置Gin模式
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 创建路由器
	router := gin.Default()

	// 配置CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 健康检查
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// API路由组
	api := router.Group("/api/v1")
	{
		// 用户相关路由
		users := api.Group("/users")
		{
			users.POST("/register", userHandler.Register)
			users.GET("/:wallet", userHandler.GetProfile)
			users.PUT("/:wallet", userHandler.UpdateProfile)
			users.GET("/:wallet/stats", userHandler.GetStats)
			users.GET("/leaderboard", userHandler.GetLeaderboard)
			// 新增：游戏相关路由（开始会话、敲击、状态）
			users.POST("/game/start", userHandler.StartGame)
			users.POST("/game/tap", userHandler.Tap)
			users.GET("/game/status/:wallet_address", userHandler.GetGameStatus)
		}

		// 捕捉相关路由
		captures := api.Group("/captures")
		{
			captures.POST("/attempt", captureHandler.AttemptCapture)
			captures.GET("/:wallet/history", captureHandler.GetUserCaptures)
			captures.GET("/recent", captureHandler.GetRecentCaptures)
			captures.GET("/stats", captureHandler.GetCaptureStats)
		}

		// NFT相关路由
		nfts := api.Group("/nfts")
		{
			nfts.POST("/mint", nftHandler.MintNFT)
			nfts.GET("/:wallet", nftHandler.GetUserNFTs)
			nfts.GET("/token/:tokenId", nftHandler.GetNFTByTokenID)
			nfts.GET("/rarity/:rarity", nftHandler.GetNFTsByRarity)
			nfts.GET("/recent", nftHandler.GetRecentNFTs)
			nfts.GET("/stats", nftHandler.GetNFTStats)
		}
	}
	
	// 创建HTTP服务器
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Server.Port),
		Handler: router,
	}

	// 启动服务器
	go func() {
		log.Printf("Server starting on port %s", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to start server:", err)
		}
	}()

	// 等待中断信号以优雅关闭服务器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// 优雅关闭，超时时间为30秒
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited")
}