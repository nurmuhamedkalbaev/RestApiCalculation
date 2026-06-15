package main

import (
	"awesomeProject/internals/calculatorService"
	"awesomeProject/internals/db"
	"awesomeProject/internals/handlers"

	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	dataBade, err := db.InitDB()
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	e := echo.New()

	calcRepo := calculatorService.NewCalcRepository(dataBade)
	calcService := calculatorService.NewCalculationService(calcRepo)
	calcHandlers := handlers.NewCalculationHandler(calcService)
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/calculations", calcHandlers.GetCalculations)
	e.POST("/calculations", calcHandlers.PostCalculation)
	e.PATCH("/calculations/:id", calcHandlers.PatchCalculation)
	e.DELETE("/calculations/:id", calcHandlers.DeleteCalculation)
	e.Start("localhost:8080")
}
