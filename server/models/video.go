package models

type Video struct {
	ID           string `json:"id"`
	Status       string `json:"status"`
	Name         string `json:"name"`
	PreviewImage string `json:"preview_image"`
}
