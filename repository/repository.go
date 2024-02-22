package repository

import (
	"context"
	"gin/types"

	_ "github.com/lib/pq"
)

type Repository interface {
	CreateQuestion(ctx context.Context, question *types.Question) error
	ReadQuestion(ctx context.Context, id *int) error
	UpdateQuestion(ctx context.Context, id *int, dataToChange *types.Question) error
	DeleteQuestion(ctx context.Context, id *int) error
}
