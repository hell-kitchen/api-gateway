package http

import (
	"context"
	"github.com/google/uuid"
	"github.com/hell-kitchen/api-gateway/internal/config"
	"github.com/hell-kitchen/api-gateway/internal/controller/http/mw"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

// Server is HTTP server struct. Includes logger, router, service, etc.
type Server struct {
	router *echo.Echo
	config *config.Server
	log    *zap.Logger
}

// NewServer return new HTTP Server.
func NewServer(config *config.Server, logger *zap.Logger) (*Server, error) {
	logger = logger.With(
		zap.String("layer", "http"),
		zap.String("running_at", config.Addr()),
	)
	logger.Debug("initializing new server")

	s := &Server{
		router: echo.New(),
		config: config,
		log:    logger,
	}

	s.configureRouter()

	logger.Info("server initialized successfully")

	return s, nil
}

func (srv *Server) OnStart(_ context.Context) error {
	srv.log.Debug("OnStart hook called")
	go func() {
		err := srv.router.Start(srv.config.Addr())
		if err != nil {
			srv.log.Error("unknown error while starting http server", zap.Error(err))
		}
	}()
	return nil
}

func (srv *Server) OnStop(ctx context.Context) error {
	srv.log.Debug("OnStop hook called")
	err := srv.router.Shutdown(ctx)
	srv.log.Info("stopped HTTP server", zap.Error(err))
	return err
}

func (srv *Server) configureRouter() {
	srv.router.HideBanner = true
	srv.configureMiddlewares()
	srv.configureRoutes()
}

func (srv *Server) configureMiddlewares() {
	srv.router.Use(
		mw.LogMiddleware(srv.log),
		middleware.CORSWithConfig(middleware.DefaultCORSConfig),
		middleware.CSRFWithConfig(middleware.DefaultCSRFConfig),
		middleware.RequestIDWithConfig(middleware.RequestIDConfig{
			Skipper:      middleware.DefaultSkipper,
			Generator:    uuid.NewString,
			TargetHeader: echo.HeaderXRequestID,
		}),
		middleware.Recover(),
	)
}

func (srv *Server) configureRoutes() {
	apiRouter := srv.router.Group("/api")
	usersRouter := apiRouter.Group("/users")
	{
		// Users subgroup
		usersRouter.GET("/", srv.UsersGetAll)
		usersRouter.POST("/", srv.UsersCreate)
		usersRouter.GET("/me", srv.UsersGetMe)
		usersRouter.GET("/:id", srv.UsersGetByID)
		usersRouter.POST("/set_password", srv.UsersSetPassword)

		// Subscriptions subgroup
		usersRouter.GET("/subscriptions", srv.UsersGetSubscriptions)
		usersRouter.POST("/:id/subscribe", srv.UsersSubscribe)
		usersRouter.POST("/:id/subscribe", srv.UsersUnsubscribe)
	}
	tokensRouter := apiRouter.Group("/auth/token")
	{
		tokensRouter.POST("/login", srv.TokensLogin)
		tokensRouter.POST("/logout", srv.TokensLogin)
	}
	tagsRouter := apiRouter.Group("/tags")
	{
		tagsRouter.GET("/", srv.TagsGetAll)
		tagsRouter.GET("/:id", srv.TagsGetOne)
	}
	recipesRouter := apiRouter.Group("/recipes")
	{
		recipesRouter.GET("/", srv.RecipesGetAll)
		recipesRouter.POST("/", srv.RecipesCreate)
		recipesRouter.GET("/:id", srv.RecipesGetByID)
		recipesRouter.PATCH("/:id", srv.RecipesUpdateByID)
		recipesRouter.DELETE("/:id", srv.RecipesDeleteByID)

		recipesRouter.GET("/download_shopping_cart", srv.RecipesDownloadShoppingCart)
		recipesRouter.POST("/:id/shopping_cart", srv.RecipesAddRecipeToShoppingCart)
		recipesRouter.DELETE("/:id/shopping_cart", srv.RecipesDeleteRecipeFromShoppingCart)

		recipesRouter.POST("/:id/favorite", srv.RecipesAddRecipeToFavorite)
		recipesRouter.DELETE("/:id/favorite", srv.RecipesDeleteRecipeFromFavorite)
	}

	ingredientsRouter := apiRouter.Group("/ingredients")
	{
		ingredientsRouter.GET("/", srv.IngredientsGetAll)
		ingredientsRouter.GET("/:id", srv.IngredientsGetByID)
	}
}
