package pagesfonctions

import (
	"context"
	"database/sql"
	"fmt"
	s "forum/structs"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func Profil(w http.ResponseWriter, r *http.Request) {

	database, err_base := sql.Open("sqlite3", "./data_base.sqlite")
	if err_base != nil {
		log.Fatal(err_base)
	}
	defer database.Close()

	ctx := context.Background()

	ProfilPage := s.ProfilPage{}
	var username string

	user_profil := r.URL.Query().Get("usr")

	username, ProfilPage.Isconnected = IsConnected(w, r)

	ProfilPage.ItsMe = false
	if ProfilPage.Isconnected && username == user_profil {
		GetUserData(w, r, &ProfilPage.CurrentUser, username)
		GetTopicFollowBy(username, &ProfilPage.Topics)
		ProfilPage.ItsMe = true
	} else if user_profil != "" {
		re := database.QueryRowContext(ctx, `SELECT COUNT(*) FROM Users WHERE username="`+user_profil+`"`)
		var res int
		re.Scan(&res)
		if res == 0 {
			http.Redirect(w, r, "/404", http.StatusFound)
		}
		GetUserData(w, r, &ProfilPage.CurrentUser, user_profil)
		GetTopicFollowBy(user_profil, &ProfilPage.Topics)
	} else {
		http.Redirect(w, r, "/404", http.StatusFound)
	}

	switch r.Method {
	case "POST":
		newUsername := r.FormValue("username")
		newMail := r.FormValue("mail")
		newDesc := r.FormValue("bio")
		file, _, err := r.FormFile("profilpic")
		if err != nil {
			fmt.Println(err)
			UpdateUser(w, r, newUsername, newMail, newDesc, "")
		}else {
			f, err := os.OpenFile("./server/img/profil_pic/"+newUsername+".png", os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err)
				UpdateUser(w, r, newUsername, newMail, newDesc, "")
			}else {
				io.Copy(f, file)
				defer f.Close()
				UpdateUser(w, r, newUsername, newMail, newDesc, "./server/img/profil_pic/"+newUsername+".png")
			}
			defer file.Close()
		}

		http.Redirect(w, r, "/profil?usr="+newUsername, http.StatusFound)
	}

	GetPostOfUser(&ProfilPage.Posts, user_profil)

	var tmplt = template.Must(template.ParseFiles("server/html/profil.html"))
	err := tmplt.Execute(w, ProfilPage)
	if err != nil {
		fmt.Println(err)
	}
}
