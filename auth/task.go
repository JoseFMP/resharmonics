package auth

import (
	"log"
)

type AuthTask struct {
	WrongToken string
}

func Auth(authTasks <-chan *AuthTask, username string, password string, tokens chan<- string) {

	lastToken := ""
	for {

		authTask, channOpen := <-authTasks
		if !channOpen {
			return
		}
		if authTask == nil {
			continue
		}
		if authTask.WrongToken != lastToken {
			tokens <- lastToken
		}

		log.Printf("Received auth task. Wrong token: %s", authTask.WrongToken)
		for {
			newToken, errFetching := fetchToken(username, password)
			if errFetching != nil {
				log.Printf("Error fetching token. Retrying: %v", errFetching)
				continue
			}
			lastToken = newToken
			tokens <- newToken
			break
		}
	}

}
