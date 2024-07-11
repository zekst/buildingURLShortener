package workers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const userID = "fajfdkfdjfdfdsfdjfdf-dffdsfd-fdfdfdfkdf--fdd"

func TestShortLinkGenerator(t *testing.T){
	
	initialLink_1 := "wwww.whateveryouwantittobehappenwillbeahppendi.com" 

	shortURL_1 := generateShortURL(string(initialLink_1), userID)

	assert.Equal(t, shortURL_1, "ummdontknow")
}