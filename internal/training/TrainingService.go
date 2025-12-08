package training

import "context"

type PersonalUser struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	// TODO(追加必要であれば追加する)
}

type RankingUser struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Score int    `json:"score"`
}

type PostTrainingRecordsInput struct {
	ExerciseID int64   `json:"exerciseId"`
	Date       string  `json:"date"`
	Amount     float64 `json:"amount"`
	ID         int64   `json:"id"`
}

type PostTrainingRecordsResult struct {
	CreatedID int64 `json:"createdId"`
}

type Service interface {
	GetPersonalInfo(ctx context.Context, id int64) (*PersonalUser, error)
	GetRanking(ctx context.Context) ([]RankingUser, error)
	PostTrainingRecords(ctx context.Context, in PostTrainingRecordsInput) (*PostTrainingRecordsResult, error)
}
