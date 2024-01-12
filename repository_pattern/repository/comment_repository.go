package repository

import (
	"context"
	"golang-clean-architecture/repository_pattern/entity"
)

type CommentRepostitory interface {
	Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	FindById(ctx context.Context, id int32) (entity.Comment, error)
	FindAll(ctx context.Context) ([]entity.Comment, error)
}