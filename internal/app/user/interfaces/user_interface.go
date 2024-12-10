package interfaces

type UserDTO struct {
	ID      string  `json:"id"`
	Email   string  `json:"email"`
	Name    string  `json:"name"`
	Surname *string `json:"surname,omitempty"`
}

type UserAddressDTO struct {
	Id           string `json:"id"`
	UserId       string `json:"user_id"`
	Cep          string `json:"cep"`
	Street       string `json:"street"`
	Neighborhood string `json:"neighborhood"`
	Uf           string `json:"uf"`
	City         string `json:"city"`
	Number       string `json:"number"`
	Complement   string `json:"complement,omitempty"`
}

type CreateUserPayload struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Surname string `json:"surname,omitempty"`
}

type CreateUserAddressPayload struct {
	UserId       string `json:"user_id"`
	Cep          string `json:"cep"`
	Street       string `json:"street"`
	Neighborhood string `json:"neighborhood"`
	Uf           string `json:"uf"`
	City         string `json:"city"`
	Number       string `json:"number"`
	Complement   string `json:"complement,omitempty"`
}
