package handlers

import (
	"crud_server/internal/user/application/interfaces/input"
	"crud_server/internal/user/domain"
	"crud_server/internal/user/insfrastructure/dtos"
	mappers_input "crud_server/internal/user/insfrastructure/mappers/input"
	mappers_output "crud_server/internal/user/insfrastructure/mappers/output"
	"crud_server/internal/user/insfrastructure/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	UserService input.UserService
}

func NewUserHandler(service input.UserService) *UserHandler {
	return &UserHandler{UserService: service}
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users := h.UserService.List()

	if len(users) == 0 {
		utils.JSON(w, http.StatusOK, true, "No users found", []*domain.User{})
		return
	}

	usersResult := mappers_output.ToUserResponses(users)

	utils.JSON(w, http.StatusOK, true, "Users listed successfully", usersResult)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, false, "Invalid user ID", nil)
		return
	}

	user := h.UserService.Read(id)
	if user == nil {
		utils.JSON(w, http.StatusOK, false, "User not found", nil)
		return
	}

	usersResult := mappers_output.ToUserResponse(user) 

	utils.JSON(w, http.StatusOK, true, "User found successfully", usersResult)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userRequest dtos.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		utils.JSON(w, http.StatusBadRequest, false, "Invalid request body", nil)
		return
	}

	if err := userRequest.Validate(); err != nil {
		utils.JSON(w, http.StatusBadRequest, false, err.Error(), nil)
		return
	}

	user := mappers_input.ToUserDomain(&userRequest)

	userCreate, err := h.UserService.Create(user)

	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "Failed to create user", nil)
		return
	}

	usersResult := mappers_output.ToUserResponse(userCreate) 

	utils.JSON(w, http.StatusCreated, true, "User created successfully", usersResult)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, false, "Invalid user ID", nil)
		return
	}

	var userRequest dtos.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		utils.JSON(w, http.StatusBadRequest, false, "Invalid Request body", nil)
		return
	}

	if err := userRequest.Validate(); err != nil {
		utils.JSON(w, http.StatusBadRequest, false, "Invalid Request body", nil)
		return
	}

	user := mappers_input.ToUserDomain(&userRequest)
	user.Id = id
	userUpdated, err := h.UserService.Update(id, user)

	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "Failed to update user", nil)
		return
	}

	usersResult := mappers_output.ToUserResponse(userUpdated)

	utils.JSON(w, http.StatusOK, true, "User updated successfully", usersResult)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, false, "Invalid user ID", nil)
		return
	}

	if err := h.UserService.Delete(id); err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "Failed to delete user", nil)
		return
	}

	utils.JSON(w, http.StatusOK, true, "User deleted successfully", nil)
}
