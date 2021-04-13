package resharmonics

import "fmt"

const hostName = `api.rerumapp.uk`
const apiV1Subpath = `api/v1`

func GetBaseUrl() string {
	return fmt.Sprintf("https://%s/%s", hostName, apiV1Subpath)
}
