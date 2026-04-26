package model

type PaymentState string

const (
	Pending    PaymentState = "PENDING"
	Successful PaymentState = "SUCCESS"
	Failed     PaymentState = "FAILED"
	Refunded   PaymentState = "REFUNDED"
)

func (ps PaymentState) IsFinal() bool {
	return ps == Failed || ps == Refunded
}
