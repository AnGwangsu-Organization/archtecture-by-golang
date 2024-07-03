package repository

import (
	"eCommerce/config"
	"eCommerce/repository/mongo"
	"eCommerce/repository/mysql"
)

type Repository struct {
	config *config.Config

	Mongo *mongo.Mongo // * 복사된 구조체를 사용
	Mysql *mysql.Mysql // * 복사된 구조체를 사용
}

func NewRepository(config *config.Config) (*Repository, error) {
	r := &Repository{
		config: config,
	}
	var err error

	if r.Mongo, err = mongo.NewMongo(config); err != nil {
		panic(err)
	}

	if r.Mysql, err = mysql.NewMysql(config); err != nil {
		panic(err)
	}

	return r, nil
}
