package config

import "time"

var (
	JWTSecret     = []byte("CHANGE_ME_SUPER_SECRET")
	JWTIssuer     = "api-gateway"
	JWTAudience   = "public-web"
	JWTExpireTime = time.Minute * 10
)
