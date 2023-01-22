package routes

import (
	"xyz_test/config"
	"xyz_test/controller"
	"xyz_test/middlewares"
	"xyz_test/repository"
	"xyz_test/service"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	db := config.Connect()

	// USER
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	// TRANSACTION
	transactionRepository := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepository)
	transactionController := controller.NewTransactionController(transactionService)

	e := echo.New()
	e.HideBanner = true

	v1 := e.Group("/api/v1")
	v1.Static("/public", "public/")

	user := v1.Group("/user")
	user.POST("", userController.CreateUserController)
	user.POST("/login", userController.LoginUserController)

	middleware := v1.Group("")
	middlewares.SetJwtMiddlewares(middleware)

	transaction := middleware.Group("/transaction")
	transaction.POST("", transactionController.CreateTransactionController)
	transaction.GET("", transactionController.ReadTransactionController)

	return e
}
