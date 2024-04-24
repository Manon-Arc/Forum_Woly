package structs

type Topic struct {
	Id        int
	Creator   Users
	Name      string
	Picture   string
	Content   string
	Categorie []string
}
