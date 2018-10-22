package gmplace_test

import (
	"testing"

	"github.com/ramadani/gmplace"
)

func TestAutocomplete(t *testing.T) {
	gmplace := gmplace.New("AIzaSyD-IGxkCSiS_EcEnjbi-KTriNKc-ATolTA")
	gmplace.Autocomplete("uii")
}
