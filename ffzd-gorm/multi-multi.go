package main

func main() {
	// DB.AutoMigrate(&Tag{}, &Video{})

	DB.Create(&Video{
		Title: "font_end-study",
		Tags: []Tag{
			{
				Name: "http",
			},
			{
				Name: "JS",
			},
		},
	})
}
