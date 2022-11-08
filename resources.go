package go_reepay

import (
	"encoding/json"
	"time"
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
	Configuration            *string      `json:"configuration,omitempty"`
	Locale                   *string      `json:"locale,omitempty"`
	TTL                      *string      `json:"ttl,omitempty"`
	Phone                    *string      `json:"phone,omitempty"`
	Invoice                  *string      `json:"invoice,omitempty"`
	Settle                   *bool        `json:"settle,omitempty"`
	Order                    *Order       `json:"order,omitempty"`
	Recurring                *bool        `json:"recurring,omitempty"`
	AcceptURL                *string      `json:"accept_url,omitempty"`
	CancelURL                *string      `json:"cancel_url,omitempty"`
	PaymentMethods           []string     `json:"payment_methods,omitempty"`
	CardOnFile               *string      `json:"card_on_file,omitempty"`
	CardOnFileRequireCvv     *bool        `json:"card_on_file_require_cvv,omitempty"`
	CardOnFileRequireExpDate *bool        `json:"card_on_file_require_exp_date,omitempty"`
	ButtonText               *string      `json:"button_text,omitempty"`
	RecurringAverageAmount   *int64       `json:"recurring_average_amount,omitempty"`
	ScaData                  *ScaData     `json:"sca_data,omitempty"`
	SessionData              *SessionData `json:"session_data,omitempty"`
	PaymentMethodReference   *string      `json:"payment_method_reference,omitempty"`
	TextOnStatement          *string      `json:"text_on_statement,omitempty"`
	RecurringOptional        *bool        `json:"recurring_optional,omitempty"`
	RecurringOptionalText    *string      `json:"recurring_optional_text,omitempty"`
}

func (r *SettleDTO) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SettleDTO struct {
	Key        *string     `json:"key,omitempty"`
	Amount     *int        `json:"amount,omitempty"`
	Ordertext  *string     `json:"ordertext,omitempty"`
	OrderLines []OrderLine `json:"order_lines,omitempty"`
}

type Order struct {
	Handle          *string     `json:"handle,omitempty"`
	Key             *string     `json:"key,omitempty"`
	Amount          *int64      `json:"amount,omitempty"`
	Currency        *string     `json:"currency,omitempty"`
	Customer        *Customer   `json:"customer,omitempty"`
	Metadata        *string     `json:"metadata,omitempty"`
	Ordertext       *string     `json:"ordertext,omitempty"`
	OrderLines      []OrderLine `json:"order_lines,omitempty"`
	CustomerHandle  *string     `json:"customer_handle,omitempty"`
	BillingAddress  *IngAddress `json:"billing_address,omitempty"`
	ShippingAddress *IngAddress `json:"shipping_address,omitempty"`
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
	Email    string `json:"email"`
	Address  string `json:"address"`
	Address2 string `json:"address2"`
	City     string `json:"city"`
	Country  string `json:"country"`
	Phone    string `json:"phone"`
	Company  string `json:"company"`
	Vat      string `json:"vat"`
	Handle   string `json:"handle"`
	Test     bool   `json:"test"`
	//Metadata       interface{} `json:"metadata"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	PostalCode     string `json:"postal_code"`
	GenerateHandle bool   `json:"generate_handle"`
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

type ScaData struct {
	Name               *string                `json:"name,omitempty"`
	Email              *string                `json:"email,omitempty"`
	HomePhone          *Phone                 `json:"home_phone,omitempty"`
	MobilePhone        *Phone                 `json:"mobile_phone,omitempty"`
	WorkPhone          *Phone                 `json:"work_phone,omitempty"`
	BillingAddress     *ScaDataBillingAddress `json:"billing_address,omitempty"`
	ShippingAddress    *ScaDataBillingAddress `json:"shipping_address,omitempty"`
	AddressMatch       *bool                  `json:"address_match,omitempty"`
	AccountID          *string                `json:"account_id,omitempty"`
	ChallengeIndicator *string                `json:"challenge_indicator,omitempty"`
	RiskIndicator      *RiskIndicator         `json:"risk_indicator,omitempty"`
	AccountInfo        *AccountInfo           `json:"account_info,omitempty"`
}
type Phone struct {
	Cc         *string `json:"cc,omitempty"`
	Subscriber *string `json:"subscriber,omitempty"`
}

type ScaDataBillingAddress struct {
	Address1        *string `json:"address1,omitempty"`
	Address2        *string `json:"address2,omitempty"`
	Address3        *string `json:"address3,omitempty"`
	City            *string `json:"city,omitempty"`
	Country         *string `json:"country,omitempty"`
	PostalCode      *string `json:"postal_code,omitempty"`
	StateOrProvince *string `json:"state_or_province,omitempty"`
}

type AccountInfo struct {
	Created                       *string `json:"created,omitempty"`
	Changed                       *string `json:"changed,omitempty"`
	AgeIndicator                  *string `json:"age_indicator,omitempty"`
	ChangeIndicator               *string `json:"change_indicator,omitempty"`
	PasswordChange                *string `json:"password_change,omitempty"`
	PasswordChangeIndicator       *string `json:"password_change_indicator,omitempty"`
	PurchaseCount                 *int64  `json:"purchase_count,omitempty"`
	AddCardAttempts               *int64  `json:"add_card_attempts,omitempty"`
	TransactionsDay               *int64  `json:"transactions_day,omitempty"`
	TransactionsYear              *int64  `json:"transactions_year,omitempty"`
	CardAge                       *string `json:"card_age,omitempty"`
	CardAgeIndicator              *string `json:"card_age_indicator,omitempty"`
	ShippingAddressUsage          *string `json:"shipping_address_usage,omitempty"`
	ShippingAddressUsageIndicator *string `json:"shipping_address_usage_indicator,omitempty"`
	ShippingNameIndicator         *bool   `json:"shipping_name_indicator,omitempty"`
	SuspiciousActivity            *bool   `json:"suspicious_activity,omitempty"`
}

type RiskIndicator struct {
	DeliveryEmail             *string `json:"delivery_email,omitempty"`
	DeliveryTimeframe         *string `json:"delivery_timeframe,omitempty"`
	GiftCardAmount            *int64  `json:"gift_card_amount,omitempty"`
	GiftCardCount             *int64  `json:"gift_card_count,omitempty"`
	GiftCardCurrency          *string `json:"gift_card_currency,omitempty"`
	PreOrderDate              *string `json:"pre_order_date,omitempty"`
	PreOrderPurchaseIndicator *string `json:"pre_order_purchase_indicator,omitempty"`
	ReorderItemsIndicator     *string `json:"reorder_items_indicator,omitempty"`
	ShippingIndicator         *string `json:"shipping_indicator,omitempty"`
}

type SessionData struct {
	Ssn                   *string `json:"ssn,omitempty"`
	AccountHolderName     *string `json:"account_holder_name,omitempty"`
	MpsAmount             *int64  `json:"mps_amount,omitempty"`
	MpsPlan               *string `json:"mps_plan,omitempty"`
	MpsDescription        *string `json:"mps_description,omitempty"`
	MpsNextPaymentDate    *string `json:"mps_next_payment_date,omitempty"`
	MpsFrequency          *string `json:"mps_frequency,omitempty"`
	MpsExternalID         *string `json:"mps_external_id,omitempty"`
	MpsPaymentDescription *string `json:"mps_payment_description,omitempty"`
	MpsCancelRedirectURL  *string `json:"mps_cancel_redirect_url,omitempty"`
	AlternativeReturnURL  *string `json:"alternative_return_url,omitempty"`
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

func (r *RefundDTO) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type RefundDTO struct {
	Invoice        *string         `json:"invoice,omitempty"`
	Key            *string         `json:"key,omitempty"`
	Amount         *int            `json:"amount,omitempty"`
	Text           *string         `json:"text,omitempty"`
	NoteLines      []NoteLine      `json:"note_lines,omitempty"`
	ManualTransfer *ManualTransfer `json:"manual_transfer,omitempty"`
}

type NoteLine struct {
	Amount      int    `json:"amount"`
	Text        string `json:"text"`
	Quantity    int    `json:"quantity"`
	OrderLineId string `json:"order_line_id"`
}

type ManualTransfer struct {
	Comment     string `json:"comment"`
	Reference   string `json:"reference"`
	Method      string `json:"method"`
	PaymentDate string `json:"payment_date"`
}

type RefundResponse struct {
	Id              string    `json:"id"`
	State           string    `json:"state"`
	Invoice         string    `json:"invoice"`
	Amount          int       `json:"amount"`
	Currency        string    `json:"currency"`
	Transaction     string    `json:"transaction"`
	Error           string    `json:"error"`
	Type            string    `json:"type"`
	Created         time.Time `json:"created"`
	CreditNoteId    string    `json:"credit_note_id"`
	RefTransaction  string    `json:"ref_transaction"`
	ErrorState      string    `json:"error_state"`
	AcquirerMessage string    `json:"acquirer_message"`
}
