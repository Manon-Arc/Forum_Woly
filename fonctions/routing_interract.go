package pagesfonctions

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func Add_Ban(w http.ResponseWriter, r *http.Request) {

	username, err := GetCookieHandler(w, r, "username")

	if err != nil || username == "" {
		fmt.Println(err)
		http.Redirect(w, r, "/", http.StatusFound)
	}
	topic_id, _ := strconv.Atoi(r.URL.Query().Get("idtopic"))

	database, err_base := sql.Open("sqlite3", "./data_base.sqlite")
	if err_base != nil {
		log.Fatal(err_base)
	}
	defer database.Close()

	ctx := context.Background()

	creator := r.URL.Query().Get("creator")

	request := `SELECT COUNT(*) FROM Ban WHERE username='` + fmt.Sprint(creator) + `' AND id_topic=` + fmt.Sprint(topic_id)

	rows, err := database.QueryContext(ctx, request)

	if err != nil {
		fmt.Println(err)
		return
	}

	var n int
	for rows.Next() {
		err := rows.Scan(&n)
		if err != nil {
			fmt.Println(err)
		}
		defer rows.Close()
	}
	if n == 0 {
		Add_ban(creator, topic_id)

		err := database.PingContext(ctx)
		if err != nil {
			fmt.Println(err)
		}
		stmts := `DELETE FROM Follow WHERE username = "` + creator + `" AND id_topic = ` + fmt.Sprint(topic_id) + ``
		_, err = database.ExecContext(ctx, stmts)
		if err != nil {
			fmt.Println(err)
		}
		http.Redirect(w, r, "/topic?id="+fmt.Sprint(topic_id), http.StatusFound)
	} else {
		stmts := `DELETE FROM Ban WHERE username = "` + creator + `" AND id_topic = ` + fmt.Sprint(topic_id) + ``
		_, err = database.ExecContext(ctx, stmts)
		if err != nil {
			fmt.Println(err)
		}
		http.Redirect(w, r, "/topic?id="+fmt.Sprint(topic_id), http.StatusFound)
	}
}

func Add_dislike(w http.ResponseWriter, r *http.Request) {
	database, err_base := sql.Open("sqlite3", "./data_base.sqlite")
	if err_base != nil {
		log.Fatal(err_base)
	}
	defer database.Close()

	ctx := context.Background()

	username, isConnected := IsConnected(w, r)
	post_id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	topic_id, _ := strconv.Atoi(r.URL.Query().Get("idtopic"))
	if isConnected {
		request := `SELECT COUNT(*) FROM Dislike WHERE username='` + fmt.Sprint(username) + `' AND id_post=` + fmt.Sprint(post_id)

		rows, err := database.QueryContext(ctx, request)

		if err != nil {
			fmt.Println(err)
			return
		}
		var n int
		for rows.Next() {
			err := rows.Scan(&n)
			if err != nil {
				fmt.Println(err)
			}
			defer rows.Close()
		}
		if n == 0 {
			Dislike(username, post_id)

			database, err_base := sql.Open("sqlite3", "./data_base.sqlite")
			if err_base != nil {
				fmt.Println(err_base)
			}

			ctx := context.Background()

			// Check if database is alive.
			err := database.PingContext(ctx)
			if err != nil {
				fmt.Println(err)
			}
			stmts := `DELETE FROM Jaime WHERE username = "` + username + `" AND id_post = ` + fmt.Sprint(post_id) + ``
			_, err = database.ExecContext(ctx, stmts)
			if err != nil {
				fmt.Println(err)
			}

			http.Redirect(w, r, "/topic?id="+fmt.Sprint(topic_id), http.StatusFound)

		} else {
			database, err_base := sql.Open("sqlite3", "./data_base.sqlite")
			if err_base != nil {
				fmt.Println(err_base)
			}

			ctx := context.Background()

			// Check if database is alive.
			err := database.PingContext(ctx)
			if err != nil {
				fmt.Println(err)
			}
			stmts := `DELETE FROM Dislike WHERE username = "` + username + `" AND id_post = ` + fmt.Sprint(post_id) + ``
			_, err = database.ExecContext(ctx, stmts)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	http.Redirect(w, r, "/topic?id="+fmt.Sprint(topic_id), http.StatusFound)
}

func Add_like(w http.ResponseWriter, r *http.Request) {
	database, err_base := sql.Open("sqlite3", "./data_base.sqlite")
	if err_base != nil {
		log.Fatal(err_base)
	}
	defer database.Close()

	ctx := context.Background()
	username, isConnected := IsConnected(w, r)
	post_id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	topic_id, _ := strconv.Atoi(r.URL.Query().Get("idtopic"))
	if isConnected {
		request := `SELECT COUNT(*) FROM Jaime WHERE username='` + fmt.Sprint(username) + `' AND id_post=` + fmt.Sprint(post_id)

		rows, err := database.QueryContext(ctx, request)

		if err != nil {
			fmt.Println(err)
			return
		}
		var n int
		for rows.Next() {
			err := rows.Scan(&n)
			if err != nil {
				fmt.Println(err)
			}
			defer rows.Close()
		}
		if n == 0 {
			Like(username, post_id)

			database, err_base := sql.Open("sqlite3", "./data_base.sqlite")
			if err_base != nil {
				fmt.Println(err_base)
			}

			ctx := context.Background()

			// Check if database is alive.
			err := database.PingContext(ctx)
			if err != nil {
				fmt.Println(err)
			}
			stmts := `DELETE FROM Dislike WHERE username = "` + username + `" AND id_post = ` + fmt.Sprint(post_id) + ``
			_, err = database.ExecContext(ctx, stmts)
			if err != nil {
				fmt.Println(err)
			}

			http.Redirect(w, r, "/topic?id="+fmt.Sprint(topic_id), http.StatusFound)

		} else {
			database, err_base := sql.Open("sqlite3", "./data_base.sqlite")
			if err_base != nil {
				fmt.Println(err_base)
			}

			ctx := context.Background()

			// Check if database is alive.
			err := database.PingContext(ctx)
			if err != nil {
				fmt.Println(err)
			}
			stmts := `DELETE FROM Jaime WHERE username = "` + username + `" AND id_post = ` + fmt.Sprint(post_id) + ``
			_, err = database.ExecContext(ctx, stmts)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	http.Redirect(w, r, "/topic?id="+fmt.Sprint(topic_id), http.StatusFound)
}

func Add_Mode(w http.ResponseWriter, r *http.Request) {
	database, err_base := sql.Open("sqlite3", "./data_base.sqlite")
	if err_base != nil {
		log.Fatal(err_base)
	}
	defer database.Close()

	ctx := context.Background()

	username, err := GetCookieHandler(w, r, "username")

	if err != nil || username == "" {
		fmt.Println(err)
		http.Redirect(w, r, "/", http.StatusFound)
	}

	topic_id, _ := strconv.Atoi(r.URL.Query().Get("idtopic"))
	creator := r.URL.Query().Get("creator")

	request := `SELECT COUNT(*) FROM Moderateur WHERE username='` + fmt.Sprint(creator) + `' AND id_topic=` + fmt.Sprint(topic_id)

	rows, err := database.QueryContext(ctx, request)

	if err != nil {
		fmt.Println(err)
		return
	}
	var n int
	for rows.Next() {
		err := rows.Scan(&n)
		if err != nil {
			fmt.Println(err)
		}
		defer rows.Close()

	}
	if n == 0 {
		Add_mode(creator, topic_id)
		http.Redirect(w, r, "/topic?id="+fmt.Sprint(topic_id), http.StatusFound)

	} else {
		database, err_base := sql.Open("sqlite3", "./data_base.sqlite")
		if err_base != nil {
			fmt.Println(err_base)
		}

		ctx := context.Background()

		// Check if database is alive.
		err := database.PingContext(ctx)
		if err != nil {
			fmt.Println(err)
		}
		stmts := `DELETE FROM Moderateur WHERE username = "` + creator + `" AND id_topic = ` + fmt.Sprint(topic_id) + ``
		_, err = database.ExecContext(ctx, stmts)
		if err != nil {
			fmt.Println(err)
		}
		http.Redirect(w, r, "/topic?id="+fmt.Sprint(topic_id), http.StatusFound)
	}
}

func Add_Follow(w http.ResponseWriter, r *http.Request) {
	database, err_base := sql.Open("sqlite3", "./data_base.sqlite")
	if err_base != nil {
		log.Fatal(err_base)
	}
	defer database.Close()

	ctx := context.Background()

	username, err := GetCookieHandler(w, r, "username")

	if err != nil || username == "" {
		fmt.Println(err)
		http.Redirect(w, r, "/", http.StatusFound)
	}
	topic_id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	request := `SELECT COUNT(*) FROM Follow WHERE username='` + fmt.Sprint(username) + `' AND id_topic=` + fmt.Sprint(topic_id)

	rows, err := database.QueryContext(ctx, request)

	if err != nil {
		fmt.Println(err)
		return
	}

	var n int
	for rows.Next() {
		err := rows.Scan(&n)
		if err != nil {
			fmt.Println(err)
		}
		defer rows.Close()
	}
	if n == 0 {
		Follow(username, topic_id)
		http.Redirect(w, r, "/topic?id="+fmt.Sprint(topic_id), http.StatusFound)
	} else {
		database, err_base := sql.Open("sqlite3", "./data_base.sqlite")
		if err_base != nil {
			fmt.Println(err_base)
		}

		ctx := context.Background()

		// Check if database is alive.
		err := database.PingContext(ctx)
		if err != nil {
			fmt.Println(err)
		}
		stmts := `DELETE FROM Follow WHERE username = "` + username + `" AND id_topic = ` + fmt.Sprint(topic_id) + ``
		_, err = database.ExecContext(ctx, stmts)
		if err != nil {
			fmt.Println(err)
		}
		http.Redirect(w, r, "/topic?id="+fmt.Sprint(topic_id), http.StatusFound)
	}

	http.Redirect(w, r, "/topic?id="+fmt.Sprint(topic_id), http.StatusFound)
}

func Creat_Post(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		post_content := r.FormValue("post_content")
		creator_name, err := GetCookieHandler(w, r, "username")
		if err != nil {
			fmt.Println(err)
		}

		topic_id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		picture, _, err := r.FormFile("postpic")

		id, _ := AddPost(w, topic_id, creator_name, post_content)

		if err != nil {
			fmt.Println(err)
		} else {
			f, err := os.OpenFile("./server/img/post_pic/"+fmt.Sprint(id)+".png", os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err)
			} else {
				defer f.Close()
				io.Copy(f, picture)
			}
			defer picture.Close()
		}

		Add_Pictopost(w, r, int(id), "./server/img/post_pic/"+fmt.Sprint(id)+".png")

		http.Redirect(w, r, "/topic?id="+fmt.Sprint(topic_id), http.StatusFound)
	default:
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func Creat_Topic(w http.ResponseWriter, r *http.Request) {

	username, err := GetCookieHandler(w, r, "username")

	if err != nil || username == "" {
		fmt.Println(err)
		http.Redirect(w, r, "/", http.StatusFound)
	}

	switch r.Method {
	case "POST":
		topic_name := r.FormValue("topic_name")
		topic_description := r.FormValue("topic_description")
		topic_category := r.FormValue("topic_category")
		topic_username, err := GetCookieHandler(w, r, "username")
		if err != nil {
			fmt.Println(err)
		}
		topic_picture := "./server/img/wolf.png"
		username, err := GetCookieHandler(w, r, "username")

		if err != nil {
			fmt.Println(err)
		}
		nId, _ := CreateTopic(topic_name, topic_description, topic_category, topic_username, topic_picture)
		Follow(username, int(nId))
		http.Redirect(w, r, "/topic?id="+fmt.Sprint(nId), http.StatusFound)
	}
}

func Edit_post(w http.ResponseWriter, r *http.Request) {

	topic_id, _ := strconv.Atoi(r.URL.Query().Get("idtopic"))
	post_id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	table := "Post"
	changeValue := make(map[string]interface{})
	changeValue["content"] = r.FormValue("new_content")
	where := make(map[string]interface{})
	where["id"] = post_id
	Update(table, changeValue, where)

	http.Redirect(w, r, "/topic?id="+fmt.Sprint(topic_id), http.StatusFound)
}

func Supp_post(w http.ResponseWriter, r *http.Request) {

	post_id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	topic_id, _ := strconv.Atoi(r.URL.Query().Get("idtopic"))

	database, err_base := sql.Open("sqlite3", "./data_base.sqlite")
	if err_base != nil {
		fmt.Println(err_base)
	}
	ctx := context.Background()

	// Check if database is alive.
	err := database.PingContext(ctx)
	if err != nil {
		fmt.Println(err)
	}
	stmts := `DELETE FROM Post WHERE id = "` + fmt.Sprint(post_id) + `"`
	_, err = database.ExecContext(ctx, stmts)
	if err != nil {
		fmt.Println(err)
	}
	http.Redirect(w, r, "/topic?id="+fmt.Sprint(topic_id), http.StatusFound)
}

func Supp_topic(w http.ResponseWriter, r *http.Request) {
	topic_id, _ := strconv.Atoi(r.URL.Query().Get("idtopic"))

	database, err_base := sql.Open("sqlite3", "./data_base.sqlite")
	if err_base != nil {
		fmt.Println(err_base)
	}
	ctx := context.Background()

	// Check if database is alive.
	err := database.PingContext(ctx)
	if err != nil {
		fmt.Println(err)
	}
	stmts := `DELETE FROM Topic WHERE id = "` + fmt.Sprint(topic_id) + `"`
	_, err = database.ExecContext(ctx, stmts)
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/home "+fmt.Sprint(topic_id), http.StatusFound)

}
