package customer

type CustomerPreview struct {
	ID       uint              `json:"id"`
	Name     string            `json:"name"`
	Contacts map[string]string `json:"contacts"`
}

