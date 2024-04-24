package pagesfonctions

import (
	"context"
	"database/sql"
	"fmt"
	s "forum/structs"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func Topic(w http.ResponseWriter, r *http.Request) {

	database, err_base := sql.Open("sqlite3", "./data_base.sqlite")
	if err_base != nil {
		log.Fatal(err_base)
	}
	defer database.Close()

	ctx := context.Background()

	TopicPage := s.TopicPage{}
	TopicPage.ImMember = false
	TopicPage.ImMod = false
	var username string

	topic_id := r.URL.Query().Get("id")

	username, TopicPage.Isconnected = IsConnected(w, r)

	if !TopicExist(topic_id) {
		http.Redirect(w, r, "/404", http.StatusFound)
	}else {
		GetTopic(&TopicPage.Topic, topic_id)
		GetUsersMemberOf(&TopicPage.Members, topic_id)
		TopicPage.IsMine = TopicPage.Topic.Creator.Username == username
	}

	if TopicPage.Isconnected {
		GetUserData(w, r, &TopicPage.CurrentUser, username)
		var mres int
		mr := database.QueryRowContext(ctx, `SELECT COUNT(*) FROM Follow WHERE id_topic=`+fmt.Sprint(topic_id)+` AND username='`+username+`'`)
		mr.Scan(&mres)
		TopicPage.ImMember = mres == 1
		var mores int
		mor := database.QueryRowContext(ctx, `SELECT COUNT(*) FROM Moderateur WHERE id_topic=`+fmt.Sprint(topic_id)+` AND username='`+username+`'`)
		mor.Scan(&mores)
		TopicPage.ImMod = mores == 1
		if TopicPage.ImMod  || TopicPage.IsMine {
			GetUsersBanOf(&TopicPage.Bans, topic_id)
		}else {	
			GetUsersModOf(&TopicPage.Moderateurs, topic_id)
		}
	}

	switch r.Method {
	case "POST":
		like := r.FormValue("like") == "isCheck"
		dlike := r.FormValue("dlike") == "isCheck"
		mine := r.FormValue("mine") == "isCheck"
		GetPostWithFilter(&TopicPage.Posts, username, topic_id, like, dlike, mine)
	default:
		GetAllPostOfTopic(&TopicPage.Posts, username, topic_id)
	}

	var tmplt = template.Must(template.ParseFiles("server/html/topic.html"))
	err := tmplt.Execute(w, TopicPage)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, err)
	}
}
