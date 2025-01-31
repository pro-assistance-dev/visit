package doctors

import (
	handler "portal/handlers/m2m"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.POST("/create", h.Create)
	r.POST("/delete", h.Delete)
	r.POST("/update", h.Update)
}
