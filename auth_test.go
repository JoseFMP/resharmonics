package resharmonics

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type tokenTransitionTestCase struct {
	tokenBefore *string
	tokenAfter  *string
	expected    bool
}

func TestTokenTransition(t *testing.T) {

	// arrange

	tokenMock := "aaaa"
	tokenMock2 := "sdfsdfgdfg"

	testCases := []tokenTransitionTestCase{
		{
			nil,
			nil,
			false,
		},
		{
			nil,
			&tokenMock,
			true,
		},
		{

			&tokenMock,
			nil,
			false,
		},
		{
			&tokenMock,
			&tokenMock,
			false,
		},
		{
			&tokenMock2,
			&tokenMock,
			true,
		},
	}

	// act

	for _, tc := range testCases {

		actual := tokenTransitionLooksValid(tc.tokenBefore, tc.tokenAfter)

		// verify
		require.Equal(t, tc.expected, actual)
	}

}
