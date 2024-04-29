package routes

import (
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	u_controllers "go-wishlist-api/controllers/user"
	w_controllers "go-wishlist-api/controllers/wishlist"
)

type RouteController struct {
	WishlistController *w_controllers.WishlistController
	UserController     *u_controllers.UserController
}

func (r *RouteController) InitRoute(e *echo.Echo) {

	wishlist := e.Group("/wishlists")
	wishlist.Use(echojwt.JWT([]byte(os.Getenv("SECRET_TOKEN"))))
	wishlist.GET("", r.WishlistController.GetAll)
	wishlist.POST("", r.WishlistController.Create)

	e.POST("/users/register", r.UserController.Register)
	e.POST("/users/login", r.UserController.Login)
}
