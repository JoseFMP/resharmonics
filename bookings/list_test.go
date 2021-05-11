package bookings

import (
	"testing"

	"github.com/JoseFMP/resharmonics"
	"github.com/stretchr/testify/assert"
)

func TestCanParse(t *testing.T) {

	bookings, errParsing := parseListResponse([]byte(mockPayload1))

	assert.Nil(t, errParsing)
	assert.NotNil(t, bookings)
	assert.Len(t, bookings, 1)
	assert.Equal(t, bookings[0].Reference, resharmonics.BookingReference("20191129-00003/805"))
	assert.Equal(t, bookings[0].Status, GetAllBookingStatuses().CheckedOut)
	assert.Equal(t, bookings[0].StartDate, "2019-11-29")
	assert.Equal(t, bookings[0].EndDate, "2020-01-01")
}

func TestCanConvert(t *testing.T) {

	br := BookingData{
		Id: Identifier("ABC"),

		Status:    BookingStatus(GetAllBookingStatuses().CheckedIn),
		StartDate: "2019-01-01",
		EndDate:   "2019-01-03",
	}

	asBooking, errConverting := br.toBooking()

	assert.Nil(t, errConverting)
	assert.NotNil(t, asBooking)

	assert.Equal(t, Identifier("ABC"), asBooking.Id)
	assert.NotNil(t, asBooking.Period)
	assert.Equal(t, 2019, asBooking.Period.From.Year)
	assert.Equal(t, 1, asBooking.Period.From.Day)

	assert.Equal(t, 2019, asBooking.Period.To.Year)
	assert.Equal(t, 3, asBooking.Period.To.Day)

}

const mockPayload1 = `
{
	"results": [
	  {
		"bookingReference": "20191129-00003/805",
		"status": "CHECKED_OUT",
		"startDate": "2019-11-29",
		"endDate": "2020-01-01",
		"nightlyAverageRate": 99.26,
		"property": {
		  "id": "dc35a9ad-008b-4d41-b7b4-9e0a4ed64375",
		  "name": "Apt 1.2A",
		  "unavailable": false,
		  "propertyType": "STUDIO",
		  "propertyTypeName": "Studio Apartment",
		  "propertyTypeDescription": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse eget quam tortor. Donec tincidunt augue tincidunt ligula tincidunt porta",
		  "propertyTypeLongDescription": null,
		  "propertyTypeDescriptionTranslations": [
			{
			  "En": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse eget quam tortor. Donec tincidunt augue tincidunt ligula tincidunt porta"
			},
			{
			  "Fr": ""
			},
			{
			  "De": ""
			}
		  ],
		  "shortDescription": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse eget quam tortor. Donec tincidunt augue tincidunt ligula tincidunt porta",
		  "shortDescriptionTranslations": [
			{
			  "En": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse eget quam tortor. Donec tincidunt augue tincidunt ligula tincidunt porta"
			},
			{
			  "Fr": ""
			},
			{
			  "De": ""
			}
		  ],
		  "maxOccupancy": 2,
		  "floorSpace": "0m2",
		  "mainImage": {
			"type": "MAIN",
			"url": "//app.rerumapp.uk/homa/images/16bc455c07b079aad1710c4b9457978a_1562756807",
			"name": "deborah-cortelazzi-gREquCUXQLI-unsplash.jpg",
			"thirdPartyUrl": null
		  },
		  "location": "Lorem ipsum dolor, Lorem ipsum dolor, Lorem ipsum dolor, Lorem ipsum dolor, Lorem ipsum dolor",
		  "latitude": 52.484268,
		  "longitude": -1.910381,
		  "buildingName": "Da Vinci Block A",
		  "buildingDescription": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse eget quam tortor. Donec tincidunt augue tincidunt ligula tincidunt porta",
		  "buildingDescriptionTranslations": [
			{
			  "En": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse eget quam tortor. Donec tincidunt augue tincidunt ligula tincidunt porta"
			},
			{
			  "Fr": ""
			},
			{
			  "De": ""
			}
		  ],
		  "buildingClassification": "MANAGED",
		  "cityId": 1,
		  "areaId": 1,
		  "areaName": "West End",
		  "propertyAddress": {
			"addressLine1": "Lorem ipsum dolor",
			"addressLine2": "Lorem ipsum dolor",
			"addressLine3": "Lorem ipsum dolor",
			"city": "Lorem ipsum dolor",
			"postCode": "Lorem ipsum dolor",
			"country": "United Kingdom"
		  },
		  "unitTypeId": 13,
		  "buildingId": 10,
		  "currencySymbol": "£",
		  "unitCostCentre": "Angelico",
		  "ownerName": "Owner Bookings",
		  "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse eget quam tortor. Donec tincidunt augue tincidunt ligula tincidunt porta",
		  "features": [
			{
			  "name": "Self check-in",
			  "nameTranslations": null,
			  "iconClass": "soap-icon-doorman"
			},
			{
			  "name": "Laundry room",
			  "nameTranslations": null,
			  "iconClass": "soap-icon-check"
			},
			{
			  "name": "On site parking",
			  "nameTranslations": null,
			  "iconClass": "soap-icon-parking"
			},
			{
			  "name": "Elevator / Lift",
			  "nameTranslations": null,
			  "iconClass": "soap-icon-elevator"
			}
		  ],
		  "images": [
			{
			  "type": "MAIN",
			  "url": "//app.rerumapp.uk/homa/images/16bc455c07b079aad1710c4b9457978a_1562756807",
			  "name": "deborah-cortelazzi-gREquCUXQLI-unsplash.jpg",
			  "thirdPartyUrl": null
			}
		  ],
		  "unitFeatures": [],
		  "customFields": [],
		  "unitContracts": null,
		  "thirdPartyUrl": null
		},
		"guests": [],
		"numberOfGuests": null,
		"bookingIdentifier": "53874f2f-5fec-4f79-9c5a-8a80f78dd061",
		"billingAccountName": null,
		"bookingAccountName": null
	  }
	]
  }
`
