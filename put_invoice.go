package fortnox

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-fortnox/utils"
)

func (c *Client) NewPutInvoiceRequest() PutInvoiceRequest {
	r := PutInvoiceRequest{
		client:  c,
		method:  http.MethodPut,
		headers: http.Header{},
	}

	r.queryParams = r.NewPutInvoiceQueryParams()
	r.pathParams = r.NewPutInvoicePathParams()
	r.requestBody = r.NewPutInvoiceRequestBody()
	return r
}

type PutInvoiceRequest struct {
	client      *Client
	queryParams *PutInvoiceQueryParams
	pathParams  *PutInvoicePathParams
	method      string
	headers     http.Header
	requestBody PutInvoiceRequestBody
}

func (r PutInvoiceRequest) NewPutInvoiceQueryParams() *PutInvoiceQueryParams {
	return &PutInvoiceQueryParams{}
}

type PutInvoiceQueryParams struct {
}

func (p PutInvoiceQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *PutInvoiceRequest) QueryParams() *PutInvoiceQueryParams {
	return r.queryParams
}

func (r PutInvoiceRequest) NewPutInvoicePathParams() *PutInvoicePathParams {
	return &PutInvoicePathParams{}
}

type PutInvoicePathParams struct {
	DocumentNumber string
}

func (p *PutInvoicePathParams) Params() map[string]string {
	return map[string]string{
		"DocumentNumber": p.DocumentNumber,
	}
}

func (r *PutInvoiceRequest) PathParams() *PutInvoicePathParams {
	return r.pathParams
}

func (r *PutInvoiceRequest) SetMethod(method string) {
	r.method = method
}

func (r *PutInvoiceRequest) Method() string {
	return r.method
}

func (r PutInvoiceRequest) NewPutInvoiceRequestBody() PutInvoiceRequestBody {
	return PutInvoiceRequestBody{}
}

type PutInvoiceRequestBody struct {
	Invoice Invoice
}

func (r *PutInvoiceRequest) RequestBody() *PutInvoiceRequestBody {
	return &r.requestBody
}

func (r *PutInvoiceRequest) SetRequestBody(body PutInvoiceRequestBody) {
	r.requestBody = body
}

func (r *PutInvoiceRequest) NewResponseBody() *PutInvoiceResponseBody {
	return &PutInvoiceResponseBody{}
}

type PutInvoiceResponseBody struct {
	MetaInformation `json:"MetaInformation"`
	Invoice         Invoice
}

func (r *PutInvoiceRequest) URL() url.URL {
	return r.client.GetEndpointURL("/invoices/{{.DocumentNumber}}", r.PathParams())
}

func (r *PutInvoiceRequest) Do() (PutInvoiceResponseBody, error) {
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
