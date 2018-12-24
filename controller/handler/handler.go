package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Data1 struct {
	Title   string `json:"title"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func RegisterGetHandler(w http.ResponseWriter, r *http.Request) {

	// 構造体を定義
	var data1 = Data1{}
	data1.Title = "sample1"
	data1.Message = "hello, sample1"
	data1.Status = 100

	// jsonエンコード
	outputJson, err := json.Marshal(&data1)
	if err != nil {
		panic(err)
	}

	// jsonヘッダーを出力
	w.Header().Set("Content-Type", "application/json")

	// jsonデータを出力
	fmt.Fprint(w, string(outputJson))
}
