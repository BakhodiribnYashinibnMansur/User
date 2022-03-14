package serviceerror

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type serviceHandler func(ctx *gin.Context) *ErrorResponse

func Middleware(sererr serviceHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var errorService *ErrorResponse
		err := sererr(ctx)
		if err != nil {
			if errors.As(err.Error, &errorService) {
				if errors.Is(err.Error, ErrorNotFound.Error) {
					ctx.JSON(http.StatusNotFound, gin.H{
						"error": ErrorNotFound,
					})
					return
				}
				// err = sererr.(*ErrorResponse)
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": errorService,
				})
				return
			}
			ctx.JSON(http.StatusTeapot, gin.H{
				"error": systemError,
			})
		}
	}
}
