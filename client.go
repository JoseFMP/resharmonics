package resharmonics

import (
	"fmt"
	"sync"
	"time"
)

type client struct {
	credentials    Credentials
	tokenLock      *sync.Mutex
	token          string
	tokenFetchedOn *time.Time
}

// Client client
type Client interface {
	//Auth() error
	DoPost(subPath string, params map[string]string) ([]byte, error)
	DoGet(subPath string, params map[string]interface{}) ([]byte, error)
}

// Init gives you a Resharmonics client with functionality to do HTTP requests and authenticate
func Init(cred Credentials, preAuthorize bool) (Client, error) {

	errValidating := validate(cred.Username, cred.Password)
	if errValidating != nil {
		return nil, errValidating
	}

	clientResult := &client{
		credentials: cred,
		tokenLock:   &sync.Mutex{},
	}
	if preAuthorize {
		authRes := clientResult.auth("")
		if authRes != nil {
			return nil, authRes
		}
	}
	return clientResult, nil
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

type OrganizationID int
