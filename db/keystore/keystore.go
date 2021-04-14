package keystore

import (
	"errors"
	"testt/db"
	"testt/graph/model"
	"testt/lib/codegen"
)

var _ db.Datastore = &keyStore{}

type keyStore struct {
	CodeGen codegen.Codegen
	store   map[string]*model.URL
}

//New returns a new instance of keystore that implements datastore interface
func New() *keyStore {
	store := make(map[string]*model.URL, 0)
	return &keyStore{
		store:   store,
		CodeGen: codegen.New(),
	}
}

func (k keyStore) Get(code string) (*model.URL, error) {
	url := &model.URL{}
	for _, u := range k.store {
		if u.Code == code {
			url = u
			return url, nil
		}

	}
	return nil, errors.New("short code not found")
}

func (k keyStore) Save(url *model.URL) (*model.URL, error) {
	code := k.genCode()
	url.Code = code
	k.store[url.Code] = url
	return url, nil
}

func (k keyStore) GetByUrl(longUrl string) (*model.URL, error) {
	url := &model.URL{}
	for _, u := range k.store {
		if u.URL == longUrl {
			url = u
			return url, nil
		}
	}

	return nil, errors.New("short code not found")
}

//genCode generate a shortcode and verify that it doesn't exist
func (k keyStore) genCode() string {
	code := k.CodeGen.Generate()
	_, err := k.Get(code)
	if err != nil {
		return code
	}
	return k.genCode()
}
