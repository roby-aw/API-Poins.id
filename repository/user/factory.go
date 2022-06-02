package user

import (
	"api-redeem-point/business/user"
	"api-redeem-point/utils"
)

func RepositoryFactory(dbCon *utils.DatabaseConnection) user.Repository {
	dummyRepo := NewPostgresRepository(dbCon.Postgres)
	return dummyRepo
}
