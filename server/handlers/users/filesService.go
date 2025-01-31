package users

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/pro-assistance-dev/sprob/handlers/basehandler"
)

func (s *FilesService) Upload(_ *gin.Context, _ basehandler.Filer, _ map[string][]*multipart.FileHeader) (err error) {
	//for i, file := range files {
	//  err := s.helper.Uploader.Upload(c, file, item.SetFilePath(&i))
	//  if err != nil {
	//    return err
	//  }
	//}
	return nil
}
