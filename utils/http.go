package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func CreatePostReq(endPoint string, params map[string]string, token *string) (*http.Request, error) {

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

	req, errCreatingReq := createReq(http.MethodPost, endPoint, body, token)
	if errCreatingReq != nil {
		return nil, errCreatingReq
	}
	if params != nil {
		req.Header.Add(contentTypeHTTPHeader, xWWWFormURLEncodedHTTPContentType)
		req.Header.Add(contentLengthHTTPHeader, strconv.Itoa(urlEncodedParamsLength))
	}

	return req, nil
}

func CreateGetReq(endPoint string, params map[string]string, token string) (*http.Request, error) {
	req, errCreatingReq := createReq(http.MethodGet, endPoint, nil, &token)
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
