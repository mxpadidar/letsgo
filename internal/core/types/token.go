package types

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewTokenPair(access, refresh string) *TokenPair {
	return &TokenPair{AccessToken: access, RefreshToken: refresh}
}
