package services

// import (
// 	"time"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/google/uuid"
// 	"gitlab.com/m8851/pmo-echo-api/config"
// 	"gitlab.com/m8851/pmo-echo-api/handlers/dto"
// 	"gitlab.com/m8851/pmo-echo-api/helper"
// 	"gitlab.com/m8851/pmo-echo-api/models/entity"
// 	"golang.org/x/crypto/bcrypt"
// )

// func GenerateAuthToken(account *entity.Account) string {
// 	claims := &dto.AuthCustomClaims{
// 		StandardClaims: jwt.StandardClaims{
// 			IssuedAt:  time.Now().Unix(),
// 			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
// 		},
// 		IdAccount: account.IdAccount,
// 		Email:     account.Email,
// 		Role:      account.Role,
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return encodeJwt(token)
// }

// func GenerateVerifyEmailToken(idaccount uuid.UUID, email string) (string, expireAt string) {
// 	claims := &dto.VerifyEmailClaims{
// 		StandardClaims: jwt.StandardClaims{
// 			IssuedAt: time.Now().Unix(),
// 		},
// 		IdAccount: idaccount,
// 		Email:     email,
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return encodeJwt(token), expireAt
// }

// func GenerateResetPasswordToken(email string) (string, expireAt string) {
// 	claims := &dto.ResetPasswordClaims{
// 		StandardClaims: jwt.StandardClaims{
// 			IssuedAt:  time.Now().Unix(),
// 			ExpiresAt: time.Now().Add((7 * 24) * time.Hour).Unix(),
// 		},
// 		Email: email,
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return encodeJwt(token), expireAt
// }

// func GenerateUnsubscribeToken(email string) string {
// 	claims := &dto.UnsubscribeNewsLetterClaims{
// 		StandardClaims: jwt.StandardClaims{
// 			IssuedAt: time.Now().Unix(),
// 		},
// 		Email: email,
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return encodeJwt(token)
// }

// func encodeJwt(token *jwt.Token) string {
// 	//encoded string
// 	t, err := token.SignedString([]byte(config.AppConfig[config.ApiSecretKey]))
// 	helper.PanicIfError(err, "")

// 	return t
// }

// func CheckPasswordHash(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }

// func GeneratePassword(password string) string {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
// 	helper.PanicIfError(err, "")
// 	return string(bytes)
// }
