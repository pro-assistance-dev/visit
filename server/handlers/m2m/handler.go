package m2m

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pro-assistance-dev/sprob/helper"
)

// helper.DB.DB.RegisterModel((*models.EventToSponsor)(nil))
// helper.DB.DB.RegisterModel((*models.EventToPlace)(nil))
// helper.DB.DB.RegisterModel((*models.EventToPerfom)(nil))
// helper.DB.DB.RegisterModel((*models.PerfomToSpeaker)(nil))
// helper.DB.DB.RegisterModel((*models.SessionToSpeaker)(nil))
// helper.DB.DB.RegisterModel((*models.ScheduleToPerfom)(nil))
func fillCols(h helper.Helper, item *M2M) {
	// // schema := h.Project.GetSchema(item.Model)
	// item.ID1.Col = schema.GetCol(item.ID1.Field)
	// item.ID2.Col = schema.GetCol(item.ID2.Field)
	// item.TableName = schema.GetTableName()
	// for f := range item.Fields {
	// 	item.Fields[f].Col = schema.GetCol(item.Fields[f].Field)
	// }
}

func (h *Handler) Create(c *gin.Context) {
	var item M2M
	_, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	fillCols(*h.helper, &item)
	err = S.Create(c.Request.Context(), &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Update(c *gin.Context) {
	var item M2M
	_, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	fillCols(*h.helper, &item)
	err = S.Update(c.Request.Context(), &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Delete(c *gin.Context) {
	var item M2M
	_, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	fillCols(*h.helper, &item)
	err = S.Delete(c.Request.Context(), &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
