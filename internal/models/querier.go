// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package models

import (
	"context"
)

type Querier interface {
	CreateNode(ctx context.Context, arg CreateNodeParams) error
	GetAll(ctx context.Context) ([]*Node, error)
	GetByBlockID(ctx context.Context, blockID string) (*Marker, error)
	GetLatestBlock(ctx context.Context, limit int32) ([]*Block, error)
	GetRandomNode(ctx context.Context, limit int32) ([]*Node, error)
	MarkBlock(ctx context.Context, blockID string) error
}

var _ Querier = (*Queries)(nil)