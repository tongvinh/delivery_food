package uploadprovider

import (
	"context"
	"myapp/common"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
}
