package training

// 外部と接続されるDTOであれば、タグを記載する必要がある

type PersonalUser struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	TotalPoints int64  `json:"totalPoints"`
	TodaysPoint int64  `json:"todaysPoint"`
}

// Record あとでPostTrainingRecordsInputも統合したい
type Record struct {
	ExerciseID int64   `json:"exerciseId"`
	Date       string  `json:"date"`
	Amount     float64 `json:"amount"`
	Point      int64   `json:"point"`
	ID         int64   `json:"trainingId"`
}

type PostTrainingRecordsInput struct {
	ExerciseID int64   `json:"exerciseId"`
	Date       string  `json:"date"`
	Amount     float64 `json:"amount"`
	ID         int64   `json:"userId"`
}

type PostTrainingRecordsResult struct {
	CreatedID int64 `json:"createdId"`
}

type User struct {
	ID   int64
	Name string
}

type Ranking struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Point int64  `json:"point"`
}

type PointRecord struct {
	Amount int64
	Point  int64
}
