package pagesfonctions

import (
	"context"
	"database/sql"
	"fmt"
	s "forum/structs"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTopic(Topic_Name, Topic_content, Category, Creator_ID, Picture string) (int64, error) {
	m := make(map[string]interface{})
	m["name"] = Topic_Name
	m["content"] = Topic_content
	m["creator"] = Creator_ID
	m["categorie"] = Category
	m["picture"] = Picture
	return Insert("Topic", m, "id")
}

func AddPost(w http.ResponseWriter, Topic_ID int, Creator_name, Post_Content string) (int64, error){
	m := make(map[string]interface{})
	m["topic"] = Topic_ID
	m["creator"] = Creator_name
	m["content"] = Post_Content
	m["pic"] = ""
	return Insert("Post", m, "id")
}

func Add_Pictopost(w http.ResponseWriter, r *http.Request,id int, img string){
	changeValue := make(map[string]interface{})
	changeValue["pic"] = img
	where := make(map[string]interface{})
	where["id"] = id
	Update("Post", changeValue, where)
}

func ModifyPostOrTopic(TableName, CellName, WhereToFind, WhatToChange, NewValue string) {
	changeValue := make(map[string]interface{})
	changeValue[CellName] = NewValue
	where := make(map[string]interface{})
	where[WhereToFind] = WhatToChange

	Update(TableName, changeValue, where)
}

func Follow(username string, idTopic int) {
	m := make(map[string]interface{})
	m["username"] = username
	m["id_topic"] = idTopic
	Insert("Follow", m, "")
}

func Add_ban(username string, idTopic int) {
	m := make(map[string]interface{})
	m["username"] = username
	m["id_topic"] = idTopic
	Insert("Ban", m, "")
}

func Add_mode(username string, idTopic int) {
	m := make(map[string]interface{})
	m["username"] = username
	m["id_topic"] = idTopic
	Insert("Moderateur", m, "")
}

func Like(username string, idPost int) {
	m := make(map[string]interface{})
	m["username"] = username
	m["id_post"] = idPost
	Insert("Jaime", m, "")
}

func Dislike(username string, idPost int) {
	m := make(map[string]interface{})
	m["username"] = username
	m["id_post"] = idPost
	Insert("Dislike", m, "")
}

func AddComment(username, content string, idPost int) {
	m := make(map[string]interface{})
	m["username"] = username
	m["id_topic"] = idPost
	m["content"] = content
	Insert("Jaime", m, "")
}

func UpdateUser(w http.ResponseWriter, r *http.Request, username, mail, desc, img string) {

	currentUsername, err := GetCookieHandler(w, r, "username")

	changeValue := make(map[string]interface{})
	changeValue["username"] = username
	changeValue["mail"] = mail
	if img != "" {
		changeValue["pic"] = img
	}
	changeValue["bio"] = desc
	where := make(map[string]interface{})
	if err != nil {
		fmt.Println(err)
	}
	where["username"] = currentUsername

	Update("Users", changeValue, where)

	if img != ""{
		err = os.Rename("./server/img/profil_pic/"+username+".png", img)
	}

	changeValue = make(map[string]interface{})
	changeValue["creator"] = username
	where = make(map[string]interface{})
	if err != nil {
		fmt.Println(err)
	}
	where["creator"] = currentUsername
	Update("Comment", changeValue, where)
	Update("Post", changeValue, where)
	Update("Topic", changeValue, where)

	changeValue = make(map[string]interface{})
	changeValue["username"] = username
	where = make(map[string]interface{})
	if err != nil {
		fmt.Println(err)
	}
	where["username"] = currentUsername

	Update("Follow", changeValue, where)
	Update("Jaime", changeValue, where)
	Update("Ban", changeValue, where)
	Update("Dislike", changeValue, where)
	Update("Moderateur", changeValue, where)

	SetCookieHandler(w, r, "username", username)
}

func GetUserData(w http.ResponseWriter, r *http.Request, user *s.Users, username string) {
	database, err_base := sql.Open("sqlite3", "./data_base.sqlite")
	if err_base != nil {
		fmt.Println(err_base)
	}
	ctx := context.Background()

	request := `SELECT username, mail, status, pic, bio FROM Users WHERE username="` + username + `"`

	rows, err := database.QueryContext(ctx, request)

	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		err := rows.Scan(&user.Username, &user.Mail, &user.Status, &user.Pic, &user.Bio)
		if err != nil {
			fmt.Println(err)
		}
	}
	defer rows.Close()
}
