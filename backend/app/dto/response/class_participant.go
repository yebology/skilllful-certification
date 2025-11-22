package response

type ParticipantClassDto struct {
	Id       uint   `json:"id"`
	ClassId  uint   `json:"classId"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

type ParticipantClassDetailDto struct {
	Id          uint   `json:"id"`
	ClassId     uint   `json:"classId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Instructor  string `json:"instructor"`
	Category    string `json:"category"`
}

type ClassParticipantDto struct {
	Id            uint   `json:"id"`
	ParticipantId uint   `json:"participantId"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Gender        string `json:"gender"`
}

type ClassParticipantDetailDto struct {
	Id            uint   `json:"id"`
	ParticipantId uint   `json:"participantId"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Gender        string `json:"gender"`
	PhoneNumber   string `json:"phoneNumber"`
	BirthDate     string `json:"birthDate"`
}
