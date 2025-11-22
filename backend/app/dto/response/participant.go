package response

type ParticipantDto struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
}

type ParticipantDetailDto struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Gender      string `json:"gender"`
	PhoneNumber string `json:"phoneNumber"`
	BirthDate   string `json:"birthDate"`
}
