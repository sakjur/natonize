package natonize

import "testing"

const (
	test1regular = "liMe"
	test1nato    = "lima india capital-mike echo"
)

func TestNormalConversion(t *testing.T) {
	response := ToNatoAlphabet(test1regular)

	if response != test1nato {
		t.Logf("Expected: %s, got: %s", test1nato, response)
		t.Fail()
	}
}

func TestReverseConversion(t *testing.T) {
	response := FromNatoAlphabet(test1nato)

	if response != test1regular {
		t.Logf("Expected: %s, got: %s", test1regular, response)
		t.Fail()
	}
}
