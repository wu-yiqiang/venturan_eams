package controllers

import (
	"path/filepath"
	"venturan/app/common/response"
	"venturan/global"
	"venturan/global/serviceErrors"
	"venturan/utils"

	"github.com/gin-gonic/gin"
)

// @Summary 文件上传
// @Description 文件上传
// @Tags 公共服务
// @ID /general/upload
// @Accept  json
// @Produce  json
// @Router /general/upload [post]
func FileUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.Fail(c, serviceErrors.FileIsNotEmpty)
		return
	}
	if file.Size > global.App.Config.App.UploadFileMaxSize {
		response.Fail(c, serviceErrors.FileIsTooLarge)
		return
	}
	newFilename := utils.GenerateUniqueFilename(file.Filename)
	savePath := filepath.Join(global.App.Config.App.StorageDir, newFilename)
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		response.Fail(c, serviceErrors.FileSaveFailed)
		return
	}
	response.Success(c, newFilename)
}
