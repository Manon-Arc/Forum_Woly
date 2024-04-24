package main

import (
	"context"
	"database/sql"
	"fmt"
	page "forum/fonctions"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database, err_base := sql.Open("sqlite3", "./data_base.sqlite")
	if err_base != nil {
		log.Fatal(err_base)
	}
	defer database.Close()

	stmt, err := database.Prepare(`SELECT COUNT(name) FROM sqlite_schema WHERE type IN ('table','view') AND name NOT LIKE 'sqlite_%' ORDER BY 1;`)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	row, err := stmt.Query()
	if err != nil {
		fmt.Println(err)
	}
	var nTable int
	for row.Next() {
		err := row.Scan(&nTable)
		if err != nil {
			fmt.Println(err)
		}
	}
	defer row.Close()

	if nTable == 0 {
		CreatTable(database)
	}

	http.Handle("/server/", http.StripPrefix("/server/", http.FileServer(http.Dir("./server"))))

	http.HandleFunc("/connection", page.Login_register_page)
	http.HandleFunc("/profil", page.Profil)
	http.HandleFunc("/creat_topic", page.Creat_Topic)
	http.HandleFunc("/add_post", page.Creat_Post)
	http.HandleFunc("/topic", page.Topic)
	http.HandleFunc("/logout", page.LogOut)
	http.HandleFunc("/add_like", page.Add_like)
	http.HandleFunc("/add_dislike", page.Add_dislike)
	http.HandleFunc("/add_follow", page.Add_Follow)
	http.HandleFunc("/supp_topic", page.Supp_topic)
	http.HandleFunc("/supp_post", page.Supp_post)
	http.HandleFunc("/edit_post", page.Edit_post)
	http.HandleFunc("/add_moderateur", page.Add_Mode)
	http.HandleFunc("/add_ban", page.Add_Ban)
	http.HandleFunc("/", page.Home)

	port := "8080"
	fmt.Println("Startup Server on port " + port)
	err_serv := http.ListenAndServe(":"+port, nil)
	if err_serv != nil {
		return
	}
}

func CreatTable(database *sql.DB) {
	fmt.Println("Creat Table")

	ctx := context.Background()

	stmts := `CREATE TABLE Users (
		username TEXT PRIMARY KEY NOT NULL,
		mail TEXT NOT NULL,
		password TEXT NOT NULL,
		status BOOL,
		pic TEXT,
		bio TEXT,
		UNIQUE (username, mail)
	  );
		
	  CREATE TABLE Post (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		creator TEXT NOT NULL,
		topic INTEGER NOT NULL,
		content TEXT,
		pic TEXT,
		UNIQUE (id),
		FOREIGN KEY (creator) REFERENCES Users(username)
		FOREIGN KEY (topic)  REFERENCES Topic(id)
	  );
	  
	  CREATE TABLE Topic (
		id	INTEGER PRIMARY KEY AUTOINCREMENT,
		creator	TEXT NOT NULL,
		name TEXT NOT NULL,
		picture TEXT,
		content	TEXT,
		categorie	TEXT,
		FOREIGN KEY (creator) REFERENCES Users(username)
	  );
	  
	  CREATE TABLE Follow (
		username TEXT NOT NULL,
		id_topic INT NOT NULL,
		FOREIGN KEY (username) REFERENCES Users(username)
		FOREIGN KEY (id_topic) REFERENCES Topic(id)
	  );
	  
	  CREATE TABLE Jaime (
		username TEXT NOT NULL,
		id_post INT NOT NULL,
		FOREIGN KEY (username) REFERENCES Users(username)
		FOREIGN KEY (id_post) REFERENCES Post(id)
	  );
	  
	  CREATE TABLE Dislike (
		username TEXT NOT NULL,
		id_post INT NOT NULL,
		FOREIGN KEY (username) REFERENCES Users(username)
		FOREIGN KEY (id_post) REFERENCES Post(id)
	  );
	  
	  CREATE TABLE Comment (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		creator TEXT NOT NULL,
		topic TEXT NOT NULL,
		content TEXT,
		UNIQUE (id),
		FOREIGN KEY (creator) REFERENCES Users(username)
		FOREIGN KEY (topic)  REFERENCES Topic(name)
	  );
	  
	  CREATE TABLE Moderateur (
		username TEXT NOT NULL,
		id_topic INT NOT NULL,
		FOREIGN KEY (username) REFERENCES Users(username)
		FOREIGN KEY (id_topic) REFERENCES Topic(id)
	  );
	  
	  CREATE TABLE Ban (
		username TEXT NOT NULL,
		id_topic INT NOT NULL,
		FOREIGN KEY (username) REFERENCES Users(username)
		FOREIGN KEY (id_topic) REFERENCES Topic(id)
	  );
	   `

	_, err := database.ExecContext(ctx, stmts)

	if err != nil {
		fmt.Println(err)
	}
}
