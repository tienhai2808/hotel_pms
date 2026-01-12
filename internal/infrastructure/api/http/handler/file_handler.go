package handler

import (
	"context"
	"time"

	"github.com/InstaySystem/is_v2-be/internal/application/dto"
	fileUC "github.com/InstaySystem/is_v2-be/internal/application/usecase/file"
	"github.com/InstaySystem/is_v2-be/pkg/errors"
	"github.com/InstaySystem/is_v2-be/pkg/utils"
	"github.com/InstaySystem/is_v2-be/pkg/validator"
	"github.com/gin-gonic/gin"
)

type FileHandler struct {
	fileUC fileUC.FileUseCase
}

func NewFileHandler(fileUC fileUC.FileUseCase) *FileHandler {
	return &FileHandler{fileUC}
}

func (h *FileHandler) UploadPresignedURLs(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	var req dto.UploadPresignedURLsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		field, tag, param := validator.HandleRequestError(err)
		c.Error(errors.ErrBadRequest.WithData(gin.H{
			"field": field,
			"tag":   tag,
			"param": param,
		}))
		return
	}

	presignedURLs, err := h.fileUC.CreateUploadURLs(ctx, req)
	if err != nil {
		c.Error(err)
		return
	}

	utils.OKResponse(c, gin.H{
		"presigned_urls": presignedURLs,
	})
}

func (h *FileHandler) ViewPresignedURLs(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	var req dto.ViewPresignedURLsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		field, tag, param := validator.HandleRequestError(err)
		c.Error(errors.ErrBadRequest.WithData(gin.H{
			"field": field,
			"tag":   tag,
			"param": param,
		}))
		return
	}

	presignedURLs, err := h.fileUC.CreateViewURLs(ctx, req)
	if err != nil {
		c.Error(err)
		return
	}

	utils.OKResponse(c, gin.H{
		"presigned_urls": presignedURLs,
	})
}
