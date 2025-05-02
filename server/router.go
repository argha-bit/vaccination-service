package server

import (
	"log"
	"net/http"
	"os"
	"vaccination-service/adapters/mysql"
	vaccineControl "vaccination-service/controller/vaccinationservice"
	rp "vaccination-service/repository/vaccinedrive"
	"vaccination-service/request"
	"vaccination-service/response"
	uc "vaccination-service/usecase/vaccinedrive"
	"vaccination-service/utils/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func newRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	e.Validator = validator.NewValidator()
	dbConn, err := mysql.GetMySQLConnect()
	if err != nil {
		log.Println("error in connecting to db", err.Error())
		os.Exit(1)
	}

	vaccineDriveRequest := request.NewVaccineDriveRequestHandler()
	vaccineDriveReqpository := rp.NewVaccineRepositoryHandler(dbConn)
	vaccineDriveUsecase := uc.NewVaccineRepositoryHandler(vaccineDriveReqpository)
	vaccineResponse := response.NewVaccineDriveResponseHandler()
	vaccineControl.NewVaccineServiceController(e, vaccineDriveRequest, vaccineDriveUsecase, vaccineResponse)

	return e
}
