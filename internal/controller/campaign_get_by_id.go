package controller

import (
	"emailn/internal/internalerrors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *Handler) CampaignGetById(w http.ResponseWriter, req *http.Request) (interface{}, int, error) {
	id := req.URL.Query().Get("id")
	x := chi.URLParam(req, "id")
	fmt.Println(x)
	campaign, err := h.CampaignService.GetBy(id)
	if campaign == nil {
		return nil, http.StatusNoContent, internalerrors.ErrNoContent
	}
	return campaign, http.StatusOK, err
}
