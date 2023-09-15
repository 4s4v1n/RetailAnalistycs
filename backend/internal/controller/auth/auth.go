package auth

import (
	"APG6/config"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"time"
)

const (
	RoleUnknown uint8 = iota
	RoleVisitor
	RoleAdmin
)

const tokenKey = `fsdfjdw31213shj#2k242`

var (
	tokenTTL   time.Duration
	refreshTTL time.Duration
)

type tokenClaims struct {
	jwt.StandardClaims
	RoleId uint8 `json:"role_id"`
}

type roleData struct {
	Id       uint8
	Login    string
	Password string
}

type session struct {
	RoleId      uint8
	Login       string
	Password    string
	TimeStarted time.Time
}

var roles map[string]roleData
var sessions map[string]session

func Init(cfgRoles config.Roles, cfgJwt config.Jwt) {
	roles = make(map[string]roleData, 2)
	roles[cfgRoles.VisitorCredentials.Login] = roleData{
		Id:       RoleVisitor,
		Login:    cfgRoles.VisitorCredentials.Login,
		Password: cfgRoles.VisitorCredentials.Password,
	}
	roles[cfgRoles.AdminCredentials.Login] = roleData{
		Id:       RoleAdmin,
		Login:    cfgRoles.AdminCredentials.Login,
		Password: cfgRoles.AdminCredentials.Password,
	}

	sessions = make(map[string]session, 2)
	tokenTTL = cfgJwt.TokenTTL
	refreshTTL = cfgJwt.RefreshTTL

	go func() {
		for {
			time.Sleep(time.Minute)
			for key, s := range sessions {
				if s.TimeStarted.Add(refreshTTL).Unix() >= time.Now().Unix() {
					delete(sessions, key)
				}
			}
		}
	}()
}

func roleId(login, password string) (uint8, error) {
	role, exist := roles[login]
	if !exist {
		return RoleUnknown, errors.New("role " + login + " doesn't exists")
	}
	if role.Password != password {
		return RoleUnknown, errors.New("invalid password")
	}
	return role.Id, nil
}

func DecodeToken(accessToken string) (uint8, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(tokenKey), nil
	})
	if err != nil {
		return RoleUnknown, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return RoleUnknown, errors.New("invalid claims data type")
	}
	return uint8(claims.RoleId), nil
}

func EncodeToken(login, password string) (string, error) {
	id, err := roleId(login, password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		RoleId: id,
	})

	return token.SignedString([]byte(tokenKey))
}

func RefreshToken() (string, error) {
	b := make([]byte, 32)
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	_, err := r.Read(b)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", b), err
}

func SaveSession(key, login, password string) error {
	id, err := roleId(login, password)
	if err != nil {
		return err
	}

	for k, role := range sessions {
		if role.RoleId == id {
			delete(sessions, k)
			break
		}
	}
	sessions[key] = session{
		RoleId:   id,
		Login:    login,
		Password: password,
	}
	return nil
}

func GetSession(key string) (string, string, error) {
	role, exists := sessions[key]
	if !exists {
		return "", "", errors.New("session doesn't exists")
	}
	return role.Login, role.Password, nil
}
