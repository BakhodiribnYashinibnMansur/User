package users

import (
	"fmt"
	"net/http"
	"users/internal/handlers"
	"users/internal/serviceerror"

	"github.com/gin-gonic/gin"
)

// var _ handlers.Handler = &handler{}

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

type handler struct{}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *gin.Engine) {
	router.GET(usersURL, serviceerror.Middleware(h.GetList))
	router.GET(userURL, serviceerror.Middleware(h.GetUserByID))
	router.POST(usersURL, serviceerror.Middleware(h.CreateUser))
	router.PUT(userURL, serviceerror.Middleware(h.UpdateUser))
	router.PATCH(userURL, serviceerror.Middleware(h.PartiallyUpdateUser))
	router.DELETE(userURL, serviceerror.Middleware(h.DeleteUser))
}

func (h *handler) GetList(ctx *gin.Context) *serviceerror.ErrorResponse {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": "Hello",
	})
	return serviceerror.ErrorNotFound
}
func (h *handler) GetUserByID(ctx *gin.Context) *serviceerror.ErrorResponse {
	return serviceerror.NewServiceError(nil, "test", "test", "t13")
}
func (h *handler) CreateUser(ctx *gin.Context) *serviceerror.ErrorResponse {
	err := fmt.Errorf("this is API error")
	return serviceerror.NewServiceError(err, "", "", "")
}
func (h *handler) UpdateUser(ctx *gin.Context) *serviceerror.ErrorResponse {
	return nil
}
func (h *handler) PartiallyUpdateUser(ctx *gin.Context) *serviceerror.ErrorResponse {
	return nil
}
func (h *handler) DeleteUser(ctx *gin.Context) *serviceerror.ErrorResponse {
	return nil
}
