package resharmonics

import (
	"context"
	"log"
	"time"

	"github.com/JoseFMP/resharmonics/auth"
)

var authContext = context.TODO()

func (clt *client) auth(tokenBeforeReq string) error {

	log.Println("Waiting for lock")
	clt.tokenLock.Lock()
	log.Println("Lock acquired")

	defer clt.tokenLock.Unlock()
	transitionLooksValid := auth.IsTokenUpdatedToValid(tokenBeforeReq, clt.token)
	if transitionLooksValid {
		log.Println("Token was updated to a valid token in the mean time.")
		return nil
	}
	log.Println("No token update, proceeding to fetch now.")

	token, errFetchingToken := auth.FetchToken(clt.credentials.Username, clt.credentials.Password)
	log.Println("Fetched new token.")
	now := time.Now()

	if errFetchingToken != nil {
		return errFetchingToken
	}

	clt.token = token
	clt.tokenFetchedOn = &now
	return nil
}
