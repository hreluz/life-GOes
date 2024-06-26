package main

type class struct {
	Title    string
	Duration string
}

type module struct {
	Title       string
	Description string
	Classes     []class
}

type course struct {
	Slug         string
	Img          string
	Title        string
	Name         string
	Description  string
	Average      float64
	Professor    string
	ProfessorImg string
	Price        float64
	Modules      []module
}

type gridPage struct {
	InternalTemplate string
	Courses          []course
}

type coursePage struct {
	InternalTemplate string
	Course           course
}

func loadGrid() []course {
	return []course{
		{
			Slug:         "go",
			Img:          "/data/go.png",
			Title:        "Go from scratch",
			Name:         "Go from scratch",
			Description:  "Start learning Go from scratch to start programming",
			Average:      4.7,
			Professor:    "Professor Frink",
			ProfessorImg: "/data/professor_frink.png",
			Price:        30,
			Modules: []module{
				{
					Title:       "What is Go",
					Description: "Learning what is Go",
					Classes: []class{
						{
							Title:    "History",
							Duration: "05:33",
						},
						{
							Title:    "Creators",
							Duration: "03:02",
						},
					},
				},
				{
					Title:       "Syntaxis",
					Description: "Language basics",
					Classes: []class{
						{
							Title:    "Variables",
							Duration: "05:33",
						},
						{
							Title:    "Constants",
							Duration: "03:02",
						},
					},
				},
			},
		},
		{
			Slug:         "go-poo",
			Img:          "/data/go_poo.jpg",
			Title:        "POO with Go",
			Name:         "POO with Go",
			Description:  "Learning POO with Go Learning POO with Go Learning POO with Go",
			Average:      4.8,
			Professor:    "Comic Guy",
			ProfessorImg: "/data/comic_guy.png",
			Price:        30,
		},
		{
			Slug:         "go-database",
			Img:          "/data/go_db.jpg",
			Title:        "Databases with Go",
			Name:         "Databases with Go",
			Description:  " Learn database integration with Go Learn database integration with Go",
			Average:      4.8,
			Professor:    "Professor Skinner",
			ProfessorImg: "/data/professor_skinner.png",
			Price:        30,
		},
		{
			Slug:         "go-testing",
			Img:          "/data/go_testing.png",
			Title:        "Testing with Go",
			Name:         "Testing with Go",
			Description:  "Learning to create tests on your Go projects",
			Average:      4.7,
			Professor:    "Comic Guy",
			ProfessorImg: "/data/comic_guy.png",
			Price:        24,
		},
	}
}

func getCourse(slug string) course {
	data := loadGrid()
	for _, v := range data {
		if v.Slug == slug {
			return v
		}
	}

	return course{}
}
