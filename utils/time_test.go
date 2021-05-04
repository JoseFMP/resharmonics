package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type parseDateTestCase struct {
	date     string
	expected *BookingDate
}

func TestCanParseDates(t *testing.T) {

	// arrange

	testCases := []parseDateTestCase{
		{
			"2020-01-05",
			&BookingDate{
				Year: 2020,
				Day:  5,
			},
		},
		{
			"2020-05-05",
			&BookingDate{
				Year: 2020,
				Day:  126,
			},
		},
	}

	for _, tc := range testCases {

		actual, errParsing := FromDateString(tc.date)
		require.Nil(t, errParsing)

		require.NotNil(t, actual)
		require.Equal(t, tc.expected.Year, actual.Year)
		require.Equal(t, tc.expected.Day, actual.Day)

	}

}
