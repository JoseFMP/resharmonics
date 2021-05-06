package resharmonics

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

type getUsedTokenTestCase struct {
	header        http.Header
	expectedToken string
}

func TestCanScrapToken(t *testing.T) {

	tcs := []getUsedTokenTestCase{
		{
			map[string][]string{},
			"",
		},
		{
			map[string][]string{"Authentication": {"Bearer lalalala"}},
			"lalalala",
		},
	}

	for _, tc := range tcs {

		actual := getUsedToken(tc.header)

		require.Equal(t, tc.expectedToken, actual)
	}

}
