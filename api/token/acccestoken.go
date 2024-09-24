package token

import (
	"api/config"
	pb "api/genproto/user"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GeneratedAccessJWTToken(req *pb.LoginResponse) error {
	conf := config.Load()
	token := *jwt.New(jwt.SigningMethodHS256)

	//payload
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = req.Id
	claims["role"] = req.Role
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(30 * time.Minute).Unix()

	newToken, err := token.SignedString([]byte(conf.ACCES_KEY))
	if err != nil {
		log.Println(err)
		return err
	}

	req.Access = newToken
	return nil
}

func ValidateAccesToken(tokenStr string) (bool, error) {
	_, err := ExtractAccesClaim(tokenStr)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractAccesClaim(tokenStr string) (*jwt.MapClaims, error) {
	conf := config.Load()
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(conf.ACCES_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, err
	}

	return &claims, nil
}

func GetUserInfoFromAccessToken(accessTokenString string) (string, string, error) {
	conf := config.Load()
	refreshToken, err := jwt.Parse(accessTokenString, func(token *jwt.Token) (interface{}, error) { return []byte(conf.ACCES_KEY), nil })
	if err != nil || !refreshToken.Valid {
		return "", "", err
	}
	claims, ok := refreshToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", err
	}
	userID := claims["user_id"].(string)
	Role := claims["role"].(string)

	return userID, Role, nil
}
