package sage

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sageone-za/odata"
	"github.com/omniboost/go-sageone-za/utils"
)

func (c *Client) NewGetCustomersRequest() GetCustomersRequest {
	r := GetCustomersRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewGetCustomersQueryParams()
	r.pathParams = r.NewGetCustomersPathParams()
	r.requestBody = r.NewGetCustomersRequestBody()
	return r
}

type GetCustomersRequest struct {
	client      *Client
	queryParams *GetCustomersQueryParams
	pathParams  *GetCustomersPathParams
	method      string
	headers     http.Header
	requestBody GetCustomersRequestBody
}

func (r GetCustomersRequest) NewGetCustomersQueryParams() *GetCustomersQueryParams {
	selectFields, _ := utils.Fields(&Customer{})
	return &GetCustomersQueryParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type GetCustomersQueryParams struct {
	CompanyID int `schema:"CompanyId"`

	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p GetCustomersQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetCustomersRequest) QueryParams() *GetCustomersQueryParams {
	return r.queryParams
}

func (r GetCustomersRequest) NewGetCustomersPathParams() *GetCustomersPathParams {
	return &GetCustomersPathParams{}
}

type GetCustomersPathParams struct {
}

func (p *GetCustomersPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetCustomersRequest) PathParams() *GetCustomersPathParams {
	return r.pathParams
}

func (r *GetCustomersRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetCustomersRequest) Method() string {
	return r.method
}

func (r GetCustomersRequest) NewGetCustomersRequestBody() GetCustomersRequestBody {
	return GetCustomersRequestBody{}
}

type GetCustomersRequestBody struct{}

func (r *GetCustomersRequest) RequestBody() *GetCustomersRequestBody {
	return &r.requestBody
}

func (r *GetCustomersRequest) SetRequestBody(body GetCustomersRequestBody) {
	r.requestBody = body
}

func (r *GetCustomersRequest) NewResponseBody() *GetCustomersResponseBody {
	return &GetCustomersResponseBody{}
}

type GetCustomersResponseBody struct {
	TotalResults    int `json:"TotalResults"`
	ReturnedResults int `json:"ReturnedResults"`
	Results         Customers
}

func (r *GetCustomersRequest) URL() url.URL {
	return r.client.GetEndpointURL("/Customer/Get", r.PathParams())
}

func (r *GetCustomersRequest) Do() (GetCustomersResponseBody, error) {
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

type Customers []Customer

type Customer struct {
	Name     string `json:"Name"`
	Category struct {
		Description string `json:"Description"`
		ID          int    `json:"ID"`
		Modified    string `json:"Modified"`
		Created     string `json:"Created"`
	} `json:"Category"`
	SalesRepresentativeID int `json:"SalesRepresentativeId"`
	SalesRepresentative   struct {
		ID        int    `json:"ID"`
		FirstName string `json:"FirstName"`
		LastName  string `json:"LastName"`
		Name      string `json:"Name"`
		Active    bool   `json:"Active"`
		Email     string `json:"Email"`
		Mobile    string `json:"Mobile"`
		Telephone string `json:"Telephone"`
		Created   string `json:"Created"`
		Modified  string `json:"Modified"`
	} `json:"SalesRepresentative"`
	TaxReference                string  `json:"TaxReference"`
	ContactName                 string  `json:"ContactName"`
	Telephone                   string  `json:"Telephone"`
	Fax                         string  `json:"Fax"`
	Mobile                      string  `json:"Mobile"`
	Email                       string  `json:"Email"`
	WebAddress                  string  `json:"WebAddress"`
	Active                      bool    `json:"Active"`
	IsObfuscated                bool    `json:"IsObfuscated"`
	Balance                     float64 `json:"Balance"`
	CreditLimit                 float64 `json:"CreditLimit"`
	CommunicationMethod         int     `json:"CommunicationMethod"`
	PostalAddress01             string  `json:"PostalAddress01"`
	PostalAddress02             string  `json:"PostalAddress02"`
	PostalAddress03             string  `json:"PostalAddress03"`
	PostalAddress04             string  `json:"PostalAddress04"`
	PostalAddress05             string  `json:"PostalAddress05"`
	DeliveryAddress01           string  `json:"DeliveryAddress01"`
	DeliveryAddress02           string  `json:"DeliveryAddress02"`
	DeliveryAddress03           string  `json:"DeliveryAddress03"`
	DeliveryAddress04           string  `json:"DeliveryAddress04"`
	DeliveryAddress05           string  `json:"DeliveryAddress05"`
	AutoAllocateToOldestInvoice bool    `json:"AutoAllocateToOldestInvoice"`
	EnableCustomerZone          bool    `json:"EnableCustomerZone"`
	CustomerZoneGUID            string  `json:"CustomerZoneGuid"`
	CashSale                    bool    `json:"CashSale"`
	TextField1                  string  `json:"TextField1"`
	TextField2                  string  `json:"TextField2"`
	TextField3                  string  `json:"TextField3"`
	NumericField1               float64 `json:"NumericField1"`
	NumericField2               float64 `json:"NumericField2"`
	NumericField3               float64 `json:"NumericField3"`
	YesNoField1                 bool    `json:"YesNoField1"`
	YesNoField2                 bool    `json:"YesNoField2"`
	YesNoField3                 bool    `json:"YesNoField3"`
	DateField1                  string  `json:"DateField1"`
	DateField2                  string  `json:"DateField2"`
	DateField3                  string  `json:"DateField3"`
	DefaultPriceListID          int     `json:"DefaultPriceListId"`
	DefaultPriceList            struct {
		ID          int    `json:"ID"`
		Description string `json:"Description"`
		IsDefault   bool   `json:"IsDefault"`
	} `json:"DefaultPriceList"`
	DefaultPriceListName       string  `json:"DefaultPriceListName"`
	AcceptsElectronicInvoices  bool    `json:"AcceptsElectronicInvoices"`
	Modified                   string  `json:"Modified"`
	Created                    string  `json:"Created"`
	BusinessRegistrationNumber string  `json:"BusinessRegistrationNumber"`
	TaxStatusVerified          string  `json:"TaxStatusVerified"`
	CurrencyID                 int     `json:"CurrencyId"`
	CurrencySymbol             string  `json:"CurrencySymbol"`
	HasActivity                bool    `json:"HasActivity"`
	DefaultDiscountPercentage  float64 `json:"DefaultDiscountPercentage"`
	DefaultTaxTypeID           int     `json:"DefaultTaxTypeId"`
	DefaultTaxType             struct {
		ID                int     `json:"ID"`
		Name              string  `json:"Name"`
		Percentage        float64 `json:"Percentage"`
		IsDefault         bool    `json:"IsDefault"`
		HasActivity       bool    `json:"HasActivity"`
		IsManualTax       bool    `json:"IsManualTax"`
		Active            bool    `json:"Active"`
		Created           string  `json:"Created"`
		Modified          string  `json:"Modified"`
		TaxTypeDefaultUID string  `json:"TaxTypeDefaultUID"`
	} `json:"DefaultTaxType"`
	DueDateMethodID              int  `json:"DueDateMethodId"`
	DueDateMethodValue           int  `json:"DueDateMethodValue"`
	CityID                       int  `json:"CityId"`
	HasSpecialCountryTax         bool `json:"HasSpecialCountryTax"`
	AccountingAgreement          bool `json:"AccountingAgreement"`
	HasSpecialCountryTaxActivity bool `json:"HasSpecialCountryTaxActivity"`
	ID                           int  `json:"ID"`
}
