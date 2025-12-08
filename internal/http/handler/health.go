package handler

import "net/http"

// Health
// wはクライアントへ返すレスポンスを書き込むためのもの
// rはリクエストの情報。クライアントから送られてきた情報が全て入っている
func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}
