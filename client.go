package resharmonics

import (
	"fmt"
	"log"
	"time"

	"github.com/JoseFMP/resharmonics/auth"
)

type client struct {
	credentials    Credentials
	reqTokenChan   chan *auth.AuthTask
	tokensChan     chan string
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
		credentials:  cred,
		reqTokenChan: make(chan *auth.AuthTask),
		tokensChan:   make(chan string),
	}
	go auth.Auth(clientResult.reqTokenChan, cred.Username, cred.Password, clientResult.tokensChan)

	if preAuthorize {
		clientResult.reqTokenChan <- &auth.AuthTask{}
		token := <-clientResult.tokensChan
		if token == "" {
			return nil, fmt.Errorf("Could not secure initial token")
		}
		clientResult.token = token
		log.Printf("Got initial token.")
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
