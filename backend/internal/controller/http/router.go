package http

import (
	"APG6/internal/repository"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/ironstar-io/chizerolog"
	"github.com/rs/zerolog/log"
)

type handler struct {
	mux            *chi.Mux
	tableManager   repository.TableManager
	viewManager    repository.ViewManager
	dataManager    repository.DataManager
	functionManger repository.FunctionManager
}

func New(mux *chi.Mux, tableManager repository.TableManager, viewManager repository.ViewManager,
	dataManager repository.DataManager, functionManger repository.FunctionManager) *handler {
	return &handler{
		mux:            mux,
		tableManager:   tableManager,
		viewManager:    viewManager,
		dataManager:    dataManager,
		functionManger: functionManger,
	}
}

func (h *handler) Run() {
	h.mux.Use(chizerolog.LoggerMiddleware(&log.Logger))
	h.mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8080", "http://localhost:4500"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	}))

	h.mux.Route("/api/v1", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(Authorizer())

			r.Route("/table", func(r chi.Router) {
				r.Route("/personal_information", func(r chi.Router) {
					r.Post("/", h.AddPersonalInformation)
					r.Get("/", h.GetPersonalInformation)
					r.Patch("/", h.UpdatePersonalInformation)
					r.Delete("/{customer_id}", h.DeletePersonalInformation)
				})

				r.Route("/cards", func(r chi.Router) {
					r.Post("/", h.AddCard)
					r.Get("/", h.GetCard)
					r.Patch("/", h.UpdateCard)
					r.Delete("/{customer_card_id}", h.DeleteCard)
				})

				r.Route("/sku_group", func(r chi.Router) {
					r.Post("/", h.AddSkuGroup)
					r.Get("/", h.GetSkuGroup)
					r.Patch("/", h.UpdateSkuGroup)
					r.Delete("/{group_id}", h.DeleteSkuGroup)
				})

				r.Route("/product_grid", func(r chi.Router) {
					r.Post("/", h.AddProductGrid)
					r.Get("/", h.GetProductGrid)
					r.Patch("/", h.UpdateProductGrid)
					r.Delete("/{sku_id}", h.DeleteProductGrid)
				})

				r.Route("/stores", func(r chi.Router) {
					r.Post("/", h.AddStore)
					r.Get("/", h.GetStore)
					r.Patch("/", h.UpdateStore)
					r.Delete("/{transaction_store_id}/{sku_id}", h.DeleteStore)
				})

				r.Route("/transactions", func(r chi.Router) {
					r.Post("/", h.AddTransaction)
					r.Get("/", h.GetTransaction)
					r.Patch("/", h.UpdateTransaction)
					r.Delete("/{transaction_id}", h.DeleteTransaction)
				})

				r.Route("/checks", func(r chi.Router) {
					r.Post("/", h.AddCheck)
					r.Get("/", h.GetCheck)
					r.Patch("/", h.UpdateCheck)
					r.Delete("/{transaction_id}/{sku_id}", h.DeleteCheck)
				})

				r.Route("/date_of_analysing_formation", func(r chi.Router) {
					r.Get("/", h.GetDateOfAnalysingFormation)
					r.Patch("/", h.UpdateDateOfAnalysingFormation)
				})
			})

			r.Route("/view", func(r chi.Router) {
				r.Route("/purchase_history", func(r chi.Router) {
					r.Get("/", h.GetPurchaseHistory)
				})

				r.Route("/periods", func(r chi.Router) {
					r.Get("/", h.GetPeriods)
				})

				r.Route("/groups", func(r chi.Router) {
					r.Get("/", h.GetGroups)
				})

				r.Route("/customers", func(r chi.Router) {
					r.Get("/", h.GetCustomers)
				})
			})

			r.Post("/import/{table}", h.Import)
			r.Get("/export/{table}", h.Export)

			r.Route("/function", func(r chi.Router) {
				r.Get("/growth_of_average_check", h.FncGrowthOfAverageCheck)
				r.Get("/defining_offer_increasing_frequency_visits", h.FncDefiningOfferIncreasingFrequencyVisits)
				r.Get("/defining_offer_increasing_margin", h.FncDefiningOfferIncreasingMargin)
			})
		})

		r.Route("/auth", func(r chi.Router) {
			r.Post("/", h.Authorization)
			r.Post("/refresh", h.Refresh)
		})
	})
}
