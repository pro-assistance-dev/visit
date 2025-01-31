package userseventssactivities

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pro-assistance-dev/sprob/middleware"
)

func (h *Handler) Upsert(c *gin.Context) {
	userID, err := h.helper.Token.ExtractTokenMetadata(c.Request, middleware.ClaimUserID)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.service.Upsert(uuid.MustParse(userID))
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (h *Handler) UsersActivitiesTimesSums(c *gin.Context) {
	items, err := h.service.UsersActivitiesTimesSums()
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) UniqueUsersByHours(c *gin.Context) {
	items, err := h.service.UniqueUsersByHours()
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, items)
}
