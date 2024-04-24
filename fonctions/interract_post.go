package pagesfonctions

import (
	"context"
	"database/sql"
	"fmt"
	s "forum/structs"
)

func GetAllPostOfTopic(posts *[]s.Post, username, topic_id string) {

	database, err_base := sql.Open("sqlite3", "./data_base.sqlite")
	if err_base != nil {
		fmt.Println(err_base)
	}
	ctx := context.Background()

	request := `SELECT DISTINCT Post.id, Post.topic, Post.content, Post.pic, Topic.name, Topic.picture, Users.username, Users.pic FROM Post INNER JOIN Topic ON Post.topic = Topic.id INNER JOIN Users ON Post.creator = Users.username WHERE Topic.id=` + topic_id

	rows, err := database.QueryContext(ctx, request)

	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		var Post s.Post
		var Topic s.Topic
		var User s.Users
		err := rows.Scan(&Post.Id, &Topic.Id, &Post.Content, &Post.Pic, &Topic.Name, &Topic.Picture, &User.Username, &User.Pic)
		if err != nil {
			fmt.Println(err)
		}

		row := database.QueryRowContext(ctx, `SELECT COUNT(*) FROM Jaime WHERE id_post=`+fmt.Sprint(Post.Id))
		row.Scan(&Post.Like)
		row = database.QueryRowContext(ctx, `SELECT COUNT(*) FROM Dislike WHERE id_post=`+fmt.Sprint(Post.Id))
		row.Scan(&Post.Dislike)
		Post.Topics = Topic
		Post.Creator = User
		r := database.QueryRowContext(ctx, `SELECT COUNT(*) FROM Jaime WHERE id_post=`+fmt.Sprint(Post.Id)+` AND username='`+username+`'`)
		var res int
		r.Scan(&res)
		Post.ILike = res == 1
		r = database.QueryRowContext(ctx, `SELECT COUNT(*) FROM Dislike WHERE id_post=`+fmt.Sprint(Post.Id)+` AND username='`+username+`'`)
		r.Scan(&res)
		Post.IDislike = res == 1
		r = database.QueryRowContext(ctx, `SELECT COUNT(*) FROM Post WHERE id=`+fmt.Sprint(Post.Id)+` AND creator='`+username+`'`)
		r.Scan(&res)
		Post.IsMine = res == 1

		*posts = append(*posts, Post)
	}
	defer rows.Close()
}

func GetPostOfUser(posts *[]s.Post, username string){
	database, err_base := sql.Open("sqlite3", "./data_base.sqlite")
	if err_base != nil {
		fmt.Println(err_base)
	}
	ctx := context.Background()

	request := `SELECT DISTINCT Post.id, Post.topic, Post.content, Post.pic, Topic.name, Topic.picture, Users.username, Users.pic FROM Post INNER JOIN Topic ON Post.topic = Topic.id INNER JOIN Users ON Post.creator = Users.username WHERE Post.creator="`+username+`"`

	rows, err := database.QueryContext(ctx, request)

	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		var Post s.Post
		var Topic s.Topic
		var User s.Users
		err := rows.Scan(&Post.Id, &Topic.Id, &Post.Content, &Post.Pic, &Topic.Name, &Topic.Picture, &User.Username, &User.Pic)
		if err != nil {
			fmt.Println(err)
		}

		row := database.QueryRowContext(ctx, `SELECT COUNT(*) FROM Jaime WHERE id_post=`+fmt.Sprint(Post.Id))
		row.Scan(&Post.Like)
		row = database.QueryRowContext(ctx, `SELECT COUNT(*) FROM Dislike WHERE id_post=`+fmt.Sprint(Post.Id))
		row.Scan(&Post.Dislike)
		Post.Topics = Topic
		Post.Creator = User
		r := database.QueryRowContext(ctx, `SELECT COUNT(*) FROM Jaime WHERE id_post=`+fmt.Sprint(Post.Id)+` AND username='`+username+`'`)
		var res int
		r.Scan(&res)
		Post.ILike = res == 1
		r = database.QueryRowContext(ctx, `SELECT COUNT(*) FROM Dislike WHERE id_post=`+fmt.Sprint(Post.Id)+` AND username='`+username+`'`)
		r.Scan(&res)
		Post.IDislike = res == 1
		*posts = append(*posts, Post)
	}
	defer rows.Close()
}

func GetPostWithFilter(posts *[]s.Post, username, topic_id string, like, dislike, mine bool) {

	database, err_base := sql.Open("sqlite3", "./data_base.sqlite")
	if err_base != nil {
		fmt.Println(err_base)
	}
	ctx := context.Background()

	request := `SELECT DISTINCT Post.id, Post.topic, Post.content, Post.pic, Topic.name, Topic.picture, Users.username, Users.pic FROM Post INNER JOIN Topic ON Post.topic = Topic.id INNER JOIN Users ON Post.creator = Users.username`
	if like {
		request += ` INNER JOIN Jaime ON Jaime.id_post = Post.id AND Jaime.username ='` + username + `'`
	}
	if dislike {
		request += ` INNER JOIN Dislike ON Dislike.id_post = Post.id AND Dislike.username ='` + username + `'`
	}
	request += ` WHERE Topic.id=` + topic_id

	if mine {
		request += ` AND Post.creator ='` + username + `'`
	}
	rows, err := database.QueryContext(ctx, request)

	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		var Post s.Post
		var Topic s.Topic
		var User s.Users
		err := rows.Scan(&Post.Id, &Topic.Id, &Post.Content, &Post.Pic, &Topic.Name, &Topic.Picture, &User.Username, &User.Pic)

		if err != nil {
			fmt.Println(err)
		}

		row := database.QueryRowContext(ctx, `SELECT COUNT(*) FROM Jaime WHERE id_post=`+fmt.Sprint(Post.Id))
		row.Scan(&Post.Like)
		row = database.QueryRowContext(ctx, `SELECT COUNT(*) FROM Dislike WHERE id_post=`+fmt.Sprint(Post.Id))
		row.Scan(&Post.Dislike)
		Post.Topics = Topic
		Post.Creator = User
		r := database.QueryRowContext(ctx, `SELECT COUNT(*) FROM Jaime WHERE id_post=`+fmt.Sprint(Post.Id)+` AND username='`+username+`'`)
		var res int
		r.Scan(&res)
		Post.ILike = res == 1
		r = database.QueryRowContext(ctx, `SELECT COUNT(*) FROM Dislike WHERE id_post=`+fmt.Sprint(Post.Id)+` AND username='`+username+`'`)
		r.Scan(&res)
		Post.IDislike = res == 1
		r = database.QueryRowContext(ctx, `SELECT COUNT(*) FROM Post WHERE id=`+fmt.Sprint(Post.Id)+` AND creator='`+username+`'`)
		r.Scan(&res)
		Post.IsMine = res == 1
		*posts = append(*posts, Post)
	}
}
