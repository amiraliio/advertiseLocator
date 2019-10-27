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

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const tokenSplitter string = "_"

//EncryptedToken data model
type EncryptedToken struct {
	Token      string             `json:"token"`
	ExpireDate primitive.DateTime `json:"expireDate"`
}

//DecryptedToken data model
type DecryptedToken struct {
	ID         string             `json:"id"`
	TokenType  string             `json:"tokenType"`
	ExpireDate primitive.DateTime `json:"expireDate"`
}

//EncodeToken encode
func EncodeToken(id, tokenType string, expireDate primitive.DateTime) (*EncryptedToken, error) {
	time := expireDate.Time().Unix()
	token := []byte(id + tokenSplitter + tokenType + tokenSplitter + strconv.FormatInt(time, 10))
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
	encryptedToken := new(EncryptedToken)
	encryptedToken.Token = tokenInBase64
	encryptedToken.ExpireDate = expireDate
	return encryptedToken, nil
}

//DecodeToken decode
func DecodeToken(token string) (*DecryptedToken, error) {
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
		decryptedToken := new(DecryptedToken)
		decryptedToken.ID = splittedToken[0]
		decryptedToken.TokenType = splittedToken[1]
		timestamp, err := strconv.ParseInt(splittedToken[2], 10, 64)
		if err != nil {
			return nil, err
		}
		decryptedToken.ExpireDate = primitive.NewDateTimeFromTime(time.Unix(timestamp, 0))
		return decryptedToken, nil
	}
	return nil, errors.New("data of token is not right")
}
