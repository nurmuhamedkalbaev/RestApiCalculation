package handlers

import (
	"awesomeProject/internals/calculatorService"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CalculationHandler struct {
	service calculatorService.CalculationService
}

func NewCalculationHandler(s calculatorService.CalculationService) *CalculationHandler {
	return &CalculationHandler{service: s}
}
func (h *CalculationHandler) GetCalculations(c echo.Context) error {
	calculations, err := h.service.GetAllCalculation()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "could not get calculation"})
	}
	return c.JSON(http.StatusOK, calculations)

}
func (h *CalculationHandler) PostCalculation(c echo.Context) error {
	var req calculatorService.CalculationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	calc, err := h.service.CreateCalculation(req.Expression)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "could not create calculation"})
	}
	return c.JSON(http.StatusCreated, calc)
}

func (h *CalculationHandler) PatchCalculation(c echo.Context) error {
	id := c.Param("id")

	var req calculatorService.CalculationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	updateCalc, err := h.service.UpdateCalculation(id, req.Expression)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "could not update calculation"})
	}
	return c.JSON(http.StatusOK, updateCalc)

}

func (h *CalculationHandler) DeleteCalculation(c echo.Context) error {
	id := c.Param("id")
	if err := h.service.DeleteCalculation(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "could not delete calculation"})
	}
	return c.NoContent(http.StatusNoContent)
}
