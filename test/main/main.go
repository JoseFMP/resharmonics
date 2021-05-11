package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/JoseFMP/resharmonics"
	"github.com/JoseFMP/resharmonics/financials"
)

// main Just do some basic tests without needing to import the library into another Go package
func main() {

	// Do not forget to create your credentials.json file
	creds := getCreds()
	log.Println("Got creds")

	//bookingsClient, errInitializing := bookings.Init(*creds, true)
	//if errInitializing != nil {
	//	log.Panicf("Err initializing bookings client\n%v", errInitializing)
	//}
	//smokeTestBookings(bookingsClient)

	finClient, errInitializingFinClient := financials.Init(*creds, true)
	if errInitializingFinClient != nil {
		log.Panicf("Err initializing financials client\n%v", errInitializingFinClient)
	}

	smokeTestInvoices(finClient)
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
