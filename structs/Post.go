package structs

type Post struct {
	Id       int
	Topics   Topic
	Creator  Users
	Content  string
	Pic 	 string
	Like     int
	Dislike  int
	ILike    bool
	IDislike bool
	IsMine   bool
}
