package blog

type Post struct {
	Id       string
	Title    string
	Subtitle string
}

func GetPosts() []Post {
	return []Post{
		{
			Id:       "active-record-pattern",
			Title:    "Active Record Pattern",
			Subtitle: "Benefits & Downsides",
		},
	}
}
