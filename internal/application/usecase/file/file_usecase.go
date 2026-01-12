package usecase

import (
	"context"

	"github.com/InstaySystem/is_v2-be/internal/application/dto"
)

type FileUseCase interface {
	CreateUploadURLs(ctx context.Context, req dto.UploadPresignedURLsRequest) ([]*dto.UploadPresignedURLResponse, error)

	CreateViewURLs(ctx context.Context, req dto.ViewPresignedURLsRequest) ([]*dto.ViewPresignedURLResponse, error)
}
