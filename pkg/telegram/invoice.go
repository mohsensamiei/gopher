package telegram

// Invoice
// This object contains basic information about an invoice.
type Invoice struct {
	Title          string `json:"title"`           // Product name
	Description    string `json:"description"`     // Product description
	StartParameter string `json:"start_parameter"` // Unique bot deep-linking parameter that can be used to generate this invoice
	Currency       string `json:"currency"`        // Three-letter ISO 4217 currency code
	TotalAmount    int    `json:"total_amount"`    // Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
}

// SuccessfulPayment
// This object contains basic information about a successful payment.
type SuccessfulPayment struct {
	Currency                string     `json:"currency"`                   // Three-letter ISO 4217 currency code
	TotalAmount             int        `json:"total_amount"`               // Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	InvoicePayload          string     `json:"invoice_payload"`            // Bot specified invoice payload
	ShippingOptionID        string     `json:"shipping_option_id"`         // Optional. Identifier of the shipping option chosen by the user
	OrderInfo               *OrderInfo `json:"order_info"`                 // Optional. Order information provided by the user
	TelegramPaymentChargeID string     `json:"telegram_payment_charge_id"` // Telegram payment identifier
	ProviderPaymentChargeID string     `json:"provider_payment_charge_id"` // Provider payment identifier
}

// PreCheckoutQuery
// This object contains information about an incoming pre-checkout query.
type PreCheckoutQuery struct {
	ID               string     `json:"id"`                 // Unique query identifier
	From             *User      `json:"from"`               // User who sent the query
	Currency         string     `json:"currency"`           // Three-letter ISO 4217 currency code
	TotalAmount      int        `json:"total_amount"`       // Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	InvoicePayload   string     `json:"invoice_payload"`    // Bot specified invoice payload
	ShippingOptionID string     `json:"shipping_option_id"` // Optional. Identifier of the shipping option chosen by the user
	OrderInfo        *OrderInfo `json:"order_info"`         // Optional. Order information provided by the user
}
