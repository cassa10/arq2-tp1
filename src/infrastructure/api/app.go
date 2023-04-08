package api

import (
	"fmt"
	swaggerDocs "github.com/cassa10/arq2-tp1/docs"
	"github.com/cassa10/arq2-tp1/src/domain/action/command"
	"github.com/cassa10/arq2-tp1/src/domain/action/query"
	v1 "github.com/cassa10/arq2-tp1/src/infrastructure/api/v1"
	"github.com/cassa10/arq2-tp1/src/infrastructure/config"
	"github.com/cassa10/arq2-tp1/src/infrastructure/logger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io"
	"net/http"
)

// Application
// @title arq2-tp1 API
// @version 1.0
// @description api for tp arq2-tp1
// @contact.name API SUPPORT
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @BasePath /
// @query.collection.format multi
type Application interface {
	Run() error
}

type application struct {
	logger logger.Logger
	config config.Config
	*ApplicationUseCases
}

type ApplicationUseCases struct {
	CreateCustomerCmd *command.CreateCustomer
	UpdateCustomerCmd *command.UpdateCustomer
	DeleteCustomerCmd *command.DeleteCustomer
	FindCustomerQuery *query.FindCustomer
}

func NewApplication(l logger.Logger, conf config.Config, applicationUseCases *ApplicationUseCases) Application {
	return &application{
		logger:              l,
		config:              conf,
		ApplicationUseCases: applicationUseCases,
	}
}

func (app *application) Run() error {
	swaggerDocs.SwaggerInfo.Host = fmt.Sprintf("localhost:%v", app.config.Port)

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard

	router := gin.Default()
	router.GET("/", HealthCheck)

	rv1 := router.Group("/api/v1")
	//customer
	{
		rv1Customer := rv1.Group("/customer")
		rv1Customer.POST("", v1.CreateCustomerHandler(app.logger, app.CreateCustomerCmd))
		rv1Customer.GET("/:customerId", v1.FindCustomerHandler(app.logger, app.FindCustomerQuery))
		rv1Customer.DELETE("/:customerId", v1.DeleteCustomerHandler(app.logger, app.DeleteCustomerCmd))
		rv1Customer.PUT("/:customerId", v1.UpdateCustomerHandler(app.logger, app.UpdateCustomerCmd))
	}

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app.logger.Infof("running http server on port %d", app.config.Port)
	return router.Run(fmt.Sprintf(":%d", app.config.Port))
}

// HealthCheck
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags Health check
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{"data": "Server is up and running"})
}
