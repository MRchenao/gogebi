package Services

import "gebi/app/Http/Serializer"

func GetToken(id int) string {
	tokenString, err := jwtService.SignedData(id, "private_key")
	if err != nil {
		Serializer.Err(403, "获取token失败", err)
	}

	return tokenString
}
