package wishlist

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"go-wishlist-api/controllers/base"
	"go-wishlist-api/controllers/wishlist/request"
	"go-wishlist-api/controllers/wishlist/response"
	"go-wishlist-api/entities"
	"go-wishlist-api/utils"
)

type WishlistController struct {
	wishlistUseCase entities.IWishlistUseCase
}

func NewWishlistController(wishlistUseCase entities.IWishlistUseCase) *WishlistController {
	return &WishlistController{
		wishlistUseCase: wishlistUseCase,
	}
}

func (w *WishlistController) GetAll(c echo.Context) error {
	wishlists, err := w.wishlistUseCase.GetAll()
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	wishlistResponses := response.FromUseCases(wishlists)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get All Wishlists", wishlistResponses))
}

func (w *WishlistController) Create(c echo.Context) error {
	var wishlistCreate request.WishlistCreate
	c.Bind(&wishlistCreate)

	wishlist, err := w.wishlistUseCase.Create(wishlistCreate.ToEntities())

	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	wishlistResponse := response.FromUseCase(wishlist)

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Create Wishlist", wishlistResponse))
}
