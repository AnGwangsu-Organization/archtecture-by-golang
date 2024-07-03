package mongo

import "eCommerce/config"

type Mongo struct {
	config *config.Config
}

func NewMongo(config *config.Config) (*Mongo, error) {
	m := &Mongo{
		config: config,
	}

	return m, nil
}
