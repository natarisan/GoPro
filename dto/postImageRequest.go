package dto

type PostImageRequest struct {
	CustomerId  string  `json:"customer_id"`
	Image       string  `json:"image"`
}