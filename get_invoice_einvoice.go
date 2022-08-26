package fortnox

import (
	"net/http"
	"net/url"

	"github.com/maokomioko/go-fortnox/utils"
)

func (c *Client) NewGetInvoiceEInvoiceRequest() GetInvoiceEInvoiceRequest {
	r := GetInvoiceEInvoiceRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewGetInvoiceEInvoiceQueryParams()
	r.pathParams = r.NewGetInvoiceEInvoicePathParams()
	r.requestBody = r.NewGetInvoiceEInvoiceRequestBody()
	return r
}

type GetInvoiceEInvoiceRequest struct {
	client      *Client
	queryParams *GetInvoiceEInvoiceQueryParams
	pathParams  *GetInvoiceEInvoicePathParams
	method      string
	headers     http.Header
	requestBody GetInvoiceEInvoiceRequestBody
}

func (r GetInvoiceEInvoiceRequest) NewGetInvoiceEInvoiceQueryParams() *GetInvoiceEInvoiceQueryParams {
	return &GetInvoiceEInvoiceQueryParams{}
}

type GetInvoiceEInvoiceQueryParams struct {
	Pagination Pagination
}

func (p GetInvoiceEInvoiceQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetInvoiceEInvoiceRequest) QueryParams() *GetInvoiceEInvoiceQueryParams {
	return r.queryParams
}

func (r GetInvoiceEInvoiceRequest) NewGetInvoiceEInvoicePathParams() *GetInvoiceEInvoicePathParams {
	return &GetInvoiceEInvoicePathParams{}
}

type GetInvoiceEInvoicePathParams struct {
	DocumentNumber string
}

func (p *GetInvoiceEInvoicePathParams) Params() map[string]string {
	return map[string]string{
		"DocumentNumber": p.DocumentNumber,
	}
}

func (r *GetInvoiceEInvoiceRequest) PathParams() *GetInvoiceEInvoicePathParams {
	return r.pathParams
}

func (r *GetInvoiceEInvoiceRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetInvoiceEInvoiceRequest) Method() string {
	return r.method
}

func (r GetInvoiceEInvoiceRequest) NewGetInvoiceEInvoiceRequestBody() GetInvoiceEInvoiceRequestBody {
	return GetInvoiceEInvoiceRequestBody{}
}

type GetInvoiceEInvoiceRequestBody struct{}

func (r *GetInvoiceEInvoiceRequest) RequestBody() *GetInvoiceEInvoiceRequestBody {
	return &r.requestBody
}

func (r *GetInvoiceEInvoiceRequest) SetRequestBody(body GetInvoiceEInvoiceRequestBody) {
	r.requestBody = body
}

func (r *GetInvoiceEInvoiceRequest) NewResponseBody() *GetInvoiceEInvoiceResponseBody {
	return &GetInvoiceEInvoiceResponseBody{}
}

type GetInvoiceEInvoiceResponseBody struct {
	MetaInformation `json:"MetaInformation"`
	Invoice         Invoice
}

func (r *GetInvoiceEInvoiceRequest) URL() url.URL {
	return r.client.GetEndpointURL("/invoices/{{.DocumentNumber}}/einvoice", r.PathParams())
}

func (r *GetInvoiceEInvoiceRequest) Do() (GetInvoiceEInvoiceResponseBody, error) {
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
