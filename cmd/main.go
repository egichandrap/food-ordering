package main

import (
	ucResto "food-ordering/internal/application/usecase/resto"
	"food-ordering/internal/application/usecase/user"
	"food-ordering/internal/config"
	"food-ordering/internal/infrastructure/repository"
	"food-ordering/internal/presentation/handler"
	"food-ordering/internal/presentation/middlewares"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Warningln(err)
	}

	db, err := config.PSQL{
		DSN:                os.Getenv("DSN"),
		MaxConnections:     os.Getenv("MAX_CONNECTIONS"),
		MaxIdleConnections: os.Getenv("MAX_IDLE_CONNECTIONS"),
		ConnMaxLifetime:    os.Getenv("CONNECTION_MAX_LIFE_TIME"),
	}.Connect()
	checkErr(err)

	err = db.Ping()
	checkErr(err)

	log.Info("database is connected")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	registerRepo := repository.NewUserRepository()
	registerUseCase := user.NewUserUseCase(registerRepo)
	registerHandler := handler.NewUserHandler(registerUseCase)
	e.POST("/register", registerHandler.HandleRegister)

	menuRepo := repository.NewCatalogRepository(db)
	menuUseCase := ucResto.NewCatalogUseCase(menuRepo)
	menuHandler := handler.NewCatalogHandler(menuUseCase)
	e.GET("/menus", menuHandler.HandleCatalogMenu, middlewares.JWTMiddleware)

	cartRepo := repository.NewCartRepository()
	cartUseCase := ucResto.NewCartUseCase(cartRepo, nil)
	cartHandler := handler.NewCartHandler(cartUseCase)
	e.POST("/add-cart", cartHandler.HandleCrtItem, middlewares.JWTMiddleware)
	e.POST("/checkout", cartHandler.HandleCheckout, middlewares.JWTMiddleware)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8002"
	}

	log.Infof("Starting server on port %s", port)
	e.Logger.Fatal(e.Start(":" + port))

}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
