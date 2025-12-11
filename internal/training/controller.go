package training

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Controller これはDIしないため直接structを使用する
type Controller struct {
	service Service
}

func NewTrainingController(s Service) *Controller {
	return &Controller{service: s}
}

// GetPersonalInfo (c *TrainingController)のように書くことで、TrainingControllerの実装みたいな感じになる
func (c *Controller) GetPersonalInfo(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	person, err := c.service.GetPersonalInfo(r.Context(), id)
	if err != nil {
		http.Error(w, "failed to get personal info", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, person)
}

func (c *Controller) GetRanking(w http.ResponseWriter, r *http.Request) {
	ranking, err := c.service.GetRanking(r.Context())
	if err != nil {
		http.Error(w, "failed to get ranking", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, ranking)
}

func (c *Controller) PostTrainingRecords(w http.ResponseWriter, r *http.Request) {
	var body PostTrainingRecordsInput
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	result, err := c.service.PostTrainingRecords(r.Context(), PostTrainingRecordsInput{
		ExerciseID: body.ExerciseID,
		Date:       body.Date,
		Amount:     body.Amount,
		ID:         body.ID,
	})
	if err != nil {
		http.Error(w, "failed to create training records", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, result)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
