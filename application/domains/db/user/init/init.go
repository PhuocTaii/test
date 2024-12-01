package init

import (
	"ia-03-backend/application/domains/db/user/repository"
	"ia-03-backend/application/domains/db/user/service"
	commonmodels "ia-03-backend/application/models"
)

type Init struct {
	Service service.IService
}

func NewInit(lib *commonmodels.Lib) *Init {
	return &Init{
		Service: service.InitService(lib, repository.InitRepo(lib)),
	}
}
