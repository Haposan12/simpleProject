package controller

import (
	"github.com/go-chi/chi"
	"net/http"
	"project/freelance/danspro/simpleProject/app/constants"
	"project/freelance/danspro/simpleProject/app/models"
	"project/freelance/danspro/simpleProject/app/services"
	"project/freelance/danspro/simpleProject/shared/response"
)

func GetJobList(w http.ResponseWriter, r *http.Request) {
	var (
		err error
	)

	descParam := r.URL.Query().Get(constants.DescriptionSearchParam)
	locParam := r.URL.Query().Get(constants.LocationSearchParam)
	fullTimeParam := r.URL.Query().Get(constants.FullTimeSearchParam)
	pageParam := r.URL.Query().Get(constants.PageParam)

	resp, err := services.GetJobList(models.JobListRequestParam{
		Description: descParam,
		Location:    locParam,
		FullTime:    fullTimeParam,
		Page:        pageParam,
	})

	if err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response.SendSuccessResponse(w, resp, http.StatusText(http.StatusOK))
}

func GetJobDetail(w http.ResponseWriter, r *http.Request) {
	var (
		err error
	)

	positionId := chi.URLParam(r, "id")

	resp, err := services.GetJobDetail(positionId)
	if err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response.SendSuccessResponse(w, resp, http.StatusText(http.StatusOK))
}
