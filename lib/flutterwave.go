package lib

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type Payment struct {
	client *resty.Client

	encryptionKey string
	encrypter     ThressDSEncrypter
}

func (p Payment) DirectCharge(args DirectChargePayload) (response *DirectChargeResponse, err error) {
	stringPayload, err := json.Marshal(&args)
	if err != nil {
		return nil, err
	}

	encyptedPayload, err := p.encrypter.EncryptData(p.encryptionKey, string(stringPayload))
	if err != nil {
		return nil, err
	}

	request, err := p.client.R().
		SetBody(map[string]string{
			"client": encyptedPayload,
		}).
		Post("/v3/charges?type=card")

	json.Unmarshal(request.Body(), &response)

	// BUG: fix error handling
	if request.IsError() {
		return nil, HttpError{
			Message:    response.Message,
			StatusCode: request.StatusCode(),
		}
	}

	return response, err
}

func (p Payment) ValidateCharge(args ValidateChargePayload) (response *ValidateChargeResponse, err error) {
	request, err := p.client.R().
		SetBody(args).
		Post("/v3/validate-charge")

	json.Unmarshal(request.Body(), &response)

	// BUG: fix error handling
	if request.IsError() {
		return nil, HttpError{
			Message:    response.Message,
			StatusCode: request.StatusCode(),
		}
	}

	return response, err
}

func (p Payment) RecurringCharge(args RecurringChargePayload) (response *RecurringChargeResponse, err error) {
	request, err := p.client.R().
		SetBody(args).
		Post("/v3/tokenized-charges")

	json.Unmarshal(request.Body(), &response)

	// BUG: fix error handling
	if request.IsError() {
		return nil, HttpError{
			Message:    response.Message,
			StatusCode: request.StatusCode(),
		}
	}

	return response, err
}

func (p Payment) RefundCharge(args RefundChargePayload) (response *RefundChargePayload, err error) {
	request, err := p.client.R().
		SetBody(map[string]interface{}{
			"amount": args.Amount,
		}).
		Post(fmt.Sprintf("/v3/charges/%s/refund", args.FlwRef))

	json.Unmarshal(request.Body(), &response)

	return response, err
}

func (p Payment) FetchTransaction(args FetchTransactionPayload) (response *FetchTransactionResponse, err error) {
	request, err := p.client.R().
		Get(fmt.Sprintf("/v3/transactions/%d/verify", args.ID))

	json.Unmarshal(request.Body(), &response)

	// BUG: fix error handling
	if request.IsError() {
		return nil, HttpError{
			Message:    response.Message,
			StatusCode: request.StatusCode(),
		}
	}

	return response, err
}

func NewPayment(
	client *resty.Client,
	env Env,
	encrypter ThressDSEncrypter,
) Payment {
	return Payment{
		client:        client,
		encryptionKey: env.FLUTTERWAVE_ENCRYTION_KEY,
		encrypter:     encrypter,
	}
}

func NewFlutterwaveClient(env Env) *resty.Client {
	client := resty.New()
	client.SetBaseURL(env.FLUTTERWAVE_API_URL)
	client.SetHeader("Authorization", fmt.Sprintf("Bearer %s", env.FLUTTERWAVE_SECRET_KEY))
	client.SetHeader("Content-Type", "application/json")

	return client
}
