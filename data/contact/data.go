package contact

type PageData struct {
	Title      string
	Heading    string
	Message    string
	ButtonText string
}

func GetPageData() PageData {
	return PageData{
		Title:      "Contact Us",
		Heading:    "Get in Touch!",
		Message:    "This is the contact page of our simple web app.",
		ButtonText: "Send Message",
	}
}
