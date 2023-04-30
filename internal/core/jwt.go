package core

type JWT struct {
	Token        string `json:"jwt_token" binding:"required"`
	IsRegistered bool   `json:"-"`
}

func (t *JWT) ToResponse() JWTResponse {
	return JWTResponse{Token: t.Token, IsRegistered: t.IsRegistered}
}

type JWTResponse struct {
	Token        string `json:"jwt_token" binding:"required"`
	IsRegistered bool   `json:"is_registered"`
}
