package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"go_project/src/common/controllers"
	"go_project/src/users/models"
	"go_project/src/users/models/services"
	"net/http"
	"strconv"
)

type UserController struct {
	AppController controllers.AppController
	UserService services.UserServiceInterface
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

func (a *UserController) UserRegister(w http.ResponseWriter, r *http.Request) {
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

func (a *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
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
