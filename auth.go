package resharmonics

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/sync/semaphore"
)

const authEndpointSubpat = `auth`

var authContext = context.TODO()

func (clt *client) auth() error {

	for {
		hasToDo, errAcquiring := hasToDoAuth(clt.tokenSemaphoere)
		if errAcquiring != nil {
			continue
		}
		if !hasToDo {
			return nil
		} else {
			break
		}
	}
	defer clt.tokenSemaphoere.Release(1)

	token, errFetchingToken := fetchToken(clt.credentials)
	now := time.Now()

	if errFetchingToken != nil {
		return errFetchingToken
	}

	clt.token = &token
	clt.tokenFetchedOn = &now
	return nil
}

func fetchToken(creds Credentials) (string, error) {

	params := map[string]string{
		"username": creds.Username,
		"password": creds.Password,
	}

	req, errCreatingReq := createPostReq(authEndpointSubpat, params, nil)

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

func hasToDoAuth(semaphore *semaphore.Weighted) (bool, error) {
	log.Println("Waiting for semaphore not blocking")
	semaphoreAcquired := semaphore.TryAcquire(1)
	if semaphoreAcquired {
		//defer semaphore.Release(1)
		return true, nil
	} else {
		log.Println("Waiting for semaphore Blocking")
		errAcquiring := semaphore.Acquire(authContext, 1)
		if errAcquiring != nil {
			log.Printf("Err acquiring token: %v", errAcquiring)
			return false, errAcquiring
		}
		log.Println("Got token, means we're good to go", errAcquiring)
		defer semaphore.Release(1)
		return false, nil
	}
}
