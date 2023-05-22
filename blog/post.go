package blog

type Post struct {
	Id       string
	Title    string
	Subtitle string
	IsActive bool
}

func GetPosts() []Post {
	return []Post{
		{
			Id:       "tailwind",
			Title:    "Tailwind",
			Subtitle: "... and a brief history of CSS",
			IsActive: true,
		},
		{
			Id:       "active-record-pattern",
			Title:    "Active Record Pattern",
			Subtitle: "Benefits & Downsides",
			IsActive: true,
		},
	}
}
