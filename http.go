package resharmonics

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/JoseFMP/resharmonics/urls"
	"github.com/JoseFMP/resharmonics/utils"
)

var baseUrl = urls.GetBaseUrl()

func (clt *client) DoGet(subPath string, params map[string]string) ([]byte, error) {

	if clt.token == nil {
		errAuthing := clt.auth()
		if errAuthing != nil {
			return nil, errAuthing
		}
	}
	endpoint := fmt.Sprintf("%s/%s", baseUrl, subPath)
	req, errCreatingReq := utils.CreateGetReq(endpoint, params, *clt.token)
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
		log.Printf("[doReq] Doing req: %s %s", req.Method, req.URL.Path)
		res, errDoingReq = httpClient.Do(req)
		if errDoingReq != nil {
			return nil, errDoingReq
		}
		if res.StatusCode != http.StatusForbidden {
			log.Printf("[doReq] Req result: %d -- %s %s", res.StatusCode, req.Method, req.URL.Path)
			break
		}
		log.Println("[doReq] Needs to authenticate...")
		errAuthenticating := clt.auth()
		if errAuthenticating != nil {
			return nil, errAuthenticating
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

	if clt.token == nil {
		errAuthing := clt.auth()
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
