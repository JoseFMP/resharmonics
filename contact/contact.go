package contact

type Details struct {
	Type                    ContactType `json:"type"`
	CompanyName             string      `json:"companyName"`
	FirstName               string      `json:"contactFirstName"`
	LastName                string      `json:"contactLastName"`
	Email                   string      `json:"email"`
	Telephone               string      `json:"telephone"`
	Address                 Address     `json:"address"`
	PassportNumber          string      `json:"passportNumber"`
	Nationality             string      `json:"nationality"`
	SocialSecurity          string      `json:"socialSecurity"`
	Referrer                string      `json:"referrer"`
	CorporateCode           string      `json:"corporateCode"`
	CompanyNumber           string      `json:"companyNumber"`
	CompanyEmail            string      `json:"companyEmail"`
	CompanyAddress          string      `json:"companyAddress"`
	BookingContactFirstName string      `json:"bookingContactFirstName"`
	BookingContactLastName  string      `json:"bookingContactLastName"`
	BookingContactEmail     string      `json:"bookingContactEmail"`
}

type ContactType string
