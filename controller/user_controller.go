package controller

import (
	"strconv"

	"github.com/1layar/clibank/app/responses"
	_users "github.com/1layar/clibank/app/users"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

type UserController struct {
	service *_users.UserService
}

func (ctrlUser *UserController) Init(service *_users.UserService) {
	ctrlUser.service = service
}

// GetUsers
// @Summary Get all users
// @Description show all users with wallet
// @Tags users
// @ID get-users
// @Accept  json
// @Produce  json
// @Success 200 {object} responses.UsersApiResponse	"ok"
// @Failure 500 {object} responses.ApiResponse "Server Error"
// @Router /users [get]
func (ctrl *UserController) GetUsers(c echo.Context) error {
	users, err := ctrl.service.GetUsers()

	if err != nil {
		return c.JSON(500, responses.ApiResponse{
			Message: "failed to get users",
			Data:    []_users.User{},
			Error:   err.Error(),
		})
	}

	usersResponse := lo.Map(users, func(t _users.User, i int) responses.UserResponse {
		data := responses.UserResponse{}
		data.FromModel(t)
		return data
	})

	return c.JSON(200, responses.UsersApiResponse{
		Message: "success",
		Data:    usersResponse,
	})
}

// GetUsers
// @Summary get user by it id
// @Description get string by ID
// @Tags users
// @ID get-user
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} responses.UserApiResponse
// @Failure 402 {object} responses.ApiResponse "Server Error"
// @Failure 404 {object} responses.ApiResponse "Server Error"
// @Router /users/{id} [get]
func (ctrl *UserController) GetUser(c echo.Context) error {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(402, responses.ApiResponse{
			Message: "invalid id",
			Data:    nil,
		})
	}
	user, err := ctrl.service.GetById(userId)

	if err != nil {
		return c.JSON(404, responses.ApiResponse{
			Message: "user not found",
			Data:    nil,
		})
	}
	data := &responses.UserResponse{}
	data.FromModel(*user)
	return c.JSON(200, responses.UserApiResponse{
		Message: "success",
		Data:    *data,
	})
}

// UpdateUser
// @Summary Update user By ID
// @Description Update string by ID
// @ID update-user
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param data body users.UpdateUserDto true "data"
// @Success 200 {object} responses.UserApiResponse
// @Failure 402 {object} responses.ApiResponse "Server Error"
// @Failure 404 {object} responses.ApiResponse "Server Error"
// @Router /users/{id} [put]
func (ctrl *UserController) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(402, responses.ApiResponse{
			Message: "invalid id",
			Data:    nil,
		})
	}

	userData := new(_users.UpdateUserDto)
	if err := c.Bind(userData); err != nil {
		return c.JSON(400, responses.ApiResponse{
			Message: "invalid body",
			Data:    nil,
		})
	}

	user, err := ctrl.service.GetById(userId)

	if err != nil {
		return c.JSON(404, responses.ApiResponse{
			Message: "user not found",
		})
	}
	updatedUser, err := ctrl.service.UpdateUser(user.PhoneNumber, user.Pin, *userData)

	if err != nil {
		return c.JSON(500, responses.ApiResponse{
			Message: "failed to update user",
		})
	}
	data := &responses.UserResponse{}
	data.FromModel(updatedUser)
	return c.JSON(200, responses.UserApiResponse{
		Message: "success",
		Data:    *data,
	})
}

// CreateUser
// @Summary Add a new user
// @Description add a new user
// @Tags users
// @ID create-users
// @Accept  json
// @Produce  json
// @Param data body _users.CreateUserDto true "data"
// @Success 200 {object} responses.UserApiResponse
// @Failure 400 {object} responses.ApiResponse "Not valid data"
// @Failure 500 {object} responses.ApiResponse "Server Error"
// @Router /users [post]
func (ctrl *UserController) CreateUser(c echo.Context) error {
	userData := new(_users.CreateUserDto)
	if err := c.Bind(userData); err != nil {
		return c.JSON(400, responses.ApiResponse{
			Message: "invalid body",
			Data:    nil,
		})
	}

	user, err := ctrl.service.CreateUser(*userData)

	if err != nil {
		return c.JSON(500, responses.ApiResponse{
			Message: "failed to create user",
		})
	}
	data := &responses.UserResponse{}
	data.FromModel(user)
	return c.JSON(200, responses.UserApiResponse{
		Message: "success",
		Data:    *data,
	})
}

// DeleteUser
// @Summary delete user by id
// @Description delete string by ID
// @Tags users
// @ID delete-users
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} responses.UsersApiResponse	"ok"
// @Failure 500 {object} responses.ApiResponse "Server Error"
// @Router /users/{id} [delete]
func (ctrl *UserController) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(402, responses.ApiResponse{
			Message: "invalid id",
			Data:    nil,
		})
	}

	user, err := ctrl.service.GetById(userId)

	if err != nil {
		return c.JSON(404, responses.ApiResponse{
			Message: "user not found",
		})
	}

	err = ctrl.service.DeleteUser(user.PhoneNumber)

	if err != nil {
		return c.JSON(500, responses.ApiResponse{
			Message: "failed to delete user",
		})
	}
	return c.JSON(204, responses.ApiResponse{
		Message: "success",
		Data:    nil,
	})
}
