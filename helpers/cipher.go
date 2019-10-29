package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/amiraliio/advertiselocator/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const tokenSplitter string = "_"

//EncodeToken encode
func EncodeToken(id, tokenType, expire string) (*models.API, error) {
	createdAt := time.Now()
	token := []byte(id + tokenSplitter + tokenType + tokenSplitter + strconv.FormatInt(createdAt.Unix(), 10))
	block, err := aes.NewCipher([]byte(os.Getenv("APP_KEY")))
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
	encryptedToken := new(models.API)
	encryptedToken.Token = tokenInBase64
	encryptedToken.CreatedAt = primitive.NewDateTimeFromTime(createdAt)
	expireTime, err := strconv.Atoi(expire)
	if err != nil {
		return nil, err
	}
	encryptedToken.ExpireDate = primitive.NewDateTimeFromTime(createdAt.AddDate(0, 0, expireTime))
	return encryptedToken, nil
}

//DecodeToken decode
func DecodeToken(token string) (*models.API, error) {
	tokenInByte, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}
	if len(tokenInByte) < aes.BlockSize {
		return nil, errors.New("token length is too short")
	}
	block, err := aes.NewCipher([]byte(os.Getenv("APP_KEY")))
	if err != nil {
		return nil, err
	}
	iv := tokenInByte[:aes.BlockSize]
	tokenInByte = tokenInByte[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(tokenInByte, tokenInByte)
	token = string(tokenInByte)
	splittedToken := strings.Split(token, tokenSplitter)
	if splittedToken != nil && len(splittedToken) == 3 {
		decryptedToken := new(models.API)
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
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}
