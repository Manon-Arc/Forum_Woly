CREATE TABLE Users (
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