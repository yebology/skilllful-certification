package request

type AddClassParticipantDto struct {
	ParticipantId uint `json:"participantId" validate:"required"`
	ClassId       uint `json:"classId" validate:"required"`
}
