package training

import (
	"context"
	"time"
)

type Service interface {
	GetPersonalInfo(ctx context.Context, id int64) (*PersonalUser, error)
	GetRanking(ctx context.Context) ([]Ranking, error)
	PostTrainingRecords(ctx context.Context, in PostTrainingRecordsInput) (*PostTrainingRecordsResult, error)
}

// Repositoryインターフェースはそれ自体が参照型
type service struct {
	repository Repository
	now        func() time.Time
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
		now:        time.Now,
	}
}

func (s *service) GetPersonalInfo(ctx context.Context, id int64) (*PersonalUser, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) GetRanking(ctx context.Context) ([]Ranking, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) PostTrainingRecords(ctx context.Context, in PostTrainingRecordsInput) (*PostTrainingRecordsResult, error) {
	//TODO implement me
	panic("implement me")
}
