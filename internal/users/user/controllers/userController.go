package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"go_project/internal/users/user/models"
	"go_project/internal/users/user/models/services"
	"go_project/pkg/common/controllers"
	"net/http"
	"strconv"
)

type UserController struct {
	AppController controllers.AppController
	UserService   services.UserServiceInterface
}

func (a *UserController) Home (w http.ResponseWriter, r *http.Request)  {
	c := map[string][]string{"message": {"how are you?", "redudant"}}
	a.AppController.RespondWithJSON(w, http.StatusOK, c)
}

func (a *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	users, err := a.UserService.GetUsers(start, count)
	if err != nil {
		a.AppController.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	a.AppController.RespondWithJSON(w, http.StatusOK, users)
}

func (a *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		a.AppController.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	p := models.User{ID: id}
	if err := a.UserService.GetUser(&p); err != nil {
		switch err {
		case sql.ErrNoRows:
			a.AppController.RespondWithError(w, http.StatusNotFound, "Product not found")
		default:
			a.AppController.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	a.AppController.RespondWithJSON(w, http.StatusOK, p)
}

func (a *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		a.AppController.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var p models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		a.AppController.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	p.ID = id

	if err := a.UserService.UpdateUser(&p); err != nil {
		a.AppController.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	a.AppController.RespondWithJSON(w, http.StatusOK, p)
}

func (a *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		a.AppController.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	p := models.User{ID: id}
	if err := a.UserService.DeleteUser(&p); err != nil {
		a.AppController.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	a.AppController.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
