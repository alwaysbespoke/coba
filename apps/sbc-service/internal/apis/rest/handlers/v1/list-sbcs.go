package handlers

import (
	"net/http"

	"github.com/alwaysbespoke/coba/apps/sbc-service/pkg/models"
	"github.com/alwaysbespoke/coba/pkg/api/responses"
)

// ListSbcs returns the current SBCs list
// todo: populate with logic. currently a stub
func (h *Handlers) ListSbcs(w http.ResponseWriter, r *http.Request) {
	sbcs := &models.ListSbcsResponse{}

	if err := responses.WriteJson(w, sbcs, http.StatusOK); err != nil {
		h.logger.Error("failed to write JSON: %w", err)
	}
}
