package repository

import (
	"context"
	"gin/types"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Repository interface {
	CreateQuestion(ctx context.Context, question *types.Question) error
	ReadQuestion(ctx context.Context) (*types.Question, error)
	UpdateQuestion(ctx context.Context, id uuid.UUID) error
	DeleteQuestion(ctx context.Context, id uuid.UUID) error
	CreateUser(ctx context.Context, user *types.User) error
	ReadUser(ctx context.Context, id uuid.UUID) (*types.User, error)
	UpdateUser(ctx context.Context, id uuid.UUID) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}
