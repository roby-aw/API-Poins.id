package modules

import (
	"api-redeem-point/api"
	adminApi "api-redeem-point/api/admin"
	customerApi "api-redeem-point/api/customer"
	mitraApi "api-redeem-point/api/mitra"
	adminBusiness "api-redeem-point/business/admin"
	customerBusiness "api-redeem-point/business/customer"
	mitraBusiness "api-redeem-point/business/mitra"
	"api-redeem-point/config"
	adminRepo "api-redeem-point/repository/admin"
	customerRepo "api-redeem-point/repository/customer"
	mitraRepo "api-redeem-point/repository/mitra"
	"api-redeem-point/utils"
)

func RegistrationModules(dbCon *utils.DatabaseConnection, config *config.AppConfig) api.Controller {

	adminPermitRepository := adminRepo.RepositoryFactory(dbCon)
	adminPermitService := adminBusiness.NewService(adminPermitRepository)
	adminPermitController := adminApi.NewController(adminPermitService)

	customerPermitRepository := customerRepo.RepositoryFactory(dbCon)
	customerPermitService := customerBusiness.NewService(customerPermitRepository)
	customerPermitController := customerApi.NewController(customerPermitService)

	mitraPermitRepository := mitraRepo.RepositoryFactory(dbCon)
	mitraPermitService := mitraBusiness.NewService(mitraPermitRepository)
	mitraPermitController := mitraApi.NewController(mitraPermitService)

	controller := api.Controller{
		AdminControlller:        adminPermitController,
		CustomerMitraController: customerPermitController,
		MitraController:         mitraPermitController,
	}
	return controller
}
