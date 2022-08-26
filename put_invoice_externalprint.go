package fortnox

import (
	"net/http"
	"net/url"

	"github.com/maokomioko/go-fortnox/utils"
)

func (c *Client) NewPutInvoiceExternalPrintRequest() PutInvoiceExternalPrintRequest {
	r := PutInvoiceExternalPrintRequest{
		client:  c,
		method:  http.MethodPut,
		headers: http.Header{},
	}

	r.queryParams = r.NewPutInvoiceExternalPrintQueryParams()
	r.pathParams = r.NewPutInvoiceExternalPrintPathParams()
	r.requestBody = r.NewPutInvoiceExternalPrintRequestBody()
	return r
}

type PutInvoiceExternalPrintRequest struct {
	client      *Client
	queryParams *PutInvoiceExternalPrintQueryParams
	pathParams  *PutInvoiceExternalPrintPathParams
	method      string
	headers     http.Header
	requestBody PutInvoiceExternalPrintRequestBody
}

func (r PutInvoiceExternalPrintRequest) NewPutInvoiceExternalPrintQueryParams() *PutInvoiceExternalPrintQueryParams {
	return &PutInvoiceExternalPrintQueryParams{}
}

type PutInvoiceExternalPrintQueryParams struct {
}

func (p PutInvoiceExternalPrintQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *PutInvoiceExternalPrintRequest) QueryParams() *PutInvoiceExternalPrintQueryParams {
	return r.queryParams
}

func (r PutInvoiceExternalPrintRequest) NewPutInvoiceExternalPrintPathParams() *PutInvoiceExternalPrintPathParams {
	return &PutInvoiceExternalPrintPathParams{}
}

type PutInvoiceExternalPrintPathParams struct {
	DocumentNumber string
}

func (p *PutInvoiceExternalPrintPathParams) Params() map[string]string {
	return map[string]string{
		"DocumentNumber": p.DocumentNumber,
	}
}

func (r *PutInvoiceExternalPrintRequest) PathParams() *PutInvoiceExternalPrintPathParams {
	return r.pathParams
}

func (r *PutInvoiceExternalPrintRequest) SetMethod(method string) {
	r.method = method
}

func (r *PutInvoiceExternalPrintRequest) Method() string {
	return r.method
}

func (r PutInvoiceExternalPrintRequest) NewPutInvoiceExternalPrintRequestBody() PutInvoiceExternalPrintRequestBody {
	return PutInvoiceExternalPrintRequestBody{}
}

type PutInvoiceExternalPrintRequestBody struct {
	Invoice Invoice
}

func (r *PutInvoiceExternalPrintRequest) RequestBody() *PutInvoiceExternalPrintRequestBody {
	return &r.requestBody
}

func (r *PutInvoiceExternalPrintRequest) SetRequestBody(body PutInvoiceExternalPrintRequestBody) {
	r.requestBody = body
}

func (r *PutInvoiceExternalPrintRequest) NewResponseBody() *PutInvoiceExternalPrintResponseBody {
	return &PutInvoiceExternalPrintResponseBody{}
}

type PutInvoiceExternalPrintResponseBody struct {
	MetaInformation `json:"MetaInformation"`
	Invoice         Invoice
}

func (r *PutInvoiceExternalPrintRequest) URL() url.URL {
	return r.client.GetEndpointURL("/invoices/{{.DocumentNumber}}/externalprint", r.PathParams())
}

func (r *PutInvoiceExternalPrintRequest) Do() (PutInvoiceExternalPrintResponseBody, error) {
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
