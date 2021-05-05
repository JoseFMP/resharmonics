package resharmonics

import (
	"context"
	"time"

	"github.com/JoseFMP/resharmonics/auth"
)

var authContext = context.TODO()

func (clt *client) auth(tokenBeforeReq *string) error {

	clt.tokenLock.Lock()
	defer clt.tokenLock.Unlock()
	transitionLooksValid := tokenTransitionLooksValid(tokenBeforeReq, clt.token)
	if transitionLooksValid {
		return nil
	}

	token, errFetchingToken := auth.FetchToken(clt.credentials.Username, clt.credentials.Password)
	now := time.Now()

	if errFetchingToken != nil {
		return errFetchingToken
	}

	clt.token = &token
	clt.tokenFetchedOn = &now
	return nil
}

func tokenTransitionLooksValid(tokenBeforeReq *string, tokenAfterReq *string) bool {

	if tokenBeforeReq == nil && tokenAfterReq == nil {
		return false
	}

	if tokenBeforeReq == nil && tokenAfterReq != nil {
		return true
	}

	if tokenBeforeReq != nil && tokenAfterReq == nil {
		return false
	}

	return *tokenBeforeReq != *tokenAfterReq
}
