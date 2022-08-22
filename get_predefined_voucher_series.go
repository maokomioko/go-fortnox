package fortnox

import (
	"net/http"
	"net/url"

	"github.com/maokomioko/go-fortnox/utils"
)

func (c *Client) NewGetPredefinedVoucherSeriesRequest() GetPredefinedVoucherSeriesRequest {
	r := GetPredefinedVoucherSeriesRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewGetPredefinedVoucherSeriesQueryParams()
	r.pathParams = r.NewGetPredefinedVoucherSeriesPathParams()
	r.requestBody = r.NewGetPredefinedVoucherSeriesRequestBody()
	return r
}

type GetPredefinedVoucherSeriesRequest struct {
	client      *Client
	queryParams *GetPredefinedVoucherSeriesQueryParams
	pathParams  *GetPredefinedVoucherSeriesPathParams
	method      string
	headers     http.Header
	requestBody GetPredefinedVoucherSeriesRequestBody
}

func (r GetPredefinedVoucherSeriesRequest) NewGetPredefinedVoucherSeriesQueryParams() *GetPredefinedVoucherSeriesQueryParams {
	return &GetPredefinedVoucherSeriesQueryParams{}
}

type GetPredefinedVoucherSeriesQueryParams struct {
	//Page   *Page
	//Limit  *Limit
	//Offset *Offset
}

func (p GetPredefinedVoucherSeriesQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetPredefinedVoucherSeriesRequest) QueryParams() *GetPredefinedVoucherSeriesQueryParams {
	return r.queryParams
}

func (r GetPredefinedVoucherSeriesRequest) NewGetPredefinedVoucherSeriesPathParams() *GetPredefinedVoucherSeriesPathParams {
	return &GetPredefinedVoucherSeriesPathParams{}
}

type GetPredefinedVoucherSeriesPathParams struct {
}

func (p *GetPredefinedVoucherSeriesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetPredefinedVoucherSeriesRequest) PathParams() *GetPredefinedVoucherSeriesPathParams {
	return r.pathParams
}

func (r *GetPredefinedVoucherSeriesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetPredefinedVoucherSeriesRequest) Method() string {
	return r.method
}

func (r GetPredefinedVoucherSeriesRequest) NewGetPredefinedVoucherSeriesRequestBody() GetPredefinedVoucherSeriesRequestBody {
	return GetPredefinedVoucherSeriesRequestBody{}
}

type GetPredefinedVoucherSeriesRequestBody struct{}

func (r *GetPredefinedVoucherSeriesRequest) RequestBody() *GetPredefinedVoucherSeriesRequestBody {
	return &r.requestBody
}

func (r *GetPredefinedVoucherSeriesRequest) SetRequestBody(body GetPredefinedVoucherSeriesRequestBody) {
	r.requestBody = body
}

func (r *GetPredefinedVoucherSeriesRequest) NewResponseBody() *GetPredefinedVoucherSeriesResponseBody {
	return &GetPredefinedVoucherSeriesResponseBody{}
}

type GetPredefinedVoucherSeriesResponseBody struct {
	MetaInformation                   `json:"MetaInformation"`
	PreDefinedVoucherSeriesCollection []struct {
		URL           string `json:"@url"`
		Name          string `json:"Name"`
		VoucherSeries string `json:"VoucherSeries"`
	} `json:"PreDefinedVoucherSeriesCollection"`
}

func (r *GetPredefinedVoucherSeriesRequest) URL() url.URL {
	return r.client.GetEndpointURL("/predefinedvoucherseries", r.PathParams())
}

func (r *GetPredefinedVoucherSeriesRequest) Do() (GetPredefinedVoucherSeriesResponseBody, error) {
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
