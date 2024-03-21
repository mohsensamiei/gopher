package authorize

type Claims struct {
	UserID   string   `json:"user_id,omitempty" map:"id"`
	UnitID   string   `json:"unit_id,omitempty"`
	Username string   `json:"username,omitempty"`
	Name     string   `json:"name,omitempty"`
	Surname  string   `json:"surname,omitempty"`
	Email    string   `json:"email,omitempty"`
	Mobile   string   `json:"mobile,omitempty"`
	Scopes   []string `json:"scopes,omitempty"`
}
