package fortnox

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-fortnox/utils"
)

func (c *Client) NewDeleteCustomerRequest() DeleteCustomerRequest {
	r := DeleteCustomerRequest{
		client:  c,
		method:  http.MethodDelete,
		headers: http.Header{},
	}

	r.queryParams = r.NewDeleteCustomerQueryParams()
	r.pathParams = r.NewDeleteCustomerPathParams()
	r.requestBody = r.NewDeleteCustomerRequestBody()
	return r
}

type DeleteCustomerRequest struct {
	client      *Client
	queryParams *DeleteCustomerQueryParams
	pathParams  *DeleteCustomerPathParams
	method      string
	headers     http.Header
	requestBody DeleteCustomerRequestBody
}

func (r DeleteCustomerRequest) NewDeleteCustomerQueryParams() *DeleteCustomerQueryParams {
	return &DeleteCustomerQueryParams{}
}

type DeleteCustomerQueryParams struct{}

func (p DeleteCustomerQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *DeleteCustomerRequest) QueryParams() *DeleteCustomerQueryParams {
	return r.queryParams
}

func (r DeleteCustomerRequest) NewDeleteCustomerPathParams() *DeleteCustomerPathParams {
	return &DeleteCustomerPathParams{}
}

type DeleteCustomerPathParams struct {
	CustomerNumber string
}

func (p *DeleteCustomerPathParams) Params() map[string]string {
	return map[string]string{
		"CustomerNumber": p.CustomerNumber,
	}
}

func (r *DeleteCustomerRequest) PathParams() *DeleteCustomerPathParams {
	return r.pathParams
}

func (r *DeleteCustomerRequest) SetMethod(method string) {
	r.method = method
}

func (r *DeleteCustomerRequest) Method() string {
	return r.method
}

func (r DeleteCustomerRequest) NewDeleteCustomerRequestBody() DeleteCustomerRequestBody {
	return DeleteCustomerRequestBody{}
}

type DeleteCustomerRequestBody struct{}

func (r *DeleteCustomerRequest) RequestBody() *DeleteCustomerRequestBody {
	return &r.requestBody
}

func (r *DeleteCustomerRequest) SetRequestBody(body DeleteCustomerRequestBody) {
	r.requestBody = body
}

func (r *DeleteCustomerRequest) NewResponseBody() *DeleteCustomerResponseBody {
	return &DeleteCustomerResponseBody{}
}

type DeleteCustomerResponseBody struct {
}

func (r *DeleteCustomerRequest) URL() url.URL {
	return r.client.GetEndpointURL("/customers/{{.CustomerNumber}}", r.PathParams())
}

func (r *DeleteCustomerRequest) Do() (*DeleteCustomerResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), nil)
	if err != nil {
		return r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return r.NewResponseBody(), err
	}

	_, err = r.client.Do(req, nil)
	return nil, err
}
