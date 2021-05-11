package resharmonics

import "strings"

type BookingReference string

func (original *BookingReference) AsCanonical() string {

	val := string(*original)

	splitted := strings.Split(val, "/")

	return splitted[0]
}
