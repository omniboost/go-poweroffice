package vismaonline_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestSalesOrderBasicV2Post(t *testing.T) {
	req := client.NewSalesOrderBasicV2Post()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
