package controller

import (
	"xyz_test/helpers"
	"xyz_test/model"
	"xyz_test/service"

	"github.com/labstack/echo"
)

type TransactionController struct {
	service service.TransactionService
}

func NewTransactionController(service service.TransactionService) *TransactionController {
	return &TransactionController{service}
}

func (transactionController *TransactionController) CreateTransactionController(c echo.Context) error {
	var (
		request model.CreateTransactionRequest
	)
	if err := c.Bind(&request); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return nil
	}
	auth, _ := helpers.ParseJWT(c)
	request.UserId = auth.Id

	if err := helpers.DoValidation(&request); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return nil
	}

	if err := transactionController.service.CreateTransaction(request); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}
	helpers.NewHandlerResponse("Berhasil input data cicilan", nil).SuccessCreate(c)
	return nil
}

func (transactionController *TransactionController) ReadTransactionController(c echo.Context) error {
	auth, _ := helpers.ParseJWT(c)
	data, err := transactionController.service.ReadTransaction(auth.Id)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}
	helpers.NewHandlerResponse("Berhasil mengambil data cicilan", data).Success(c)
	return nil
}
