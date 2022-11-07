package go_reepay

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func UnmarshalChargeDTO(data []byte) (ChargeDTO, error) {
	var r ChargeDTO
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ChargeDTO) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ChargeDTO struct {
	Handle                 string      `json:"handle"`
	Key                    string      `json:"key"`
	Amount                 int64       `json:"amount"`
	Currency               string      `json:"currency"`
	Customer               Customer    `json:"customer"`
	Metadata               Metadata    `json:"metadata"`
	Source                 string      `json:"source"`
	Settle                 bool        `json:"settle"`
	Recurring              bool        `json:"recurring"`
	Parameters             Parameters  `json:"parameters"`
	Ordertext              string      `json:"ordertext"`
	OrderLines             []OrderLine `json:"order_lines"`
	CustomerHandle         string      `json:"customer_handle"`
	BillingAddress         IngAddress  `json:"billing_address"`
	ShippingAddress        IngAddress  `json:"shipping_address"`
	UsePmForSubscription   bool        `json:"use_pm_for_subscription"`
	TextOnStatement        string      `json:"text_on_statement"`
	PaymentMethodReference string      `json:"payment_method_reference"`
	Async                  string      `json:"async"`
	AcceptUrl              string      `json:"accept_url"`
	CancelUrl              string      `json:"cancel_url"`
}

type IngAddress struct {
	Company         string `json:"company"`
	Vat             string `json:"vat"`
	Attention       string `json:"attention"`
	Address         string `json:"address"`
	Address2        string `json:"address2"`
	City            string `json:"city"`
	Country         string `json:"country"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	PostalCode      string `json:"postal_code"`
	StateOrProvince string `json:"state_or_province"`
}

type Customer struct {
	Email          string   `json:"email"`
	Address        string   `json:"address"`
	Address2       string   `json:"address2"`
	City           string   `json:"city"`
	Country        string   `json:"country"`
	Phone          string   `json:"phone"`
	Company        string   `json:"company"`
	Vat            string   `json:"vat"`
	Handle         string   `json:"handle"`
	Test           bool     `json:"test"`
	Metadata       Metadata `json:"metadata"`
	FirstName      string   `json:"first_name"`
	LastName       string   `json:"last_name"`
	PostalCode     string   `json:"postal_code"`
	GenerateHandle bool     `json:"generate_handle"`
}

type Metadata struct {
}

type OrderLine struct {
	ID                   *string  `json:"id,omitempty"`
	Ordertext            *string  `json:"ordertext,omitempty"`
	Amount               *int64   `json:"amount,omitempty"`
	Vat                  *float64 `json:"vat,omitempty"`
	Quantity             *int64   `json:"quantity,omitempty"`
	Origin               *string  `json:"origin,omitempty"`
	Timestamp            *string  `json:"timestamp,omitempty"`
	DiscountedAmount     *int64   `json:"discounted_amount,omitempty"`
	AmountInclVat        *bool    `json:"amount_incl_vat,omitempty"`
	AmountVat            *int64   `json:"amount_vat,omitempty"`
	AmountExVat          *int64   `json:"amount_ex_vat,omitempty"`
	UnitAmount           *int64   `json:"unit_amount,omitempty"`
	UnitAmountVat        *int64   `json:"unit_amount_vat,omitempty"`
	UnitAmountExVat      *int64   `json:"unit_amount_ex_vat,omitempty"`
	AmountDefinedInclVat *bool    `json:"amount_defined_incl_vat,omitempty"`
	OriginHandle         *string  `json:"origin_handle,omitempty"`
	PeriodFrom           *string  `json:"period_from,omitempty"`
	PeriodTo             *string  `json:"period_to,omitempty"`
	DiscountPercentage   *int64   `json:"discount_percentage,omitempty"`
	DiscountedOrderLine  *string  `json:"discounted_order_line,omitempty"`
}

type Parameters struct {
	MpsTTL string `json:"mps_ttl"`
}

type ChargeResponse struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}

type ChargeError struct {
	Code             int64  `json:"code"`
	Error            string `json:"error"`
	Message          string `json:"message"`
	Path             string `json:"path"`
	Timestamp        string `json:"timestamp"`
	HTTPStatus       int64  `json:"http_status"`
	HTTPReason       string `json:"http_reason"`
	RequestID        string `json:"request_id"`
	TransactionError string `json:"transaction_error"`
}

func UnmarshalSettleResponse(data []byte) (SettleResponse, error) {
	var r SettleResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SettleResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SettleResponse struct {
	Handle                 string      `json:"handle"`
	State                  string      `json:"state"`
	Customer               string      `json:"customer"`
	Amount                 int64       `json:"amount"`
	Currency               string      `json:"currency"`
	Authorized             string      `json:"authorized"`
	Settled                string      `json:"settled"`
	Cancelled              string      `json:"cancelled"`
	Created                string      `json:"created"`
	Transaction            string      `json:"transaction"`
	Error                  string      `json:"error"`
	Processing             string      `json:"processing"`
	Source                 Source      `json:"source"`
	OrderLines             []OrderLine `json:"order_lines"`
	RefundedAmount         int64       `json:"refunded_amount"`
	AuthorizedAmount       int64       `json:"authorized_amount"`
	ErrorState             string      `json:"error_state"`
	RecurringPaymentMethod string      `json:"recurring_payment_method"`
	BillingAddress         IngAddress  `json:"billing_address"`
	ShippingAddress        IngAddress  `json:"shipping_address"`
	PaymentContext         string      `json:"payment_context"`
}

type Source struct {
	Type                       string `json:"type"`
	Card                       string `json:"card"`
	Mps                        string `json:"mps"`
	Fingerprint                string `json:"fingerprint"`
	Provider                   string `json:"provider"`
	VippsRecurring             string `json:"vipps_recurring"`
	AuthTransaction            string `json:"auth_transaction"`
	CardType                   string `json:"card_type"`
	TransactionCardType        string `json:"transaction_card_type"`
	ExpDate                    string `json:"exp_date"`
	MaskedCard                 string `json:"masked_card"`
	CardCountry                string `json:"card_country"`
	StrongAuthenticationStatus string `json:"strong_authentication_status"`
	ThreeDSecureStatus         string `json:"three_d_secure_status"`
	RiskRule                   string `json:"risk_rule"`
	AcquirerCode               string `json:"acquirer_code"`
	AcquirerMessage            string `json:"acquirer_message"`
	AcquirerReference          string `json:"acquirer_reference"`
	TextOnStatement            string `json:"text_on_statement"`
	SurchargeFee               int64  `json:"surcharge_fee"`
}

func (r *Reepay) CreateChargeSession(charge ChargeDTO) (chargeResponse *ChargeResponse, chargeError *ChargeError, err error) {
	charge.AcceptUrl = r.SuccessURL
	charge.CancelUrl = r.CancelURL

	j, err := charge.Marshal()
	if err != nil {
		return
	}

	url := "https://checkout-api.reepay.com/v1/session/charge/"

	body := strings.NewReader(string(j))

	req, _ := http.NewRequest("POST", url, body)

	keyEnc := base64.StdEncoding.EncodeToString([]byte(r.Key))

	req.Header.Add("Authorization", keyEnc)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	if res.StatusCode == 200 {
		err = json.Unmarshal(b, &chargeResponse)
	} else {
		err = json.Unmarshal(b, &chargeError)
	}
	return
}

func (r *Reepay) SettleChargeSession(charge ChargeDTO) (settleResponse SettleResponse, err error) {

	url := "https://checkout-api.reepay.com/v1/session/charge/" + charge.Handle + "/settle"

	j, err := charge.Marshal()
	if err != nil {
		return
	}

	body := strings.NewReader(string(j))

	req, _ := http.NewRequest("POST", url, body)

	keyEnc := base64.StdEncoding.EncodeToString([]byte(r.Key))

	req.Header.Add("Authorization", keyEnc)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &settleResponse)
	return

}
