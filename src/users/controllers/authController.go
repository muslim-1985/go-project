package controllers

import (
	"encoding/json"
	"go_project/src/common/controllers"
	"go_project/src/users/models"
	"go_project/src/users/models/services"
	"net/http"
)

type AuthController struct {
	AppController controllers.AppController
	UserService services.AuthServiceInterface
}

func (a *AuthController) UserRegister(w http.ResponseWriter, r *http.Request) {
	var p models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		a.AppController.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	p.RoleId = 1
	defer r.Body.Close()

	if err := a.UserService.UserRegister(&p); err != nil {
		a.AppController.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	a.AppController.CreateToken(&p)
	a.AppController.RespondWithJSON(w, http.StatusCreated, p)
}

func (a *AuthController) LoginUser(w http.ResponseWriter, r *http.Request) {
	var p *models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		a.AppController.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := a.UserService.LoginUser(p); err != nil {
		a.AppController.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	a.AppController.CreateToken(p)
	a.AppController.RespondWithJSON(w, http.StatusCreated, p)
}