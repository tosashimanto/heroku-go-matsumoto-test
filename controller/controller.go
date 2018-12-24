package controller

import (
	"fmt"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/tosashimanto/heroku-go-matsumoto-test/controller/handler"
	"html/template"
	"net/http"
	"time"
)

var (
	// セッションストアの初期化
	store *sessions.CookieStore = sessions.NewCookieStore(securecookie.GenerateRandomKey(64))
	// 登録ユーザーをメモリ上で管理
	users map[string]string = make(map[string]string)
)

const (
	SessionName       = "session-name"
	ContextSessionKey = "session"
)

func Handler() {
	r := mux.NewRouter()

	// "/"の場合の処理
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// テンプレート
		tmpl := template.Must(template.New("index").ParseFiles("templates/index.html"))
		err := tmpl.Execute(w, nil)
		if err != nil {
			panic(err)
		}
		fmt.Print(time.Now())
		fmt.Printf(" path: " + r.URL.Path + "\n")
	}).Methods("GET")

	handleFunc(r, "/test", handler.RegisterGetHandler).Methods("GET")

	// "/save"の場合の処理
	r.HandleFunc("/save", func(w http.ResponseWriter, r *http.Request) {

		// form値を取得
		r.ParseForm()
		data1 := r.FormValue("test1")
		data2 := r.FormValue("test2")
		// テンプレート
		tmpl := template.Must(template.New("save").ParseFiles("templates/save.html"))
		err := tmpl.Execute(w, struct {
			Data1 string
			Data2 string
		}{
			Data1: data1,
			Data2: data2,
		})
		if err != nil {
			panic(err)
		}
		fmt.Print(time.Now())
		fmt.Printf(" path: " + r.URL.Path + "\n")
	}).Methods("POST")

	// 可変URL
	r.HandleFunc("/test1/{sample1}/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "sample1\n")
		fmt.Print(time.Now())
		fmt.Printf(" path: " + r.URL.Path + "\n")
	})

	// 可変URL2
	r.HandleFunc("/test2/{sample1}/{id:[0-9]{1,6}}/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "sample2\n")
		fmt.Print(time.Now())
		fmt.Printf(" path: " + r.URL.Path + "\n")
	})

	// handing
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

// アプリケーション共通処理を常に呼び出すための糖衣構文
func handleFunc(r *mux.Router, path string, fn http.HandlerFunc) *mux.Route {
	return r.HandleFunc(path, applicationHandler(fn))
}

// アプリケーション共通処理
func applicationHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// セッションの取得
		session, err := store.Get(r, SessionName)
		if err != nil {
			// 不正なセッションだった場合は作り直す
			session, err = store.New(r, SessionName)
			checkError(err)
		}
		context.Set(r, ContextSessionKey, session)
		// 個別のハンドラー呼び出し
		fn(w, r)
	}
}

// エラーチェック
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
