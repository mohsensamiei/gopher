package telegram

type OrderInfo struct {
	Name            string           `json:"name"`             // Optional. User name
	PhoneNumber     string           `json:"phone_number"`     // Optional. User's phone number
	Email           string           `json:"email"`            // Optional. User email
	ShippingAddress *ShippingAddress `json:"shipping_address"` // Optional. User shipping address
}

type ShippingAddress struct {
	CountryCode string `json:"country_code"` // Two-letter ISO 3166-1 alpha-2 country code
	State       string `json:"state"`        // State, if applicable
	City        string `json:"city"`         // City
	StreetLine1 string `json:"street_line1"` // First line for the address
	StreetLine2 string `json:"street_line2"` // Second line for the address
	PostCode    string `json:"post_code"`    // Address post code
}

type ShippingQuery struct {
	ID              string           `json:"id"`               // Unique query identifier
	From            *User            `json:"from"`             // User who sent the query
	InvoicePayload  string           `json:"invoice_payload"`  // Bot specified invoice payload
	ShippingAddress *ShippingAddress `json:"shipping_address"` // User specified shipping address
}
