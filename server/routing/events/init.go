package doctors

import (
	handler "portal/handlers/event"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h *handler.Handler, ws *gin.RouterGroup) {
	r.GET("", h.GetAll)
	r.POST("/ftsp", h.FTSP)
	r.GET("/by-slug/:slug", h.GetBySlug)
	r.GET("/:id", h.Get)
	// r.GET("/messages/:event-id", h.GetAllMessages)
	r.POST("", h.Create)
	r.DELETE("/:id", h.Delete)
	r.PUT("/:id", h.Update)
	// r.POST("/message", h.CreateMessage)
	// r := api.Group(path)
	ws.GET("/connect/:chatId/:userId", h.Connect)
}
