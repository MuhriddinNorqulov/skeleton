package payment

type PaymentURLParams struct {
	AccountID   uint
	Amount      int64
	RedirectURL string
}

type PaymentProvider interface {
	PaymentURL(method string, p *PaymentURLParams) string
}
