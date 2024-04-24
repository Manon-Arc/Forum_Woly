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

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		var tmplt = template.Must(template.ParseFiles("server/html/404.html"))
		err := tmplt.Execute(w, nil)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		database, err_base := sql.Open("sqlite3", "./data_base.sqlite")
		if err_base != nil {
			log.Fatal(err_base)
		}
		defer database.Close()

		HomeData := s.MainPage{}
		var username string

		username, HomeData.Isconnected = IsConnected(w, r)

		if HomeData.Isconnected {
			GetUserData(w, r, &HomeData.CurrentUser, username)
			GetTopicFollowBy(username, &HomeData.Follow_Topics)
		} else {
			HomeData.Follow_Topics = []s.Topic{}
		}

		switch r.Method {
		case "POST":
			name := r.FormValue("tpname")
			cat := r.FormValue("tpcat")
			asc := r.FormValue("asc") == "isCheck"
			desc := r.FormValue("desc") == "isCheck"
			notfollow := r.FormValue("nfollow") == "isCheck"
			order := ""
			if asc {
				order = "ASC"
			} else if desc {
				order = "DESC"
			}
			GetTopicWithFilter(&HomeData.All_Topics, order, name, cat, username, notfollow)
		default:
			GetAllTopic(&HomeData.All_Topics)
		}

		// select posts

		database, err_base = sql.Open("sqlite3", "./data_base.sqlite")
		if err_base != nil {
			fmt.Println(err_base)
		}
		ctx := context.Background()

		err := database.PingContext(ctx)
		if err != nil {
			fmt.Println(err)
		}
		request := `SELECT Post.*, Topic.name, Topic.picture, Users.pic FROM Post INNER JOIN Follow ON Post.topic = Follow.id_topic INNER JOIN Topic ON Post.topic = Topic.id INNER JOIN Users ON Post.creator = Users.username ORDER BY id Desc LIMIT 5`
		_, err = database.ExecContext(ctx, request)
		if err != nil {
			fmt.Println(err)
		}

		stmt, err := database.Prepare(request)
		if err != nil {
			fmt.Println(err)
		}
		defer stmt.Close()

		rows_post, _ := database.QueryContext(ctx, request)

		for rows_post.Next() {
			var Post s.Post

			err := rows_post.Scan(&Post.Id, &Post.Creator.Username, &Post.Topics.Id, &Post.Content, &Post.Pic, &Post.Topics.Name, &Post.Topics.Picture, &Post.Creator.Pic)
			if err != nil {
				fmt.Println(err)
			}
			HomeData.Posts = append(HomeData.Posts, Post)
		}

		var tmplt = template.Must(template.ParseFiles("server/html/index.html"))
		err = tmplt.Execute(w, HomeData)
		if err != nil {
			fmt.Println(err)
		}
	}
}
