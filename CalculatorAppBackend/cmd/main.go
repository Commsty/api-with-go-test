package main

import (
	"calc/internal/db"
	"calc/internal/handlers"
	"calc/internal/service"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	e := echo.New()

	calcRepo := service.NewCalcRepository(database)
	calcService := service.NewCalcService(calcRepo)
	calcHandlers := handlers.NewCalcHandler(calcService)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/calculations", calcHandlers.GetCalculations)
	e.POST("/calculations", calcHandlers.PostCalculations)
	e.PATCH("/calculations/:id", calcHandlers.PatchCalculations)
	e.DELETE("/calculations/:id", calcHandlers.DeleteCalculation)

	e.Start("localhost:8080")
}
