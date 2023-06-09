package repositories

import (
	"context"

	"github.com/sisu-network/interview/internal/models"
)

type BlockRepository interface {
	Create(ctx context.Context, block *models.Block) error
	GetAll(ctx context.Context) ([]*models.Block, error)
}
