package user

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"go-wishlist-api/controllers/base"
	"go-wishlist-api/controllers/user/request"
	"go-wishlist-api/controllers/user/response"
	"go-wishlist-api/entities"
	"go-wishlist-api/utils"
)

type UserController struct {
	userUseCase entities.IUserUseCase
}

func NewUserController(userUseCase entities.IUserUseCase) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}

func (u *UserController) Register(c echo.Context) error {
	var userRegister request.UserRegister
	c.Bind(&userRegister)

	user, err := u.userUseCase.Register(userRegister.ToEntities())

	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	userResponse := response.FromUseCaseRegister(user)

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Register User", userResponse))
}

func (u *UserController) Login(c echo.Context) error {
	var userLogin request.UserLogin
	c.Bind(&userLogin)

	token, err := u.userUseCase.Login(userLogin.ToEntities())

	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	userResponse := response.FromUseCaseLogin(token)

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Login User", userResponse))
}
