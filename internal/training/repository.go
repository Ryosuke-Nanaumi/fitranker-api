package training

import (
	"context"
	"database/sql"
	"time"
)

type Repository interface {
	GetUserById(ctx context.Context, id int64) (*User, error)
	GetPoint(ctx context.Context, id int64, date *time.Time) ([]PointRecord, error)
	// GetRanking sliceは参照型、配列は値
	GetRanking(ctx context.Context) ([]Ranking, error)
	PostTrainingRecords(ctx context.Context, in PostTrainingRecordsInput) (int64, error)
}

type repository struct {
	db *sql.DB
}

// NewRepository Repositoryを満たしたrepositoryを返す
// repositoryには、メソッドレシーバーにより実装が追加される
func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

// GetUserById この書き方はメソッドレシーバー
// repositoryにGetUserByIdを追加するという意味になる
func (r *repository) GetUserById(ctx context.Context, id int64) (*User, error) {
	return &User{ID: id, Name: "foo"}, nil
}

func (r *repository) GetPoint(ctx context.Context, id int64, date *time.Time) ([]PointRecord, error) {
	return []PointRecord{
		{Amount: 0, Point: 0},
	}, nil
}

func (r *repository) GetRanking(ctx context.Context) ([]Ranking, error) {
	return []Ranking{
		{ID: 1, Name: "Alice", Point: 120},
	}, nil
}

func (r *repository) PostTrainingRecords(ctx context.Context, in PostTrainingRecordsInput) (int64, error) {
	return 1, nil
}
