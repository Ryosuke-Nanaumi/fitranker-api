package training

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
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
// メソッドレシーバーにポインタを使うのは、構造体の中身を変更したいかどうか！
func (r *repository) GetUserById(ctx context.Context, id int64) (*User, error) {
	//return &User{ID: id, Name: "foo"}, nil
	const q = `SELECT id, name FROM users WHERE id = $1`
	var u User
	// 結果が1行のみなのでQueryRowContext
	// close不要(独自にカーソルを保持しない)
	// Scanはsqlの結果を構造体にコピーする役割を持つ
	err := r.db.QueryRowContext(ctx, q, id).Scan(&u.ID, &u.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *repository) GetPoint(ctx context.Context, id int64, date *time.Time) ([]PointRecord, error) {
	baseQuery := `
		SELECT tr.amount, e.point
		FROM training_records AS tr
		JOIN exercises AS e ON tr.exercise_id = e.id
		WHERE tr.user_id = $1
	`

	args := []any{id}

	if date != nil {
		baseQuery += " AND DATE(tr.date) = DATE($2)"
		args = append(args, *date)
	}

	// baseQueryをargsで呼び出す。
	// closeが必要。DBコネクション上でカーソルが開かれる
	// カーソルとは、まだ読み終わっていない結果セット(DBの用語)。closeしないとDBとの接続が解放されない
	// 処理が終わったら解放するように予約する
	rows, err := r.db.QueryContext(ctx, baseQuery, args...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			_ = fmt.Errorf("getPoint Error")
		}
	}(rows)

	var result []PointRecord
	// Nextはカーソルを次の行に進めるだけ
	for rows.Next() {
		var pr PointRecord
		// rows.Scan()は今カーソルが指している1行分のデータを構造体にコピーする
		// ポインタを渡すのはScanに構造体の中身を変更してもらうため
		// ポインタを渡すときは書き換えを実施する時と考える
		if err := rows.Scan(&pr.Amount, &pr.Point); err != nil {
			return nil, err
		}
		result = append(result, pr)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, err
}

func (r *repository) GetRanking(ctx context.Context) ([]Ranking, error) {
	const q = `
		SELECT tr.user_id, u.name, SUM(e.point * tr.amount) AS point
		FROM training_records AS tr
		JOIN users AS u ON tr.user_id = u.id
		JOIN exercises AS e ON tr.exercise_id = e.id
		GROUP BY tr.user_id, u.name
		ORDER BY point DESC 
	`
	// 行を返すクエリを実行する
	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Printf("rows close error: %v", err)
		}
	}()

	var result []Ranking
	for rows.Next() {
		var rank Ranking
		if err := rows.Scan(&rank.ID, &rank.Name, &rank.Point); err != nil {
			return nil, err
		}
		result = append(result, rank)
	}
	return result, err
}

func (r *repository) PostTrainingRecords(ctx context.Context, in PostTrainingRecordsInput) (int64, error) {
	const q = `
		INSERT INTO training_records (user_id, exercise_id, date, amount) 
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	var id int64
	err := r.db.QueryRowContext(ctx, q, in.ID, in.ExerciseID, in.Date, in.Amount).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
