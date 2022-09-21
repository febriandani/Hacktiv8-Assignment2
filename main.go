package main

import (
	"assign2-new/customers"
	"assign2-new/handler"
	"assign2-new/orders"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=junior34 dbname=orders_by port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// db.AutoMigrate(&customers.Customer{})

	//repository
	orderRepository := orders.NewRepository(db)
	customerRepository := customers.NewRepository(db)
	//service
	orderService := orders.NewService(orderRepository)
	customerService := customers.NewCustomerService(customerRepository)
	//handler
	orderHandler := handler.NewOrderHandler(orderService)
	customerHandler := handler.NewCustomerHandler(customerService)

	//router
	router := gin.Default()

	api := router.Group("/api/v1")
	api.POST("/order", orderHandler.CreateOrder)
	api.POST("/customer", customerHandler.CreateCustomer)
	api.GET("/orders", orderHandler.GetOrders)
	api.PUT("/order/:id", orderHandler.UpdateOrder)
	api.DELETE("/delete/:id", orderHandler.DeleteOrder)
	router.Run()
}
