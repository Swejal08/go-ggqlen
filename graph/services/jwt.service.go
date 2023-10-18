package services

func Add() string {
	return "1"
}

// func SignAccessToken(payload TokenMetadata) (string, error) {
// 	tokenSecret := GoDotEnv("TOKEN_SECRET")
// 	return signToken(map[string]interface{}{
// 		"id":    payload.ID,
// 		"email": payload.Email,
// 	}, tokenSecret, 200*time.Minute)

// }

// func signToken(Data map[string]interface{}, secret string, ExpiredAt time.Duration) (string, error) {
// 	claims := jwt.MapClaims{}

// 	for key, value := range Data {
// 		claims[key] = value
// 	}
// 	claims["exp"] = time.Now().Add(ExpiredAt).Unix()
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString([]byte(secret))
// }

// func VerifyToken(accessToken, secret string) (*TokenMetadata, error) {
// 	claims := jwt.MapClaims{}
// 	token, err := jwt.ParseWithClaims(accessToken, &claims, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(secret), nil
// 	})

// 	if err != nil || !token.Valid {
// 		return nil, err
// 	}
// 	return &TokenMetadata{ID: fmt.Sprint(claims["id"]), Email: fmt.Sprint(claims["email"])}, nil
// }
