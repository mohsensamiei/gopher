package authorize

type Claims struct {
	ID       string   `json:"id,omitempty"`
	Unit     string   `json:"unit,omitempty"`
	Username string   `json:"username,omitempty"`
	Name     string   `json:"name,omitempty"`
	Surname  string   `json:"surname,omitempty"`
	Email    string   `json:"email,omitempty"`
	Phone    string   `json:"phone,omitempty"`
	Scopes   []string `json:"scopes,omitempty"`
}
