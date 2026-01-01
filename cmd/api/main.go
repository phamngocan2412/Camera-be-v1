package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/phamngocan2412/camera-be-v1/internal/config"
	"github.com/phamngocan2412/camera-be-v1/internal/handlers"
	"github.com/phamngocan2412/camera-be-v1/internal/middleware"
	"github.com/phamngocan2412/camera-be-v1/internal/platform/db"
	"github.com/phamngocan2412/camera-be-v1/internal/platform/logger"
	"github.com/phamngocan2412/camera-be-v1/internal/repository"
	"github.com/phamngocan2412/camera-be-v1/internal/service"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	zapLogger, err := logger.NewLogger(cfg.Log.Level)
	if err != nil {
		panic(err)
	}
	defer zapLogger.Sync()

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(logger.GinLogger(zapLogger))

	dbConn, err := db.NewDatabase(cfg.Database.URL)
	if err != nil {
		zapLogger.Fatal("database connection failed", zap.Error(err))
	}

	userRepo := repository.NewGORMUserRepository(dbConn)
	authService := service.NewAuthService(userRepo, cfg.JWT.Secret)
	userService := service.NewUserService(userRepo)

	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)

	// Public
	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	// Protected
	api := r.Group("/api")
	api.Use(middleware.JWTAuth(cfg.JWT.Secret))
	{
		users := api.Group("/users")
		{
			users.GET("/me", userHandler.GetMe)
			users.PUT("/me", userHandler.UpdateMe)
			users.PUT("/me/password", userHandler.ChangePassword)
		}
	}

	zapLogger.Info("Server starting", zap.String("port", cfg.Server.Port))
	r.Run(cfg.Server.Port)
}
