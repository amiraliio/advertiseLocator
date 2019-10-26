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
	token := id + tokenSplitter + tokenType + tokenSplitter + string(time)
	block, err := aes.NewCipher([]byte(os.Getenv("APP_KEY")))
	if err != nil {
		return nil, err
	}
	tokenInBase64 := base64.StdEncoding.EncodeToString([]byte(token))
	cipherText := make([]byte, aes.BlockSize+len(tokenInBase64))
	iv := cipherText[:aes.BlockSize] //identifier vector
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cipherText[aes.BlockSize:], []byte(tokenInBase64))
	encryptedToken := new(EncryptedToken)
	encryptedToken.Token = string(cipherText)
	encryptedToken.ExpireDate = expireDate
	return encryptedToken, nil
}

//DecodeToken decode
func DecodeToken(token string) (*DecryptedToken, error) {
	block, err := aes.NewCipher([]byte(os.Getenv("APP_KEY")))
	if err != nil {
		return nil, err
	}
	tokenText := []byte(token)
	if len(tokenText) < aes.BlockSize {
		return nil, errors.New("token length is too short")
	}
	iv := tokenText[:aes.BlockSize]
	tokenText = tokenText[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(tokenText, tokenText)
	data, err := base64.StdEncoding.DecodeString(string(tokenText))
	if err != nil {
		return nil, err
	}
	token = string(data)
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
