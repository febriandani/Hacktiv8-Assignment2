package handler

import (
	"assign2-new/customers"
	"assign2-new/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type customerHandler struct {
	s customers.Service
}

func NewCustomerHandler(s customers.Service) *customerHandler {
	return &customerHandler{s}
}

func (h *customerHandler) CreateCustomer(c *gin.Context) {
	var input customers.CustomerInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create customer", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	newCustomer, err := h.s.CreateCustomer(input)
	if err != nil {
		response := helper.APIResponse("Failed to create customer", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create customer", http.StatusOK, "success", customers.FormatCustomer(newCustomer))
	c.JSON(http.StatusOK, response)
}
