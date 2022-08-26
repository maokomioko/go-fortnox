package fortnox

import (
	"net/http"
	"net/url"

	"github.com/maokomioko/go-fortnox/utils"
)

func (c *Client) NewPutInvoiceCancelRequest() PutInvoiceCancelRequest {
	r := PutInvoiceCancelRequest{
		client:  c,
		method:  http.MethodPut,
		headers: http.Header{},
	}

	r.queryParams = r.NewPutInvoiceCancelQueryParams()
	r.pathParams = r.NewPutInvoiceCancelPathParams()
	r.requestBody = r.NewPutInvoiceCancelRequestBody()
	return r
}

type PutInvoiceCancelRequest struct {
	client      *Client
	queryParams *PutInvoiceCancelQueryParams
	pathParams  *PutInvoiceCancelPathParams
	method      string
	headers     http.Header
	requestBody PutInvoiceCancelRequestBody
}

func (r PutInvoiceCancelRequest) NewPutInvoiceCancelQueryParams() *PutInvoiceCancelQueryParams {
	return &PutInvoiceCancelQueryParams{}
}

type PutInvoiceCancelQueryParams struct {
}

func (p PutInvoiceCancelQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *PutInvoiceCancelRequest) QueryParams() *PutInvoiceCancelQueryParams {
	return r.queryParams
}

func (r PutInvoiceCancelRequest) NewPutInvoiceCancelPathParams() *PutInvoiceCancelPathParams {
	return &PutInvoiceCancelPathParams{}
}

type PutInvoiceCancelPathParams struct {
	DocumentNumber string
}

func (p *PutInvoiceCancelPathParams) Params() map[string]string {
	return map[string]string{
		"DocumentNumber": p.DocumentNumber,
	}
}

func (r *PutInvoiceCancelRequest) PathParams() *PutInvoiceCancelPathParams {
	return r.pathParams
}

func (r *PutInvoiceCancelRequest) SetMethod(method string) {
	r.method = method
}

func (r *PutInvoiceCancelRequest) Method() string {
	return r.method
}

func (r PutInvoiceCancelRequest) NewPutInvoiceCancelRequestBody() PutInvoiceCancelRequestBody {
	return PutInvoiceCancelRequestBody{}
}

type PutInvoiceCancelRequestBody struct {
	Invoice Invoice
}

func (r *PutInvoiceCancelRequest) RequestBody() *PutInvoiceCancelRequestBody {
	return &r.requestBody
}

func (r *PutInvoiceCancelRequest) SetRequestBody(body PutInvoiceCancelRequestBody) {
	r.requestBody = body
}

func (r *PutInvoiceCancelRequest) NewResponseBody() *PutInvoiceCancelResponseBody {
	return &PutInvoiceCancelResponseBody{}
}

type PutInvoiceCancelResponseBody struct {
	MetaInformation `json:"MetaInformation"`
	Invoice         Invoice
}

func (r *PutInvoiceCancelRequest) URL() url.URL {
	return r.client.GetEndpointURL("/invoices/{{.DocumentNumber}}/cancel", r.PathParams())
}

func (r *PutInvoiceCancelRequest) Do() (PutInvoiceCancelResponseBody, error) {
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
