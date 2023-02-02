package authorize

type Claims struct {
	UserID   string `json:"user_id" map:"id"`
	UnitID   string `json:"unit_id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
}
