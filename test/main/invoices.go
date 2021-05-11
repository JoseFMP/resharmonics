package main

import (
	"log"
	"math/rand"

	"github.com/JoseFMP/resharmonics"
	"github.com/JoseFMP/resharmonics/financials"
	"github.com/JoseFMP/resharmonics/utils"
)

func smokeTestInvoices(finClient financials.FinancialsClient) {

	period := getPeriod()
	invs, errGettingInvs := finClient.ListInvoices(*period, resharmonics.OrganizationID(1), nil)

	if errGettingInvs != nil {
		panic(errGettingInvs)
	}

	log.Printf("Got %d invoices", len(invs))

	for _, inv := range invs {
		inv, errGettingInv := finClient.GetInvoice(inv.ID.ID)
		if errGettingInv != nil {
			panic(errGettingInv)
		}
		log.Printf("Got inv %d right!", int(inv.ID.ID))
	}

}

func getPeriod() *utils.BookingPeriod {

	from := utils.BookingDate{
		Year: 2020,
		Day:  40,
	}
	to := utils.BookingDate{
		Year: 2020,
		Day:  40 + int(rand.Float64()*300.0),
	}

	return &utils.BookingPeriod{
		From: &from,
		To:   &to,
	}
}
