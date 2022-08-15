package fortnox

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-fortnox/utils"
)

func (c *Client) NewPutCustomerRequest() PutCustomerRequest {
	r := PutCustomerRequest{
		client:  c,
		method:  http.MethodPut,
		headers: http.Header{},
	}

	r.queryParams = r.NewPutCustomerQueryParams()
	r.pathParams = r.NewPutCustomerPathParams()
	r.requestBody = r.NewPutCustomerRequestBody()
	return r
}

type PutCustomerRequest struct {
	client      *Client
	queryParams *PutCustomerQueryParams
	pathParams  *PutCustomerPathParams
	method      string
	headers     http.Header
	requestBody PutCustomerRequestBody
}

func (r PutCustomerRequest) NewPutCustomerQueryParams() *PutCustomerQueryParams {
	return &PutCustomerQueryParams{
		// Pagination: odata.NewPagination(),
	}
}

type PutCustomerQueryParams struct {
}

func (p PutCustomerQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *PutCustomerRequest) QueryParams() *PutCustomerQueryParams {
	return r.queryParams
}

func (r PutCustomerRequest) NewPutCustomerPathParams() *PutCustomerPathParams {
	return &PutCustomerPathParams{}
}

type PutCustomerPathParams struct {
	CustomerNumber string
}

func (p *PutCustomerPathParams) Params() map[string]string {
	return map[string]string{
		"CustomerNumber": p.CustomerNumber,
	}
}

func (r *PutCustomerRequest) PathParams() *PutCustomerPathParams {
	return r.pathParams
}

func (r *PutCustomerRequest) SetMethod(method string) {
	r.method = method
}

func (r *PutCustomerRequest) Method() string {
	return r.method
}

func (r PutCustomerRequest) NewPutCustomerRequestBody() PutCustomerRequestBody {
	return PutCustomerRequestBody{}
}

type PutCustomerRequestBody struct {
	Customer Customer
}

func (r *PutCustomerRequest) RequestBody() *PutCustomerRequestBody {
	return &r.requestBody
}

func (r *PutCustomerRequest) SetRequestBody(body PutCustomerRequestBody) {
	r.requestBody = body
}

func (r *PutCustomerRequest) NewResponseBody() *PutCustomerResponseBody {
	return &PutCustomerResponseBody{}
}

type PutCustomerResponseBody struct {
	MetaInformation `json:"MetaInformation"`
	Customer        Customer
}

func (r *PutCustomerRequest) URL() url.URL {
	return r.client.GetEndpointURL("/customers/{{.CustomerNumber}}", r.PathParams())
}

func (r *PutCustomerRequest) Do() (PutCustomerResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), r.RequestBody())
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
