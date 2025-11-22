package request

type ParticipantDto struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required"`
	GenderId    uint   `json:"genderId" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	BirthDate   string `json:"birthDate" validate:"required"`
}
