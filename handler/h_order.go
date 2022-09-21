package handler

import (
	"assign2-new/helper"
	"assign2-new/orders"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	s orders.Service
}

func NewOrderHandler(s orders.Service) *orderHandler {
	return &orderHandler{s}
}

func (h *orderHandler) CreateOrder(c *gin.Context) {
	var input orders.OrderInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create order 1", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newOrder, err := h.s.Create(input)
	if err != nil {
		response := helper.APIResponse("Failed to create order, customer id not found", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create order", http.StatusOK, "success", orders.FormatOrder(newOrder))
	c.JSON(http.StatusOK, response)
}

func (h *orderHandler) GetOrders(c *gin.Context) {
	order, err := h.s.GetOrders()
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to get All Orders", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Orders", http.StatusOK, "success", orders.FormatOrders(order))
	c.JSON(http.StatusOK, response)
}

func (h *orderHandler) UpdateOrder(c *gin.Context) {
	var inputID orders.GetOrderByID

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update Order", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData orders.OrderUpdate

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update Order", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// currentUser := c.MustGet("currentUser").(user.User)
	// inputData.User = currentUser

	updatedOrder, err := h.s.UpdateOrder(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update Order", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update Order", http.StatusOK, "success", orders.FormatOrder(updatedOrder))
	c.JSON(http.StatusOK, response)
}

func (h *orderHandler) GetOrderByID(c *gin.Context) {
	var inputID orders.GetOrderByID

	err := c.ShouldBindUri("id")
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to get Order, order id not found", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	order, err := h.s.GetOrderByID(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to get order", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success Get Order", http.StatusOK, "success", orders.FormatOrder(order))
	c.JSON(http.StatusOK, response)
}

func (h *orderHandler) DeleteOrder(c *gin.Context) {

	idparam := c.Param("id")
	id, err := strconv.Atoi(idparam)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to delete Order, order id not found", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	order, err := h.s.DeleteOrder(id)
	if err != nil {
		response := helper.APIResponse("Failed to delete order", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete Order", http.StatusOK, "success", orders.FormatOrder(order))
	c.JSON(http.StatusOK, response)
}
