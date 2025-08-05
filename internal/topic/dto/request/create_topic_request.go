package request

type CreateTopicRequest struct {
	Title string `json:"title" binding:"required"`
	Icon  string `json:"icon" binding:"required"`
}
