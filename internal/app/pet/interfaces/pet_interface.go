package interfaces

type PetDTO struct {
	ID           int     `json:"id"`
	UserId       int     `json:"user_id"`
	BreedId      int     `json:"breed_id"`
	TypeId       int     `json:"type_id"`
	Name         string  `json:"name"`
	Gender       string  `json:"gender"`
	Castrated    bool    `json:"castrated"`
	Weight       float32 `json:"weight,omitempty"`
	ProfileImage string  `json:"profile_image,omitempty"`
}

type CreatePetPayload struct {
	UserId       string  `json:"user_id"`
	BreedId      int     `json:"breed_id"`
	TypeId       int     `json:"type_id"`
	Name         string  `json:"name"`
	Gender       string  `json:"gender"`
	Castrated    bool    `json:"castrated"`
	Weight       float32 `json:"weight,omitempty"`
	ProfileImage string  `json:"profile_image,omitempty"`
}
