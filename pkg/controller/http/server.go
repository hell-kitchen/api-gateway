package http

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/hell-kitchen/api-gateway/pkg/config"
	"github.com/hell-kitchen/api-gateway/pkg/controller/http/mw"
	"github.com/hell-kitchen/api-gateway/pkg/model"
	"github.com/hell-kitchen/api-gateway/pkg/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

// Server is HTTP server struct. Includes logger, router, service, etc.
type Server struct {
	router *echo.Echo
	config *config.Server
	log    *zap.Logger
	srv    service.Interface
}

// NewServer return new HTTP Server.
func NewServer(config *config.Server, logger *zap.Logger, srv service.Interface) (*Server, error) {
	logger = logger.With(
		zap.String("layer", "http"),
		zap.String("running_at", config.Addr()),
	)
	logger.Debug("initializing new server")

	s := &Server{
		router: echo.New(),
		config: config,
		log:    logger,
		srv:    srv,
	}

	s.configureRouter()

	logger.Info("server initialized successfully")

	return s, nil
}

// OnStart is function which starts http server.
//
// It must be called only once on starting of server life cycle.
//
// Could be used in fx Hooks.
func (srv *Server) OnStart(ctx context.Context) error {
	srv.log.Info("OnStart hook called")
	ctx, cancel := context.WithCancelCause(ctx)
	go func() {
		srv.log.Info("started server")
		err := srv.router.Start(srv.config.Addr())
		if err != nil {
			cancel(err)
			srv.log.Error("unknown error while starting http server", zap.Error(err))
		}
	}()
	time.Sleep(300 * time.Millisecond)
	return ctx.Err()
}

// OnStop stops server.
//
// Must be called only once.
//
// May be used in fx Hooks.
func (srv *Server) OnStop(ctx context.Context) error {
	srv.log.Info("OnStop hook called")
	err := srv.router.Shutdown(ctx)
	srv.log.Info("stopped HTTP server", zap.Error(err))
	return err
}

// configureRouter configures echo router.
//
// It calls Middleware and Routes configuration functions.
func (srv *Server) configureRouter() {
	srv.router.HideBanner = true
	srv.configureMiddlewares()
	srv.configureRoutes()
	srv.router.HTTPErrorHandler = srv.errorHandler
}

// errorHandler is custom http errors handler.
func (srv *Server) errorHandler(err error, c echo.Context) {
	resp := model.BaseErrorResponse{Detail: err.Error()}

	if err = c.JSON(http.StatusInternalServerError, resp); err != nil {
		srv.log.Error("got unexpected error while handling error", zap.Error(err))
	}
}

// configureMiddlewares configures Middlewares.
//
// Adds CORS, RequestID, Recover middlewares.
func (srv *Server) configureMiddlewares() {
	srv.router.Use(
		mw.LogMiddleware(srv.log),
		middleware.CORSWithConfig(middleware.DefaultCORSConfig),
		middleware.RequestIDWithConfig(middleware.RequestIDConfig{
			Skipper:      middleware.DefaultSkipper,
			Generator:    uuid.NewString,
			TargetHeader: echo.HeaderXRequestID,
		}),
		middleware.Gzip(),
		middleware.Recover(),
	)
}

func (srv *Server) configureRoutes() {
	apiRouter := srv.router.Group("/api")
	usersRouter := apiRouter.Group("/users")
	{
		// Users subgroup
		usersRouter.GET("", srv.UsersGetAll)
		usersRouter.GET("/", srv.UsersGetAll)
		usersRouter.POST("", srv.UsersCreate)
		usersRouter.POST("/", srv.UsersCreate)
		usersRouter.GET("/me", srv.UsersGetMe)
		usersRouter.GET("/me/", srv.UsersGetMe)
		usersRouter.POST("/set_password", srv.UsersSetPassword)
		usersRouter.POST("/set_password/", srv.UsersSetPassword)

		// Subscriptions subgroup
		usersRouter.GET("/subscriptions", srv.UsersGetSubscriptions)
		usersRouter.GET("/subscriptions/", srv.UsersGetSubscriptions)
		usersRouter.GET("/:id", srv.UsersGetByID)
		usersRouter.GET("/:id/", srv.UsersGetByID)
		usersRouter.POST("/:id/subscribe", srv.UsersSubscribe)
		usersRouter.POST("/:id/subscribe/", srv.UsersSubscribe)
		usersRouter.DELETE("/:id/subscribe", srv.UsersUnsubscribe)
		usersRouter.DELETE("/:id/subscribe/", srv.UsersUnsubscribe)
	}
	tokensRouter := apiRouter.Group("/auth/token")
	{
		tokensRouter.POST("/login", srv.TokensLogin)
		tokensRouter.POST("/login/", srv.TokensLogin)
		tokensRouter.POST("/logout", srv.TokensLogin)
		tokensRouter.POST("/logout/", srv.TokensLogin)
	}
	tagsRouter := apiRouter.Group("/tags")
	{
		tagsRouter.GET("", srv.TagsGetAll)
		tagsRouter.GET("/", srv.TagsGetAll)
		tagsRouter.GET("/:id", srv.TagsGetByID)
		tagsRouter.GET("/:id/", srv.TagsGetByID)
	}
	recipesRouter := apiRouter.Group("/recipes")
	{
		recipesRouter.GET("", srv.RecipesGetAll)
		recipesRouter.GET("/", srv.RecipesGetAll)
		recipesRouter.POST("", srv.RecipesCreate)
		recipesRouter.POST("/", srv.RecipesCreate)
		recipesRouter.GET("/:id", srv.RecipesGetByID)
		recipesRouter.GET("/:id/", srv.RecipesGetByID)
		recipesRouter.PATCH("/:id", srv.RecipesUpdateByID)
		recipesRouter.PATCH("/:id/", srv.RecipesUpdateByID)
		recipesRouter.DELETE("/:id", srv.RecipesDeleteByID)
		recipesRouter.DELETE("/:id/", srv.RecipesDeleteByID)

		recipesRouter.GET("/download_shopping_cart", srv.RecipesDownloadShoppingCart)
		recipesRouter.GET("/download_shopping_cart/", srv.RecipesDownloadShoppingCart)
		recipesRouter.POST("/:id/shopping_cart", srv.RecipesAddRecipeToShoppingCart)
		recipesRouter.POST("/:id/shopping_cart/", srv.RecipesAddRecipeToShoppingCart)
		recipesRouter.DELETE("/:id/shopping_cart", srv.RecipesRemoveRecipeFromShoppingCart)
		recipesRouter.DELETE("/:id/shopping_cart/", srv.RecipesRemoveRecipeFromShoppingCart)

		recipesRouter.POST("/:id/favorite", srv.RecipesAddRecipeToFavorite)
		recipesRouter.POST("/:id/favorite/", srv.RecipesAddRecipeToFavorite)
		recipesRouter.DELETE("/:id/favorite", srv.RecipesRemoveRecipeFromFavorite)
		recipesRouter.DELETE("/:id/favorite/", srv.RecipesRemoveRecipeFromFavorite)
	}

	ingredientsRouter := apiRouter.Group("/ingredients")
	{
		ingredientsRouter.GET("", srv.IngredientsGetAll)
		ingredientsRouter.GET("/", srv.IngredientsGetAll)
		ingredientsRouter.GET("/:id", srv.IngredientsGetByID)
		ingredientsRouter.GET("/:id/", srv.IngredientsGetByID)
	}
}
