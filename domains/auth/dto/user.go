package dto

type User struct {
	ID        string `json:"id" `
	Email     string `json:"email" `
	Name      string `json:"name" `
	AvatarUrl string `json:"avatar_url"`
}
