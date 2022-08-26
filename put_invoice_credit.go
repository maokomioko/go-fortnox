package fortnox

import (
	"net/http"
	"net/url"

	"github.com/maokomioko/go-fortnox/utils"
)

func (c *Client) NewPutInvoiceCreditRequest() PutInvoiceCreditRequest {
	r := PutInvoiceCreditRequest{
		client:  c,
		method:  http.MethodPut,
		headers: http.Header{},
	}

	r.queryParams = r.NewPutInvoiceCreditQueryParams()
	r.pathParams = r.NewPutInvoiceCreditPathParams()
	r.requestBody = r.NewPutInvoiceCreditRequestBody()
	return r
}

type PutInvoiceCreditRequest struct {
	client      *Client
	queryParams *PutInvoiceCreditQueryParams
	pathParams  *PutInvoiceCreditPathParams
	method      string
	headers     http.Header
	requestBody PutInvoiceCreditRequestBody
}

func (r PutInvoiceCreditRequest) NewPutInvoiceCreditQueryParams() *PutInvoiceCreditQueryParams {
	return &PutInvoiceCreditQueryParams{}
}

type PutInvoiceCreditQueryParams struct {
}

func (p PutInvoiceCreditQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *PutInvoiceCreditRequest) QueryParams() *PutInvoiceCreditQueryParams {
	return r.queryParams
}

func (r PutInvoiceCreditRequest) NewPutInvoiceCreditPathParams() *PutInvoiceCreditPathParams {
	return &PutInvoiceCreditPathParams{}
}

type PutInvoiceCreditPathParams struct {
	DocumentNumber string
}

func (p *PutInvoiceCreditPathParams) Params() map[string]string {
	return map[string]string{
		"DocumentNumber": p.DocumentNumber,
	}
}

func (r *PutInvoiceCreditRequest) PathParams() *PutInvoiceCreditPathParams {
	return r.pathParams
}

func (r *PutInvoiceCreditRequest) SetMethod(method string) {
	r.method = method
}

func (r *PutInvoiceCreditRequest) Method() string {
	return r.method
}

func (r PutInvoiceCreditRequest) NewPutInvoiceCreditRequestBody() PutInvoiceCreditRequestBody {
	return PutInvoiceCreditRequestBody{}
}

type PutInvoiceCreditRequestBody struct {
	Invoice Invoice
}

func (r *PutInvoiceCreditRequest) RequestBody() *PutInvoiceCreditRequestBody {
	return &r.requestBody
}

func (r *PutInvoiceCreditRequest) SetRequestBody(body PutInvoiceCreditRequestBody) {
	r.requestBody = body
}

func (r *PutInvoiceCreditRequest) NewResponseBody() *PutInvoiceCreditResponseBody {
	return &PutInvoiceCreditResponseBody{}
}

type PutInvoiceCreditResponseBody struct {
	MetaInformation `json:"MetaInformation"`
	Invoice         Invoice
}

func (r *PutInvoiceCreditRequest) URL() url.URL {
	return r.client.GetEndpointURL("/invoices/{{.DocumentNumber}}/credit", r.PathParams())
}

func (r *PutInvoiceCreditRequest) Do() (PutInvoiceCreditResponseBody, error) {
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
