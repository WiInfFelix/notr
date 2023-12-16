package handlers

import (
	"net/http"
	"strconv"

	"github.com/WiInfFelix/notr/models"
	"github.com/WiInfFelix/notr/services"
	"github.com/WiInfFelix/notr/utils"
)

type UserHandler struct {
	UserService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (h *UserHandler) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	itemsPerPage, _ := strconv.Atoi(r.URL.Query().Get("itemsPerPage"))

	users, err := h.UserService.GetAllUsers(page, itemsPerPage)

	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, users)
}

func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.NewUserInput

	err := utils.DecodeJSONBody(w, r, &user)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
	}

	createdUser, err := h.UserService.CreateUser(user)

	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, createdUser)
}
