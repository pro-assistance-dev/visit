package userseventsactivities

import (
	handler "portal/handlers/userseventssactivities"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("ping", h.Upsert)
	r.GET("users-activities-times-sums", h.UsersActivitiesTimesSums)
	r.GET("unique-users-by-hours", h.UniqueUsersByHours)
}
