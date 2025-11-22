package response

type ClassDto struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

type ClassDetailDto struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Instructor  string `json:"instructor"`
	Category    string `json:"category"`
}
