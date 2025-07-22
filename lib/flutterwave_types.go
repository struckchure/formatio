package lib

import (
	"time"
)

type DirectChargePayloadAuthorization struct {
	Mode string `json:"mode"`
	Pin  string `json:"pin"`
}

type DirectChargePayloadMeta struct {
	CardId string `json:"cardId"`
}

type DirectChargePayload struct {
	CardNumber    string                           `json:"card_number"`
	Cvv           string                           `json:"cvv"`
	ExpiryMonth   string                           `json:"expiry_month"`
	ExpiryYear    string                           `json:"expiry_year"`
	Amount        int32                            `json:"amount"`
	Fullname      string                           `json:"fullname"`
	TxRef         string                           `json:"tx_ref"`
	Currency      string                           `json:"currency"`
	Country       string                           `json:"country"`
	Email         string                           `json:"email"`
	RedirectURL   string                           `json:"redirect_url"`
	Authorization DirectChargePayloadAuthorization `json:"authorization"`
	Meta          DirectChargePayloadMeta          `json:"meta"`
}

type DirectChargeResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		ID                int         `json:"id"`
		TxRef             string      `json:"tx_ref"`
		FlwRef            string      `json:"flw_ref"`
		DeviceFingerprint string      `json:"device_fingerprint"`
		Amount            int         `json:"amount"`
		ChargedAmount     int         `json:"charged_amount"`
		AppFee            float64     `json:"app_fee"`
		MerchantFee       int         `json:"merchant_fee"`
		ProcessorResponse string      `json:"processor_response"`
		AuthModel         string      `json:"auth_model"`
		Currency          string      `json:"currency"`
		IP                string      `json:"ip"`
		Narration         string      `json:"narration"`
		Status            string      `json:"status"`
		AuthURL           string      `json:"auth_url"`
		PaymentType       string      `json:"payment_type"`
		Plan              interface{} `json:"plan"`
		FraudStatus       string      `json:"fraud_status"`
		ChargeType        string      `json:"charge_type"`
		CreatedAt         time.Time   `json:"created_at"`
		AccountID         int         `json:"account_id"`
		Customer          struct {
			ID          int         `json:"id"`
			PhoneNumber interface{} `json:"phone_number"`
			Name        string      `json:"name"`
			Email       string      `json:"email"`
			CreatedAt   time.Time   `json:"created_at"`
		} `json:"customer"`
		Card struct {
			First6Digits string `json:"first_6digits"`
			Last4Digits  string `json:"last_4digits"`
			Issuer       string `json:"issuer"`
			Country      string `json:"country"`
			Type         string `json:"type"`
			Expiry       string `json:"expiry"`
		} `json:"card"`
	} `json:"data"`
	Meta struct {
		Authorization struct {
			Mode     string `json:"mode"`
			Endpoint string `json:"endpoint"`
		} `json:"authorization"`
	} `json:"meta"`
}

type ValidateChargePayload struct {
	Otp    string `json:"otp"`
	FlwRef string `json:"flw_ref"`
	Type   string `json:"type"` // optional
}

type ValidateChargeResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		ID                int       `json:"id"`
		TxRef             string    `json:"tx_ref"`
		FlwRef            string    `json:"flw_ref"`
		DeviceFingerprint string    `json:"device_fingerprint"`
		Amount            int       `json:"amount"`
		ChargedAmount     int       `json:"charged_amount"`
		AppFee            float64   `json:"app_fee"`
		MerchantFee       int       `json:"merchant_fee"`
		ProcessorResponse string    `json:"processor_response"`
		AuthModel         string    `json:"auth_model"`
		Currency          string    `json:"currency"`
		IP                string    `json:"ip"`
		Narration         string    `json:"narration"`
		Status            string    `json:"status"`
		AuthURL           string    `json:"auth_url"`
		PaymentType       string    `json:"payment_type"`
		FraudStatus       string    `json:"fraud_status"`
		ChargeType        string    `json:"charge_type"`
		CreatedAt         time.Time `json:"created_at"`
		AccountID         int       `json:"account_id"`
		Customer          struct {
			ID          int         `json:"id"`
			PhoneNumber interface{} `json:"phone_number"`
			Name        string      `json:"name"`
			Email       string      `json:"email"`
			CreatedAt   time.Time   `json:"created_at"`
		} `json:"customer"`
		Card struct {
			First6Digits string `json:"first_6digits"`
			Last4Digits  string `json:"last_4digits"`
			Issuer       string `json:"issuer"`
			Country      string `json:"country"`
			Type         string `json:"type"`
			Expiry       string `json:"expiry"`
		} `json:"card"`
	} `json:"data"`
}

type RecurringChargePayload struct {
	Token     string `json:"token"`
	Currency  string `json:"currency"`
	Country   string `json:"country"`
	Amount    int    `json:"amount"`
	Email     string `json:"email"`
	TxRef     string `json:"tx_ref"`
	Narration string `json:"narration"`
}

type RecurringChargeResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		ID                int       `json:"id"`
		TxRef             string    `json:"tx_ref"`
		FlwRef            string    `json:"flw_ref"`
		RedirectURL       string    `json:"redirect_url"`
		DeviceFingerprint string    `json:"device_fingerprint"`
		Amount            int       `json:"amount"`
		ChargedAmount     int       `json:"charged_amount"`
		AppFee            int       `json:"app_fee"`
		MerchantFee       int       `json:"merchant_fee"`
		ProcessorResponse string    `json:"processor_response"`
		AuthModel         string    `json:"auth_model"`
		Currency          string    `json:"currency"`
		IP                string    `json:"ip"`
		Narration         string    `json:"narration"`
		Status            string    `json:"status"`
		PaymentType       string    `json:"payment_type"`
		CreatedAt         time.Time `json:"created_at"`
		AccountID         int       `json:"account_id"`
		Customer          struct {
			ID          int         `json:"id"`
			PhoneNumber interface{} `json:"phone_number"`
			Name        string      `json:"name"`
			Email       string      `json:"email"`
			CreatedAt   time.Time   `json:"created_at"`
		} `json:"customer"`
		Card struct {
			First6Digits string `json:"first_6digits"`
			Last4Digits  string `json:"last_4digits"`
			Issuer       string `json:"issuer"`
			Country      string `json:"country"`
			Type         string `json:"type"`
			Expiry       string `json:"expiry"`
		} `json:"card"`
		Meta struct {
			Authorization struct {
				Mode     string `json:"mode"`
				Redirect string `json:"redirect"`
			} `json:"authorization"`
		} `json:"meta"`
	} `json:"data"`
}

type RefundChargePayload struct {
	FlwRef string `json:"flw_ref"`
	Amount int    `json:"amount"`
}

type RefundChargeResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		ID                int       `json:"id"`
		TxRef             string    `json:"tx_ref"`
		FlwRef            string    `json:"flw_ref"`
		DeviceFingerprint string    `json:"device_fingerprint"`
		Amount            int       `json:"amount"`
		ChargedAmount     int       `json:"charged_amount"`
		AppFee            int       `json:"app_fee"`
		MerchantFee       int       `json:"merchant_fee"`
		ProcessorResponse string    `json:"processor_response"`
		AuthModel         string    `json:"auth_model"`
		Currency          string    `json:"currency"`
		IP                string    `json:"ip"`
		Narration         string    `json:"narration"`
		Status            string    `json:"status"`
		AuthURL           string    `json:"auth_url"`
		PaymentType       string    `json:"payment_type"`
		FraudStatus       string    `json:"fraud_status"`
		ChargeType        string    `json:"charge_type"`
		CreatedAt         time.Time `json:"created_at"`
		AccountID         int       `json:"account_id"`
		Customer          struct {
			ID            int         `json:"id"`
			Phone         string      `json:"phone"`
			FullName      string      `json:"fullName"`
			Customertoken interface{} `json:"customertoken"`
			Email         string      `json:"email"`
			CreatedAt     time.Time   `json:"createdAt"`
			UpdatedAt     time.Time   `json:"updatedAt"`
			DeletedAt     interface{} `json:"deletedAt"`
			AccountID     int         `json:"AccountId"`
		} `json:"customer"`
		Card struct {
			First6Digits string `json:"first_6digits"`
			Last4Digits  string `json:"last_4digits"`
			Issuer       string `json:"issuer"`
			Country      string `json:"country"`
			Type         string `json:"type"`
			Expiry       string `json:"expiry"`
		} `json:"card"`
	} `json:"data"`
}

type FetchTransactionPayload struct {
	ID int `json:"id"`
}

type FetchTransactionResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		ID                int       `json:"id"`
		TxRef             string    `json:"tx_ref"`
		FlwRef            string    `json:"flw_ref"`
		DeviceFingerprint string    `json:"device_fingerprint"`
		Amount            int       `json:"amount"`
		Currency          string    `json:"currency"`
		ChargedAmount     int       `json:"charged_amount"`
		AppFee            float64   `json:"app_fee"`
		MerchantFee       int       `json:"merchant_fee"`
		ProcessorResponse string    `json:"processor_response"`
		AuthModel         string    `json:"auth_model"`
		IP                string    `json:"ip"`
		Narration         string    `json:"narration"`
		Status            string    `json:"status"`
		PaymentType       string    `json:"payment_type"`
		CreatedAt         time.Time `json:"created_at"`
		AccountID         int       `json:"account_id"`
		Card              struct {
			First6Digits string `json:"first_6digits"`
			Last4Digits  string `json:"last_4digits"`
			Issuer       string `json:"issuer"`
			Country      string `json:"country"`
			Type         string `json:"type"`
			Token        string `json:"token"`
			Expiry       string `json:"expiry"`
		} `json:"card"`
		Meta          DirectChargePayloadMeta `json:"meta"`
		AmountSettled float64                 `json:"amount_settled"`
		Customer      struct {
			ID          int       `json:"id"`
			Name        string    `json:"name"`
			PhoneNumber string    `json:"phone_number"`
			Email       string    `json:"email"`
			CreatedAt   time.Time `json:"created_at"`
		} `json:"customer"`
	} `json:"data"`
}
