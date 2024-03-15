package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)


func SignAccessToken(userId int) (token string, err error) {
	err = godotenv.Load("app.env")
	if err != nil {
		return "", err
	}
	var SecretKey = os.Getenv("ACCESS_KEY_SECRET")


	var jwtToken = jwt.New(jwt.SigningMethodHS256)

	claims := jwtToken.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := jwtToken.SignedString([]byte(SecretKey))
	if err != nil {
		return "", errors.New("Error in generating token")
		
	}

	return tokenString, nil


}

func SignRefreshToken(userId int) (token string, err error) {
	err = godotenv.Load("app.env")
	if err != nil {
		return "", err
	}
	var SecretKey = os.Getenv("REFRESH_KEY_SECRET")


	var jwtToken = jwt.New(jwt.SigningMethodHS256)

	claims := jwtToken.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := jwtToken.SignedString([]byte(SecretKey))
	if err != nil {
		return "", errors.New("Error in generating token")
		
	}

	return tokenString, nil


}

func VerifyToken(tokenString string) (intUserId *int, err error) {
	err = godotenv.Load("app.env")
	if err != nil {
		return nil, err
	}
	var SecretKey = os.Getenv("ACCESS_KEY_SECRET")

	claims := jwt.MapClaims{}
	
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Lấy thời gian hết hạn từ claims
	expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)

	// Kiểm tra xem token có hết hạn chưa
	if time.Now().After(expirationTime) {
		return nil, errors.New("token is expired")
	} 

	userId, ok := claims["userId"].(float64)
	if !ok {
		return nil, errors.New("invalid token")
	
	}

	intuserId := int(userId)

	return &intuserId, nil


}


func RefreshToken(tokenString string) (accessToken string, err error) {
	err = godotenv.Load("app.env")
	if err != nil {
		return "", err
	}
	var SecretKey = os.Getenv("REFRESH_KEY_SECRET")

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("invalid token")
	}

	// Lấy thời gian hết hạn từ claims
	expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
	userID, ok := claims["userId"].(float64)
	
	// Kiểm tra xem token có hết hạn chưa
	if time.Now().After(expirationTime) {
		return "", errors.New("token is expired")
	}
	if !ok {
		return "", errors.New("invalid token")
	
	}
 
	accessToken, err = SignAccessToken(int(userID))

	if err != nil {
		return "", err
	}

	return accessToken, nil


}