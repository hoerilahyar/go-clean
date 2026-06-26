package entity

type Token struct {
	AccessToken  string
	RefreshToken string
	TokenType    string
	ExpiresIn    int64
}
