package jwthelper

import (
	"github.com/gbrlsnchs/jwt/v3"
)

var JwtKey = jwt.NewHS256([]byte("helmdashboard"))
