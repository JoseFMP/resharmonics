package resharmonics

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"

	"github.com/JoseFMP/resharmonics/urls"
	"github.com/JoseFMP/resharmonics/utils"
)

var baseUrl = urls.GetBaseUrl()

func (clt *client) DoGet(subPath string, params map[string]interface{}) ([]byte, error) {

	if clt.token == "" {
		errAuthing := clt.auth("")
		if errAuthing != nil {
			return nil, errAuthing
		}
	}
	endpoint := fmt.Sprintf("%s/%s", baseUrl, subPath)

	req, errCreatingReq := utils.CreateGetReq(endpoint, params, clt.token)
	if errCreatingReq != nil {
		return nil, errCreatingReq
	}
	resPayload, errReadingBody := clt.doReq(req)

	if errReadingBody != nil {
		return nil, errReadingBody
	}
	return resPayload, nil
}

func (clt *client) doReq(req *http.Request) ([]byte, error) {

	var res *http.Response
	var errDoingReq error
	for {
		httpClient := http.Client{}
		log.Printf("[doReq] Doing req: %s %s %s", req.Method, req.URL, req.URL.Query().Encode())

		res, errDoingReq = httpClient.Do(req)
		if errDoingReq != nil {
			return nil, errDoingReq
		}

		hasToRedoAuth := rand.Float64() > 0.5

		if !hasToRedoAuth && res.StatusCode == http.StatusOK {
			break
		}

		if !hasToRedoAuth && res.StatusCode != http.StatusForbidden {
			return nil, fmt.Errorf("[doReq] Req result: %d -- %s %s", res.StatusCode, req.Method, req.URL.Path)
		}

		log.Println("[doReq] Needs to authenticate...")

		usedToken := getUsedToken(req.Header)
		errAuthenticating := clt.auth(usedToken)
		if errAuthenticating != nil {
			return nil, errAuthenticating
		} else {

			injectToken(clt.token, req.Header)
			log.Println("Authentication returned no error, should have new token")
		}
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d - %s", res.StatusCode, res.Status)
	}

	respPayload, errReadingBody := ioutil.ReadAll(res.Body)

	if errReadingBody != nil {
		return nil, errReadingBody
	}
	return respPayload, nil
}

func (clt *client) DoPost(subPath string, params map[string]string) ([]byte, error) {

	if clt.token == "" {
		errAuthing := clt.auth("")
		if errAuthing != nil {
			return nil, errAuthing
		}
	}
	endpoint := fmt.Sprintf("%s/%s", baseUrl, subPath)

	req, errCreatingReq := utils.CreatePostReq(endpoint, params, clt.token)
	if errCreatingReq != nil {
		return nil, errCreatingReq
	}
	resPayload, errReadingBody := clt.doReq(req)

	if errReadingBody != nil {
		return nil, errReadingBody
	}
	return resPayload, nil
}

func getUsedToken(header http.Header) string {

	if header == nil {
		return ""
	}

	authHeader := header.Get(authHeaderName)
	if authHeader == "" {
		return ""
	}

	//cleaned := strings.ToLower(authHeader)
	cleaned := strings.ReplaceAll(authHeader, "Bearer", "")
	cleaned = strings.ReplaceAll(cleaned, " ", "")

	return cleaned
}

func injectToken(token string, header http.Header) error {

	if header == nil {
		return fmt.Errorf("Header is nil")
	}

	header.Del(authHeaderName)
	header.Add(authHeaderName, token)

	return nil
}

const authHeaderName = `Authentication`
const bearerKeyword = `Bearer`
