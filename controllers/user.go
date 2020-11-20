package controllers

import (
	"go-template/models"
	"go-template/services"
	"go-template/utils"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/gofiber/fiber/v2"
)

// UserController struct
type UserController struct {
	Base
}

// CreateUserRequest params to create user
type CreateUserRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// UserResponseSuccess swagger response
type UserResponseSuccess struct {
	HTTPResponse
	Data models.User `json:"data"`
}

// UserResponseError swagger response
type UserResponseError struct {
	HTTPResponse
	Data string `json:",omitempty"`
}

// ValidateCreateRequest method
func (p CreateUserRequest) ValidateCreateRequest() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required, validation.Length(5, 20)),
		validation.Field(&p.Password, validation.Required, validation.Length(5, 50)),
		validation.Field(&p.Email, validation.Required, validation.Length(5, 50), is.Email),
	)
}

var service = (*services.UserService)(nil)

// GetByID godoc
// @Summary Get detail user by ID
// @Description Get detail user by ID
// @Tags User
// @Accept  json
// @Produce  json
// @Param ID path int true "User ID"
// @Success 200 {object} UserResponseSuccess
// @Success 400 {object} UserResponseError
// @Success 419 {string} string "Too many request"
// @Router /api/user/{ID} [get]
func (u *UserController) GetByID(c *fiber.Ctx) error {
	ID, err := strconv.Atoi(c.Params("ID"))
	if err != nil {
		return u.BadRequest(c, err.Error())
	}
	usr, err := service.GetByID(ID)
	if err != nil {
		return u.BadRequest(c, err.Error())
	}
	return u.Success(c, usr)
}

// Create godoc
// @Summary Create an user
// @Description create an user
// @Tags User
// @Accept  json
// @Produce  json
// @Param user body CreateUserRequest true "Add user"
// @Success 200 {object} UserResponseSuccess
// @Success 400 {object} UserResponseError
// @Success 419 {string} string "Too many request"
// @Router /api/user/create [post]
func (u *UserController) Create(c *fiber.Ctx) error {
	params := CreateUserRequest{}
	if err := c.BodyParser(&params); err != nil {
		return u.BadRequest(c, err.Error())
	}
	if err := params.ValidateCreateRequest(); err != nil {
		return u.BadRequest(c, err.Error())
	}
	pwd, err := utils.HashPassword(params.Password)
	if err != nil {
		return u.BadRequest(c, err.Error())
	}
	usr, err := service.Create(params.Name, params.Email, pwd)
	if err != nil {
		return u.BadRequest(c, err.Error())
	}
	return u.Success(c, usr)
}
