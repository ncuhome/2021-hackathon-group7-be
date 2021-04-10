package dao

type EmailBindCache struct {
	ID    uint   `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
	Key   string `json:"key,omitempty"`
}

type EmailPasswordCache struct {
	ID  uint   `json:"id,omitempty"`
	Key string `json:"key,omitempty"`
}
