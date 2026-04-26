package main

import (
	"mock_payment_gateway/internal/handler"
	"mock_payment_gateway/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	paymentService := service.NewPaymentService()

	paymentHandler := handler.NewPaymentHandler(paymentService)

	r := gin.Default()

	r.POST("/payments", paymentHandler.CreatePayment)
	r.GET("/payments/:id", paymentHandler.GetPayment)
	r.PUT("/payments/:id", paymentHandler.UpdatePayment)

	r.Run(":8080")
}
