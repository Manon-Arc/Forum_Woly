package pagesfonctions

import (
	"context"
	"database/sql"
	"fmt"
	s "forum/structs"
	"log"
	"strings"
)

func TopicExist(id string) bool {
	database, err_base := sql.Open("sqlite3", "./data_base.sqlite")
	if err_base != nil {
		log.Fatal(err_base)
	}
	defer database.Close()

	ctx := context.Background()
	re := database.QueryRowContext(ctx, `SELECT COUNT(*) FROM Topic WHERE id=`+id)
	var res int
	re.Scan(&res)
	return res == 1
}

func GetTopic(topic *s.Topic, topic_id string) {
	database, err_base := sql.Open("sqlite3", "./data_base.sqlite")
	if err_base != nil {
		log.Fatal(err_base)
	}
	defer database.Close()

	ctx := context.Background()
	request := `SELECT Topic.id, Topic.creator, Topic.name, Topic.content, Users.pic FROM Topic INNER JOIN Users ON Users.username = Topic.creator WHERE Topic.id=` + fmt.Sprint(topic_id)

	rows, err := database.QueryContext(ctx, request)

	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		var Topic s.Topic
		err := rows.Scan(&Topic.Id, &Topic.Creator.Username, &Topic.Name, &Topic.Content, &Topic.Creator.Pic)
		if err != nil {
			fmt.Println(err)
		}
		*topic = Topic
	}
	defer rows.Close()
}

func GetAllTopic(topics *[]s.Topic) {
	database, err_base := sql.Open("sqlite3", "./data_base.sqlite")
	if err_base != nil {
		log.Fatal(err_base)
	}
	defer database.Close()

	ctx := context.Background()
	request := `SELECT DISTINCT * FROM Topic`

	rows, err := database.QueryContext(ctx, request)

	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		var Topic s.Topic
		var catTopicStr string
		err := rows.Scan(&Topic.Id, &Topic.Creator.Username, &Topic.Name, &Topic.Picture, &Topic.Content, &catTopicStr)
		if err != nil {
			fmt.Println(err)
		}
		Topic.Categorie = strings.Split(catTopicStr, " ")
		*topics = append(*topics, Topic)
	}
	defer rows.Close()
}

func GetTopicFollowBy(username string, topics *[]s.Topic) {
	database, err_base := sql.Open("sqlite3", "./data_base.sqlite")
	if err_base != nil {
		log.Fatal(err_base)
	}
	defer database.Close()

	ctx := context.Background()
	request := `SELECT DISTINCT Topic.id, Topic.creator, Topic.name, Topic.picture, Topic.content, Topic.categorie FROM Topic INNER JOIN Follow ON Topic.id = Follow.id_topic WHERE Follow.username = "` + username + `"`

	rows, err := database.QueryContext(ctx, request)

	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		var Topic s.Topic
		var catTopicStr string
		err := rows.Scan(&Topic.Id, &Topic.Creator.Username, &Topic.Name, &Topic.Picture, &Topic.Content, &catTopicStr)
		if err != nil {
			fmt.Println(err)
		}
		Topic.Categorie = strings.Split(catTopicStr, " ")
		*topics = append(*topics, Topic)
	}
	defer rows.Close()
}

func GetTopicWithFilter(topics *[]s.Topic, order, name, cat, username string, notfollow bool) {

	database, err_base := sql.Open("sqlite3", "./data_base.sqlite")
	if err_base != nil {
		fmt.Println(err_base)
	}
	ctx := context.Background()

	var categorie []string

	categorie = strings.Split(cat, " ")

	request := `SELECT DISTINCT Topic.id, Topic.name, Topic.picture, Topic.content, Topic.categorie FROM Topic`
	if notfollow {
		request += ` INNER JOIN Follow ON Topic.id = Follow.id_topic AND NOT Follow.username ='` + username + `'`
	}
	if name != "" {
		request += ` WHERE Topic.name LIKE '%` + name + `%'`
	}

	if len(categorie) != 0 {
		for i, element := range categorie {
			if element != "" {
				if i == 0 {
					if name == "" {
						request += ` WHERE`
					}else {
						request += ` AND`
					}
					request += ` (Topic.categorie LIKE "%`+element+`%"`
				}else {
					request += ` OR Topic.categorie LIKE "%`+element+`%"`
				}
			}
			if i == len(categorie)-1{
				request += `)`
			}
		}
	}

	if order != "" {
		request += ` ORDER BY Topic.name ` + order
	}

	fmt.Println(request)

	rows, err := database.QueryContext(ctx, request)

	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		var Topic s.Topic
		var catTopicStr string

		err := rows.Scan(&Topic.Id, &Topic.Name, &Topic.Picture, &Topic.Content, &catTopicStr)
		if err != nil {
			fmt.Println(err)
			return
		}
		Topic.Categorie = strings.Split(catTopicStr, " ")
		*topics = append(*topics, Topic)
	}
}
