package user

type Driver struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
	TaxiType string `json:"taxiType"`
	Status   string `json:"status"`
}

type AuthDetails struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
type ResponseAuthDetails struct {
	Phone        string `json:"phone"`
	HashPassword string `json:"HashPassword"`
}
