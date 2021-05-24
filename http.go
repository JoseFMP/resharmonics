package resharmonics

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/JoseFMP/resharmonics/auth"
	"github.com/JoseFMP/resharmonics/urls"
	"github.com/JoseFMP/resharmonics/utils"
)

var baseUrl = urls.GetBaseUrl()

func (clt *client) ensureTokenNotEmpty() {
	if clt.token == "" {
		clt.reqTokenChan <- &auth.AuthTask{}
		newToken := <-clt.tokensChan
		clt.token = newToken
	}
}

func (clt *client) DoGet(subPath string, params map[string]interface{}) ([]byte, error) {

	clt.ensureTokenNotEmpty()
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

		if res.StatusCode == http.StatusOK {
			break
		}

		if res.StatusCode != http.StatusForbidden {
			return nil, fmt.Errorf("[doReq] Req result: %d -- %s %s", res.StatusCode, req.Method, req.URL.Path)
		}

		log.Println("[doReq] Needs to authenticate...")

		usedToken := getUsedToken(req.Header)
		clt.reqTokenChan <- &auth.AuthTask{WrongToken: usedToken}
		newToken := <-clt.tokensChan
		clt.token = newToken

		injectToken(clt.token, req.Header)
		log.Println("Authentication returned no error, should have new token")

	}

	respPayload, errReadingBody := ioutil.ReadAll(res.Body)

	if errReadingBody != nil {
		return nil, errReadingBody
	}
	return respPayload, nil
}

func (clt *client) DoPost(subPath string, params map[string]string) ([]byte, error) {

	clt.ensureTokenNotEmpty()
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
