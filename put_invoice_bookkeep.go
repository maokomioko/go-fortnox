package fortnox

import (
	"net/http"
	"net/url"

	"github.com/maokomioko/go-fortnox/utils"
)

func (c *Client) NewPutInvoiceBookKeepRequest() PutInvoiceBookKeepRequest {
	r := PutInvoiceBookKeepRequest{
		client:  c,
		method:  http.MethodPut,
		headers: http.Header{},
	}

	r.queryParams = r.NewPutInvoiceBookKeepQueryParams()
	r.pathParams = r.NewPutInvoiceBookKeepPathParams()
	r.requestBody = r.NewPutInvoiceBookKeepRequestBody()
	return r
}

type PutInvoiceBookKeepRequest struct {
	client      *Client
	queryParams *PutInvoiceBookKeepQueryParams
	pathParams  *PutInvoiceBookKeepPathParams
	method      string
	headers     http.Header
	requestBody PutInvoiceBookKeepRequestBody
}

func (r PutInvoiceBookKeepRequest) NewPutInvoiceBookKeepQueryParams() *PutInvoiceBookKeepQueryParams {
	return &PutInvoiceBookKeepQueryParams{}
}

type PutInvoiceBookKeepQueryParams struct {
}

func (p PutInvoiceBookKeepQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *PutInvoiceBookKeepRequest) QueryParams() *PutInvoiceBookKeepQueryParams {
	return r.queryParams
}

func (r PutInvoiceBookKeepRequest) NewPutInvoiceBookKeepPathParams() *PutInvoiceBookKeepPathParams {
	return &PutInvoiceBookKeepPathParams{}
}

type PutInvoiceBookKeepPathParams struct {
	DocumentNumber string
}

func (p *PutInvoiceBookKeepPathParams) Params() map[string]string {
	return map[string]string{
		"DocumentNumber": p.DocumentNumber,
	}
}

func (r *PutInvoiceBookKeepRequest) PathParams() *PutInvoiceBookKeepPathParams {
	return r.pathParams
}

func (r *PutInvoiceBookKeepRequest) SetMethod(method string) {
	r.method = method
}

func (r *PutInvoiceBookKeepRequest) Method() string {
	return r.method
}

func (r PutInvoiceBookKeepRequest) NewPutInvoiceBookKeepRequestBody() PutInvoiceBookKeepRequestBody {
	return PutInvoiceBookKeepRequestBody{}
}

type PutInvoiceBookKeepRequestBody struct {
	Invoice Invoice
}

func (r *PutInvoiceBookKeepRequest) RequestBody() *PutInvoiceBookKeepRequestBody {
	return &r.requestBody
}

func (r *PutInvoiceBookKeepRequest) SetRequestBody(body PutInvoiceBookKeepRequestBody) {
	r.requestBody = body
}

func (r *PutInvoiceBookKeepRequest) NewResponseBody() *PutInvoiceBookKeepResponseBody {
	return &PutInvoiceBookKeepResponseBody{}
}

type PutInvoiceBookKeepResponseBody struct {
	MetaInformation `json:"MetaInformation"`
	Invoice         Invoice
}

func (r *PutInvoiceBookKeepRequest) URL() url.URL {
	return r.client.GetEndpointURL("/invoices/{{.DocumentNumber}}/bookkeep", r.PathParams())
}

func (r *PutInvoiceBookKeepRequest) Do() (PutInvoiceBookKeepResponseBody, error) {
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
