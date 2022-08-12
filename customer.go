package fortnox

import "github.com/omniboost/go-fortnox/utils"

type Customers []Customer

type Customer struct {
	URL                      utils.URL            `json:"@url,omitempty"`            // Direct URL to the record.
	Active                   bool                 `json:"Active"`                    // If the customer is active or not
	Address1                 string               `json:"Address1"`                  // Address 1 of the customer.
	Address2                 string               `json:"Address2"`                  // Address 2 of the customer.
	City                     string               `json:"City"`                      // City of the customer.
	Country                  string               `json:"Country,omitempty"`         // Country name for the customer.
	Comments                 string               `json:"Comments"`                  // Comments of the customer.
	Currency                 string               `json:"Currency"`                  // Code of the currency for the customer. This will be used as the predefined currency for documents for the customer. The code must be of an existing currency.
	CostCenter               string               `json:"CostCenter"`                // Code of the cost center for the customer. The code must be of a an existing currency.
	CountryCode              string               `json:"CountryCode"`               // Code of the country for the customer. The code must be of an existing country according to ISO 3166-1 Alpha-2.
	CustomerNumber           string               `json:"CustomerNumber"`            // Customer number of the customer. If no customer number is provided, the next number in the series will be used. Only alpha numeric characters, with the addition of – + / and _, are allowed.
	DefaultDeliveryTypes     DefaultDeliveryTypes `json:"DefaultDeliveryTypes"`      // The properties for this object is listed in the table for “Default Delivery Types”.
	DefaultTemplates         DefaultTemplates     `json:"DefaultTemplates"`          // The properties for this object is listed in the table for “Default Templates”.
	DeliveryAddress1         string               `json:"DeliveryAddress1"`          // Delivery address 1 for the customer.
	DeliveryAddress2         string               `json:"DeliveryAddress2"`          // Delivery address 2 for the customer.
	DeliveryCity             string               `json:"DeliveryCity"`              // Delivery city for the customer.
	DeliveryCountry          string               `json:"DeliveryCountry,omitempty"` // Delivery country for the customer.
	DeliveryCountryCode      string               `json:"DeliveryCountryCode"`       // Code of the delivery country for the customer. The code must be of an existing country according to ISO 3166-1 Alpha-2.
	DeliveryFax              string               `json:"DeliveryFax"`               // Delivery fax number of the customer.
	DeliveryName             string               `json:"DeliveryName"`              // Delivery name for the customer.
	DeliveryPhone1           string               `json:"DeliveryPhone1"`            // Delivery phone number 1 for the customer.
	DeliveryPhone2           string               `json:"DeliveryPhone2"`            // Delivery phone number 2 for the customer.
	DeliveryZipCode          string               `json:"DeliveryZipCode"`           // Delivery zip code for the customer.
	Email                    string               `json:"Email"`                     // Email address for the customer. This must be a valid email address.
	EmailInvoice             string               `json:"EmailInvoice"`              // Specific email address used for invoices sent to the customer. This must be a valid email address.
	EmailInvoiceBCC          string               `json:"EmailInvoiceBCC"`           // Specific blind carbon copy email address used for invoices sent to the customer. This must be a valid email address.
	EmailInvoiceCC           string               `json:"EmailInvoiceCC"`            // Specific carbon copy email address used for invoices sent to the customer. This must be a valid email address.
	EmailOffer               string               `json:"EmailOffer"`                // Specific email address used for offers sent to the customer. This must be a valid email address.
	EmailOfferBCC            string               `json:"EmailOfferBCC"`             // Specific blind carbon copy email address used for offers sent to the customer. This must be a valid email address.
	EmailOfferCC             string               `json:"EmailOfferCC"`              // Specific carbon copy email address used for offers sent to the customer. This must be a valid email address.
	EmailOrder               string               `json:"EmailOrder"`                // Specific email address used for orders sent to the customer. This must be a valid email address.
	EmailOrderBCC            string               `json:"EmailOrderBCC"`             // Specific blind carbon copy email address used for orders sent to the customer. This must be a valid email address.
	EmailOrderCC             string               `json:"EmailOrderCC"`              // Specific carbon copy email address used for orders sent to the customer. This must be a valid email address.
	ExternalReference        string               `json:"ExternalReference"`
	Fax                      string               `json:"Fax"`                       // Fax number for the customer.
	GLN                      string               `json:"GLN"`                       // Global Location Number of the customer
	GLNDelivery              string               `json:"GLNDelivery"`               // Global Location Delivery Number
	InvoiceAdministrationFee string               `json:"InvoiceAdministrationFee"`  // Predefined invoice administration fee for the customer.
	InvoiceDiscount          float64              `json:"InvoiceDiscount"`           // Predefined invoice discount for the customer.
	InvoiceFreight           string               `json:"InvoiceFreight"`            // Predefined invoice freight fee for the customer.
	InvoiceRemark            string               `json:"InvoiceRemark"`             // Predefined invoice remark for the customer.
	Name                     string               `json:"Name"`                      // Name of the customer.
	OrganisationNumber       string               `json:"OrganisationNumber"`        // Organisation number of the customer. It needs to be a valid organisation numer.
	OurReference             string               `json:"OurReference"`              // Our reference of the customer.
	Phone                    string               `json:"Phone,omitempty"`           // Phone number of the customer.
	Phone1                   string               `json:"Phone1"`                    // Phone number 1 of the customer.
	Phone2                   string               `json:"Phone2"`                    // Phone number 2 of the customer.
	PriceList                string               `json:"PriceList"`                 // Code of the price list for the customer. The code must be of a an existing price list.
	Project                  string               `json:"Project"`                   // Number of the project for the customer. The number must be of a an existing project.
	SalesAccount             string               `json:"SalesAccount"`              // Predefined sales account of the customer.
	ShowPriceVATIncluded     bool                 `json:"ShowPriceVATIncluded"`      // If prices should be displayed with VAT included.
	TermsOfDelivery          string               `json:"TermsOfDelivery"`           // Code of the terms of delivery for the customer. The code must be of a an existing terms of delivery.
	TermsOfPayment           string               `json:"TermsOfPayment"`            // Code of the terms of payment for the customer. The code must be of a an existing terms of payment.
	Type                     CustomerType         `json:"Type"`                      // Type of the customer. Can be PRIVATE or COMPANY.
	VATNumber                string               `json:"VATNumber"`                 // VAT number for the customer.
	VATType                  VATType              `json:"VATType"`                   // VAT type of the customer. Can be SEVAT SEREVERSEDVAT EUREVERSEDVAT EUVAT or EXPORT.
	VisitingAddress          string               `json:"VisitingAddress"`           // Visiting address of the customer.
	VisitingCity             string               `json:"VisitingCity"`              // Visiting city of the customer.
	VisitingCountry          string               `json:"VisitingCountry,omitempty"` // Visiting country of the customer.
	VisitingCountryCode      string               `json:"VisitingCountryCode"`       // Code of the visiting country for the customer. The code must be of an existing country according to ISO 3166-1 Alpha-2.
	VisitingZipCode          string               `json:"VisitingZipCode"`           // Visiting zip code of the customer.
	WWW                      string               `json:"WWW"`                       // Website of the customer.
	WayOfDelivery            string               `json:"WayOfDelivery"`             // Code of the way of delivery for the customer. The code must be of a an existing way of delivery.
	YourReference            string               `json:"YourReference"`             // Your reference of the customer.
	ZipCode                  string               `json:"ZipCode"`                   // Zip code of the customers.
}

type DefaultDeliveryTypes struct {
	Invoice DeliveryType
	Order   DeliveryType
	Offer   DeliveryType
}

type DeliveryType string

var (
	DeliveryTypePrint        DeliveryType = "PRINT"
	DeliveryTypeEmail        DeliveryType = "EMAIL"
	DeliveryTypePrintService DeliveryType = "PRINTSERVICE"
)

type DefaultTemplates struct {
	Order       string `json:"Order"`
	Offer       string `json:"Offer"`
	Invoice     string `json:"Invoice"`
	CashInvoice string `json:"CashInvoice"`
}

type CustomerType string

var (
	CustomerTypePrivate CustomerType = "PRIVATE"
	CustomerTypeCompany CustomerType = "COMPANY"
)

// VATType of the customer. Can be SEVAT SEREVERSEDVAT EUREVERSEDVAT EUVAT or EXPORT.
type VATType string

var (
	VATTypeSEVAT         VATType = "SEVAT"
	VATTypeSEReversedVAT VATType = "SEREVERSEDVAT"
	VATTypeEUReversedVAT VATType = "EUREVERSEDVAT"
	VATTypeEUVat         VATType = "EUVAT"
	VATTypeExport        VATType = "EXPORT"
)
