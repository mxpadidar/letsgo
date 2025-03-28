package dtos

type TokenDTO struct {
	Token string `json:"token"`
}

func NewTokenDTO(token string) *TokenDTO {
	return &TokenDTO{Token: token}
}

type TokenPayload struct {
	UserID int `json:"user_id"`
}

func NewTokenPayload(userID int) *TokenPayload {
	return &TokenPayload{UserID: userID}
}
