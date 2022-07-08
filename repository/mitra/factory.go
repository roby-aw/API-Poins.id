package mitra

import (
	"api-redeem-point/business/mitra"
	"api-redeem-point/utils"
)

func RepositoryFactory(dbCon *utils.DatabaseConnection) mitra.Repository {
	dummyRepo := NewPostgresRepository(dbCon.Postgres)
	return dummyRepo
}
