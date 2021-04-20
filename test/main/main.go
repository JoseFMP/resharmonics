package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/JoseFMP/resharmonics"
	"github.com/JoseFMP/resharmonics/bookings"
	"github.com/JoseFMP/resharmonics/utils"
)

// main Just do some basic tests without needing to import the library into another Go package
func main() {

	// Do not forget to create your credentials.json file
	creds := getCreds()
	log.Println("Got creds")

	bookingsClient, errInitializing := bookings.Init(*creds, true)
	if errInitializing != nil {
		log.Panicf("Err initializing bookings client\n%v", errInitializing)
	}

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
	bookingsList, errGettingList := bookingsClient.List(bookingPeriod, nil, nil)

	if errGettingList != nil {
		log.Panicf("Error getting the list... agg\n%v", errGettingList)
	}

	for _, b := range bookingsList {
		//log.Printf("Booking: %+v", b)
		singleBooking, errGettingSingle := bookingsClient.Get(b.Identifier)
		if errGettingSingle != nil {
			log.Panicf("Error getting single booking... agg\n%v", errGettingSingle)
		}

		marshalled, errMarhsalling := json.MarshalIndent(singleBooking, "", " ")
		if errMarhsalling != nil {
			panic("Err marshalling")
		}
		log.Printf("singleBooking: %s", string(marshalled))
	}
}

func getCreds() *resharmonics.Credentials {
	fileContent, errReading := ioutil.ReadFile("./test/credentials.json")
	if errReading != nil {
		log.Panicf("Error reading cred file: %v", errReading)
	}
	var creds resharmonics.Credentials
	errUnmarshalling := json.Unmarshal(fileContent, &creds)
	if errUnmarshalling != nil {
		log.Panicf("Error unmarshalling creds: %v", errUnmarshalling)
	}

	return &creds
}
