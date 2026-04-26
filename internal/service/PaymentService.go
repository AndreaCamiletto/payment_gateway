package service

import (
	"errors"
	"mock_payment_gateway/internal/model"
	"sync"
	"time"

	"github.com/google/uuid"
)

var ErrPaymentNotFound = errors.New("payment not found")

type PaymentService struct {
	paymentsMap map[uuid.UUID]model.Payment
	mu          sync.RWMutex
}

func NewPaymentService() *PaymentService {
	return &PaymentService{
		paymentsMap: make(map[uuid.UUID]model.Payment),
	}
}

func (s *PaymentService) CreatePayment(amount int64, currency string) model.Payment {
	payment := model.Payment{ID: uuid.New(),
		Amount:    amount,
		Currency:  currency,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status:    model.Pending}
	s.mu.Lock()
	s.paymentsMap[payment.ID] = payment
	s.mu.Unlock()
	return payment
}

func (s *PaymentService) GetPayment(id uuid.UUID) (model.Payment, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	payment, ok := s.paymentsMap[id]
	if !ok {
		return model.Payment{}, ErrPaymentNotFound
	}
	return payment, nil
}

func (s *PaymentService) UpdatePayment(id uuid.UUID) (model.Payment, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	payment, ok := s.paymentsMap[id]
	if !ok {
		return model.Payment{}, ErrPaymentNotFound
	}

	if payment.Status.IsFinal() {
		return model.Payment{}, errors.New("cannot update status of a finalized payment")
	}

	payment.UpdatedAt = time.Now()
	payment.Status = model.Successful

	s.paymentsMap[id] = payment
	return payment, nil
}
