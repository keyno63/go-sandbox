package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	//"github.com/keyno/go-sandbox/src/main/app/internal/libs"
)

var db *sql.DB

// 実行関数
func main() {
	// Http Handler の設定
	mux := http.NewServeMux()
	handler := NewHandler()
	mux.Handle("/api/go-app/handle", handler)

	// server の起動設定
	server := http.Server{
		Addr:    ":8180",
		Handler: mux,
	}

	// server の起動
	_ = server.ListenAndServe()
}

type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	app *App
}

func NewHandler() Handler {
	// DI
	repository := UserRepository{db}
	service := UserService{&repository}
	controller := UserController{&service}
	app := App{&controller}

	return handler{&app}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	userId := r.FormValue("user_id")
	user := h.app.ExecGetUser(userId)

	jsonUser, _ := json.Marshal(user)

	printValue := fmt.Sprintf("param value is = [%s]", user)
	fmt.Println(printValue)

	w.Header().Set("Content-Type", "application/json")
	//_, _ = fmt.Fprintf(w, "sample handle")
	ret, err := w.Write(jsonUser)
	if err != nil {
		fmt.Printf("occured error: " + err.Error())
	}

	fmt.Println("write ret: " + string(rune(ret)))
}

/**
App.
handler と Controller が 1:1 になっているので、この層は不要？
TODO: 要不要の検討, 修正.
*/

type App struct {
	userController *UserController
}

func (app App) ExecGetUser(id string) UserAccount {
	return app.userController.GetUser(id)
}

/**
Controller
*/

type UserController struct {
	userService *UserService
}

func (uc UserController) GetUser(id string) UserAccount {
	return uc.userService.GetUser(id)
}

/**
Service
*/

type UserService struct {
	userRepository *UserRepository
}

func (us UserService) GetUser(id string) UserAccount {
	return us.userRepository.GetUser(id)
}

/**
Repository
*/

type UserRepository struct {
	dbConnector *sql.DB
}

func (us UserRepository) GetUser(id string) UserAccount {
	// 仮
	// TODO: DBとの接続の実装
	return UserAccount{
		Id:        id,
		FirstName: "first",
		LastName:  "last",
	}
}

/**
  data struct
*/

type UserAccount struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
