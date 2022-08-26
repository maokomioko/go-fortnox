package fortnox

import (
	"net/http"
	"net/url"

	"github.com/maokomioko/go-fortnox/utils"
)

func (c *Client) NewPutInvoicePrintRequest() PutInvoicePrintRequest {
	r := PutInvoicePrintRequest{
		client:  c,
		method:  http.MethodPut,
		headers: http.Header{},
	}

	r.queryParams = r.NewPutInvoicePrintQueryParams()
	r.pathParams = r.NewPutInvoicePrintPathParams()
	r.requestBody = r.NewPutInvoicePrintRequestBody()
	return r
}

type PutInvoicePrintRequest struct {
	client      *Client
	queryParams *PutInvoicePrintQueryParams
	pathParams  *PutInvoicePrintPathParams
	method      string
	headers     http.Header
	requestBody PutInvoicePrintRequestBody
}

func (r PutInvoicePrintRequest) NewPutInvoicePrintQueryParams() *PutInvoicePrintQueryParams {
	return &PutInvoicePrintQueryParams{}
}

type PutInvoicePrintQueryParams struct {
}

func (p PutInvoicePrintQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *PutInvoicePrintRequest) QueryParams() *PutInvoicePrintQueryParams {
	return r.queryParams
}

func (r PutInvoicePrintRequest) NewPutInvoicePrintPathParams() *PutInvoicePrintPathParams {
	return &PutInvoicePrintPathParams{}
}

type PutInvoicePrintPathParams struct {
	DocumentNumber string
}

func (p *PutInvoicePrintPathParams) Params() map[string]string {
	return map[string]string{
		"DocumentNumber": p.DocumentNumber,
	}
}

func (r *PutInvoicePrintRequest) PathParams() *PutInvoicePrintPathParams {
	return r.pathParams
}

func (r *PutInvoicePrintRequest) SetMethod(method string) {
	r.method = method
}

func (r *PutInvoicePrintRequest) Method() string {
	return r.method
}

func (r PutInvoicePrintRequest) NewPutInvoicePrintRequestBody() PutInvoicePrintRequestBody {
	return PutInvoicePrintRequestBody{}
}

type PutInvoicePrintRequestBody struct {
	Invoice Invoice
}

func (r *PutInvoicePrintRequest) RequestBody() *PutInvoicePrintRequestBody {
	return &r.requestBody
}

func (r *PutInvoicePrintRequest) SetRequestBody(body PutInvoicePrintRequestBody) {
	r.requestBody = body
}

func (r *PutInvoicePrintRequest) NewResponseBody() *PutInvoicePrintResponseBody {
	return &PutInvoicePrintResponseBody{}
}

type PutInvoicePrintResponseBody struct {
	MetaInformation `json:"MetaInformation"`
	Invoice         Invoice
}

func (r *PutInvoicePrintRequest) URL() url.URL {
	return r.client.GetEndpointURL("/invoices/{{.DocumentNumber}}/print", r.PathParams())
}

func (r *PutInvoicePrintRequest) Do() (PutInvoicePrintResponseBody, error) {
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
