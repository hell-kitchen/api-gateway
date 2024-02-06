package http

import (
	"github.com/hell-kitchen/api-gateway/internal/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (srv *Server) IngredientsGetAll(ctx echo.Context) error {
	var req model.IngredientsGetAllRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, req)
}

func (srv *Server) IngredientsGetByID(ctx echo.Context) error {
	var req model.IngredientsGetByIDRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, req)
}

func (srv *Server) RecipesGetAll(ctx echo.Context) error {
	var req model.RecipesGetAllRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, req)
}

func (srv *Server) RecipesCreate(ctx echo.Context) error {
	var req model.RecipesCreateRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, req)
}

func (srv *Server) RecipesGetByID(ctx echo.Context) error {
	var req model.RecipesGetByIDRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, req)
}

func (srv *Server) RecipesUpdateByID(ctx echo.Context) error {
	var req model.RecipesUpdateByIDRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, req)
}

func (srv *Server) RecipesDeleteByID(ctx echo.Context) error {
	var req model.RecipesDeleteByIDRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, req)
}

func (srv *Server) RecipesDownloadShoppingCart(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, struct{}{})
}

func (srv *Server) RecipesAddRecipeToShoppingCart(ctx echo.Context) error {
	var req model.RecipesAddRecipeToShoppingCartRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, req)
}

func (srv *Server) RecipesRemoveRecipeFromShoppingCart(ctx echo.Context) error {
	var req model.RecipesRemoveRecipeFromShoppingCartRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, req)
}

func (srv *Server) RecipesAddRecipeToFavorite(ctx echo.Context) error {
	var req model.RecipesAddRecipeToFavoriteRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, req)
}

func (srv *Server) RecipesRemoveRecipeFromFavorite(ctx echo.Context) error {
	var req model.RecipesRemoveRecipeFromFavoriteRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, req)
}

func (srv *Server) TagsGetAll(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, []struct{}{})
}

func (srv *Server) TagsGetByID(ctx echo.Context) error {
	var req model.TagsGetByIDRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, req)
}

func (srv *Server) TokensLogin(ctx echo.Context) error {
	var req model.TokensLoginRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, req)
}

func (srv *Server) TokensLogout(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, struct{}{})
}

func (srv *Server) UsersGetAll(ctx echo.Context) error {
	var req model.UsersGetAllRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, req)
}

func (srv *Server) UsersCreate(ctx echo.Context) error {
	var req model.UsersCreateRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, req)
}

func (srv *Server) UsersGetByID(ctx echo.Context) error {
	var req model.UsersGetByIDRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, req)
}

func (srv *Server) UsersGetMe(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, struct{}{})
}

func (srv *Server) UsersSetPassword(ctx echo.Context) error {
	var req model.UsersSetPasswordRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, req)
}

func (srv *Server) UsersGetSubscriptions(ctx echo.Context) error {
	var req model.UsersGetSubscriptionsRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, req)
}

func (srv *Server) UsersSubscribe(ctx echo.Context) error {
	var req model.UsersSubscribeRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, req)
}

func (srv *Server) UsersUnsubscribe(ctx echo.Context) error {
	var req model.UsersUnsubscribeRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, req)
}
