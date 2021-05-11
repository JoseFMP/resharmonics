package main

import (
	"log"

	"github.com/JoseFMP/resharmonics/bookings"
	"github.com/JoseFMP/resharmonics/utils"
)

func smokeTestBookings(bookingsClient bookings.BookingsClient) {

	bookingPeriod := utils.BookingPeriod{
		From: &utils.BookingDate{
			Year: 2018,
			Day:  1,
		},
		To: &utils.BookingDate{
			Year: 2021,
			Day:  20,
		},
	}

	log.Println("Getting list...")
	bookingsList, errGettingList := bookingsClient.List(bookingPeriod, nil, nil, nil)

	if errGettingList != nil {
		log.Panicf("Error getting the list... agg\n%v", errGettingList)
	}

	for _, b := range bookingsList {
		sb, errGettingSingle := bookingsClient.Get(b.Id)
		if errGettingSingle != nil {
			log.Panicf("Error getting single booking... agg\n%v", errGettingSingle)
		}
		log.Printf("Did %s", string(sb.Identifier))

		//marshalled, errMarhsalling := json.MarshalIndent(singleBooking, "", " ")
		//if errMarhsalling != nil {
		//	panic("Err marshalling")
		//}
		//log.Printf("singleBooking: %s", string(marshalled))
	}
}
