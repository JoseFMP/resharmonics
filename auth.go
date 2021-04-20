package resharmonics

import (
	"context"
	"log"
	"time"

	"github.com/JoseFMP/resharmonics/auth"
	"golang.org/x/sync/semaphore"
)

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

	token, errFetchingToken := auth.FetchToken(clt.credentials.Username, clt.credentials.Password)
	now := time.Now()

	if errFetchingToken != nil {
		return errFetchingToken
	}

	clt.token = &token
	clt.tokenFetchedOn = &now
	return nil
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
