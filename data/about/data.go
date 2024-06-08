package about

type PageData struct {
	Title      string
	Heading    string
	Message    string
	ButtonText string
}

func GetPageData() PageData {
	return PageData{
		Title:      "About Us",
		Heading:    "Learn More About Us!",
		Message:    "This is the about page of our simple web app.",
		ButtonText: "Contact Us",
	}
}
