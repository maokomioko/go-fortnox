package fortnox

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-fortnox/utils"
)

func (c *Client) NewGetInvoiceRequest() GetInvoiceRequest {
	r := GetInvoiceRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewGetInvoiceQueryParams()
	r.pathParams = r.NewGetInvoicePathParams()
	r.requestBody = r.NewGetInvoiceRequestBody()
	return r
}

type GetInvoiceRequest struct {
	client      *Client
	queryParams *GetInvoiceQueryParams
	pathParams  *GetInvoicePathParams
	method      string
	headers     http.Header
	requestBody GetInvoiceRequestBody
}

func (r GetInvoiceRequest) NewGetInvoiceQueryParams() *GetInvoiceQueryParams {
	return &GetInvoiceQueryParams{}
}

type GetInvoiceQueryParams struct {
	//Page   *Page
	//Limit  *Limit
	//Offset *Offset
}

func (p GetInvoiceQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetInvoiceRequest) QueryParams() *GetInvoiceQueryParams {
	return r.queryParams
}

func (r GetInvoiceRequest) NewGetInvoicePathParams() *GetInvoicePathParams {
	return &GetInvoicePathParams{}
}

type GetInvoicePathParams struct {
	DocumentNumber string
}

func (p *GetInvoicePathParams) Params() map[string]string {
	return map[string]string{
		"DocumentNumber": p.DocumentNumber,
	}
}

func (r *GetInvoiceRequest) PathParams() *GetInvoicePathParams {
	return r.pathParams
}

func (r *GetInvoiceRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetInvoiceRequest) Method() string {
	return r.method
}

func (r GetInvoiceRequest) NewGetInvoiceRequestBody() GetInvoiceRequestBody {
	return GetInvoiceRequestBody{}
}

type GetInvoiceRequestBody struct{}

func (r *GetInvoiceRequest) RequestBody() *GetInvoiceRequestBody {
	return &r.requestBody
}

func (r *GetInvoiceRequest) SetRequestBody(body GetInvoiceRequestBody) {
	r.requestBody = body
}

func (r *GetInvoiceRequest) NewResponseBody() *GetInvoiceResponseBody {
	return &GetInvoiceResponseBody{}
}

type GetInvoiceResponseBody struct {
	MetaInformation `json:"MetaInformation"`
	Invoice         Invoice
}

func (r *GetInvoiceRequest) URL() url.URL {
	return r.client.GetEndpointURL("/invoices/{{.DocumentNumber}}", r.PathParams())
}

func (r *GetInvoiceRequest) Do() (GetInvoiceResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), nil)
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
