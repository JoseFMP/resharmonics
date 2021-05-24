package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/JoseFMP/resharmonics/urls"
	"github.com/JoseFMP/resharmonics/utils"
)

const authEndpointSubpat = `auth`

func fetchToken(userName string, password string) (string, error) {

	params := map[string]string{
		"username": userName,
		"password": password,
	}

	endPoint := fmt.Sprintf("%s/%s", urls.GetBaseUrl(), authEndpointSubpat)
	req, errCreatingReq := utils.CreatePostReq(endPoint, params, "")

	if errCreatingReq != nil {
		return "", errCreatingReq
	}

	httpClient := http.Client{}
	res, errDoingReq := httpClient.Do(req)
	if errDoingReq != nil {
		return "", errDoingReq
	}

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Executing auth req returned not OK HTTP code: %s", res.Status)
	}

	payload, errReadingBody := ioutil.ReadAll(res.Body)
	if errReadingBody != nil {
		return "", errReadingBody
	}
	var responseParsed APITokenResponse
	errorUnmarshalling := json.Unmarshal(payload, &responseParsed)
	if errorUnmarshalling != nil {
		return "", errorUnmarshalling
	}

	if responseParsed.APIToken == "" {
		return "", fmt.Errorf("Token empty received")
	}

	return responseParsed.APIToken, nil
}

type APITokenResponse struct {
	APIToken string `json:"apiToken"`
}
