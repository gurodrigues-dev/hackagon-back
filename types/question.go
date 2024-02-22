package types

type Question struct {
	Title       string
	Description string
	Date        string
	Points      int
	Language    string
}

type Response struct {
	Response string
	Output   string
	Date     string
}
