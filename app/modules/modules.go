package modules

import (
	"api-redeem-point/api"
	adminApi "api-redeem-point/api/admin"
	userApi "api-redeem-point/api/user"
	adminBusiness "api-redeem-point/business/admin"
	userBusiness "api-redeem-point/business/user"
	"api-redeem-point/config"
	adminRepo "api-redeem-point/repository/admin"
	userRepo "api-redeem-point/repository/user"
	"api-redeem-point/utils"
)

func RegistrationModules(dbCon *utils.DatabaseConnection, config *config.AppConfig) api.Controller {

	adminPermitRepository := adminRepo.RepositoryFactory(dbCon)
	adminPermitService := adminBusiness.NewService(adminPermitRepository)
	adminPermitController := adminApi.NewController(adminPermitService)

	userPermitRepository := userRepo.RepositoryFactory(dbCon)
	userPermitService := userBusiness.NewService(userPermitRepository)
	userPermitController := userApi.NewController(userPermitService)

	controller := api.Controller{
		AdminControlller: adminPermitController,
		UserController:   userPermitController,
	}
	return controller
}
