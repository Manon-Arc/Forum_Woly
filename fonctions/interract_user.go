package pagesfonctions

import (
	"context"
	"database/sql"
	"fmt"
	s "forum/structs"
	"log"
)

func GetUsersMemberOf(users *[]s.Users, topic_id string) {

	database, err_base := sql.Open("sqlite3", "./data_base.sqlite")
	if err_base != nil {
		log.Fatal(err_base)
	}
	defer database.Close()

	ctx := context.Background()

	request := `SELECT DISTINCT Users.username, Users.status , Users.pic FROM Users INNER JOIN Topic ON Follow.id_topic = Topic.id INNER JOIN Follow ON Follow.username = Users.username WHERE Topic.id=` + topic_id

	rows, err := database.QueryContext(ctx, request)

	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		var User s.Users
		err := rows.Scan(&User.Username, &User.Status, &User.Pic)
		if err != nil {
			fmt.Println(err)
		}

		re := `SELECT COUNT(*) FROM Moderateur WHERE id_topic =` + topic_id + ` AND username="`+User.Username+`"`

		r, err := database.QueryContext(ctx, re)

		if err != nil {
			fmt.Println(err)
			return
		}

		for r.Next() {
			var res int
			err = r.Scan(&res)
			if err != nil {
				fmt.Println(err)
				return
			}

			if res == 1 {
				User.State = "Moderator"
			}else {
				User.State = "Member"
			}
		}

		*users = append(*users, User)
	}
	defer rows.Close()
}

func GetUsersBanOf(users *[]s.Users, topic_id string) {
	database, err_base := sql.Open("sqlite3", "./data_base.sqlite")
	if err_base != nil {
		log.Fatal(err_base)
	}
	defer database.Close()

	ctx := context.Background()

	request := `SELECT DISTINCT Users.username, Users.status , Users.pic FROM Users INNER JOIN Topic ON Ban.id_topic = Topic.id INNER JOIN Ban ON Ban.username = Users.username WHERE Topic.id=` + topic_id

	rows, err := database.QueryContext(ctx, request)

	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		var User s.Users

		err := rows.Scan(&User.Username, &User.Status, &User.Pic)
		if err != nil {
			fmt.Println(err)
		}
		*users = append(*users, User)
	}
	defer rows.Close()
}

func GetUsersModOf(users *[]s.Users, topic_id string) {

	database, err_base := sql.Open("sqlite3", "./data_base.sqlite")
	if err_base != nil {
		log.Fatal(err_base)
	}
	defer database.Close()

	ctx := context.Background()

	request := `SELECT DISTINCT Users.username, Users.pic FROM Users INNER JOIN Topic ON Moderateur.id_topic = Topic.id INNER JOIN Moderateur ON Moderateur.username = Users.username WHERE Topic.id=` + topic_id

	rows, err := database.QueryContext(ctx, request)

	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		var User s.Users

		err := rows.Scan(&User.Username, &User.Pic)
		if err != nil {
			fmt.Println(err)
		}
		*users = append(*users, User)
	}
	defer rows.Close()
}
