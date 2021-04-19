package resharmonics

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

const authEndpointSubpat = `auth`

var authContext = context.TODO()

func (clt *client) auth() error {

	log.Println("Waiting for semaphore not blocking")
	semaphoreAcquired := clt.tokenSemaphoere.TryAcquire(1)
	if !semaphoreAcquired {
		for {
			log.Println("Waiting for semaphore Blocking")
			errAcquiring := clt.tokenSemaphoere.Acquire(authContext, 1)
			if errAcquiring != nil {
				return errAcquiring
			}
			clt.tokenSemaphoere.Release(1)
			return nil
		}
	}
	defer clt.tokenSemaphoere.Release(1)

	params := map[string]string{
		"username": clt.credentials.Username,
		"password": clt.credentials.Password,
	}
	res, errDoingPost := clt.DoPost(authEndpointSubpat, params)
	now := time.Now()

	if errDoingPost != nil {
		return errDoingPost
	}

	var responseParsed APITokenResponse
	errorUnmarshalling := json.Unmarshal(res, &responseParsed)
	if errorUnmarshalling != nil {
		return errorUnmarshalling
	}

	if responseParsed.APIToken == "" {
		return fmt.Errorf("Token empty received")
	}

	clt.token = &responseParsed.APIToken
	clt.tokenFetchedOn = &now
	return nil
}

type APITokenResponse struct {
	APIToken string `json:"apiToken"`
}
