package characters

// Character defines the info of a character
type Character struct {
	Name             string `json:"name"`
	PhontName        string `json:"phontName"`
	Description      string `json:"description"`
	ProfileImagePath string `json:"profileImagePath"`
}
