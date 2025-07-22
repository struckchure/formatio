package types

import (
	"time"

	"github.com/shopspring/decimal"
	"pkg.formatio/lib"
	"pkg.formatio/prisma/db"
)

const (
	CREATE_INVOICE_QUEUE = "create-invoice-queue"
)

type ListInvoiceArgs struct {
	lib.BaseListFilterArgs

	UserId *string
}

type CreateInvoiceArgs struct {
	UserId      string
	ProductId   string
	Reference   string
	Description string
	From        time.Time
	To          time.Time
	Quantity    int
	UnitPrice   decimal.Decimal
	TotalPrice  decimal.Decimal
	Currency    db.Currency
	Status      db.InvoiceStatus
}

type GetInvoiceArgs struct {
	Id string
}

type UpdateInvoiceArgs struct {
	Id          string
	ProductId   *string
	Reference   *string
	Description *string
	From        *time.Time
	To          *time.Time
	Quantity    *int
	UnitPrice   *decimal.Decimal
	TotalPrice  *decimal.Decimal
	Currency    *db.Currency
	Status      *db.InvoiceStatus
}

type DeleteInvoiceArgs struct {
	GetInvoiceArgs
}

type LastInvoiceArgs struct {
	ProductId string
}

type ListCardArgs struct {
	lib.BaseListFilterArgs

	UserId     *string `swaggerignore:"true"`
	IsApproved *bool   `query:"isApproved"`
}

type CreateCardArgs struct {
	UserId      string `swaggerignore:"true"`
	AuthToken   string `json:"authToken"`
	IsDefault   bool   `json:"isDefault"`
	IsApproved  bool   `json:"-" swag-validate:"optional" swaggerignore:"true"`
	CardType    string `json:"cardType"`
	LastDigits  string `json:"lastDigits"`
	ExpiryMonth string `json:"expiryMonth"`
	ExpiryYear  string `json:"expiryYear"`
}

type GetCardArgs struct {
	Id         string
	UserId     *string
	IsDefault  *bool
	IsApproved *bool
}

type UpdateCardArgs struct {
	Id          string  `swaggerignore:"true"`
	AuthToken   *string `json:"AuthToken" swag-validate:"optional"`
	IsDefault   *bool   `json:"IsDefault" swag-validate:"optional"`
	IsApproved  *bool   `json:"IsApproved" swag-validate:"optional"`
	CardType    *string `json:"CardType" swag-validate:"optional"`
	LastDigits  *string `json:"LastDigits" swag-validate:"optional"`
	ExpiryMonth *string `json:"ExpiryMonth" swag-validate:"optional"`
	ExpiryYear  *string `json:"ExpiryYear" swag-validate:"optional"`
}

type DeleteCardArgs struct {
	Id string
}

type PreAuthorizeCardArgs struct {
	UserId      string `swaggerignore:"true"`
	CardNumber  string `json:"cardNumber"`
	Cvv         string `json:"cvv"`
	Pin         string `json:"pin"`
	ExpiryMonth string `json:"expiryMonth"`
	ExpiryYear  string `json:"expiryYear"`
}

type PreAuthorizeCardResult struct {
	Reference string `json:"reference"`
}

type AuthorizeCardArgs struct {
	Reference string `json:"reference"`
	Otp       string `json:"otp"`
}
