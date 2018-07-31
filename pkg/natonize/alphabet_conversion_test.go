package natonize

import "testing"

const (
	test1regular = "liMe"
	test1nato = "lima india capital-mike echo"
)

func TestNormalConversion(t *testing.T) {
	response := ToNatoAlphabet(test1regular)

	if response != test1nato {
		t.Log("Got: ", response)
		t.Fail()
	}
}

func TestReverseConversion(t *testing.T) {
	println(FromNatoAlphabet(test1nato))
}
