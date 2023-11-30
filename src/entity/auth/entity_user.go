package entity

import (
	"main/utils"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

const SECRET_KEY = "9an0afx$thw)k9#y*_d9-ch^r&a6ndi#x#dwu^52zbqw=hso(9"

var validate *validator.Validate = validator.New()

type SignedDetails struct {
	ID    string
	Name  string
	Email string
	jwt.StandardClaims
}

type EntityUser struct {
	ID        string    `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,min=6"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateUser(userParam EntityUser) (*EntityUser, error) {
	now := time.Now()
	var password string
	var err error

	password, err = utils.GeneratePassword(userParam.Password)
	if err != nil {
		return nil, err
	}
	u := &EntityUser{
		ID:        utils.GenerateID(),
		Name:      userParam.Name,
		Password:  password,
		CreatedAt: now,
		UpdatedAt: now,
	}
	return u, nil
}

func (u *EntityUser) ValidatePassword(p string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))

	if err != nil {
		return err
	}

	return nil
}

func (u *EntityUser) Validate() error {
	return validate.Struct(u)
}

func (u *EntityUser) JWTTokenGenerator() (signedToken string, signedRefreshToken string, err error) {

	claims := SignedDetails{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	refreshClaims := SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7 * 365).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}

func (u *EntityUser) ValidateToken(signedToken string) (claims *SignedDetails, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {

		return nil, err
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {

		return nil, err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {

		return nil, err
	}

	return claims, nil
}

func (u *EntityUser) UpdatePassword(newPassword string) error {
	hash, err := utils.GeneratePassword(newPassword)
	if err != nil {
		return err
	}

	u.Password = hash

	return nil
}
