package auth

import "log"

func IsTokenUpdatedToValid(tokenBeforeReq string, tokenAfterReq string) bool {

	if tokenBeforeReq == tokenAfterReq {
		log.Println("Tokens did not change")
		return false
	}
	if tokenAfterReq == "" {
		log.Println("New token is empty")
		return false
	}

	return true
}
