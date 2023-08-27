package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
	"github.com/laurentpoirierfr/ms-mongodb-api/internal/handlers"
	"github.com/laurentpoirierfr/ms-mongodb-api/internal/repositories"
	"github.com/laurentpoirierfr/ms-mongodb-api/util"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		slog.Fatal("cannot load config")
	}

	slog.Configure(func(logger *slog.SugaredLogger) {
		f := logger.Formatter.(*slog.TextFormatter)
		f.EnableColor = true
	})

	router := gin.Default()
	router.GET("/ops/ping", handlers.Ping)

	// Instanciate Repository
	repo := repositories.NewMongoRepository(&config)
	hdls := handlers.NewApiHandler(repo)

	api := router.Group("/api")
	{
		api.GET("/:documents", hdls.FindDocuments)
		api.GET("/:documents/:id", hdls.FindOneDocument)
		api.POST("/:documents", hdls.CreateDocument)
		api.PUT("/:documents/:id", hdls.UpdateDocument)
		api.DELETE("/:documents/:id", hdls.DeleteDocument)
	}

	slog.Info("Server started.")
	router.Run(":" + config.Port)
	slog.Info("Server stopped.")
}
