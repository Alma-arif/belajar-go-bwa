package auth

import "github.com/golang-jwt/jwt"

type Service interface {
	GenerateToken(userId int) (string, error)
}

type jwtService struct {
}

var SECRET_KEY = []byte("alma_dan_dia_kenapa_3456")

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
