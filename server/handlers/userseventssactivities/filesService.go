package userseventssactivities

import (
	"portal/models"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

func (s *FilesService) Upload(_ *gin.Context, _ *models.UserEventActivity, _ map[string][]*multipart.FileHeader) (err error) {
	//for i, file := range files {
	//	err = s.uploader.Upload(c, file, item.SetFilePath(&i))
	//	if err != nil {
	//		return err
	//	}
	//}
	return nil
}
