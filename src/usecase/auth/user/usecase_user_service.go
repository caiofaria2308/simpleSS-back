package usecase_user

import (
	entity "main/entity/auth"
	"time"

	"github.com/golang-jwt/jwt"
)

const SECRET_KEY = "9an0afx$thw)k9#y*_d9-ch^r&a6ndi#x#dwu^52zbqw=hso(9"

type SignedDetails struct {
	ID    string
	Name  string
	Email string
	jwt.StandardClaims
}

type UseCaseUser struct {
	repo IRepositoryUser
}

func NewService(repository IRepositoryUser) *UseCaseUser {
	return &UseCaseUser{repo: repository}
}

func (s *UseCaseUser) LoginUser(email string, password string) (*entity.EntityUser, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	err = user.ValidatePassword(password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UseCaseUser) Create(user entity.EntityUser) (*entity.EntityUser, error) {
	err := user.Validate()
	if err != nil {
		return nil, err
	}
	return s.repo.Create(user)
}

func (s *UseCaseUser) Update(user entity.EntityUser) (*entity.EntityUser, error) {
	err := user.Validate()
	if err != nil {
		return nil, err
	}
	return s.repo.Update(user)
}

func (s *UseCaseUser) Delete(id string) error { return s.repo.Delete(id) }

func (s *UseCaseUser) GetByEmail(email string) (*entity.EntityUser, error) {
	return s.repo.GetByEmail(email)
}

func (s *UseCaseUser) GetUserByToken(token string) (*entity.EntityUser, error) {
	claims, err := ValidateToken(token)
	if err != nil {
		return nil, err
	}
	user, err := s.repo.GetByEmail(claims.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func JWTTokenGenerator(u entity.EntityUser) (signedToken string, signedRefreshToken string, err error) {

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

func ValidateToken(signedToken string) (claims *SignedDetails, err error) {
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
