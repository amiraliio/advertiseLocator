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
)

//Token interface
type Token interface {
	Encode(id, tokenType string) (*EncryptedToken, error)
	Decode(token string) (*DecryptedToken, error)
}

const tokenSplitter string = "_"

//EncryptedToken data model
type EncryptedToken struct {
	Token string `json:"token"`
	Time  int64  `json:"time"`
}

//DecryptedToken data model
type DecryptedToken struct {
	ID        string `json:"id"`
	TokenType string `json:"tokenType"`
	Time      int64  `json:"time"`
}

type tokenService struct{}

//Encode encode
func (tokenService *tokenService) Encode(id, tokenType string) (*EncryptedToken, error) {
	time := time.Now().Unix()
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
	tokenModel := new(EncryptedToken)
	tokenModel.Token = string(cipherText)
	tokenModel.Time = time
	return tokenModel, nil
}

//Decode decode
func (tokenService *tokenService) Decode(token string) (*DecryptedToken, error) {
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
		time, err := strconv.ParseInt(splittedToken[2], 10, 64)
		if err != nil {
			return nil, err
		}
		decryptedToken.Time = time
		return decryptedToken, nil
	}
	return nil, errors.New("data of token is not right")
}
