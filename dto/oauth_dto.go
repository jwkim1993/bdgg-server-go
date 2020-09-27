package dto

type OAuthDTO struct {
	UserId       int64  `json:"userId,omitempty"`
	NativeKey    string `json:"nativeKey,omitempty"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken,omitempty"`
	Provider     string `json:"provider,omitempty"`
}
