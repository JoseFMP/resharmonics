package resharmonics

import (
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
)

type client struct {
	credentials     Credentials
	tokenSemaphoere *semaphore.Weighted
	token           *string
	tokenFetchedOn  *time.Time
}

// Client client
type Client interface {
	//Auth() error
	DoPost(subPath string, params map[string]string) ([]byte, error)
	DoGet(subPath string, params map[string]string) ([]byte, error)
}

// Init gives you a Resharmonics client with functionality to do HTTP requests and authenticate
func Init(cred Credentials) (Client, error) {

	errValidating := validate(cred.Username, cred.Password)
	if errValidating != nil {
		return nil, errValidating
	}
	tokenSemaphore := semaphore.NewWeighted(1)
	return &client{
		credentials:     cred,
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

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
