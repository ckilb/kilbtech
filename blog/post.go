package blog

type Post struct {
	Id       string
	Title    string
	Subtitle string
	Date     string
	IsActive bool
}

func GetPosts() []Post {
	return []Post{
		{
			Id:       "ship-first",
			Title:    "Ship First, Design Later",
			Subtitle: "Why I stopped waiting for Figma files.",
			Date:     "2026-05-31",
			IsActive: true,
		},
		{
			Id:       "orm-hate",
			Title:    "Don't hate ORMs",
			Subtitle: "... if you mean Active Records instead.",
			Date:     "2023-12-11",
			IsActive: true,
		},
		{
			Id:       "golang-templates",
			Title:    "Go's template engine: A quick guide",
			Subtitle: "How to work with Go's template engine",
			Date:     "2023-12-11",
			IsActive: true,
		},
		{
			Id:       "pocketbook-inkpad-color-2",
			Title:    "Pocketbook InkPad Color 2 Review",
			Subtitle: "Testing the InkPad Color 2 with KOReader",
			Date:     "2023-07-12",
			IsActive: true,
		},
		{
			Id:       "spryker",
			Title:    "Spryker",
			Subtitle: "... what it is & if you should use it",
			Date:     "2023-05-22",
			IsActive: false,
		},
		{
			Id:       "tailwind",
			Title:    "Tailwind",
			Subtitle: "... and a brief history of CSS",
			Date:     "2023-05-19",
			IsActive: true,
		},
		{
			Id:       "active-record-pattern",
			Title:    "Active Record Pattern",
			Subtitle: "Benefits & Downsides",
			Date:     "2023-05-19",
			IsActive: true,
		},
	}
}
