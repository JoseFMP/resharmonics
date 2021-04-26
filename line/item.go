package line

import (
	"github.com/JoseFMP/resharmonics/bookings"
	"github.com/JoseFMP/resharmonics/product"
)

type LineItem struct {
	ID          ItemID `json:"itemId"`
	Description string `json:"description"`

	// Code a code that references the product that this line item corresponds to
	Code ItemCode `json:"itemCode"`

	// TaxType a mapping element used to specify the tax code used for this line item
	TaxType     string       `json:"taxType"`
	ProductID   product.ID   `json:"productId"`
	ProductCode product.Code `json:"productCode"`

	Quantity     float64 `json:"quantity"`
	UnitPriceNet float64 `json:"unitPriceNet"`
	TotalNet     float64 `json:"totalNet"`
	TotalTax     float64 `json:"totalTax"`

	CostCenter string `json:"costCenter"`

	UnitReference string `json:"unitReference"`

	BookingReference bookings.BookingReference `json:"bookingReference"`

	// ChargeFrom $date
	ChargeFrom string `json:"chargeFrom"`

	// ChargeTo $date
	ChargeTo string `json:"chargeTo"`
}

type ItemID int
type ItemCode string
