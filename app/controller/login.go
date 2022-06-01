package controller

import (
	"encoding/json"
	"net/http"
	"project/freelance/danspro/simpleProject/app/models"
	"project/freelance/danspro/simpleProject/app/services"
	"project/freelance/danspro/simpleProject/shared/response"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var (
		err          error
		loginRequest models.LoginRequest
	)

	//decode request body
	if err = json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	//call service
	resp, err := services.LoginService(loginRequest)
	if err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response.SendSuccessResponse(w, resp, http.StatusText(http.StatusOK))
}
