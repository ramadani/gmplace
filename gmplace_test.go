package gmplace_test

import (
	"testing"

	"github.com/ramadani/gmplace"
	"github.com/stretchr/testify/assert"
)

func TestAutocomplete(t *testing.T) {
	assert := assert.New(t)
	gmplace := gmplace.New("YOUR_GOOGLE_API_KEY")
	res, err := gmplace.Autocomplete(map[string]string{
		"input":    "uii",
		"language": "id",
	})

	assert.Nil(err)
	assert.Equal(5, len(res.Predictions))
	assert.Equal("OK", res.Status)
}
