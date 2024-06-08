package home

// define data on page
type PageData struct {
	Title      string
	Heading    string
	Message    string
	ButtonText string
}

func GetPageData() PageData {
	return PageData{
		Title:      "Home Page",
		Heading:    "Welcome to my page about my work",
		Message:    "My work will be here",
		ButtonText: "Learn more",
	}
}
