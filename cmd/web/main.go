package main

import (
	"boiler-plate-clean/config"
	"boiler-plate-clean/internal/delivery/http"
	"boiler-plate-clean/internal/delivery/http/route"
	"boiler-plate-clean/internal/repository"
	services "boiler-plate-clean/internal/services"
	"boiler-plate-clean/migration"
	"boiler-plate-clean/pkg/database"
	"boiler-plate-clean/pkg/httpclient"
	"boiler-plate-clean/pkg/logger"
	"boiler-plate-clean/pkg/server"
	"boiler-plate-clean/pkg/xvalidator"
	"log/slog"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

var (
	httpClient    httpclient.Client
	sqlClientRepo *database.Database
)

// @title           Pigeon
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/notificationsvc/api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	validate, _ := xvalidator.NewValidator()
	conf := config.InitAppConfig(validate)
	logger.SetupLogger(&logger.Config{
		AppENV:  conf.AppEnvConfig.AppEnv,
		LogPath: conf.AppEnvConfig.LogFilePath,
		Debug:   conf.AppEnvConfig.AppDebug,
	})
	initInfrastructure(conf)
	ginServer := server.NewGinServer(&server.GinConfig{
		HttpPort:     conf.AppEnvConfig.HttpPort,
		AllowOrigins: conf.AppEnvConfig.AllowOrigins,
		AllowMethods: conf.AppEnvConfig.AllowMethods,
		AllowHeaders: conf.AppEnvConfig.AllowHeaders,
	})

	// repository
	productRepository := repository.NewProductSQLRepository()
	stockRepository := repository.NewStockSQLRepository()
	saleRepository := repository.NewSaleSQLRepository()

	// external api
	//gotifySvcExternalAPI := externalapi.NewProductExternalImpl(conf, httpClient)

	// service
	productService := services.NewProductService(sqlClientRepo.GetDB(), productRepository, validate)
	stockService := services.NewStockService(sqlClientRepo.GetDB(), stockRepository, productRepository, validate)
	saleService := services.NewSaleService(sqlClientRepo.GetDB(), saleRepository, stockRepository, productRepository, validate)
	// Handler
	productHandler := http.NewProductHTTPHandler(productService)
	stockHandler := http.NewStockHTTPHandler(stockService)
	saleHandler := http.NewSaleHTTPHandler(saleService)

	router := route.Router{
		App:            ginServer.App,
		ProductHandler: productHandler,
		StockHandler:   stockHandler,
		SaleHandler:    saleHandler,
	}
	router.Setup()
	router.SwaggerRouter()
	echan := make(chan error)
	go func() {
		echan <- ginServer.Start()
	}()

	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)

	select {
	case <-term:
		slog.Info("signal terminated detected")
	case err := <-echan:
		slog.Error("Failed to start http server", err)
	}
}

func initInfrastructure(config *config.Config) {
	//initPostgreSQL()

	sqlClientRepo = initSQL(config)

	httpClient = initHttpclient()
}

func initSQL(conf *config.Config) *database.Database {
	db := database.NewDatabase(conf.DatabaseConfig.Dbservice, &database.Config{
		DbHost: conf.DatabaseConfig.Dbhost,
		DbUser: conf.DatabaseConfig.Dbuser,
		DbPass: conf.DatabaseConfig.Dbpassword,
		DbName: conf.DatabaseConfig.Dbname,
		DbPort: strconv.Itoa(conf.DatabaseConfig.Dbport),
	})
	if conf.IsStaging() {
		migration.AutoMigration(db)
	}
	return db
}

func initHttpclient() httpclient.Client {
	httpClientFactory := httpclient.New()
	httpClient := httpClientFactory.CreateClient()
	return httpClient
}
