package place

import (
	handler "portal/handlers/speakers"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("", h.GetAll)
	r.GET("/:id", h.Get)
	r.POST("", h.Create)
	r.DELETE("/:id", h.Delete)
	r.PUT("/:id", h.Update)
	r.POST("/ftsp", h.FTSP)
}
