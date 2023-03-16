package api

import (
	"fmt"
	"github.com/cassa10/arq2-tp1/src/infrastructure/config"
	"github.com/cassa10/arq2-tp1/src/infrastructure/logger"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type Application interface {
	Run() error
}

type application struct {
	logger logger.Logger
	config config.Config
}

func NewApplication(l logger.Logger, conf config.Config) Application {
	return &application{
		logger: l,
		config: conf,
	}
}

func (app *application) Run() error {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	rv1 := router.Group("/api/v1")
	rv1.GET("/asd", func(c *gin.Context) { c.String(http.StatusOK, "asd ok") })

	app.logger.Infof("running http server on port %d", app.config.Port)
	return router.Run(fmt.Sprintf(":%d", app.config.Port))
}
