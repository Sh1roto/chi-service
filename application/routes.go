package application

import (
	"chi-service/handler"
	"chi-service/repository/order"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (a *App) loadRoutes() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Route("/orders", a.loadOrderRoutes)

	a.router = router
}

func (a *App) loadOrderRoutes(router chi.Router) { //	subrouter
	orderHandler := &handler.Order{
		Repo: &order.RedisRepo{
			Client: a.rdb,
		},
	}

	router.Get("/", orderHandler.List)
	router.Post("/", orderHandler.Create)
	router.Get("/{id}", orderHandler.GetByID)
	router.Put("/{id}", orderHandler.UpdateByID)
	router.Delete("/{id}", orderHandler.DeleteByID)
}
