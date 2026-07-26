package payment

import (
	"github.com/muhriddinnorqulov/skeleton/src/core/domain/ports/payment"
	"github.com/muhriddinnorqulov/skeleton/src/infrastructure/env"
)

type PaymentAdapter struct {
	env *env.Env
}

// @inject
func NewPaymentAdapter(env *env.Env) payment.PaymentProvider {
	return &PaymentAdapter{env: env}
}

func (this *PaymentAdapter) PaymentURL(method string, p *payment.PaymentURLParams) string {
	// TODO: to'lov tizimi URL'ini generatsiya qilish
	return ""
}
