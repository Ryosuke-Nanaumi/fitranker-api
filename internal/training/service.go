package training

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

type Service interface {
	GetPersonalInfo(ctx context.Context, id int64) (*PersonalUser, error)
	GetRanking(ctx context.Context) ([]Ranking, error)
	GetTrainingRecords(ctx context.Context, id int64) ([]Record, error)
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
	user, err := s.repository.GetUserById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	now := s.now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	var (
		totalRecords  []PointRecord
		todaysRecords []PointRecord
	)

	grp, grpctx := errgroup.WithContext(ctx)

	grp.Go(func() error {
		var err error
		totalRecords, err = s.repository.GetPoint(grpctx, user.ID, nil)
		return err
	})
	grp.Go(func() error {
		var err error
		todaysRecords, err = s.repository.GetPoint(grpctx, user.ID, &today)
		return err
	})

	if err := grp.Wait(); err != nil {
		return nil, err
	}

	return &PersonalUser{
		user.ID,
		user.Name,
		calcPoint(totalRecords),
		calcPoint(todaysRecords),
	}, nil
}

func calcPoint(records []PointRecord) int64 {
	var total int64
	for _, r := range records {
		total += r.Amount * r.Point
	}
	return total
}

func (s *service) GetTrainingRecords(ctx context.Context, id int64) ([]Record, error) {
	return s.repository.GetTrainingRecords(ctx, id)
}

func (s *service) GetRanking(ctx context.Context) ([]Ranking, error) {
	return s.repository.GetRanking(ctx)
}

func (s *service) PostTrainingRecords(ctx context.Context, in PostTrainingRecordsInput) (*PostTrainingRecordsResult, error) {
	id, err := s.repository.PostTrainingRecords(ctx, in)
	if err != nil {
		return nil, err
	}

	return &PostTrainingRecordsResult{id}, nil
}
