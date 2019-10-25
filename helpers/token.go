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

//TODO create a interface for each helpers

const tokenSplitter string = "_"

//EncryptedToken data model
type EncryptedToken struct {
	Token string
	time  int64
}

//DecryptedToken data model
type DecryptedToken struct {
	ID   string
	Type string
	Time int64
}

//GenerateToken helper
func GenerateToken(id, tokenType string) (*EncryptedToken, error) {
	time := time.Now().Unix()
	init := id + tokenSplitter + tokenType + tokenSplitter + string(time)
	token, err := encode(init)
	if err != nil {
		return nil, err
	}
	tokenModel := new(EncryptedToken)
	tokenModel.Token = token
	tokenModel.time = time
	return tokenModel, nil
}

//token encode
func encode(token string) (string, error) {
	block, err := aes.NewCipher([]byte(os.Getenv("APP_KEY")))
	if err != nil {
		return "", err
	}
	tokenInBase64 := base64.StdEncoding.EncodeToString([]byte(token))
	cipherText := make([]byte, aes.BlockSize+len(tokenInBase64))
	iv := cipherText[:aes.BlockSize] //identifier vector
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cipherText[aes.BlockSize:], []byte(tokenInBase64))
	return string(cipherText), nil
}

//token decode
func decode(token string) (*DecryptedToken, error) {
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
	decryptedToken := new(DecryptedToken)
	splittedToken := strings.Split(token, tokenSplitter)
	if splittedToken != nil && len(splittedToken) == 3 {
		decryptedToken.ID = splittedToken[0]
		decryptedToken.Type = splittedToken[1]
		time, err := strconv.ParseInt(splittedToken[2], 10, 64)
		if err != nil {
			return nil, err
		}
		decryptedToken.Time = time
		return decryptedToken, nil
	}
	return nil, errors.New("data of token is not right")
}
