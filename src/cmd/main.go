package main

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/action/command"
	"github.com/cassa10/arq2-tp1/src/domain/action/query"
	"github.com/cassa10/arq2-tp1/src/domain/usecase"
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
	mongoDB := mongo.Connect(context.Background(), baseLogger, conf.MongoURI, conf.MongoDatabase)

	//repositories
	customerRepo := mongo.NewCustomerRepository(baseLogger, mongoDB, conf.MongoTimeout)
	sellerRepo := mongo.NewSellerRepository(baseLogger, mongoDB, conf.MongoTimeout)
	productRepo := mongo.NewProductRepository(baseLogger, mongoDB, conf.MongoTimeout)
	orderRepo := mongo.NewOrderRepository(baseLogger, mongoDB, conf.MongoTimeout, conf.MongoDatabase)

	//customer
	findCustomerByIdQuery := query.NewFindCustomerById(customerRepo)
	createCustomerCmd := command.NewCreateCustomer(customerRepo)
	updateCustomerCmd := command.NewUpdateCustomer(customerRepo, *findCustomerByIdQuery)
	deleteCustomerCmd := command.NewDeleteCustomer(customerRepo, *findCustomerByIdQuery)

	//seller
	findSellerByIdQuery := query.NewFindSellerById(sellerRepo, productRepo)
	createSellerCmd := command.NewCreateSeller(sellerRepo)
	updateSellerCmd := command.NewUpdateSeller(sellerRepo, *findSellerByIdQuery)
	deleteSellerCmd := command.NewDeleteSeller(sellerRepo, *findSellerByIdQuery)

	//product
	findProductByIdQuery := query.NewFindProductById(productRepo)
	createProductCmd := command.NewCreateProduct(productRepo, *findSellerByIdQuery)
	updateProductCmd := command.NewUpdateProduct(productRepo, *findProductByIdQuery)
	deleteProductCmd := command.NewDeleteProduct(productRepo, *findProductByIdQuery)
	searchProductQuery := query.NewSearchProduct(productRepo)

	//order
	createOrderUseCase := usecase.NewCreateOrder(baseLogger, orderRepo, *findProductByIdQuery, *findCustomerByIdQuery)
	findOrderByIdQuery := query.NewFindOrderById(orderRepo)

	app := api.NewApplication(baseLogger, conf, &api.ApplicationUseCases{
		FindCustomerQuery: findCustomerByIdQuery,
		CreateCustomerCmd: createCustomerCmd,
		UpdateCustomerCmd: updateCustomerCmd,
		DeleteCustomerCmd: deleteCustomerCmd,

		FindSellerQuery: findSellerByIdQuery,
		CreateSellerCmd: createSellerCmd,
		UpdateSellerCmd: updateSellerCmd,
		DeleteSellerCmd: deleteSellerCmd,

		FindProductQuery:   findProductByIdQuery,
		CreateProductCmd:   createProductCmd,
		UpdateProductCmd:   updateProductCmd,
		DeleteProductCmd:   deleteProductCmd,
		SearchProductQuery: searchProductQuery,

		CreateOrderUseCase: createOrderUseCase,
		FindOrderQuery:     findOrderByIdQuery,
	})
	baseLogger.Fatal(app.Run())
}
