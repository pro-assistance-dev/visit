package event

import (
	"net/http"
	"portal/models"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/uptrace/bun"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	var item models.Event
	_, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	// err := c.Bind(&item)
	// if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
	// 	return
	// }
	err = S.Create(c.Request.Context(), &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) FTSP(c *gin.Context) {
	data, err := S.GetAll(c.Request.Context())
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, models.FTSPAnswer{Data: data, FTSP: *h.helper.SQL.ExtractFTSP(c.Request.Context())})
}

func (h *Handler) GetAll(c *gin.Context) {
	items, err := S.GetAll(c.Request.Context())
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) Get(c *gin.Context) {
	item, err := S.Get(c.Request.Context(), c.Param("id"))
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Delete(c *gin.Context) {
	err := S.Delete(c.Request.Context(), c.Param("id"))
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) Update(c *gin.Context) {
	var item models.Event
	_, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = S.Update(c.Request.Context(), &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetBySlug(c *gin.Context) {
	item, err := S.GetBySlug(c.Request.Context(), c.Param("slug"))
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

type WithId struct {
	ID uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
}
type ChatMessage[UserT any] struct {
	bun.BaseModel `bun:"chat_messages,alias:chat_messages"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	User          UserT         `bun:"rel:belongs-to,join:user_id=id,scanonly,default:null" json:"user"`
	UserID        uuid.NullUUID `bun:"type:uuid" json:"userId"`
	ChatID        uuid.NullUUID `bun:"type:uuid" json:"chatId"`
	Message       string        `json:"message"`
	CreatedAt     time.Time     `bun:",nullzero,notnull" json:"createdAt"`
}

func (h *Handler) CreateMessage(c *gin.Context) {
	var item ChatMessage[WithId]

	_, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = S.CreateMessage(c.Request.Context(), &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetAllMessages(c *gin.Context) {
	items, err := S.GetAllMessages(c.Request.Context(), c.Param("event-id"))
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, items)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) Connect(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	chatMember := models.NewChatMember(conn, c.Param("userId"))
	chatMember.StartChat(c.Param("chatId"))
}
