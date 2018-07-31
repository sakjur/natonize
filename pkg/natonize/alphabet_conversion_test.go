package natonize

import "testing"

func TestConversion(t *testing.T) {
	response := ToNatoAlphabet("liMe")

	if response != "lima india capital-mike echo" {
		t.Log("Got: ", response)
		t.Fail()
	}
}
