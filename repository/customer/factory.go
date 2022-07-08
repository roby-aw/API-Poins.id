package customer

import (
	"api-redeem-point/business/customer"
	"api-redeem-point/utils"
)

func RepositoryFactory(dbCon *utils.DatabaseConnection) customer.Repository {
	dummyRepo := NewPostgresRepository(dbCon.Postgres)
	return dummyRepo
}
