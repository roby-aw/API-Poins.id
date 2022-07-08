package modules

import (
	"api-redeem-point/api"
	adminApi "api-redeem-point/api/admin"
	customerApi "api-redeem-point/api/customer"
	storeApi "api-redeem-point/api/store"
	adminBusiness "api-redeem-point/business/admin"
	customerBusiness "api-redeem-point/business/customer"
	storeBusiness "api-redeem-point/business/store"
	"api-redeem-point/config"
	adminRepo "api-redeem-point/repository/admin"
	customerRepo "api-redeem-point/repository/customer"
	storeRepo "api-redeem-point/repository/store"
	"api-redeem-point/utils"
)

func RegistrationModules(dbCon *utils.DatabaseConnection, config *config.AppConfig) api.Controller {

	adminPermitRepository := adminRepo.RepositoryFactory(dbCon)
	adminPermitService := adminBusiness.NewService(adminPermitRepository)
	adminPermitController := adminApi.NewController(adminPermitService)

	customerPermitRepository := customerRepo.RepositoryFactory(dbCon)
	customerPermitService := customerBusiness.NewService(customerPermitRepository)
	customerPermitController := customerApi.NewController(customerPermitService)

	storePermitRepository := storeRepo.RepositoryFactory(dbCon)
	storePermitService := storeBusiness.NewService(storePermitRepository)
	storePermitController := storeApi.NewController(storePermitService)

	controller := api.Controller{
		AdminControlller:   adminPermitController,
		CustomerController: customerPermitController,
		StoreController:    storePermitController,
	}
	return controller
}
