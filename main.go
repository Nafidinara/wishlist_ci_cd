package main

import (
	"github.com/labstack/echo/v4"

	"go-wishlist-api/config"
	cl_user "go-wishlist-api/controllers/user"
	cl_wishlist "go-wishlist-api/controllers/wishlist"
	"go-wishlist-api/drivers/mysql"
	"go-wishlist-api/repositories"
	"go-wishlist-api/routes"
	uc_user "go-wishlist-api/usecases/user"
	uc_wishlist "go-wishlist-api/usecases/wishlist"
)

func main() {
	config.LoadEnv()
	config.InitConfigMySQL()
	db := mysql.ConnectDB(config.InitConfigMySQL())

	e := echo.New()

	wishlistRepo := repositories.NewWishlistRepo(db)
	wishlistUseCase := uc_wishlist.NewWishlistUseCase(wishlistRepo)
	wishlistController := cl_wishlist.NewWishlistController(wishlistUseCase)

	userRepo := repositories.NewUserRepo(db)
	userUseCase := uc_user.NewUserUseCase(userRepo)
	userController := cl_user.NewUserController(userUseCase)

	routes := routes.RouteController{
		WishlistController: wishlistController,
		UserController:     userController,
	}

	routes.InitRoute(e)
	e.Logger.Fatal(e.Start(":8080"))
}
