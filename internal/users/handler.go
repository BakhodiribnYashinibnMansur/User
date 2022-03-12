package users

import (
	"users/internal/handlers"

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
	router.GET(usersURL, h.GetList)
	router.GET(userURL, h.GetUserByID)
	router.POST(usersURL, h.CreateUser)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PartiallyUpdateUser)
	router.DELETE(userURL, h.DeleteUser)
}

func (h *handler) GetList(ctx *gin.Context) {

}
func (h *handler) GetUserByID(ctx *gin.Context) {

}
func (h *handler) CreateUser(ctx *gin.Context) {

}
func (h *handler) UpdateUser(ctx *gin.Context) {

}
func (h *handler) PartiallyUpdateUser(ctx *gin.Context) {

}
func (h *handler) DeleteUser(ctx *gin.Context) {

}
