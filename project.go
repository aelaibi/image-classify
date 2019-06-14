package image_classify

type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Text string `json:"text"`
	Done bool   `json:"done"`
	UserID string  `json:"user"`
}