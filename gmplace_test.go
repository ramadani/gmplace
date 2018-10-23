package gmplace_test

import (
	"testing"

	"github.com/ramadani/gmplace"
	"github.com/stretchr/testify/assert"
)

func TestAutocomplete(t *testing.T) {
	assert := assert.New(t)
	gmplace := gmplace.New("GOOGLE_API_KEY")
	res, err := gmplace.Autocomplete("uii")

	assert.Nil(err)
	assert.Equal(5, len(res.Predictions))
	assert.Equal("OK", res.Status)
}
