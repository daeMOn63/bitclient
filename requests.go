package bitclient

import (
	"fmt"
	"net/http"
)

const BASE_URI = "/rest/api/1.0"

type PagedRequest struct {
	Limit uint `url:"limit,omitempty"`
	Start uint `url:"start,omitempty"`
}

type RequestError struct {
	Code    int
	Message string
}

type ErrorResponse struct {
	Errors []Error
}

type PagedResponse struct {
	Size       uint `json:"size"`
	Limit      uint `json:"limit"`
	IsLastPage bool `json:"isLastPage"`
	Start      uint `json:"start"`
}

func (r RequestError) Error() string {
	return r.Message
}

func (bc *BitClient) checkReponse(resp *http.Response, errorResponse *ErrorResponse) (*http.Response, error) {

	if resp != nil && resp.StatusCode > 299 {

		message := fmt.Sprintf("%s - %s\n", resp.Status, resp.Request.URL.String())
		for _, e := range errorResponse.Errors {
			message += e.Message
		}

		return nil, RequestError{
			Code:    resp.StatusCode,
			Message: message,
		}
	}
	return resp, nil
}

func (bc *BitClient) DoGet(uri string, params interface{}, rData interface{}) (*http.Response, error) {

	rError := new(ErrorResponse)

	resp, _ := bc.sling.New().Get(BASE_URI+uri).QueryStruct(params).Receive(rData, rError)

	return bc.checkReponse(resp, rError)
}

func (bc *BitClient) DoPostUrl(uri string, params interface{}, rData interface{}) (*http.Response, error) {

	rError := new(ErrorResponse)

	resp, _ := bc.sling.New().Post(BASE_URI+uri).QueryStruct(params).Receive(rData, rError)

	return bc.checkReponse(resp, rError)
}

func (bc *BitClient) DoPost(uri string, data interface{}, rData interface{}) (*http.Response, error) {

	rError := new(ErrorResponse)

	resp, _ := bc.sling.New().Post(BASE_URI+uri).BodyJSON(data).Receive(rData, rError)

	return bc.checkReponse(resp, rError)
}

func (bc *BitClient) DoPut(uri string, data interface{}, rData interface{}) (*http.Response, error) {

	rError := new(ErrorResponse)

	resp, _ := bc.sling.New().Put(BASE_URI+uri).BodyJSON(data).Receive(rData, rError)

	return bc.checkReponse(resp, rError)
}

func (bc *BitClient) DoPutUrl(uri string, data interface{}, rData interface{}) (*http.Response, error) {

	rError := new(ErrorResponse)

	resp, _ := bc.sling.New().Put(BASE_URI+uri).QueryStruct(data).Receive(rData, rError)

	return bc.checkReponse(resp, rError)
}

func (bc *BitClient) DoDeleteUrl(uri string, params interface{}, rData interface{}) (*http.Response, error) {

	rError := new(ErrorResponse)

	resp, _ := bc.sling.New().Delete(BASE_URI+uri).QueryStruct(params).Receive(rData, rError)

	return bc.checkReponse(resp, rError)
}
