package posts

type post struct {
	Id              int
	Title           string
	Content         string
	Author          string
	PublicationDate string
}

func GetPost(id int) post{
	return post{
		id,
		"tagada",
		"Pouet pouet tagada",
		"Naouak",
		"2018-04-15T23:22+02:00",
	}
}

func GetPosts() []post{
	return []post{
		GetPost(1),
		GetPost(2),
		GetPost(3),
	}
}
