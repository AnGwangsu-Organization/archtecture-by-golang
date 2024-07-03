package app

import (
	"eCommerce/config"
	"eCommerce/repository"
	"eCommerce/router"
	"eCommerce/service"
)

type Application struct {
	config *config.Config // * 복사된 구조체를 사용

	router     *router.Router
	repository *repository.Repository
	service    *service.Service
}

func NewApplication(config *config.Config) *Application {
	a := &Application{
		config: config,
	}

	var err error

	if a.repository, err = repository.NewRepository(config); err != nil {
		panic(err)
	}

	if a.router, err = router.NewRouter(config); err != nil {
		panic(err)
	}

	if a.service, err = service.NewService(config); err != nil {
		panic(err)
	}

	//TODO : 서버 실행

	return a
}
