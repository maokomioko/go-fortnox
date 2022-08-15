package fortnox_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestDeleteCustomer(t *testing.T) {
	req := client.NewDeleteCustomerRequest()
	req.NewDeleteCustomerPathParams().CustomerNumber = "1"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
