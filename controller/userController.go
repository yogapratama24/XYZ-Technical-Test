package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"xyz_test/helpers"
	"xyz_test/model"
	"xyz_test/service"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	service service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{service}
}

func (userController *UserController) CreateUserController(c echo.Context) error {
	var (
		request                  model.UserCreateRequest
		fullPath, fullPathSelfie string
		err                      error
	)
	if err := c.Bind(&request); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return nil
	}

	if err := helpers.DoValidation(&request); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return nil
	}

	// UPLOAD FOTO KTP
	file, _ := c.FormFile("foto_ktp")
	if file != nil {
		path := "upload/user"
		fullPath, err = helpers.FileUploadHandler(file, path)
		if err != nil {
			helpers.NewHandlerResponse("Failed to upload image file", nil).Failed(c)
			log.Printf("Error upload file with: %v", err)
			return nil
		}
	}
	request.FotoKtp = fullPath

	// UPLOAD FOTO SELFIE
	fileSelfie, _ := c.FormFile("foto_selfie")
	if fileSelfie != nil {
		path := "upload/user_selfie"
		fullPathSelfie, err = helpers.FileUploadHandler(fileSelfie, path)
		if err != nil {
			helpers.NewHandlerResponse("Failed to upload image file", nil).Failed(c)
			log.Printf("Error upload file with: %v", err)
			return nil
		}
	}
	request.FotoSelfie = fullPathSelfie

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), 4)
	if err != nil {
		return err
	}
	request.Password = string(password)

	if err := userController.service.CreateUser(request); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}
	helpers.NewHandlerResponse("Berhasil input data diri, silahkan login!", nil).SuccessCreate(c)
	return nil
}

func (userController *UserController) LoginUserController(c echo.Context) error {
	var (
		request model.UserLoginRequest
	)
	if err := c.Bind(&request); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return nil
	}

	if err := helpers.DoValidation(&request); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return nil
	}

	userData, err := userController.service.LoginUser(request)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}

	errHash := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(request.Password))
	if errHash != nil {
		fmt.Printf("Email or Password Incorrect with err: %s\n", errHash)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Email or Password Is Incorrect",
			"success": false,
		})
	}

	tokenString, err := helpers.GenerateJWT(userData)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}
	id := strconv.Itoa(userData.Id)
	cookieToken := &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24 * 60),
		Path:     "/",
		SameSite: 2,
		HttpOnly: true,
	}
	cookieUserId := &http.Cookie{
		Name:     "user_id",
		Value:    id,
		Expires:  time.Now().Add(time.Hour * 24 * 60),
		Path:     "/",
		SameSite: 2,
		HttpOnly: true,
	}

	c.SetCookie(cookieToken)
	c.SetCookie(cookieUserId)
	helpers.NewHandlerResponse("Successfully Login", nil).Success(c)
	return nil
}
