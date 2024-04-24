package pagesfonctions

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func Login_register_page(w http.ResponseWriter, r *http.Request) {
	database, err_base := sql.Open("sqlite3", "./data_base.sqlite")
	if err_base != nil {
		log.Fatal(err_base)
	}
	defer database.Close()

	var tmplt = template.Must(template.ParseFiles("server/html/connection.html"))

	username := r.FormValue("username")
	password := r.FormValue("password")
	confirm_password := r.FormValue("confim_password")
	mail := r.FormValue("mail")
	var Mess string

	switch r.Method {
	case "POST":
		if len(username) != 0 && len(password) != 0 && len(mail) == 0 {
			_, err := Login(w, r, username, password)
			if err == nil {
				http.Redirect(w, r, "/", http.StatusFound)
			} else {
				http.Redirect(w, r, "/connection", http.StatusFound)
			}
		} else {
			if password == confirm_password {
				err := Register(w, r, username, password, mail)
				if err != nil {
					http.Redirect(w, r, "/connection", http.StatusFound)
				} else {
					http.Redirect(w, r, "/", http.StatusFound)
				}
			} else {
				Mess = "password must be the same as confirm password"
			}
		}
	}
	err := tmplt.Execute(w, Mess)
	if err != nil {
		fmt.Println(err)
	}
}

func Login(w http.ResponseWriter, r *http.Request, user, pass string) (string, error) {
	database, err_base := sql.Open("sqlite3", "./data_base.sqlite")
	if err_base != nil {
		log.Fatal(err_base)
	}
	defer database.Close()

	ctx := context.Background()
	err := database.PingContext(ctx)
	if err != nil {
		fmt.Println(err)
	}

	tsql := fmt.Sprintf("SELECT username, password from Users WHERE username ='%s'", user)

	rows, err := database.QueryContext(ctx, tsql)
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var name, password string

		err := rows.Scan(&name, &password)
		if err != nil {
			fmt.Println(err)
		}
		if ComparePassword(password, pass) == nil {
			SetCookieHandler(w, r, "username", name)
			SetCookieHandler(w, r, "password", password)
			return "login ok", nil
		}
	}
	return "", errors.New("user or password isn't good")
}

func Register(w http.ResponseWriter, r *http.Request, user, pass, mail string) error {
	database, err_base := sql.Open("sqlite3", "./data_base.sqlite")
	if err_base != nil {
		log.Fatal(err_base)
		return err_base
	}
	defer database.Close()

	ctx := context.Background()

	err := database.PingContext(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}
	insertValue := make(map[string]interface{})
	insertValue["username"] = user
	insertValue["mail"] = mail
	insertValue["password"] = HashPassword(pass)
	insertValue["status"] = true
	insertValue["pic"] = "./server/img/wolf.png"
	insertValue["bio"] = ""

	_, err = Insert("Users", insertValue, "")
	if err != nil {
		fmt.Println(err)
		return err
	}
	SetCookieHandler(w, r, "username", user)
	SetCookieHandler(w, r, "password", pass)
	return nil
}

func HashPassword(password string) string {
	pw := []byte(password)
	result, _ := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	return string(result)

}

func ComparePassword(hashPassword string, password string) error {
	pw := []byte(password)
	hw := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(hw, pw)
	return err
}

func LogOut(w http.ResponseWriter, r *http.Request) {
	SetCookieHandler(w, r, "username", "")
	SetCookieHandler(w, r, "password", "")
	http.Redirect(w, r, "/", http.StatusFound)
}

func IsConnected(w http.ResponseWriter, r *http.Request) (string, bool) {
	database, err_base := sql.Open("sqlite3", "./data_base.sqlite")
	if err_base != nil {
		log.Fatal(err_base)
	}
	defer database.Close()

	username, err := GetCookieHandler(w, r, "username")
	if err != nil {
		return "", false
	}

	ctx := context.Background()
	row := database.QueryRowContext(ctx, `SELECT COUNT(*) FROM Users WHERE username='`+username+`'`)
	var res int
	row.Scan(&res)
	if res != 1 {
		SetCookieHandler(w, r, "username", "")
		SetCookieHandler(w, r, "password", "")
		return "", false
	}

	return username, true
}
