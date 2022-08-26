package fortnox

import (
	"net/http"
	"net/url"

	"github.com/maokomioko/go-fortnox/utils"
)

func (c *Client) NewGetInvoiceEPrintRequest() GetInvoiceEPrintRequest {
	r := GetInvoiceEPrintRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewGetInvoiceEPrintQueryParams()
	r.pathParams = r.NewGetInvoiceEPrintPathParams()
	r.requestBody = r.NewGetInvoiceEPrintRequestBody()
	return r
}

type GetInvoiceEPrintRequest struct {
	client      *Client
	queryParams *GetInvoiceEPrintQueryParams
	pathParams  *GetInvoiceEPrintPathParams
	method      string
	headers     http.Header
	requestBody GetInvoiceEPrintRequestBody
}

func (r GetInvoiceEPrintRequest) NewGetInvoiceEPrintQueryParams() *GetInvoiceEPrintQueryParams {
	return &GetInvoiceEPrintQueryParams{}
}

type GetInvoiceEPrintQueryParams struct {
	Pagination Pagination
}

func (p GetInvoiceEPrintQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetInvoiceEPrintRequest) QueryParams() *GetInvoiceEPrintQueryParams {
	return r.queryParams
}

func (r GetInvoiceEPrintRequest) NewGetInvoiceEPrintPathParams() *GetInvoiceEPrintPathParams {
	return &GetInvoiceEPrintPathParams{}
}

type GetInvoiceEPrintPathParams struct {
	DocumentNumber string
}

func (p *GetInvoiceEPrintPathParams) Params() map[string]string {
	return map[string]string{
		"DocumentNumber": p.DocumentNumber,
	}
}

func (r *GetInvoiceEPrintRequest) PathParams() *GetInvoiceEPrintPathParams {
	return r.pathParams
}

func (r *GetInvoiceEPrintRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetInvoiceEPrintRequest) Method() string {
	return r.method
}

func (r GetInvoiceEPrintRequest) NewGetInvoiceEPrintRequestBody() GetInvoiceEPrintRequestBody {
	return GetInvoiceEPrintRequestBody{}
}

type GetInvoiceEPrintRequestBody struct{}

func (r *GetInvoiceEPrintRequest) RequestBody() *GetInvoiceEPrintRequestBody {
	return &r.requestBody
}

func (r *GetInvoiceEPrintRequest) SetRequestBody(body GetInvoiceEPrintRequestBody) {
	r.requestBody = body
}

func (r *GetInvoiceEPrintRequest) NewResponseBody() *GetInvoiceEPrintResponseBody {
	return &GetInvoiceEPrintResponseBody{}
}

type GetInvoiceEPrintResponseBody struct {
	MetaInformation `json:"MetaInformation"`
	Invoice         Invoice
}

func (r *GetInvoiceEPrintRequest) URL() url.URL {
	return r.client.GetEndpointURL("/invoices/{{.DocumentNumber}}/eprint", r.PathParams())
}

func (r *GetInvoiceEPrintRequest) Do() (GetInvoiceEPrintResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), nil)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, true)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
