package bitclient

import (
	"net/http"
)

// ##############################################################################
// Request
// ##############################################################################

type PagedParams struct {
	Limit uint `url:"limit,omitempty"`
	Start uint `url:"start,omitempty"`
}

type RequestError struct {
	Code    int
	Message string
}

func (r RequestError) Error() string {
	return r.Message
}

type CreateRepositoryRequest struct {
	Name     string `json:"name"`
	ScmId    string `json:"scmId"`
	Forkable bool   `json:"forkable"`
}

func (bc *BitClient) checkReponse(resp *http.Response, err error) (*http.Response, error) {

	if resp != nil && resp.StatusCode > 299 {
		return nil, RequestError{
			Code:    resp.StatusCode,
			Message: resp.Status + " - " + resp.Request.URL.String(),
		}
	}

	return resp, err
}

func (bc *BitClient) DoGetPaged(uri string, rData interface{}, rError interface{}) (*http.Response, error) {

	maxPagedParam := PagedParams{
		Limit: 10000,
		Start: 0,
	}

	resp, err := bc.sling.New().Get(uri).QueryStruct(maxPagedParam).Receive(rData, rError)

	return bc.checkReponse(resp, err)
}

func (bc *BitClient) DoPost(uri string, data interface{}, rData interface{}, rError interface{}) (*http.Response, error) {

	resp, err := bc.sling.New().Post(uri).BodyJSON(data).Receive(rData, rError)

	return bc.checkReponse(resp, err)
}

func (bc *BitClient) DoPut(uri string, data interface{}, rData interface{}, rError interface{}) (*http.Response, error) {
	resp, err := bc.sling.New().Put(uri).BodyJSON(data).Receive(rData, rError)

	return bc.checkReponse(resp, err)
}
