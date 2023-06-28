package server

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"warehouse-application/errors"
	"warehouse-application/page"
	"warehouse-application/user"
)

type card struct {
	id     int64
	name   string
	photo  string
	cost   int64
	amount int64
	have   bool
}

type warehouse struct {
	id    int
	items []int64
}

var DB *sql.DB

func StartWebServer(db *sql.DB) {
	DB = db
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		page.ShowPage(w, r)
	})

	http.HandleFunc("/RegistrationUser", handlerRegistrationUser)

	http.HandleFunc("/LoginUser", handlerLoginUser)

	http.HandleFunc("/CreateWarehouse", handlerCreateWarehouse)

	err := http.ListenAndServe("localhost:8080", nil)
	errors.CheckError(err)
}

func handlerCreateWarehouse(w http.ResponseWriter, r *http.Request) {

}

func createWarehouse() {

}

func handlerRegistrationUser(w http.ResponseWriter, r *http.Request) {
	newUser := getUser(r)
	answer := registrationUser(newUser)
	writeAnswer(w, answer)
}

func getUser(r *http.Request) user.User {
	var newUser user.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	errors.CheckWarning(err)
	return newUser
}

func registrationUser(newUser user.User) string {
	if newUser.RegisterUser(DB) == nil {
		return "Successful registration!"
	} else {
		return "Already registered..."
	}
}

func writeAnswer(w http.ResponseWriter, answer string) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(answer))
	errors.CheckWarning(err)
}

func handlerLoginUser(w http.ResponseWriter, r *http.Request) {
	alreadyUser := getUser(r)
	answer := loginUser(alreadyUser)
	writeAnswer(w, answer)
}

func loginUser(alreadyUser user.User) string {
	err := alreadyUser.LoginUser(DB)
	if err == nil {
		return "Successful login!"
	}
	return "Dont registered"
}
