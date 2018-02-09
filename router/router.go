package router

import (
	"net/http"
	"os"

	"github.com/at-vudang95/go-food-market-api/infrastructure"
	"github.com/at-vudang95/go-food-market-api/shared/handler"
	mMiddleware "github.com/at-vudang95/go-food-market-api/shared/middleware"
	"github.com/at-vudang95/go-food-market-api/shared/repository"
	"github.com/at-vudang95/go-food-market-api/shared/usecase"
	"github.com/at-vudang95/go-food-market-api/user"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Router is application struct hold Mux and db connection
type Router struct {
	Mux                *chi.Mux
	SQLHandler         *infrastructure.SQL
	S3Handler          *infrastructure.S3
	CacheHandler       *infrastructure.Cache
	LoggerHandler      *infrastructure.Logger
	TranslationHandler *infrastructure.Translation
	SearchAPIHandler   infrastructure.SearchAPI
}

// InitializeRouter initializes Mux and middleware
func (r *Router) InitializeRouter() {

	r.Mux.Use(middleware.RequestID)
	r.Mux.Use(middleware.RealIP)
	// Custom middleware(Translation)
	r.Mux.Use(r.TranslationHandler.Middleware.Middleware)
	// Custom middleware(Logger)
	r.Mux.Use(mMiddleware.Logger(r.LoggerHandler))
}

// SetupHandler set database and redis and usecase.
func (r *Router) SetupHandler() {
	// error handler set.
	eh := handler.NewHTTPErrorHandler(r.LoggerHandler.Log)
	r.Mux.NotFound(eh.StatusNotFound)
	r.Mux.MethodNotAllowed(eh.StatusMethodNotAllowed)

	// profiler
	env := os.Getenv("ENV_API")
	if env == "development" {
		r.Mux.Mount("/debug", middleware.Profiler())
	}

	r.Mux.Method(http.MethodGet, "/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// base set.
	bh := handler.NewBaseHTTPHandler(r.LoggerHandler.Log)
	// base set.
	br := repository.NewBaseRepository(r.LoggerHandler.Log)
	// base set.
	bu := usecase.NewBaseUsecase(r.LoggerHandler.Log)
	// user set.
	uh := user.NewHTTPHandler(bh, bu, br, r.SQLHandler, r.CacheHandler)
	r.Mux.Route("/v1", func(cr chi.Router) {
		cr.Get("/hello", uh.RegisterByDevice)
	})
}
