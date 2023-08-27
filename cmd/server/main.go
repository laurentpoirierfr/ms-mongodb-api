package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
	"github.com/laurentpoirierfr/ms-mongodb-api/internal/handlers"
	"github.com/laurentpoirierfr/ms-mongodb-api/internal/repositories"
	"github.com/laurentpoirierfr/ms-mongodb-api/util"

	_ "github.com/laurentpoirierfr/ms-mongodb-api/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger ms-Mongodb-api
// @version 1.0
// @description This is a ms-Mongodb-api server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host homezone.swagger.io:8080
// @BasePath /
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	// Check health
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
