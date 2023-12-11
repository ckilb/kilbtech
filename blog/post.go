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
			Id:       "orm-hate",
			Title:    "Don't hate ORMs",
			Subtitle: "... if you mean Active Records instead.",
			IsActive: true,
		},
		{
			Id:       "golang-templates",
			Title:    "Go's template engine: A quick guide",
			Subtitle: "How to work with Go's template engine",
			IsActive: true,
		},
		{
			Id:       "pocketbook-inkpad-color-2",
			Title:    "Pocketbook InkPad Color 2 Review",
			Subtitle: "Testing the InkPad Color 2 with KOReader",
			IsActive: true,
		},
		{
			Id:       "spryker",
			Title:    "Spryker",
			Subtitle: "... what it is & if you should use it",
			IsActive: false,
		},
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
