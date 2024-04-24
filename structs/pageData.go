package structs

type TopicPage struct {
	CurrentUser Users
	Topic       Topic
	Members     []Users
	Moderateurs []Users
	Bans        []Users
	Posts       []Post
	Isconnected bool 
	IsMine      bool 
	ImMember    bool
	ImMod       bool
}

type ProfilPage struct {
	CurrentUser Users   
	Topics      []Topic 
	Posts       []Post  
	Isconnected bool    
	ItsMe       bool    
}

type MainPage struct {
	CurrentUser   Users   
	All_Topics    []Topic 
	Follow_Topics []Topic 
	Posts         []Post
	Isconnected   bool 
}
