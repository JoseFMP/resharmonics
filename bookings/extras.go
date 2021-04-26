package bookings

type Extra struct {
	ID             ExtraID      `json:"id"`
	Description    string       `json:"description"`
	WebDescription string       `json:"webDescription"`
	Frequency      string       `json:"frequency"`
	WebCategory    string       `json:"webCategory"`
	Compulsory     bool         `json:"compulsory"`
	Pricing        ExtraPricing `json:"grossNetVat"`
}

type ExtraID int
type ExtraPricing struct {
	Gross float64 `json:"gross"`
	Net   float64 `json:"net"`
	Vat   float64 `json:"vat"`
}
