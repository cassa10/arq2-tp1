package main

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/action/command"
	"github.com/cassa10/arq2-tp1/src/domain/action/query"
	"github.com/cassa10/arq2-tp1/src/infrastructure/api"
	"github.com/cassa10/arq2-tp1/src/infrastructure/config"
	"github.com/cassa10/arq2-tp1/src/infrastructure/logger"
	"github.com/cassa10/arq2-tp1/src/infrastructure/repository/mongo"
)

func main() {
	conf := config.LoadConfig()
	baseLogger := logger.New(&logger.Config{
		ServiceName:     "arq2-tp1",
		EnvironmentName: conf.Environment,
		LogLevel:        conf.LogLevel,
		LogFormat:       logger.JsonFormat,
	})

	db := mongo.Connect(context.Background(), baseLogger, conf.MongoURI, conf.MongoDatabase)
	customerRepo := mongo.NewCustomerRepository(baseLogger, db, conf.MongoTimeout)
	findCustomerQuery := query.NewFindCustomer(customerRepo)
	createCustomerCmd := command.NewCreateCustomer(customerRepo)
	updateCustomerCmd := command.NewUpdateCustomer(customerRepo, *findCustomerQuery)
	deleteCustomerCmd := command.NewDeleteCustomer(customerRepo, *findCustomerQuery)

	appUseCases := &api.ApplicationUseCases{
		CreateCustomerCmd: createCustomerCmd,
		UpdateCustomerCmd: updateCustomerCmd,
		FindCustomerQuery: findCustomerQuery,
		DeleteCustomerCmd: deleteCustomerCmd,
	}
	app := api.NewApplication(baseLogger, conf, appUseCases)
	baseLogger.Fatal(app.Run())
}
