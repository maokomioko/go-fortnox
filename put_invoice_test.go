package fortnox_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestPutInvoice(t *testing.T) {
	req := client.NewPutInvoiceRequest()
	req.NewPutInvoicePathParams().DocumentNumber = "1"
	customer := req.RequestBody().Invoice
	customer.CustomerName = "Acme Inc"
	customer.CustomerNumber = "SE98956364601"
	req.RequestBody().Invoice = customer
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
