package main

import (
	uc "food-ordering/internal/application/usecase/user"
	"food-ordering/internal/config"
	"food-ordering/internal/infrastructure/repository/user"
	"food-ordering/internal/presentation/handler"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
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

	menuRepo := user.NewRepository(db)
	menuUseCase := uc.NewUC(menuRepo)

	e := echo.New()
	menuHandler := handler.NewHandler(menuUseCase)
	e.GET("/menus", menuHandler.HandleMenu)

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
