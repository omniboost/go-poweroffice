package vismaonline_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestLedgerGetAll(t *testing.T) {
	req := client.NewLedgerGetAll()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
