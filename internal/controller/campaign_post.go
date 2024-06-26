package controller

import (
	"emailn/internal/controller/dto"
	"emailn/internal/controller/utils"
	"net/http"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, req *http.Request) (interface{}, int, error) {
	var campaignDto dto.NewPostCampaignDto
	_, _ = utils.ParseJSONRequest(req, &campaignDto)
	email := req.Context().Value("email").(string)
	campaignDto.CreatedBy = email
	id, err := h.CampaignService.CreateCampaign(req.Context(), campaignDto)
	return map[string]int{"id": id}, http.StatusCreated, err
}
