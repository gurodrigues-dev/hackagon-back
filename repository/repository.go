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
	ReadUser(ctx context.Context, id *int) (*types.User, error)
	UpdateUser(ctx context.Context, id *int) error
	DeleteUser(ctx context.Context, nickname *string) error
	VerifyLogin(ctx context.Context, user *types.User) error
	CreateAnswer(ctx context.Context, answer *types.Answer) error
	DeleteAnswer(ctx context.Context, id uuid.UUID) error
	VerifyAnswer(ctx context.Context, question *types.Question, nickname *string) (*types.Answer, error)
	// GetRank(ctx context.Context, nickname *string) error
	// IncreaseScore(ctx context.Context, nickname *string) error
}
