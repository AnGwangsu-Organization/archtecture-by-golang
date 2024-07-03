package mysql

import "eCommerce/config"

type Mysql struct {
	config *config.Config
}

func NewMysql(config *config.Config) (*Mysql, error) {
	m := &Mysql{
		config: config,
	}

	return m, nil
}
