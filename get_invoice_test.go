package fortnox_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetInvoice(t *testing.T) {
	req := client.NewGetInvoiceRequest()
	req.NewGetInvoicePathParams().DocumentNumber = "1"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
