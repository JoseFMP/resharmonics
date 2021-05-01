package bookings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanonicalReference(t *testing.T) {

	// arrange
	canonicalMock := "asdfasdfasdf"
	mock1 := BookingReference(fmt.Sprintf("%s/123", canonicalMock))
	mock2 := BookingReference(fmt.Sprintf("%s", canonicalMock))

	// act
	canonical1 := mock1.AsCanonical()
	canonical2 := mock2.AsCanonical()

	// verify
	require.NotNil(t, canonical1)
	require.Equal(t, canonicalMock, canonical1)

	require.NotNil(t, canonical2)
	require.Equal(t, canonicalMock, canonical2)
}
