package fortnox_test

import (
	"encoding/json"
	"log"
	"testing"

	fortnox "github.com/omniboost/go-fortnox"
)

func TestPutCustomer(t *testing.T) {
	req := client.NewPutCustomerRequest()
	req.NewPutCustomerPathParams().CustomerNumber = "1"
	customer := req.RequestBody().Customer
	customer.Name = "TEST"
	customer.Active = true
	customer.Type = fortnox.CustomerTypePrivate
	customer.VATType = fortnox.VATTypeSEVAT
	req.RequestBody().Customer = customer
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
