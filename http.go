package resharmonics

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var baseUrl = GetBaseUrl()

func (clt *client) DoGet(subPath string, params map[string]string) ([]byte, error) {

	if clt.token == nil {
		errAuthing := clt.auth()
		if errAuthing != nil {
			return nil, errAuthing
		}
	}

	req, errCreatingReq := createGetReq(subPath, params, *clt.token)
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

	if subPath != authEndpointSubpat && clt.token == nil {
		errAuthing := clt.auth()
		if errAuthing != nil {
			return nil, errAuthing
		}
	}

	req, errCreatingReq := createPostReq(subPath, params, clt.token)
	if errCreatingReq != nil {
		return nil, errCreatingReq
	}
	resPayload, errReadingBody := clt.doReq(req)

	if errReadingBody != nil {
		return nil, errReadingBody
	}
	return resPayload, nil
}

func createGetReq(subPath string, params map[string]string, token string) (*http.Request, error) {
	endpoint := fmt.Sprintf("%s/%s", baseUrl, subPath)
	req, errCreatingReq := createReq(http.MethodGet, endpoint, nil, &token)
	if errCreatingReq != nil {
		return nil, errCreatingReq
	}

	if params != nil {
		q := req.URL.Query()
		for k, v := range params {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	return req, nil
}

func createPostReq(subPath string, params map[string]string, token *string) (*http.Request, error) {
	endpoint := fmt.Sprintf("%s/%s", baseUrl, subPath)

	var body io.Reader
	var urlEncodedParamsLength = 0
	if params != nil {
		paramsAsURL := url.Values{}
		for k, v := range params {
			paramsAsURL.Set(k, v)
		}

		encodedData := paramsAsURL.Encode()
		urlEncodedParamsLength = len(encodedData)
		body = strings.NewReader(encodedData)
	}

	req, errCreatingReq := createReq(http.MethodPost, endpoint, body, token)
	if errCreatingReq != nil {
		return nil, errCreatingReq
	}
	if params != nil {
		req.Header.Add(contentTypeHTTPHeader, xWWWFormURLEncodedHTTPContentType)
		req.Header.Add(contentLengthHTTPHeader, strconv.Itoa(urlEncodedParamsLength))
	}

	return req, nil
}

func createReq(method string, endpoint string, body io.Reader, token *string) (*http.Request, error) {
	req, errCreatingReq := http.NewRequest(method, endpoint, body)
	if errCreatingReq != nil {
		return nil, errCreatingReq
	}
	if token != nil {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", *token))
	}
	return req, nil
}

const contentTypeHTTPHeader = `Content-Type`
const contentLengthHTTPHeader = `Content-Length`
const xWWWFormURLEncodedHTTPContentType = "application/x-www-form-urlencoded"
