package contact

import "github.com/JoseFMP/resharmonics/property"

type Details struct {
	Type                    ContactType       `json:"type"`
	CompanyName             string            `json:"companyName"`
	FirstName               string            `json:"contactFirstName"`
	LastName                string            `json:"contactLastName"`
	Email                   string            `json:"email"`
	Telephone               string            `json:"telephone,omitempty"`
	Address                 *property.Address `json:"address,omitempty"`
	PassportNumber          string            `json:"passportNumber,omitempty"`
	Nationality             string            `json:"nationality,omitempty"`
	SocialSecurity          string            `json:"socialSecurity,omitempty"`
	Referrer                string            `json:"referrer,omitempty"`
	CorporateCode           string            `json:"corporateCode,omitempty"`
	CompanyNumber           string            `json:"companyNumber,omitempty"`
	CompanyEmail            string            `json:"companyEmail,omitempty"`
	CompanyAddress          string            `json:"companyAddress,omitempty"`
	BookingContactFirstName string            `json:"bookingContactFirstName,omitempty"`
	BookingContactLastName  string            `json:"bookingContactLastName,omitempty"`
	BookingContactEmail     string            `json:"bookingContactEmail,omitempty"`
}

type ContactType string
