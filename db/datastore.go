package db

import "testt/graph/model"

type Datastore interface {
	Get(code string) (*model.URL, error)
	GetByUrl(url string) (*model.URL, error)
	Save(url *model.URL) (*model.URL, error)
}
