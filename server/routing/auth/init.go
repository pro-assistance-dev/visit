package auth

import (
	"github.com/gin-gonic/gin"

	handler "portal/handlers/auth"

	baseH "github.com/pro-assistance-dev/sprob/handlers/auth"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.POST("/login", h.Login)
	r.POST("/register", h.Register)
	r.POST("/restore-password", h.RestorePassword)
	//
	r.PUT("/password-change", baseH.H.RefreshPassword)
	r.GET("/check-uuid/:id/:uuid", baseH.H.CheckUUID)
}
