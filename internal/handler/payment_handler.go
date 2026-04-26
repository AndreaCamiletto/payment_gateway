package handler

import (
	"mock_payment_gateway/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PaymentHandler struct {
	pService *service.PaymentService
}

func NewPaymentHandler(pService *service.PaymentService) *PaymentHandler {
	return &PaymentHandler{pService: pService}
}

func (p *PaymentHandler) CreatePayment(c *gin.Context) {
	var req CreatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	payment := p.pService.CreatePayment(req.Amount, req.Currency)
	c.JSON(http.StatusCreated, payment)
}

func (p *PaymentHandler) GetPayment(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid UUID",
		})
		return
	}
	payment, err := p.pService.GetPayment(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, payment)
}

func (p *PaymentHandler) UpdatePayment(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid UUID",
		})
		return
	}
	payment, err := p.pService.UpdatePayment(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, payment)
}

type CreatePaymentRequest struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}
