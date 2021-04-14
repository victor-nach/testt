package keystore

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"testt/graph/model"
)

func TestKeystore_Get(t *testing.T) {
	keyStore := New()
	url, err := keyStore.GetByUrl("someUrl")
	assert.Error(t, err)
	assert.Nil(t, url)

	//	save
	url, err = keyStore.Save(&model.URL{URL: "someurl"})
	assert.NotNil(t, url)

	//	save
	url, err = keyStore.Save(&model.URL{URL: "someurl"})
	assert.NotNil(t, url)
}

func TestKeystore_genCode(t *testing.T) {
	keyStore := New()
	url, err := keyStore.Get("someUrl")
	assert.Error(t, err)
	assert.Nil(t, url)

	//	gencode
	code := keyStore.genCode()
	assert.NotNil(t, code)
}
