package init

import (
	user "ia-03-backend/application/domains/db/user/init"
	"ia-03-backend/application/models"
)

type DbService struct {
	User *user.Init
}

func NewInit(lib *models.Lib) *DbService {
	return &DbService{
		User: user.NewInit(lib),
	}
}
