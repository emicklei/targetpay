package targetpay

import (
	"log"
	"testing"
)

func TestCheck(t *testing.T) {
	c := NewClient()
	ok, reason := c.CheckPayment(17894, 123456, "BETAAL+AA", 3010, 31, true)
	if !ok {
		log.Fatal(reason)
	} else {
		print(reason)
	}
}
