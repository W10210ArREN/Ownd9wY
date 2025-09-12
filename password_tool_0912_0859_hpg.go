// 代码生成时间: 2025-09-12 08:59:59
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/hex"
    "errors"
    "fmt"
    "io"
    "os"
    "revel"
)

// App struct for Revel framework
type App struct {
    revel.Controller
}

// Encrypt will encrypt the given plain text with AES algorithm
func (app *App) Encrypt(plainText string) (string, error) {
    key := []byte("your-secret-key-16-bytes-long")
    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }

    nonce := make([]byte, gcm.NonceSize())
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return "", err
    }

    cipherText := gcm.Seal(nonce, nonce, []byte(plainText), nil)
    return hex.EncodeToString(cipherText), nil
}

// Decrypt will decrypt the given cipher text with AES algorithm
func (app *App) Decrypt(cipherText string) (string, error) {
    key := []byte("your-secret-key-16-bytes-long")
    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }

    cipherTextBytes, err := hex.DecodeString(cipherText)
    if err != nil {
        return "", err
    }

    nonceSize := gcm.NonceSize()
    if len(cipherTextBytes) < nonceSize {
        return "", errors.New("invalid cipher text")
    }

    nonce, cipherText := cipherTextBytes[:nonceSize], cipherTextBytes[nonceSize:]
    plainText, err := gcm.Open(nil, nonce, cipherText, nil)
    if err != nil {
        return "", err
    }

    return string(plainText), nil
}

// main function to initialize Revel framework
func main() {
    revel.RunProd()
}