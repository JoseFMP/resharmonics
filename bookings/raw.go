package bookings

import (
	"github.com/JoseFMP/resharmonics/contact"
	"github.com/JoseFMP/resharmonics/invoices"
	"github.com/JoseFMP/resharmonics/property"
)

// BookingData is the raw payload format as returned by the Resharmonics API
type BookingData struct {
	Id                 Identifier            `json:"bookingIdentifier"`
	Reference          BookingReference      `json:"bookingReference"`
	Status             BookingStatus         `json:"status"`
	StartDate          string                `json:"startDate"` // just date as 2005-01-01
	EndDate            string                `json:"endDate"`   // just date as 2005-01-01
	Guests             []contact.Details     `json:"guests"`
	Property           property.PropertyData `json:"property"`
	NightlyAverageRate float64               `json:"nightlyAverageRate"`
	Invoices           []invoices.Invoice    `json:"invoices"`
	Extras             []Extra               `json:"extras"`
	BookingAccountName string                `json:"bookingAccountName"`
	BillingAccountName string                `json:"billingAccountName"`
}


