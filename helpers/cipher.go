package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

//TODO inner struct instead of modles.API for both client and apikey

const tokenSplitter string = "_"

type Cipher struct {
	Key        string
	Type       string
	Token      string
	CreatedAt  primitive.DateTime
	ExpireDate primitive.DateTime
}

//EncodeToken encode
func EncodeToken(id, tokenType string, expireTime int) (*Cipher, error) {
	createdAt := time.Now()
	token := []byte(id + tokenSplitter + tokenType + tokenSplitter + strconv.FormatInt(createdAt.Unix(), 10))
	block, err := aes.NewCipher([]byte(viper.GetString("APP.KEY")))
	if err != nil {
		return nil, err
	}
	cipherText := make([]byte, aes.BlockSize+len(token))
	//identifier vector from 0 to before index 16
	iv := cipherText[:aes.BlockSize]
	//generate random number for iv
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cipherText[aes.BlockSize:], token)
	tokenInBase64 := base64.StdEncoding.EncodeToString(cipherText)
	encryptedToken := new(Cipher)
	encryptedToken.Token = tokenInBase64
	encryptedToken.CreatedAt = primitive.NewDateTimeFromTime(createdAt)
	encryptedToken.ExpireDate = primitive.NewDateTimeFromTime(createdAt.AddDate(0, 0, expireTime))
	return encryptedToken, nil
}

//DecodeToken decode
func DecodeToken(token string) (*Cipher, error) {
	tokenInByte, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}
	if len(tokenInByte) < aes.BlockSize {
		return nil, errors.New("token length is too short")
	}
	block, err := aes.NewCipher([]byte(viper.GetString("APP.KEY")))
	if err != nil {
		return nil, err
	}
	iv := tokenInByte[:aes.BlockSize]
	tokenInByte = tokenInByte[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(tokenInByte, tokenInByte)
	token = string(tokenInByte)
	splittedToken := strings.Split(token, tokenSplitter)
	if len(splittedToken) == 3 {
		decryptedToken := new(Cipher)
		decryptedToken.Key = splittedToken[0]
		decryptedToken.Type = splittedToken[1]
		timestamp, err := strconv.ParseInt(splittedToken[2], 10, 64)
		if err != nil {
			return nil, err
		}
		decryptedToken.CreatedAt = primitive.NewDateTimeFromTime(time.Unix(timestamp, 0))
		return decryptedToken, nil
	}
	return nil, errors.New("data of token is not right")
}

//HashPassword helper
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

//CheckPasswordHash handler
func CheckPasswordHash(password, hash string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false
	}
	return true
}
