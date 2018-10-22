package gmplace_test

import (
	"log"
	"testing"

	"github.com/ramadani/gmplace"
	"github.com/stretchr/testify/assert"
)

func TestAutocomplete(t *testing.T) {
	assert := assert.New(t)
	gmplace := gmplace.New("YOUR_GOOGLE_API_KEY")
	res, err := gmplace.Autocomplete("uii")
	if err != nil {
		log.Println(err)
	}

	assert.Nil(err)
	assert.Equal(5, len(res.Predictions))
	assert.Equal("OK", res.Status)
}
