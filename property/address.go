package property

type Address struct {
	Line1       string  `json:"addressLine1"`
	Line2       string  `json:"addressLine2"`
	Line3       *string `json:"addressLine3,omitempty"`
	City        string  `json:"city"`
	ZipCode     string  `json:"postcode"`
	Country     string  `json:"country"`     // Country name
	IsoCountry2 *string `json:"isoCountry2"` //The ISO 3166-1 alpha-2 country code
}
