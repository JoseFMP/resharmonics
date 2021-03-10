package resharmonics

import (
	"fmt"

	"golang.org/x/sync/semaphore"
)

type client struct {
	username        string
	password        string
	tokenSemaphoere *semaphore.Weighted
	token           string
}

// Resharmonics client
type Resharmonics interface {
	//Auth() error
	DoPost(subPath string, params map[string]string) ([]byte, error)
	DoGet(subPath string, params map[string]string) ([]byte, error)
}

// Init gives you a Resharmonics client with functionality to do HTTP requests and authenticate
func Init(username string, password string) (Resharmonics, error) {

	errValidating := validate(username, password)
	if errValidating != nil {
		return nil, errValidating
	}
	tokenSemaphore := semaphore.NewWeighted(1)
	return &client{
		username:        username,
		password:        password,
		tokenSemaphoere: tokenSemaphore,
	}, nil
}

func validate(username string, password string) error {
	if username == "" {
		return fmt.Errorf("Username is empty")
	}

	if password == "" {
		return fmt.Errorf("Password is empty")
	}

	return nil
}
