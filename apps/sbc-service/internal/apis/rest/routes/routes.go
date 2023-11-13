package routes

import (
	"github.com/go-chi/chi/v5"

	v1handlers "github.com/alwaysbespoke/coba/apps/sbc-service/internal/apis/rest/handlers/v1"
)

func New(v1Handlers *v1handlers.Handlers) *chi.Mux {
	router := chi.NewRouter()

	router.Get("v1/sbcs/{sbc-id}/activate", v1Handlers.ActivateSbc)
	router.Get("v1/sbcs/{sbc-id}/deactivate", v1Handlers.DeactivateSbc)
	router.Get("v1/sbcs", v1Handlers.ListSbcs)
	router.Delete("v1/sbcs/{sbc-id}", v1Handlers.DeleteSbc)
	router.Post("v1/sbcs", v1Handlers.CreateSbc)

	return router
}
