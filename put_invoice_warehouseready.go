package fortnox

import (
	"net/http"
	"net/url"

	"github.com/maokomioko/go-fortnox/utils"
)

func (c *Client) NewPutInvoiceWarehouseReadyRequest() PutInvoiceWarehouseReadyRequest {
	r := PutInvoiceWarehouseReadyRequest{
		client:  c,
		method:  http.MethodPut,
		headers: http.Header{},
	}

	r.queryParams = r.NewPutInvoiceWarehouseReadyQueryParams()
	r.pathParams = r.NewPutInvoiceWarehouseReadyPathParams()
	r.requestBody = r.NewPutInvoiceWarehouseReadyRequestBody()
	return r
}

type PutInvoiceWarehouseReadyRequest struct {
	client      *Client
	queryParams *PutInvoiceWarehouseReadyQueryParams
	pathParams  *PutInvoiceWarehouseReadyPathParams
	method      string
	headers     http.Header
	requestBody PutInvoiceWarehouseReadyRequestBody
}

func (r PutInvoiceWarehouseReadyRequest) NewPutInvoiceWarehouseReadyQueryParams() *PutInvoiceWarehouseReadyQueryParams {
	return &PutInvoiceWarehouseReadyQueryParams{}
}

type PutInvoiceWarehouseReadyQueryParams struct {
}

func (p PutInvoiceWarehouseReadyQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *PutInvoiceWarehouseReadyRequest) QueryParams() *PutInvoiceWarehouseReadyQueryParams {
	return r.queryParams
}

func (r PutInvoiceWarehouseReadyRequest) NewPutInvoiceWarehouseReadyPathParams() *PutInvoiceWarehouseReadyPathParams {
	return &PutInvoiceWarehouseReadyPathParams{}
}

type PutInvoiceWarehouseReadyPathParams struct {
	DocumentNumber string
}

func (p *PutInvoiceWarehouseReadyPathParams) Params() map[string]string {
	return map[string]string{
		"DocumentNumber": p.DocumentNumber,
	}
}

func (r *PutInvoiceWarehouseReadyRequest) PathParams() *PutInvoiceWarehouseReadyPathParams {
	return r.pathParams
}

func (r *PutInvoiceWarehouseReadyRequest) SetMethod(method string) {
	r.method = method
}

func (r *PutInvoiceWarehouseReadyRequest) Method() string {
	return r.method
}

func (r PutInvoiceWarehouseReadyRequest) NewPutInvoiceWarehouseReadyRequestBody() PutInvoiceWarehouseReadyRequestBody {
	return PutInvoiceWarehouseReadyRequestBody{}
}

type PutInvoiceWarehouseReadyRequestBody struct {
	Invoice Invoice
}

func (r *PutInvoiceWarehouseReadyRequest) RequestBody() *PutInvoiceWarehouseReadyRequestBody {
	return &r.requestBody
}

func (r *PutInvoiceWarehouseReadyRequest) SetRequestBody(body PutInvoiceWarehouseReadyRequestBody) {
	r.requestBody = body
}

func (r *PutInvoiceWarehouseReadyRequest) NewResponseBody() *PutInvoiceWarehouseReadyResponseBody {
	return &PutInvoiceWarehouseReadyResponseBody{}
}

type PutInvoiceWarehouseReadyResponseBody struct {
	MetaInformation `json:"MetaInformation"`
	Invoice         Invoice
}

func (r *PutInvoiceWarehouseReadyRequest) URL() url.URL {
	return r.client.GetEndpointURL("/invoices/{{.DocumentNumber}}/warehouseready", r.PathParams())
}

func (r *PutInvoiceWarehouseReadyRequest) Do() (PutInvoiceWarehouseReadyResponseBody, error) {
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
