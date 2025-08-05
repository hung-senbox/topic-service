package response

import "time"

type TopicResponse struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Icon      string    `json:"icon"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
